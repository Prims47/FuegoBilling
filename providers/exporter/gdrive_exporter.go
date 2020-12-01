package providersExporter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"
)

type GoogleDriveExporter struct{}

func (g *GoogleDriveExporter) Save(fileName string, data []byte) error {
	if os.Getenv("DRIVE_CLIENT_ID") == "" ||
		os.Getenv("DRIVE_CLIENT_SECRET") == "" ||
		os.Getenv("DRIVE_REDIRECT_URL") == "" ||
		os.Getenv("DRIVE_AUTH_URL") == "" ||
		os.Getenv("DRIVE_TOKEN_URL") == "" {
		return errors.New("We must provide valid env var like 'DRIVE_CLIENT_ID', 'DRIVE_CLIENT_SECRET', 'DRIVE_REDIRECT_URL, DRIVE_AUTH_URL and DRIVE_TOKEN_URL'")
	}

	config := createOauthConfig(
		os.Getenv("DRIVE_CLIENT_ID"),
		os.Getenv("DRIVE_CLIENT_SECRET"),
		os.Getenv("DRIVE_REDIRECT_URL"),
		os.Getenv("DRIVE_AUTH_URL"),
		os.Getenv("DRIVE_TOKEN_URL"),
	)

	client := getClient(config)

	srv, err := drive.New(client)
	if err != nil {
		return errors.New("Unable to retrieve Drive client")
	}

	f := &drive.File{
		MimeType: "application/pdf",
		Name:     fileName,
		Parents:  []string{"root"},
	}

	body := bytes.NewReader(data)

	_, errCreate := srv.Files.Create(f).Media(body).Do()
	if errCreate != nil {
		return errors.New("Could not create file")
	}

	return nil
}

func (a *GoogleDriveExporter) CanSave(exporterProviderName string) bool {
	if "drive" == exporterProviderName ||
		"Drive" == exporterProviderName ||
		"google" == exporterProviderName ||
		"Google" == exporterProviderName {
		return true
	}

	return false
}

func createOauthConfig(clientID string, clientSecret string, redirectURL string, authURL string, tokenURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{drive.DriveFileScope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}

	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

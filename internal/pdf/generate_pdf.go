package pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/prims47/FuegoBilling/internal/model"
	"github.com/prims47/FuegoBilling/internal/services"
)

type BillingPDFInterface interface {
	CreatePDF()
}

type BillingPDF struct {
	PdfName       string
	PdfPath       string
	Account       model.Account
	Customer      model.Customer
	Service       model.Service
	FormatInt     services.FormatIntInterface
	FormatFloat   services.FormatFloatInterface
	BillingNumber string
	DateTo        string
}

func NewBillingPDF(pdfPath string,
	pdfName string,
	account model.Account,
	customer model.Customer,
	service model.Service,
	formatInt services.FormatIntInterface,
	formatFloat services.FormatFloatInterface,
	billingNumber string,
	dateTo string) BillingPDFInterface {
	return &BillingPDF{PdfName: pdfName,
		PdfPath:       pdfPath,
		Account:       account,
		Customer:      customer,
		Service:       service,
		FormatInt:     formatInt,
		FormatFloat:   formatFloat,
		BillingNumber: billingNumber,
		DateTo:        dateTo}
}

func handleHeader(pdf *gofpdf.Fpdf, b *BillingPDF) {
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetHeaderFuncMode(func() {
		// pdf.Image("logo.png", 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "", 22)
		pdf.SetTextColor(239, 114, 56)
		pdf.Cell(80, 20, "Facture")
		pdf.SetFont("Arial", "", 11)
		pdf.Ln(20)
		pdf.SetTextColor(0, 0, 0)

		pdf.Cell(10, 10, tr("Référence de facture : "+b.BillingNumber))
		pdf.Ln(8)
		pdf.Cell(10, 10, tr(fmt.Sprintf("Émise le %s", b.DateTo)))
	}, true)
}

func handleFooter(pdf *gofpdf.Fpdf, b *BillingPDF) {
	pdf.SetFooterFunc(func() {
		tr := pdf.UnicodeTranslatorFromDescriptor("")

		pdf.SetY(-55)
		pdf.SetX(10)
		pdf.SetFont("Arial", "", 8)
		pdf.MultiCell(185, 5, tr(fmt.Sprintf("%s au capital de %.2f euros\n SIRET %s - RCS %s - NAF %s\n TVA intracommunautaire : %s", b.Account.Company.Type, +b.Account.Company.Capital, b.Account.Company.Siret, b.Account.Company.RCS, b.Account.Company.NAF, b.Account.Company.Tva)), "", "C", false)
		pdf.Ln(6)
		pdf.SetX(10)
		pdf.MultiCell(185, 5, tr("La facture est payable sous 30 jours.\n Tout règlement effectué après expiration du délai donnera lieu, à titre de pénalité de retard, à\n l'application d'un intérêt égal à celui appliqué par la Banque Centrale Européenne à son opération de\n refinancement la plus récente, majoré de 10 points de pourcentage, ainsi qu'à une indemnité forfaitaire\n pour frais de recouvrement d'un montant de 40 Euros.\n Les pénalités de retard sont exigibles sans qu'un rappel soit nécessaire."), "", "C", false)

		pdf.SetDrawColor(182, 182, 182)
		pdf.SetY(-10)
		pdf.SetX(-10)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 8, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
}

func cleanTotalHT(b *BillingPDF) string {
	return fmt.Sprintf("%s €", b.FormatInt.IntToStringFrenchFormat(int(b.Service.GetTotalHT())))
}

func getTVA(b *BillingPDF) float32 {
	return b.Service.TVA.GetTVA(b.Service.GetTotalHT())
}

func cleanTVA(b *BillingPDF) string {
	return fmt.Sprintf("%s €", b.FormatInt.IntToStringFrenchFormat(int(getTVA(b))))
}

func getTotalTTC(b *BillingPDF) float32 {
	return b.Service.GetTotalTTC()
}

func cleanTotalTTC(b *BillingPDF) string {
	return fmt.Sprintf("%s €", b.FormatInt.IntToStringFrenchFormat(int(getTotalTTC(b))))
}

/*
CreatePdf func
*/
func (b *BillingPDF) CreatePDF() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTopMargin(30)
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	w, _ := pdf.GetPageSize()

	handleHeader(pdf, b)
	handleFooter(pdf, b)

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 11)
	pdf.SetTextColor(182, 182, 182)
	pdf.Cell(110, 10, "AU NOM ET POUR LE COMPTE")
	x, y := pdf.GetXY()
	pdf.SetDrawColor(182, 182, 182)
	pdf.Rect(x, y+10, 84, 0.1, "DF")
	pdf.Cell(50, 12, tr("ADRESSÉ À"))
	pdf.Ln(5)
	x, y = pdf.GetXY()
	pdf.SetDrawColor(182, 182, 182)
	pdf.Rect(x+1, y+10, 62, 0.1, "DF")
	pdf.Cell(30, 12, "DE")
	pdf.Ln(15)

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 11)

	pdf.Cell(110, 10, b.Account.Name)
	pdf.Cell(30, 0, b.Customer.Name)
	pdf.Ln(6)
	pdf.Cell(110, 10, b.Account.FirstName+" "+b.Account.LastName)
	pdf.Cell(30, 0, b.Customer.Address.Street)
	pdf.Ln(6)
	pdf.Cell(110, 10, b.Account.Address.Street)
	pdf.Cell(30, 0, b.Customer.Address.ZipCode+" "+b.Customer.Address.City+", "+b.Customer.Address.Country)
	pdf.Ln(6)
	pdf.Cell(110, 10, b.Account.Address.ZipCode+" "+b.Account.Address.City+", "+b.Account.Address.Country)
	pdf.Cell(30, 0, tr("N° SIRET: "+b.Account.Company.Siret))
	pdf.Ln(6)
	pdf.Cell(110, 10, tr("N° SIRET: "+b.Customer.Company.Siret))
	pdf.Cell(20, 0, tr("N° TVA Intracommunautaire : "+b.Customer.Company.Tva))
	pdf.Ln(6)
	pdf.Cell(110, 10, tr("N° TVA Intracommunautaire : "+b.Account.Company.Tva))
	pdf.Ln(6)
	pdf.Cell(120, 10, tr("@ : "+b.Account.Mail))
	pdf.Ln(15)

	pdf.SetTextColor(182, 182, 182)
	pdf.SetFont("Arial", "", 11)
	x, y = pdf.GetXY()

	pdf.Cell(50, 8, "PRESTATION")
	pdf.SetDrawColor(182, 182, 182)
	pdf.Rect(x+1, y+10, w-17, 0.1, "DF")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(60, 60, 60)
	pdf.Cell(95, 8, tr("Détail"))
	pdf.Cell(30, 8, tr("Quantité"))
	pdf.Cell(45, 8, tr("Prix unit. (HT)"))
	pdf.Cell(30, 8, "TOTAL (HT)")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 11)
	pdf.Cell(100, 8, tr(b.Service.Detail))
	pdf.Cell(44, 8, fmt.Sprintf("%g", b.Service.Quantity))
	pdf.Cell(32, 8, fmt.Sprintf("%g", b.Service.UnitPrice))
	pdf.Cell(30, 8, tr(cleanTotalHT(b)))
	pdf.Ln(15)

	x, y = pdf.GetXY()
	pdf.SetDrawColor(182, 182, 182)
	pdf.Rect(x+1, y, w-17, 0.1, "DF")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 11)
	pdf.SetLeftMargin(w - 60)

	x, y = pdf.GetXY()
	pdf.Cell(30, 8, "TOTAL (HT)")
	pdf.SetFont("Arial", "B", 13)
	pdf.Cell(12, 8, tr(cleanTotalHT(b)))
	pdf.SetDrawColor(182, 182, 182)
	pdf.Rect(x+1, y+10, 48, 0.1, "DF")
	pdf.Ln(13)

	pdf.SetFont("Arial", "", 11)
	x, y = pdf.GetXY()
	pdf.Cell(32, 8, "TVA ("+b.FormatFloat.Float32ToString(b.Service.TVA.Pourcent)+"%)")
	pdf.SetFont("Arial", "B", 13)
	pdf.Cell(12, 8, tr(cleanTVA(b)))
	pdf.SetDrawColor(182, 182, 182)
	pdf.Rect(x+1, y+10, 48, 0.1, "DF")

	pdf.Ln(13)

	pdf.SetFont("Arial", "", 13)
	pdf.Cell(28, 8, "TOTAL TTC")
	pdf.SetFont("Arial", "B", 15)
	pdf.Cell(12, 8, tr(cleanTotalTTC(b)))

	pdf.OutputFileAndClose(b.PdfPath + "/" + b.PdfName + ".pdf")
}

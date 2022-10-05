package pdfgenerator

type PdfGeneratorInterface interface {
	Create(htmlFile string) (string, error)
}

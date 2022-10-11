package pdfgenerator

import (
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wk struct {
	rootPath string
}

func NewWkHtmlToPdf(rootPath string) PdfGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return "", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlFile)))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}

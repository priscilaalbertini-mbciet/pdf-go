package pdfgenerator

import (
	"bytes"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type wk struct {
	rootPath string
}

func NewWkHtmlToPdf(rootPath string) PdfGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (*bytes.Buffer, error) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlFile)))

	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return pdfg.Buffer(), nil
}

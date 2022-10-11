package pdfgenerator

import "bytes"

type PdfGeneratorInterface interface {
	Create(htmlFile string) (*bytes.Buffer, error)
}

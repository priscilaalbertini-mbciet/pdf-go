package htmlparser

import (
	"bytes"
	"text/template"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (a *htmlStruct) Create(templateName string, data interface{}) (string, error) {

	templateGenerator, err := template.ParseFiles(templateName)

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	foo := new(bytes.Buffer)

	if err := templateGenerator.Execute(foo, data); err != nil {
		return "", err
	}

	return foo.String(), nil
}

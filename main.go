package main

import (
	"fmt"
	"os"
	"pdf-go/htmlparser"
	"pdf-go/pdfgenerator"

	"github.com/gofiber/fiber/v2"

	"github.com/shopspring/decimal"
)

type Data struct {
	ReportName string
	Company    Company
	Customers  []Customers
}

type Customers struct {
	Id       int64
	Name     string
	Cpf      string
	Quantity decimal.Decimal
	Price    decimal.Decimal
}

type Company struct {
	Name string
	Cnpj string
	Ie   string
}

func main() {
	app := fiber.New()

	h := htmlparser.New("tmp")
	wk := pdfgenerator.NewWkHtmlToPdf("tmp")

	dataHTML := Data{
		ReportName: "Vendas",
		Company: Company{
			Name: "Mercado Bitcoin",
			Cnpj: "45.830.915/0001-62",
			Ie:   "955.527.685.779",
		},
		Customers: []Customers{
			{Name: "Maria", Id: 1, Cpf: "123.456.789-10", Quantity: decimal.NewFromFloat(10), Price: decimal.NewFromFloat(2.5)},
			{Name: "João", Id: 2, Cpf: "546.234.645-23", Quantity: decimal.NewFromFloat(2.75), Price: decimal.NewFromFloat(1.30)},
			{Name: "Rafael", Id: 3, Cpf: "637.345.752-44", Quantity: decimal.NewFromFloat(5), Price: decimal.NewFromFloat(1.75)},
			{Name: "Letícia", Id: 4, Cpf: "321.321.321-32", Quantity: decimal.NewFromFloat(50), Price: decimal.NewFromFloat(10)},
		},
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer os.Remove(htmlGenerated)
	// fmt.Println("HTML gerado", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/pdf;base64")
		// c.Set("Content-type", "application/json")

		// t := Teste{
		// 	Order: "RRRRR",
		// 	File:  a.Bytes(),
		// }
		return c.Send(filePDFName.Bytes())
	})
	app.Listen(":3000")

}

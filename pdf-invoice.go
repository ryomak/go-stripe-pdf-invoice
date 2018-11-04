package go_stripe_pdf_invoice

import (
	"html/template"
	"fmt"
	"bytes"
	"path/filepath"
	"github.com/ryomak/go-stripe-pdf-invoice/templates"
	"github.com/jung-kurt/gofpdf"
)

func MakeDefaultPDF(stripeSecretKey ,invoiceID string,company Company,clientCompany ClientCompany)(interface{},error){
	fmt.Println("start")
	i,err := getInvoice(stripeSecretKey,invoiceID)
	if err != nil{
		return nil, err
	}
	htmlStr ,err := makeBody(TemplateData{*i,company,clientCompany},"templates/invoice.tmpl")
	if err != nil {
		return nil, err
	}
	//make html to pdf
	fmt.Println(htmlStr)
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 20)
	html := pdf.HTMLBasicNew()
	_, lineHt := pdf.GetFontSize()
	html.Write(lineHt, htmlStr)
	err = pdf.OutputFileAndClose("example.pdf")
	if err != nil{
		return nil,err
	}
	return pdf,nil
}
//go:generate go-bindata -pkg templates -o=templates/tmpl.go templates/
func makeBody(vars interface{}, files ...string) (string, error) {
	t, err := parseAssets(files...)
	if err != nil {
		return "", err
	}
	buff := &bytes.Buffer{}
	t.Execute(buff, vars)
	return buff.String(), nil
}

func parseAssets(filename ...string) (*template.Template, error) {
	var t *template.Template
	if len(filename) == 0 {
		return nil, fmt.Errorf("html/template: no files named ")
	}
	for _, f := range filename {
		name := filepath.Base(f)
		b, err := templates.Asset(f)
		if err != nil {
			return nil, err
		}
		s := string(b)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

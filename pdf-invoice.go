package go_stripe_pdf_invoice

import (
	"html/template"
	"bytes"
	"path/filepath"
	"github.com/ryomak/go-stripe-pdf-invoice/templates"
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func MakePDF(stripeSecretKey ,invoiceID string,company Company,clientCompany ClientCompany)(interface{},error){
	i,err := getInvoice(stripeSecretKey,invoiceID)
	if err != nil{
		return nil, err
	}
	htmlStr ,err := makeBody(TemplateData{*i,company,clientCompany},"templates/invoice.tmpl")
	if err != nil {
		return nil, err
	}
	//make html to pdf
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	src, err := templates.Asset("template/GenShinGothic-Regular.ttf")
	pdf.AddFontFromBytes("reglar","",src)
	pdf.SetFont("Arial", "B", 15)
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(14)
	_, lineHt := pdf.GetFontSize()
	html := pdf.HTMLBasicNew()
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

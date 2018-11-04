package go_stripe_pdf_invoice

type Company struct{
	Name string
	Address string
	ZipCode string
	City string
	Notice []string
	Logo interface{}
}

type ClientCompany struct {
	Name string
	Address string
	ZipCode string
	City string
}

type TemplateData struct {
	Invoice interface{}
	Company Company
	ClientCompany ClientCompany
}
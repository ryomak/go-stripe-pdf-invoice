package go_stripe_pdf_invoice

import(
	stripeClient "github.com/stripe/stripe-go/client"
	"github.com/stripe/stripe-go"
)

func newClient(secretKey string)*stripeClient.API{
	return stripeClient.New(secretKey, nil)
}

func getInvoice(secretKey ,invoiceID string)(*stripe.Invoice ,error){
	client := newClient(secretKey)
	i, err := client.Invoices.Get(invoiceID, nil)
	if err != nil {
		return nil,err
	}
	return i,nil
}
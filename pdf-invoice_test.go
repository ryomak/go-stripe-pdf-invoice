package go_stripe_pdf_invoice

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)
func TestMakeDefaultPDF(t *testing.T) {
	_, err := MakeDefaultPDF(
		"sk_test_eGedOsjx4ejuJAKDebnCLRve",
		"in_1DNyB2Iy5glBwHuiFXrlAr7A",
	Company{
		Name:"いい塾",
	},
	ClientCompany{
		Name:"木村",
		})
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

}
package strategy

import "log"

type Invoice struct {
	InvoiceID string
	ContractorName string
	Vendor Vendor
	GrossAmount uint32
}

type Vendor struct {
	VendorName string
}

type cilValidation struct {
	invoice *Invoice
}

func NewCILValidation(invoice *Invoice) Validation {
	return &cilValidation{
		invoice: invoice,
	}
}

func (cil *cilValidation) RunValidation() []error {
	log.Println("running cil validation")
	var errors []error

	err := NewContractorNameValidator(cil.invoice.ContractorName, cil.invoice.Vendor.VendorName).Validate()
	if err != nil {
		errors = append(errors, err)
	}

	err = NewGrossAmtValidator(cil.invoice.GrossAmount).Validate()
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}
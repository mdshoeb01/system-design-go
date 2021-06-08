package strategy

import "log"

type defaultValidation struct {
	invoice *Invoice
}

func NewDefaultValidation(inv *Invoice) Validation {
	return &defaultValidation{
		invoice: inv,
	}
}

func (defV *defaultValidation) RunValidation() []error {
	log.Println("running default validation")

	var  errors []error

	err := NewContractorNameValidator(defV.invoice.ContractorName, defV.invoice.Vendor.VendorName).Validate()
	if err != nil {
		errors = append(errors, err)
	}

	err = NewGrossAmtValidator(defV.invoice.GrossAmount).Validate()
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}
package observer

import (
	"context"
	"errors"
	"log"
	"strings"
)

type contractorValidator struct {
	ctx context.Context
	config string
	contractorName string
	vendorName string
	cilValidation cilValidation
}

func NewContractorValidator(cilValidation cilValidation) Validator {
	// should subscribe here.
	contractorValidator := &contractorValidator{
		// ctx: ctx,
		// config: config,
		// contractorName: contractorName,
		// vendorName: vendorName,
		cilValidation: cilValidation,
	}
	_ = cilValidation.validationManager.Subscribe(contractorValidator)
	return contractorValidator

}

func (v *contractorValidator) Validate() error {
	log.Println("running validate for contractor name")
	if !strings.Contains(v.contractorName, v.vendorName) {
		return errors.New("Contractor name not matching")
	}
	return nil
}


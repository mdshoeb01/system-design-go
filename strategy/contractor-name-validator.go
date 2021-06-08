package strategy

import (
	"errors"
	"log"
	"strings"
)

type contractorNameValidator struct {
	contractorName string
	vendorName string
}

func NewContractorNameValidator(cName string, vName string) Validator {
	return &contractorNameValidator{
		contractorName: cName,
		vendorName: vName,
	}
}

func (cNV *contractorNameValidator) Validate() error {
	log.Println("validate contractor name")
	if !strings.Contains(cNV.contractorName, cNV.vendorName) {
		return errors.New("contractor name does not contain vendor name")
	}
	return nil
}
package strategy

import "log"

type grossAmountValidator struct {
	amount uint32
}

func NewGrossAmtValidator(amt uint32) Validator {
	return &grossAmountValidator{
		amount: amt,
	}
} 

func (v *grossAmountValidator) Validate() error {
	log.Println("validating gross amount")
	return nil
}

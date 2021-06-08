package strategy

import (
	"errors"
	"log"
)

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
	if v.amount == 0 {
		return errors.New("gross amount cant be 0")
	}
	return nil
}

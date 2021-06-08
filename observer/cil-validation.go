package observer

// import "log"

type cilValidation struct {
	validationManager Validation
}

func NewCILValidation(vM Validation) *cilValidation {
	// validations := make([]Validator, 0)
	return &cilValidation{
		validationManager: vM,
	}
}

func (cil *cilValidation) RunValidation() {
	// contractorValidator := NewContractorValidator(&cil)
	// cil.validationManager.Subscribe()
	cil.validationManager.NotifySubscribers()
}

// func (v *cilValidation) Subscribe(validator Validator) error {
// 	v.validations = append(v.validations, validator)
// 	return nil
// }

// func (v *cilValidation) UnSubscribe(validator Validator) error {
// 	return nil
// }

// func (v *cilValidation) NotifySubscribers() error {
// 	for _, v := range v.validations {
// 		err := v.Validate()
// 		if err != nil {
// 			log.Println("err: ", err)
// 		}
// 	}
// 	return nil
// }

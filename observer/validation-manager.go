package observer

type ValidationManager struct {
	validators []Validator
}

func NewValidationManager() Validation {
	validators := make([]Validator, 0)
	return &ValidationManager{
		validators: validators,
	}
}

func (vM *ValidationManager) Subscribe(validator Validator) error {
	vM.validators = append(vM.validators, validator)
	return nil
}

func (vM *ValidationManager) UnSubscribe(validator Validator) error {
	return nil
}

func (vM *ValidationManager) NotifySubscribers() error {
	for _, v := range vM.validators {
		v.Validate()
	}
	return nil
}



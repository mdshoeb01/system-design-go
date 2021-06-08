package observer

import "log"

func StartObserverService() {
	log.Print("running observer service")
	validationManager := NewValidationManager()
	cilValidation := NewCILValidation(validationManager)
	_ = NewContractorValidator(*cilValidation)
	cilValidation.RunValidation()
}
package strategy

import "log"

type strategyService struct {
	invoice *Invoice
}

func NewStrategyService(invoice *Invoice) *strategyService{
	return &strategyService{
		invoice: invoice,
	}
}

func (s *strategyService) StartStrategyService() {
	log.Println("running strategy service")

	invoiceType := "GE"

	validationType := s.getValidation(invoiceType)
	errs := validationType.RunValidation()
	log.Println(errs)

	invoiceType = "CIL"

	validationType = s.getValidation(invoiceType)
	errs = validationType.RunValidation()
	log.Println(errs)

}

// factory can be used here to send the impl instance.
func (s *strategyService) getValidation(name string) Validation {
	if name == "CIL" {
		return NewCILValidation(s.invoice)
	} else {
		// default validation impl
		return NewDefaultValidation(s.invoice)
	}
}
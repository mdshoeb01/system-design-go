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

	invoiceType := "CIL"

	validationType := s.getValidation(invoiceType)
	errs := validationType.RunValidation()
	log.Println(errs)

}

func (s *strategyService) getValidation(name string) Validation {
	if name == "CIL" {
		return NewCILValidation(s.invoice)
	} else {
		// default validation impl
		return nil
	}
}
package strategy


type Validation interface {
	RunValidation() []error
}
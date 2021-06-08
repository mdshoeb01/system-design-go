package observer

type Validation interface {
	Subscribe(validator Validator) error
	UnSubscribe(validation Validator) error
	NotifySubscribers() error
}
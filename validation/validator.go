package validation

func NewDefaultValidator(validators map[string]ValidatorFunc) Validator {
	return &TagValidator{DefaultValidator()}
}

type Validator interface {
	Validate(data any) (ok bool, errs []ValidationError)
}

type ValidationError struct {
	FieldName string
	Error     error
}

type ValidatorFunc func(fileName string, value any, arg string) (bool, error)

func DefaultValidator() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required": required,
		"min":      min,
	}
}

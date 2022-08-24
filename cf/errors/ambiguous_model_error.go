package errors

import (
	"strings"

	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
)

type AmbiguousModelError struct {
	ModelType string
	ModelName string
}

func NewAmbiguousModelError(modelType, name string) error {
	return &AmbiguousModelError{
		ModelType: modelType,
		ModelName: name,
	}
}

func (err *AmbiguousModelError) Error() string {
	return T("Multiple ") + strings.ToLower(err.ModelType) + "s named " + err.ModelName + T(" found.")
}

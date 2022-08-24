package actionerror

import (
	"fmt"

	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

type ServiceInstanceTypeError struct {
	Name         string
	RequiredType resources.ServiceInstanceType
}

func (e ServiceInstanceTypeError) Error() string {
	return fmt.Sprintf("The service instance '%s' is not %s", e.Name, e.RequiredType)
}

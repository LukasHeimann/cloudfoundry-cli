package actionerror

import "github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccerror"

type ServiceOfferingNameAmbiguityError struct {
	ccerror.ServiceOfferingNameAmbiguityError
}

func (e ServiceOfferingNameAmbiguityError) Error() string {
	return e.ServiceOfferingNameAmbiguityError.Error() + "\nSpecify a broker by using the '-b' flag."
}

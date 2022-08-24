package formatters

import (
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
)

func Allowed(allowed bool) string {
	if allowed {
		return T("allowed")
	}
	return T("disallowed")
}

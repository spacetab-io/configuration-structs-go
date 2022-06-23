package errors

import (
	"sort"
	"strings"
)

type Errors []error

func (es Errors) Error() string {
	errs := make([]string, 0)

	for _, e := range es {
		errs = append(errs, e.Error())
	}

	sort.Strings(errs)

	return strings.Join(errs, ";")
}

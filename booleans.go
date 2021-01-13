package env

/* booleans.go adds booleans variables processing */

import (
	"strings"
)

const (
	strTrue = "true"
	strOne  = "1"
	strYes  = "yes"
	strOn   = "on"
)

func toBool(val string) bool {
	for _, strValue := range []string{strTrue, strYes, strOn, strOne} {
		if strings.EqualFold(val, strValue) {
			return true
		}
	}

	return false
}

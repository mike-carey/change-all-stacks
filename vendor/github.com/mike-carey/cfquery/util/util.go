package util

import (
	"reflect"
	"errors"
	"fmt"
)

func MapKeys(subject interface{}) []string {
	mapKeys := reflect.ValueOf(subject).MapKeys()
	keys := make([]string, len(mapKeys))

	for i, key := range mapKeys {
		keys[i] = key.String()
		i++
	}

	return keys
}

func StackErrors(errs []error) error {
	s := ""
	if len(errs) > 0 {
		s = "s"
	}

	err := fmt.Sprintf("%d error%s occured:\n", len(errs), s)
	for _, e := range errs {
		err += fmt.Sprintf("%q\n", e)
	}

	return errors.New(err)
}

func StringSliceContains(subject []string, item string) bool {
	for _, str := range subject {
		if str == item {
			return true
		}
	}

	return false
}

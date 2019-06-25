package query

import (
	"fmt"
	"strings"
)

func parseType(t interface{}) string {
	s := fmt.Sprintf("%T", t)
	u := strings.Split(s, ".")

	return u[len(u)-1]
}

type SearchCriteria struct {
	Key string
	Value string
}

type NotFoundError struct {
	subject interface{}
	searchCriteria SearchCriteria
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Could not find %s by %s: %s", parseType(e.subject), e.searchCriteria.Key, e.searchCriteria.Value)
}

func NewNotFoundError(subject interface{}, searchCriteriaKey string, searchCriteriaValue string) error {
	return &NotFoundError{
		subject: subject,
		searchCriteria: SearchCriteria{
			Key: searchCriteriaKey,
			Value: searchCriteriaValue,
		},
	}
}

var _ error = new(NotFoundError)

package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type AuthGroupType string

const (
	Eduate  AuthGroupType = "EDUATE"
	Staff   AuthGroupType = "STAFF"
	Student AuthGroupType = "STUDENT"
	Parent  AuthGroupType = "PARENT"
)

var AllAuthGroupType = []AuthGroupType{
	Eduate,
	Staff,
	Student,
	Parent,
}

func (e AuthGroupType) IsValid() bool {
	switch e {
	case Eduate, Staff, Student, Parent:
		return true
	}
	return false
}

func (e AuthGroupType) String() string {
	return string(e)
}

func (e *AuthGroupType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthGroupType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Auth group type", str)
	}
	return nil
}

func (e AuthGroupType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (AuthGroupType) Values() (kinds []string) {
	for _, s := range []AuthGroupType{Eduate, Staff, Student, Parent} {
		kinds = append(kinds, string(s))
	}
	return
}

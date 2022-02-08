package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type ParentType string

const (
	Father   ParentType = "FATHER"
	Mother   ParentType = "MOTHER"
	Guardian ParentType = "GUARDIAN"
)

var AllParentType = []ParentType{
	Father,
	Mother,
	Guardian,
}

func (e ParentType) IsValid() bool {
	switch e {
	case Father, Mother, Guardian:
		return true
	}
	return false
}

func (e ParentType) String() string {
	return string(e)
}

func (e *ParentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ParentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Auth group type", str)
	}
	return nil
}

func (e ParentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (ParentType) Values() (kinds []string) {
	for _, s := range []ParentType{Father, Mother, Guardian} {
		kinds = append(kinds, string(s))
	}
	return
}

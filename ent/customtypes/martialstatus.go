package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type MartialStatus string

const (
	Married   MartialStatus = "MARRIED"
	Unmarried MartialStatus = "UNMARRIED"
)

var AllMartialStatus = []MartialStatus{
	Married,
	Unmarried,
}

func (e MartialStatus) IsValid() bool {
	switch e {
	case Married, Unmarried:
		return true
	}
	return false
}

func (e MartialStatus) String() string {
	return string(e)
}

func (e *MartialStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MartialStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Married type", str)
	}
	return nil
}

func (e MartialStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (MartialStatus) Values() (kinds []string) {
	for _, s := range []MartialStatus{Married, Unmarried} {
		kinds = append(kinds, string(s))
	}
	return
}

package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type IsActive string

const (
	Active   IsActive = "ACTIVE"
	Inactive IsActive = "INACTIVE"
)

var AllActive = []IsActive{
	Active,
	Inactive,
}

func (e IsActive) IsValid() bool {
	switch e {
	case Active, Inactive:
		return true
	}
	return false
}

func (e IsActive) String() string {
	return string(e)
}

func (e *IsActive) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IsActive(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Active type", str)
	}
	return nil
}

func (e IsActive) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (IsActive) Values() (kinds []string) {
	for _, s := range []IsActive{Active, Inactive} {
		kinds = append(kinds, string(s))
	}
	return
}

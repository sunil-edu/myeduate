package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type Sex string

const (
	Male   Sex = "MALE"
	Female Sex = "FEMALE"
)

var AllSex = []Sex{
	Male,
	Female,
}

func (e Sex) IsValid() bool {
	switch e {
	case Male, Female:
		return true
	}
	return false
}

func (e Sex) String() string {
	return string(e)
}

func (e *Sex) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Sex(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid sex type", str)
	}
	return nil
}

func (e Sex) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (Sex) Values() (kinds []string) {
	for _, s := range []Sex{Male, Female} {
		kinds = append(kinds, string(s))
	}
	return
}

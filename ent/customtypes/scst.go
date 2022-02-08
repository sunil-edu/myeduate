package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type ScSt string

const (
	Sc ScSt = "SC"
	St ScSt = "ST"
)

var AllScSt = []ScSt{
	Sc,
	St,
}

func (e ScSt) IsValid() bool {
	switch e {
	case Sc, St:
		return true
	}
	return false
}

func (e ScSt) String() string {
	return string(e)
}

func (e *ScSt) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScSt(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ST or SC type", str)
	}
	return nil
}

func (e ScSt) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (ScSt) Values() (kinds []string) {
	for _, s := range []ScSt{Sc, St} {
		kinds = append(kinds, string(s))
	}
	return
}

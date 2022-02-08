package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type RuralUrban string

const (
	Rural RuralUrban = "RURAL"
	Urban RuralUrban = "URBAN"
)

var AllRuralUrban = []RuralUrban{
	Rural,
	Urban,
}

func (e RuralUrban) IsValid() bool {
	switch e {
	case Rural, Urban:
		return true
	}
	return false
}

func (e RuralUrban) String() string {
	return string(e)
}

func (e *RuralUrban) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RuralUrban(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Rural or ubran type", str)
	}
	return nil
}

func (e RuralUrban) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (RuralUrban) Values() (kinds []string) {
	for _, s := range []RuralUrban{Rural, Urban} {
		kinds = append(kinds, string(s))
	}
	return
}

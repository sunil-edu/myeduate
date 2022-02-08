package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type Region string

const (
	Kar    Region = "KAR"
	NonKar Region = "NON_KAR"
)

var AllRegion = []Region{
	Kar,
	NonKar,
}

func (e Region) IsValid() bool {
	switch e {
	case Kar, NonKar:
		return true
	}
	return false
}

func (e Region) String() string {
	return string(e)
}

func (e *Region) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Region(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid region type", str)
	}
	return nil
}

func (e Region) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (Region) Values() (kinds []string) {
	for _, s := range []Region{Kar, NonKar} {
		kinds = append(kinds, string(s))
	}
	return
}

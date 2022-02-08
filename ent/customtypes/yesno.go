package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type YesNo string

const (
	Yes YesNo = "YES"
	No  YesNo = "NO"
)

var AllYesNo = []YesNo{
	Yes,
	No,
}

func (e YesNo) IsValid() bool {
	switch e {
	case Yes, No:
		return true
	}
	return false
}

func (e YesNo) String() string {
	return string(e)
}

func (e *YesNo) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = YesNo(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Yes or No type", str)
	}
	return nil
}

func (e YesNo) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (YesNo) Values() (kinds []string) {
	for _, s := range []YesNo{Yes, No} {
		kinds = append(kinds, string(s))
	}
	return
}

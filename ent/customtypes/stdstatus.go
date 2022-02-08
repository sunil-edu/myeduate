package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type StdStatus string

const (
	Cur StdStatus = "CUR" //Current Student
	Tc  StdStatus = "TC"  //Transfer Taken
	Det StdStatus = "DET" //Detained
	Ne  StdStatus = "NE"  //Not eligible
)

var AllStdStatus = []StdStatus{
	Cur,
	Tc,
	Det,
	Ne,
}

func (e StdStatus) IsValid() bool {
	switch e {
	case Cur, Tc, Det, Ne:
		return true
	}
	return false
}

func (e StdStatus) String() string {
	return string(e)
}

func (e *StdStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StdStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid student status type", str)
	}
	return nil
}

func (e StdStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (StdStatus) Values() (kinds []string) {
	for _, s := range []StdStatus{Cur, Tc, Det, Ne} {
		kinds = append(kinds, string(s))
	}
	return
}

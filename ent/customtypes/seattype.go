package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type SeatType string

const (
	Free    SeatType = "FREE"
	Payment SeatType = "PAYMENT"
)

var AllSeatType = []SeatType{
	Free,
	Payment,
}

func (e SeatType) IsValid() bool {
	switch e {
	case Free, Payment:
		return true
	}
	return false
}

func (e SeatType) String() string {
	return string(e)
}

func (e *SeatType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SeatType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid seat type", str)
	}
	return nil
}

func (e SeatType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (SeatType) Values() (kinds []string) {
	for _, s := range []SeatType{Free, Payment} {
		kinds = append(kinds, string(s))
	}
	return
}

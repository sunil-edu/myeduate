package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type EntranceExam string

const (
	Cet    EntranceExam = "CET"
	Comdek EntranceExam = "COMDEK"
	Neet   EntranceExam = "NEET"
)

var AllEntranceExam = []EntranceExam{
	Cet,
	Comdek,
	Neet,
}

func (e EntranceExam) IsValid() bool {
	switch e {
	case Cet, Comdek, Neet:
		return true
	}
	return false
}

func (e EntranceExam) String() string {
	return string(e)
}

func (e *EntranceExam) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntranceExam(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not valid entrance exam type", str)
	}
	return nil
}

func (e EntranceExam) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (EntranceExam) Values() (kinds []string) {
	for _, s := range []EntranceExam{Cet, Comdek, Neet} {
		kinds = append(kinds, string(s))
	}
	return
}

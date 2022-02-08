package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type BloodGroup string

const (
	A_positive  BloodGroup = "A_POSITIVE"
	A_negative  BloodGroup = "A_NEGATIVE"
	B_positive  BloodGroup = "B_POSITIVE"
	B_negative  BloodGroup = "B_NEGATIVE"
	O_positive  BloodGroup = "O_POSITIVE"
	O_negative  BloodGroup = "O_NEGATIVE"
	AB_positive BloodGroup = "AB_POSITIVE"
	AB_negative BloodGroup = "AB_NEGATIVE"
)

var AllBloodGroup = []BloodGroup{
	A_positive,
	A_negative,
	B_positive,
	B_negative,
	O_positive,
	O_negative,
	AB_positive,
	AB_negative,
}

func (e BloodGroup) IsValid() bool {
	switch e {
	case A_positive, A_negative, B_positive, B_negative, O_positive, O_negative, AB_positive, AB_negative:
		return true
	}
	return false
}

func (e BloodGroup) String() string {
	return string(e)
}

func (e *BloodGroup) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BloodGroup(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Blood group", str)
	}
	return nil
}

func (e BloodGroup) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (BloodGroup) Values() (kinds []string) {
	for _, s := range []BloodGroup{A_positive, A_negative, B_positive, B_negative, O_positive, O_negative, AB_positive, AB_negative} {
		kinds = append(kinds, string(s))
	}
	return
}

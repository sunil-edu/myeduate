package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type EduateUserType string

const (
	Eduate_User      EduateUserType = "EDUATE_USER"
	Eduate_Manager   EduateUserType = "EDUATE_MANAGER"
	Eduate_Admin     EduateUserType = "EDUATE_ADMIN"
	Eduate_Sys_Admin EduateUserType = "EDUATE_SYS_ADMIN"
)

var AllEduateUserType = []EduateUserType{
	Eduate_User,
	Eduate_Manager,
	Eduate_Admin,
	Eduate_Sys_Admin,
}

func (e EduateUserType) IsValid() bool {
	switch e {
	case Eduate_User, Eduate_Manager, Eduate_Admin, Eduate_Sys_Admin:
		return true
	}
	return false
}

func (e EduateUserType) String() string {
	return string(e)
}

func (e *EduateUserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EduateUserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Eduate user type", str)
	}
	return nil
}

func (e EduateUserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (EduateUserType) Values() (kinds []string) {
	for _, s := range []EduateUserType{Eduate_User, Eduate_Manager, Eduate_Admin, Eduate_Sys_Admin} {
		kinds = append(kinds, string(s))
	}
	return
}

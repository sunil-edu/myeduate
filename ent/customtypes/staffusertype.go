package customtypes

import (
	"fmt"
	"io"
	"strconv"
)

type StaffUserType string

const (
	Inst_Sys_Admin  StaffUserType = "INST_SYS_ADMIN"
	Inst_Chairman   StaffUserType = "INST_CHAIRMAN"
	Inst_Principal  StaffUserType = "INST_PRINCIPAL"
	Inst_Staff      StaffUserType = "INST_STAFF"
	Inst_Member     StaffUserType = "INST_MEMBER"
	Inst_Head       StaffUserType = "INST_HEAD"
	Inst_HOD        StaffUserType = "INST_HOD"
	Inst_Group_Head StaffUserType = "INST_GROUP_HEAD"
	Inst_Faculty    StaffUserType = "INST_FACULTY"
)

var AllStaffUserType = []StaffUserType{
	Inst_Sys_Admin,
	Inst_Chairman,
	Inst_Principal,
	Inst_Staff,
	Inst_Member,
	Inst_Head,
	Inst_HOD,
	Inst_Group_Head,
	Inst_Faculty,
}

func (e StaffUserType) IsValid() bool {
	switch e {
	case Inst_Sys_Admin, Inst_Chairman, Inst_Principal, Inst_Staff, Inst_Member, Inst_Head, Inst_HOD, Inst_Group_Head, Inst_Faculty:
		return true
	}
	return false
}

func (e StaffUserType) String() string {
	return string(e)
}

func (e *StaffUserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StaffUserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid staff user type", str)
	}
	return nil
}

func (e StaffUserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Values provides list valid values for Enum.
func (StaffUserType) Values() (kinds []string) {
	for _, s := range []StaffUserType{Inst_Sys_Admin, Inst_Chairman, Inst_Principal, Inst_Staff, Inst_Member, Inst_Head, Inst_HOD, Inst_Group_Head, Inst_Faculty} {
		kinds = append(kinds, string(s))
	}
	return
}

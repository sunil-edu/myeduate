// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"myeduate/ent/customtypes"
	"myeduate/ent/mststudent"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// MstStudent is the model entity for the MstStudent schema.
type MstStudent struct {
	config `gqlgen:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// StdFirstName holds the value of the "std_first_name" field.
	StdFirstName string `json:"std_first_name,omitempty" gqlgen:"first_name"`
	// StdMiddleName holds the value of the "std_middle_name" field.
	StdMiddleName string `json:"std_middle_name,omitempty" gqlgen:"middle_name"`
	// StdLastName holds the value of the "std_last_name" field.
	StdLastName string `json:"std_last_name,omitempty" gqlgen:"last_name"`
	// StdStudying holds the value of the "std_studying" field.
	StdStudying bool `json:"std_studying,omitempty"`
	// StdStatus holds the value of the "std_status" field.
	StdStatus customtypes.StdStatus `json:"std_status,omitempty"`
	// StdSex holds the value of the "std_sex" field.
	StdSex customtypes.Sex `json:"std_sex,omitempty"`
	// StdRegNo holds the value of the "std_reg_no" field.
	StdRegNo string `json:"std_reg_no,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MstStudent) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mststudent.FieldStdStudying:
			values[i] = new(sql.NullBool)
		case mststudent.FieldID:
			values[i] = new(sql.NullInt64)
		case mststudent.FieldStdFirstName, mststudent.FieldStdMiddleName, mststudent.FieldStdLastName, mststudent.FieldStdStatus, mststudent.FieldStdSex, mststudent.FieldStdRegNo:
			values[i] = new(sql.NullString)
		case mststudent.FieldCreatedAt, mststudent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MstStudent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MstStudent fields.
func (ms *MstStudent) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mststudent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ms.ID = int(value.Int64)
		case mststudent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ms.CreatedAt = value.Time
			}
		case mststudent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ms.UpdatedAt = value.Time
			}
		case mststudent.FieldStdFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field std_first_name", values[i])
			} else if value.Valid {
				ms.StdFirstName = value.String
			}
		case mststudent.FieldStdMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field std_middle_name", values[i])
			} else if value.Valid {
				ms.StdMiddleName = value.String
			}
		case mststudent.FieldStdLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field std_last_name", values[i])
			} else if value.Valid {
				ms.StdLastName = value.String
			}
		case mststudent.FieldStdStudying:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field std_studying", values[i])
			} else if value.Valid {
				ms.StdStudying = value.Bool
			}
		case mststudent.FieldStdStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field std_status", values[i])
			} else if value.Valid {
				ms.StdStatus = customtypes.StdStatus(value.String)
			}
		case mststudent.FieldStdSex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field std_sex", values[i])
			} else if value.Valid {
				ms.StdSex = customtypes.Sex(value.String)
			}
		case mststudent.FieldStdRegNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field std_reg_no", values[i])
			} else if value.Valid {
				ms.StdRegNo = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MstStudent.
// Note that you need to call MstStudent.Unwrap() before calling this method if this MstStudent
// was returned from a transaction, and the transaction was committed or rolled back.
func (ms *MstStudent) Update() *MstStudentUpdateOne {
	return (&MstStudentClient{config: ms.config}).UpdateOne(ms)
}

// Unwrap unwraps the MstStudent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ms *MstStudent) Unwrap() *MstStudent {
	tx, ok := ms.config.driver.(*txDriver)
	if !ok {
		panic("ent: MstStudent is not a transactional entity")
	}
	ms.config.driver = tx.drv
	return ms
}

// String implements the fmt.Stringer.
func (ms *MstStudent) String() string {
	var builder strings.Builder
	builder.WriteString("MstStudent(")
	builder.WriteString(fmt.Sprintf("id=%v", ms.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(ms.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ms.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", std_first_name=")
	builder.WriteString(ms.StdFirstName)
	builder.WriteString(", std_middle_name=")
	builder.WriteString(ms.StdMiddleName)
	builder.WriteString(", std_last_name=")
	builder.WriteString(ms.StdLastName)
	builder.WriteString(", std_studying=")
	builder.WriteString(fmt.Sprintf("%v", ms.StdStudying))
	builder.WriteString(", std_status=")
	builder.WriteString(fmt.Sprintf("%v", ms.StdStatus))
	builder.WriteString(", std_sex=")
	builder.WriteString(fmt.Sprintf("%v", ms.StdSex))
	builder.WriteString(", std_reg_no=")
	builder.WriteString(ms.StdRegNo)
	builder.WriteByte(')')
	return builder.String()
}

// MstStudents is a parsable slice of MstStudent.
type MstStudents []*MstStudent

func (ms MstStudents) config(cfg config) {
	for _i := range ms {
		ms[_i].config = cfg
	}
}
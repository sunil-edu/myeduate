// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"myeduate/ent/customtypes"
	"myeduate/ent/mstcustomer"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// MstCustomer is the model entity for the MstCustomer schema.
type MstCustomer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CustCode holds the value of the "cust_code" field.
	CustCode string `json:"cust_code,omitempty"`
	// CustName holds the value of the "cust_name" field.
	CustName string `json:"cust_name,omitempty"`
	// CustAddress holds the value of the "cust_address" field.
	CustAddress string `json:"cust_address,omitempty"`
	// CustPlace holds the value of the "cust_place" field.
	CustPlace string `json:"cust_place,omitempty"`
	// CustState holds the value of the "cust_state" field.
	CustState string `json:"cust_state,omitempty"`
	// CustPin holds the value of the "cust_pin" field.
	CustPin string `json:"cust_pin,omitempty"`
	// CustContactPerson holds the value of the "cust_contact_person" field.
	CustContactPerson string `json:"cust_contact_person,omitempty"`
	// CustPhone holds the value of the "cust_phone" field.
	CustPhone string `json:"cust_phone,omitempty"`
	// CustEmail holds the value of the "cust_email" field.
	CustEmail string `json:"cust_email,omitempty"`
	// CustMobile holds the value of the "cust_mobile" field.
	CustMobile string `json:"cust_mobile,omitempty"`
	// CustURL holds the value of the "cust_url" field.
	CustURL string `json:"cust_url,omitempty"`
	// CustBanner1 holds the value of the "cust_banner_1" field.
	CustBanner1 string `json:"cust_banner_1,omitempty"`
	// CustBanner2 holds the value of the "cust_banner_2" field.
	CustBanner2 string `json:"cust_banner_2,omitempty"`
	// CustLogoURL holds the value of the "cust_logo_url" field.
	CustLogoURL string `json:"cust_logo_url,omitempty"`
	// CustIsActive holds the value of the "cust_is_active" field.
	CustIsActive customtypes.IsActive `json:"cust_is_active,omitempty"`
	// CustStatus holds the value of the "cust_status" field.
	CustStatus string `json:"cust_status,omitempty"`
	// CustNumInst holds the value of the "cust_num_inst" field.
	CustNumInst int `json:"cust_num_inst,omitempty"`
	// CustTimeZone holds the value of the "cust_time_zone" field.
	CustTimeZone time.Time `json:"cust_time_zone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MstCustomerQuery when eager-loading is set.
	Edges MstCustomerEdges `json:"edges"`
}

// MstCustomerEdges holds the relations/edges for other nodes in the graph.
type MstCustomerEdges struct {
	// Cust2Inst holds the value of the Cust2Inst edge.
	Cust2Inst []*MstInst `json:"Cust2Inst,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Cust2InstOrErr returns the Cust2Inst value or an error if the edge
// was not loaded in eager-loading.
func (e MstCustomerEdges) Cust2InstOrErr() ([]*MstInst, error) {
	if e.loadedTypes[0] {
		return e.Cust2Inst, nil
	}
	return nil, &NotLoadedError{edge: "Cust2Inst"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MstCustomer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mstcustomer.FieldID, mstcustomer.FieldCustNumInst:
			values[i] = new(sql.NullInt64)
		case mstcustomer.FieldCustCode, mstcustomer.FieldCustName, mstcustomer.FieldCustAddress, mstcustomer.FieldCustPlace, mstcustomer.FieldCustState, mstcustomer.FieldCustPin, mstcustomer.FieldCustContactPerson, mstcustomer.FieldCustPhone, mstcustomer.FieldCustEmail, mstcustomer.FieldCustMobile, mstcustomer.FieldCustURL, mstcustomer.FieldCustBanner1, mstcustomer.FieldCustBanner2, mstcustomer.FieldCustLogoURL, mstcustomer.FieldCustIsActive, mstcustomer.FieldCustStatus:
			values[i] = new(sql.NullString)
		case mstcustomer.FieldCreatedAt, mstcustomer.FieldUpdatedAt, mstcustomer.FieldCustTimeZone:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MstCustomer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MstCustomer fields.
func (mc *MstCustomer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mstcustomer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mc.ID = int(value.Int64)
		case mstcustomer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mc.CreatedAt = value.Time
			}
		case mstcustomer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mc.UpdatedAt = value.Time
			}
		case mstcustomer.FieldCustCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_code", values[i])
			} else if value.Valid {
				mc.CustCode = value.String
			}
		case mstcustomer.FieldCustName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_name", values[i])
			} else if value.Valid {
				mc.CustName = value.String
			}
		case mstcustomer.FieldCustAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_address", values[i])
			} else if value.Valid {
				mc.CustAddress = value.String
			}
		case mstcustomer.FieldCustPlace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_place", values[i])
			} else if value.Valid {
				mc.CustPlace = value.String
			}
		case mstcustomer.FieldCustState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_state", values[i])
			} else if value.Valid {
				mc.CustState = value.String
			}
		case mstcustomer.FieldCustPin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_pin", values[i])
			} else if value.Valid {
				mc.CustPin = value.String
			}
		case mstcustomer.FieldCustContactPerson:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_contact_person", values[i])
			} else if value.Valid {
				mc.CustContactPerson = value.String
			}
		case mstcustomer.FieldCustPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_phone", values[i])
			} else if value.Valid {
				mc.CustPhone = value.String
			}
		case mstcustomer.FieldCustEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_email", values[i])
			} else if value.Valid {
				mc.CustEmail = value.String
			}
		case mstcustomer.FieldCustMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_mobile", values[i])
			} else if value.Valid {
				mc.CustMobile = value.String
			}
		case mstcustomer.FieldCustURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_url", values[i])
			} else if value.Valid {
				mc.CustURL = value.String
			}
		case mstcustomer.FieldCustBanner1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_banner_1", values[i])
			} else if value.Valid {
				mc.CustBanner1 = value.String
			}
		case mstcustomer.FieldCustBanner2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_banner_2", values[i])
			} else if value.Valid {
				mc.CustBanner2 = value.String
			}
		case mstcustomer.FieldCustLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_logo_url", values[i])
			} else if value.Valid {
				mc.CustLogoURL = value.String
			}
		case mstcustomer.FieldCustIsActive:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_is_active", values[i])
			} else if value.Valid {
				mc.CustIsActive = customtypes.IsActive(value.String)
			}
		case mstcustomer.FieldCustStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cust_status", values[i])
			} else if value.Valid {
				mc.CustStatus = value.String
			}
		case mstcustomer.FieldCustNumInst:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cust_num_inst", values[i])
			} else if value.Valid {
				mc.CustNumInst = int(value.Int64)
			}
		case mstcustomer.FieldCustTimeZone:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field cust_time_zone", values[i])
			} else if value.Valid {
				mc.CustTimeZone = value.Time
			}
		}
	}
	return nil
}

// QueryCust2Inst queries the "Cust2Inst" edge of the MstCustomer entity.
func (mc *MstCustomer) QueryCust2Inst() *MstInstQuery {
	return (&MstCustomerClient{config: mc.config}).QueryCust2Inst(mc)
}

// Update returns a builder for updating this MstCustomer.
// Note that you need to call MstCustomer.Unwrap() before calling this method if this MstCustomer
// was returned from a transaction, and the transaction was committed or rolled back.
func (mc *MstCustomer) Update() *MstCustomerUpdateOne {
	return (&MstCustomerClient{config: mc.config}).UpdateOne(mc)
}

// Unwrap unwraps the MstCustomer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mc *MstCustomer) Unwrap() *MstCustomer {
	tx, ok := mc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MstCustomer is not a transactional entity")
	}
	mc.config.driver = tx.drv
	return mc
}

// String implements the fmt.Stringer.
func (mc *MstCustomer) String() string {
	var builder strings.Builder
	builder.WriteString("MstCustomer(")
	builder.WriteString(fmt.Sprintf("id=%v", mc.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(mc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(mc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", cust_code=")
	builder.WriteString(mc.CustCode)
	builder.WriteString(", cust_name=")
	builder.WriteString(mc.CustName)
	builder.WriteString(", cust_address=")
	builder.WriteString(mc.CustAddress)
	builder.WriteString(", cust_place=")
	builder.WriteString(mc.CustPlace)
	builder.WriteString(", cust_state=")
	builder.WriteString(mc.CustState)
	builder.WriteString(", cust_pin=")
	builder.WriteString(mc.CustPin)
	builder.WriteString(", cust_contact_person=")
	builder.WriteString(mc.CustContactPerson)
	builder.WriteString(", cust_phone=")
	builder.WriteString(mc.CustPhone)
	builder.WriteString(", cust_email=")
	builder.WriteString(mc.CustEmail)
	builder.WriteString(", cust_mobile=")
	builder.WriteString(mc.CustMobile)
	builder.WriteString(", cust_url=")
	builder.WriteString(mc.CustURL)
	builder.WriteString(", cust_banner_1=")
	builder.WriteString(mc.CustBanner1)
	builder.WriteString(", cust_banner_2=")
	builder.WriteString(mc.CustBanner2)
	builder.WriteString(", cust_logo_url=")
	builder.WriteString(mc.CustLogoURL)
	builder.WriteString(", cust_is_active=")
	builder.WriteString(fmt.Sprintf("%v", mc.CustIsActive))
	builder.WriteString(", cust_status=")
	builder.WriteString(mc.CustStatus)
	builder.WriteString(", cust_num_inst=")
	builder.WriteString(fmt.Sprintf("%v", mc.CustNumInst))
	builder.WriteString(", cust_time_zone=")
	builder.WriteString(mc.CustTimeZone.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MstCustomers is a parsable slice of MstCustomer.
type MstCustomers []*MstCustomer

func (mc MstCustomers) config(cfg config) {
	for _i := range mc {
		mc[_i].config = cfg
	}
}

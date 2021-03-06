// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package mstcustomer

import (
	"fmt"
	"myeduate/ent/customtypes"
	"time"

	"github.com/arsmn/fastgql/graphql"
)

const (
	// Label holds the string label denoting the mstcustomer type in the database.
	Label = "mst_customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCustCode holds the string denoting the cust_code field in the database.
	FieldCustCode = "cust_code"
	// FieldCustName holds the string denoting the cust_name field in the database.
	FieldCustName = "cust_name"
	// FieldCustAddress holds the string denoting the cust_address field in the database.
	FieldCustAddress = "cust_address"
	// FieldCustPlace holds the string denoting the cust_place field in the database.
	FieldCustPlace = "cust_place"
	// FieldCustState holds the string denoting the cust_state field in the database.
	FieldCustState = "cust_state"
	// FieldCustPin holds the string denoting the cust_pin field in the database.
	FieldCustPin = "cust_pin"
	// FieldCustContactPerson holds the string denoting the cust_contact_person field in the database.
	FieldCustContactPerson = "cust_contact_person"
	// FieldCustPhone holds the string denoting the cust_phone field in the database.
	FieldCustPhone = "cust_phone"
	// FieldCustEmail holds the string denoting the cust_email field in the database.
	FieldCustEmail = "cust_email"
	// FieldCustMobile holds the string denoting the cust_mobile field in the database.
	FieldCustMobile = "cust_mobile"
	// FieldCustURL holds the string denoting the cust_url field in the database.
	FieldCustURL = "cust_url"
	// FieldCustBanner1 holds the string denoting the cust_banner_1 field in the database.
	FieldCustBanner1 = "cust_banner_1"
	// FieldCustBanner2 holds the string denoting the cust_banner_2 field in the database.
	FieldCustBanner2 = "cust_banner_2"
	// FieldCustLogoURL holds the string denoting the cust_logo_url field in the database.
	FieldCustLogoURL = "cust_logo_url"
	// FieldCustIsActive holds the string denoting the cust_is_active field in the database.
	FieldCustIsActive = "cust_is_active"
	// FieldCustStatus holds the string denoting the cust_status field in the database.
	FieldCustStatus = "cust_status"
	// FieldCustNumInst holds the string denoting the cust_num_inst field in the database.
	FieldCustNumInst = "cust_num_inst"
	// FieldCustTimeZone holds the string denoting the cust_time_zone field in the database.
	FieldCustTimeZone = "cust_time_zone"
	// EdgeCust2Inst holds the string denoting the cust2inst edge name in mutations.
	EdgeCust2Inst = "Cust2Inst"
	// Table holds the table name of the mstcustomer in the database.
	Table = "mst_customers"
	// Cust2InstTable is the table that holds the Cust2Inst relation/edge.
	Cust2InstTable = "mst_insts"
	// Cust2InstInverseTable is the table name for the MstInst entity.
	// It exists in this package in order to avoid circular dependency with the "mstinst" package.
	Cust2InstInverseTable = "mst_insts"
	// Cust2InstColumn is the table column denoting the Cust2Inst relation/edge.
	Cust2InstColumn = "customer_id"
)

// Columns holds all SQL columns for mstcustomer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCustCode,
	FieldCustName,
	FieldCustAddress,
	FieldCustPlace,
	FieldCustState,
	FieldCustPin,
	FieldCustContactPerson,
	FieldCustPhone,
	FieldCustEmail,
	FieldCustMobile,
	FieldCustURL,
	FieldCustBanner1,
	FieldCustBanner2,
	FieldCustLogoURL,
	FieldCustIsActive,
	FieldCustStatus,
	FieldCustNumInst,
	FieldCustTimeZone,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// CustNameValidator is a validator for the "cust_name" field. It is called by the builders before save.
	CustNameValidator func(string) error
	// DefaultCustNumInst holds the default value on creation for the "cust_num_inst" field.
	DefaultCustNumInst int
)

const DefaultCustIsActive customtypes.IsActive = "ACTIVE"

// CustIsActiveValidator is a validator for the "cust_is_active" field enum values. It is called by the builders before save.
func CustIsActiveValidator(cia customtypes.IsActive) error {
	switch cia {
	case "ACTIVE", "INACTIVE":
		return nil
	default:
		return fmt.Errorf("mstcustomer: invalid enum value for cust_is_active field: %q", cia)
	}
}

var (
	// customtypes.IsActive must implement graphql.Marshaler.
	_ graphql.Marshaler = customtypes.IsActive("")
	// customtypes.IsActive must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*customtypes.IsActive)(nil)
)

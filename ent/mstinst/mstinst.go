// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package mstinst

import (
	"fmt"
	"myeduate/ent/customtypes"
	"myeduate/ent/schema/pulid"
	"time"

	"github.com/arsmn/fastgql/graphql"
)

const (
	// Label holds the string label denoting the mstinst type in the database.
	Label = "mst_inst"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldInstCode holds the string denoting the inst_code field in the database.
	FieldInstCode = "inst_code"
	// FieldInstName holds the string denoting the inst_name field in the database.
	FieldInstName = "inst_name"
	// FieldInstShortName holds the string denoting the inst_short_name field in the database.
	FieldInstShortName = "inst_short_name"
	// FieldInstAddress holds the string denoting the inst_address field in the database.
	FieldInstAddress = "inst_address"
	// FieldInstPlace holds the string denoting the inst_place field in the database.
	FieldInstPlace = "inst_place"
	// FieldInstState holds the string denoting the inst_state field in the database.
	FieldInstState = "inst_state"
	// FieldInstPin holds the string denoting the inst_pin field in the database.
	FieldInstPin = "inst_pin"
	// FieldInstContactPerson holds the string denoting the inst_contact_person field in the database.
	FieldInstContactPerson = "inst_contact_person"
	// FieldInstPhone holds the string denoting the inst_phone field in the database.
	FieldInstPhone = "inst_phone"
	// FieldInstEmail holds the string denoting the inst_email field in the database.
	FieldInstEmail = "inst_email"
	// FieldInstMobile holds the string denoting the inst_mobile field in the database.
	FieldInstMobile = "inst_mobile"
	// FieldInstURL holds the string denoting the inst_url field in the database.
	FieldInstURL = "inst_url"
	// FieldInstBanner1 holds the string denoting the inst_banner_1 field in the database.
	FieldInstBanner1 = "inst_banner_1"
	// FieldInstBanner2 holds the string denoting the inst_banner_2 field in the database.
	FieldInstBanner2 = "inst_banner_2"
	// FieldInstLogoURL holds the string denoting the inst_logo_url field in the database.
	FieldInstLogoURL = "inst_logo_url"
	// FieldInstIsActive holds the string denoting the inst_is_active field in the database.
	FieldInstIsActive = "inst_is_active"
	// FieldInstStatus holds the string denoting the inst_status field in the database.
	FieldInstStatus = "inst_status"
	// FieldInstTimeZone holds the string denoting the inst_time_zone field in the database.
	FieldInstTimeZone = "inst_time_zone"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// EdgeInstfromCust holds the string denoting the instfromcust edge name in mutations.
	EdgeInstfromCust = "InstfromCust"
	// Table holds the table name of the mstinst in the database.
	Table = "mst_insts"
	// InstfromCustTable is the table that holds the InstfromCust relation/edge.
	InstfromCustTable = "mst_insts"
	// InstfromCustInverseTable is the table name for the MstCustomer entity.
	// It exists in this package in order to avoid circular dependency with the "mstcustomer" package.
	InstfromCustInverseTable = "mst_customers"
	// InstfromCustColumn is the table column denoting the InstfromCust relation/edge.
	InstfromCustColumn = "customer_id"
)

// Columns holds all SQL columns for mstinst fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldInstCode,
	FieldInstName,
	FieldInstShortName,
	FieldInstAddress,
	FieldInstPlace,
	FieldInstState,
	FieldInstPin,
	FieldInstContactPerson,
	FieldInstPhone,
	FieldInstEmail,
	FieldInstMobile,
	FieldInstURL,
	FieldInstBanner1,
	FieldInstBanner2,
	FieldInstLogoURL,
	FieldInstIsActive,
	FieldInstStatus,
	FieldInstTimeZone,
	FieldCustomerID,
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
	// InstNameValidator is a validator for the "inst_name" field. It is called by the builders before save.
	InstNameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

const DefaultInstIsActive customtypes.IsActive = "ACTIVE"

// InstIsActiveValidator is a validator for the "inst_is_active" field enum values. It is called by the builders before save.
func InstIsActiveValidator(iia customtypes.IsActive) error {
	switch iia {
	case "ACTIVE", "INACTIVE":
		return nil
	default:
		return fmt.Errorf("mstinst: invalid enum value for inst_is_active field: %q", iia)
	}
}

var (
	// customtypes.IsActive must implement graphql.Marshaler.
	_ graphql.Marshaler = customtypes.IsActive("")
	// customtypes.IsActive must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*customtypes.IsActive)(nil)
)

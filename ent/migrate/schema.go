// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthParentsColumns holds the columns for the "auth_parents" table.
	AuthParentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "parent_first_name", Type: field.TypeString},
		{Name: "parent_middle_name", Type: field.TypeString},
		{Name: "parent_last_name", Type: field.TypeString},
	}
	// AuthParentsTable holds the schema information for the "auth_parents" table.
	AuthParentsTable = &schema.Table{
		Name:       "auth_parents",
		Columns:    AuthParentsColumns,
		PrimaryKey: []*schema.Column{AuthParentsColumns[0]},
	}
	// AuthStaffsColumns holds the columns for the "auth_staffs" table.
	AuthStaffsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "staff_first_name", Type: field.TypeString},
		{Name: "staff_middle_name", Type: field.TypeString},
		{Name: "staff_last_name", Type: field.TypeString},
	}
	// AuthStaffsTable holds the schema information for the "auth_staffs" table.
	AuthStaffsTable = &schema.Table{
		Name:       "auth_staffs",
		Columns:    AuthStaffsColumns,
		PrimaryKey: []*schema.Column{AuthStaffsColumns[0]},
	}
	// MsgChannelMessagesColumns holds the columns for the "msg_channel_messages" table.
	MsgChannelMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "msg_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "DATE"}},
		{Name: "msg_is_expiry", Type: field.TypeBool, Default: false},
		{Name: "msg_expiry_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "DATE"}},
		{Name: "msg_is_text", Type: field.TypeBool, Default: true},
		{Name: "msg_content", Type: field.TypeString},
		{Name: "msg_media_type", Type: field.TypeString},
		{Name: "msg_media_content", Type: field.TypeString},
		{Name: "msg_active", Type: field.TypeBool, Default: false},
		{Name: "msg_is_individual", Type: field.TypeBool},
		{Name: "msg_recv_or_sent", Type: field.TypeString},
	}
	// MsgChannelMessagesTable holds the schema information for the "msg_channel_messages" table.
	MsgChannelMessagesTable = &schema.Table{
		Name:       "msg_channel_messages",
		Columns:    MsgChannelMessagesColumns,
		PrimaryKey: []*schema.Column{MsgChannelMessagesColumns[0]},
	}
	// MstCustomersColumns holds the columns for the "mst_customers" table.
	MstCustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "cust_code", Type: field.TypeString},
		{Name: "cust_name", Type: field.TypeString, Unique: true},
		{Name: "cust_address", Type: field.TypeString},
		{Name: "cust_place", Type: field.TypeString},
		{Name: "cust_state", Type: field.TypeString},
		{Name: "cust_pin", Type: field.TypeString},
		{Name: "cust_contact_person", Type: field.TypeString},
		{Name: "cust_phone", Type: field.TypeString},
		{Name: "cust_email", Type: field.TypeString},
		{Name: "cust_mobile", Type: field.TypeString},
		{Name: "cust_url", Type: field.TypeString},
		{Name: "cust_banner_1", Type: field.TypeString},
		{Name: "cust_banner_2", Type: field.TypeString},
		{Name: "cust_logo_url", Type: field.TypeString},
		{Name: "cust_is_active", Type: field.TypeEnum, Enums: []string{"ACTIVE", "INACTIVE"}, Default: "ACTIVE"},
		{Name: "cust_status", Type: field.TypeString},
		{Name: "cust_num_inst", Type: field.TypeInt, Default: 0},
		{Name: "cust_time_zone", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "DATE"}},
	}
	// MstCustomersTable holds the schema information for the "mst_customers" table.
	MstCustomersTable = &schema.Table{
		Name:       "mst_customers",
		Columns:    MstCustomersColumns,
		PrimaryKey: []*schema.Column{MstCustomersColumns[0]},
	}
	// MstInstsColumns holds the columns for the "mst_insts" table.
	MstInstsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "inst_code", Type: field.TypeString},
		{Name: "inst_name", Type: field.TypeString},
		{Name: "inst_short_name", Type: field.TypeString},
		{Name: "inst_address", Type: field.TypeString},
		{Name: "inst_place", Type: field.TypeString},
		{Name: "inst_state", Type: field.TypeString},
		{Name: "inst_pin", Type: field.TypeString},
		{Name: "inst_contact_person", Type: field.TypeString},
		{Name: "inst_phone", Type: field.TypeString},
		{Name: "inst_email", Type: field.TypeString},
		{Name: "inst_mobile", Type: field.TypeString},
		{Name: "inst_url", Type: field.TypeString},
		{Name: "inst_banner_1", Type: field.TypeString},
		{Name: "inst_banner_2", Type: field.TypeString},
		{Name: "inst_logo_url", Type: field.TypeString},
		{Name: "inst_is_active", Type: field.TypeEnum, Enums: []string{"ACTIVE", "INACTIVE"}, Default: "ACTIVE"},
		{Name: "inst_status", Type: field.TypeString},
		{Name: "inst_time_zone", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "DATE"}},
		{Name: "customer_id", Type: field.TypeInt},
	}
	// MstInstsTable holds the schema information for the "mst_insts" table.
	MstInstsTable = &schema.Table{
		Name:       "mst_insts",
		Columns:    MstInstsColumns,
		PrimaryKey: []*schema.Column{MstInstsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mst_insts_mst_customers_Cust2Inst",
				Columns:    []*schema.Column{MstInstsColumns[21]},
				RefColumns: []*schema.Column{MstCustomersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MstStudentsColumns holds the columns for the "mst_students" table.
	MstStudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "std_first_name", Type: field.TypeString},
		{Name: "std_middle_name", Type: field.TypeString},
		{Name: "std_last_name", Type: field.TypeString},
	}
	// MstStudentsTable holds the schema information for the "mst_students" table.
	MstStudentsTable = &schema.Table{
		Name:       "mst_students",
		Columns:    MstStudentsColumns,
		PrimaryKey: []*schema.Column{MstStudentsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthParentsTable,
		AuthStaffsTable,
		MsgChannelMessagesTable,
		MstCustomersTable,
		MstInstsTable,
		MstStudentsTable,
	}
)

func init() {
	MstInstsTable.ForeignKeys[0].RefTable = MstCustomersTable
}

// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"myeduate/ent/authparent"
	"myeduate/ent/authstaff"
	"myeduate/ent/msgchannelmessage"
	"myeduate/ent/mstcustomer"
	"myeduate/ent/mstinst"
	"myeduate/ent/mststudent"
	"myeduate/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authparentMixin := schema.AuthParent{}.Mixin()
	authparentMixinFields0 := authparentMixin[0].Fields()
	_ = authparentMixinFields0
	authparentFields := schema.AuthParent{}.Fields()
	_ = authparentFields
	// authparentDescCreatedAt is the schema descriptor for created_at field.
	authparentDescCreatedAt := authparentMixinFields0[0].Descriptor()
	// authparent.DefaultCreatedAt holds the default value on creation for the created_at field.
	authparent.DefaultCreatedAt = authparentDescCreatedAt.Default.(func() time.Time)
	// authparentDescUpdatedAt is the schema descriptor for updated_at field.
	authparentDescUpdatedAt := authparentMixinFields0[1].Descriptor()
	// authparent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	authparent.DefaultUpdatedAt = authparentDescUpdatedAt.Default.(func() time.Time)
	// authparent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	authparent.UpdateDefaultUpdatedAt = authparentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// authparentDescParentFirstName is the schema descriptor for parent_first_name field.
	authparentDescParentFirstName := authparentFields[0].Descriptor()
	// authparent.ParentFirstNameValidator is a validator for the "parent_first_name" field. It is called by the builders before save.
	authparent.ParentFirstNameValidator = authparentDescParentFirstName.Validators[0].(func(string) error)
	authstaffMixin := schema.AuthStaff{}.Mixin()
	authstaffMixinFields0 := authstaffMixin[0].Fields()
	_ = authstaffMixinFields0
	authstaffFields := schema.AuthStaff{}.Fields()
	_ = authstaffFields
	// authstaffDescCreatedAt is the schema descriptor for created_at field.
	authstaffDescCreatedAt := authstaffMixinFields0[0].Descriptor()
	// authstaff.DefaultCreatedAt holds the default value on creation for the created_at field.
	authstaff.DefaultCreatedAt = authstaffDescCreatedAt.Default.(func() time.Time)
	// authstaffDescUpdatedAt is the schema descriptor for updated_at field.
	authstaffDescUpdatedAt := authstaffMixinFields0[1].Descriptor()
	// authstaff.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	authstaff.DefaultUpdatedAt = authstaffDescUpdatedAt.Default.(func() time.Time)
	// authstaff.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	authstaff.UpdateDefaultUpdatedAt = authstaffDescUpdatedAt.UpdateDefault.(func() time.Time)
	// authstaffDescStaffFirstName is the schema descriptor for staff_first_name field.
	authstaffDescStaffFirstName := authstaffFields[0].Descriptor()
	// authstaff.StaffFirstNameValidator is a validator for the "staff_first_name" field. It is called by the builders before save.
	authstaff.StaffFirstNameValidator = authstaffDescStaffFirstName.Validators[0].(func(string) error)
	msgchannelmessageMixin := schema.MsgChannelMessage{}.Mixin()
	msgchannelmessageMixinFields0 := msgchannelmessageMixin[0].Fields()
	_ = msgchannelmessageMixinFields0
	msgchannelmessageFields := schema.MsgChannelMessage{}.Fields()
	_ = msgchannelmessageFields
	// msgchannelmessageDescCreatedAt is the schema descriptor for created_at field.
	msgchannelmessageDescCreatedAt := msgchannelmessageMixinFields0[0].Descriptor()
	// msgchannelmessage.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgchannelmessage.DefaultCreatedAt = msgchannelmessageDescCreatedAt.Default.(func() time.Time)
	// msgchannelmessageDescUpdatedAt is the schema descriptor for updated_at field.
	msgchannelmessageDescUpdatedAt := msgchannelmessageMixinFields0[1].Descriptor()
	// msgchannelmessage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	msgchannelmessage.DefaultUpdatedAt = msgchannelmessageDescUpdatedAt.Default.(func() time.Time)
	// msgchannelmessage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	msgchannelmessage.UpdateDefaultUpdatedAt = msgchannelmessageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// msgchannelmessageDescMsgDate is the schema descriptor for msg_date field.
	msgchannelmessageDescMsgDate := msgchannelmessageFields[0].Descriptor()
	// msgchannelmessage.DefaultMsgDate holds the default value on creation for the msg_date field.
	msgchannelmessage.DefaultMsgDate = msgchannelmessageDescMsgDate.Default.(func() time.Time)
	// msgchannelmessageDescMsgIsExpiry is the schema descriptor for msg_is_expiry field.
	msgchannelmessageDescMsgIsExpiry := msgchannelmessageFields[1].Descriptor()
	// msgchannelmessage.DefaultMsgIsExpiry holds the default value on creation for the msg_is_expiry field.
	msgchannelmessage.DefaultMsgIsExpiry = msgchannelmessageDescMsgIsExpiry.Default.(bool)
	// msgchannelmessageDescMsgExpiryDate is the schema descriptor for msg_expiry_date field.
	msgchannelmessageDescMsgExpiryDate := msgchannelmessageFields[2].Descriptor()
	// msgchannelmessage.DefaultMsgExpiryDate holds the default value on creation for the msg_expiry_date field.
	msgchannelmessage.DefaultMsgExpiryDate = msgchannelmessageDescMsgExpiryDate.Default.(func() time.Time)
	// msgchannelmessageDescMsgIsText is the schema descriptor for msg_is_text field.
	msgchannelmessageDescMsgIsText := msgchannelmessageFields[3].Descriptor()
	// msgchannelmessage.DefaultMsgIsText holds the default value on creation for the msg_is_text field.
	msgchannelmessage.DefaultMsgIsText = msgchannelmessageDescMsgIsText.Default.(bool)
	// msgchannelmessageDescMsgActive is the schema descriptor for msg_active field.
	msgchannelmessageDescMsgActive := msgchannelmessageFields[7].Descriptor()
	// msgchannelmessage.DefaultMsgActive holds the default value on creation for the msg_active field.
	msgchannelmessage.DefaultMsgActive = msgchannelmessageDescMsgActive.Default.(bool)
	mstcustomerMixin := schema.MstCustomer{}.Mixin()
	mstcustomerMixinFields0 := mstcustomerMixin[0].Fields()
	_ = mstcustomerMixinFields0
	mstcustomerFields := schema.MstCustomer{}.Fields()
	_ = mstcustomerFields
	// mstcustomerDescCreatedAt is the schema descriptor for created_at field.
	mstcustomerDescCreatedAt := mstcustomerMixinFields0[0].Descriptor()
	// mstcustomer.DefaultCreatedAt holds the default value on creation for the created_at field.
	mstcustomer.DefaultCreatedAt = mstcustomerDescCreatedAt.Default.(func() time.Time)
	// mstcustomerDescUpdatedAt is the schema descriptor for updated_at field.
	mstcustomerDescUpdatedAt := mstcustomerMixinFields0[1].Descriptor()
	// mstcustomer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mstcustomer.DefaultUpdatedAt = mstcustomerDescUpdatedAt.Default.(func() time.Time)
	// mstcustomer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	mstcustomer.UpdateDefaultUpdatedAt = mstcustomerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mstcustomerDescCustName is the schema descriptor for cust_name field.
	mstcustomerDescCustName := mstcustomerFields[1].Descriptor()
	// mstcustomer.CustNameValidator is a validator for the "cust_name" field. It is called by the builders before save.
	mstcustomer.CustNameValidator = mstcustomerDescCustName.Validators[0].(func(string) error)
	// mstcustomerDescCustNumInst is the schema descriptor for cust_num_inst field.
	mstcustomerDescCustNumInst := mstcustomerFields[16].Descriptor()
	// mstcustomer.DefaultCustNumInst holds the default value on creation for the cust_num_inst field.
	mstcustomer.DefaultCustNumInst = mstcustomerDescCustNumInst.Default.(int)
	mstinstMixin := schema.MstInst{}.Mixin()
	mstinstMixinFields0 := mstinstMixin[0].Fields()
	_ = mstinstMixinFields0
	mstinstFields := schema.MstInst{}.Fields()
	_ = mstinstFields
	// mstinstDescCreatedAt is the schema descriptor for created_at field.
	mstinstDescCreatedAt := mstinstMixinFields0[0].Descriptor()
	// mstinst.DefaultCreatedAt holds the default value on creation for the created_at field.
	mstinst.DefaultCreatedAt = mstinstDescCreatedAt.Default.(func() time.Time)
	// mstinstDescUpdatedAt is the schema descriptor for updated_at field.
	mstinstDescUpdatedAt := mstinstMixinFields0[1].Descriptor()
	// mstinst.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mstinst.DefaultUpdatedAt = mstinstDescUpdatedAt.Default.(func() time.Time)
	// mstinst.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	mstinst.UpdateDefaultUpdatedAt = mstinstDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mstinstDescInstName is the schema descriptor for inst_name field.
	mstinstDescInstName := mstinstFields[1].Descriptor()
	// mstinst.InstNameValidator is a validator for the "inst_name" field. It is called by the builders before save.
	mstinst.InstNameValidator = mstinstDescInstName.Validators[0].(func(string) error)
	mststudentMixin := schema.MstStudent{}.Mixin()
	mststudentMixinFields0 := mststudentMixin[0].Fields()
	_ = mststudentMixinFields0
	mststudentFields := schema.MstStudent{}.Fields()
	_ = mststudentFields
	// mststudentDescCreatedAt is the schema descriptor for created_at field.
	mststudentDescCreatedAt := mststudentMixinFields0[0].Descriptor()
	// mststudent.DefaultCreatedAt holds the default value on creation for the created_at field.
	mststudent.DefaultCreatedAt = mststudentDescCreatedAt.Default.(func() time.Time)
	// mststudentDescUpdatedAt is the schema descriptor for updated_at field.
	mststudentDescUpdatedAt := mststudentMixinFields0[1].Descriptor()
	// mststudent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mststudent.DefaultUpdatedAt = mststudentDescUpdatedAt.Default.(func() time.Time)
	// mststudent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	mststudent.UpdateDefaultUpdatedAt = mststudentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mststudentDescStdFirstName is the schema descriptor for std_first_name field.
	mststudentDescStdFirstName := mststudentFields[0].Descriptor()
	// mststudent.StdFirstNameValidator is a validator for the "std_first_name" field. It is called by the builders before save.
	mststudent.StdFirstNameValidator = mststudentDescStdFirstName.Validators[0].(func(string) error)
}

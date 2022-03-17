// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"myeduate/ent/authparent"
	"myeduate/ent/authstaff"
	"myeduate/ent/msgchannelmessage"
	"myeduate/ent/mstcustomer"
	"myeduate/ent/mstinst"
	"myeduate/ent/mststudent"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (ap *AuthParent) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ap.ID,
		Type:   "AuthParent",
		Fields: make([]*Field, 11),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(ap.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentFirstName); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "parent_first_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentMiddleName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "parent_middle_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentLastName); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "parent_last_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentAddress); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "parent_address",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentPlace); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "parent_place",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentState); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "parent_state",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentPin); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "parent_pin",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentEmail); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "parent_email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ap.ParentMobile); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "parent_mobile",
		Value: string(buf),
	}
	return node, nil
}

func (as *AuthStaff) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     as.ID,
		Type:   "AuthStaff",
		Fields: make([]*Field, 11),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(as.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffFirstName); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "staff_first_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffMiddleName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "staff_middle_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffLastName); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "staff_last_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffAddress); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "staff_address",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffPlace); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "staff_place",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffState); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "staff_state",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffPin); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "staff_pin",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffEmail); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "staff_email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.StaffMobile); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "staff_mobile",
		Value: string(buf),
	}
	return node, nil
}

func (mcm *MsgChannelMessage) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     mcm.ID,
		Type:   "MsgChannelMessage",
		Fields: make([]*Field, 12),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(mcm.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgDate); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "msg_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgIsExpiry); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "bool",
		Name:  "msg_is_expiry",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgExpiryDate); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "msg_expiry_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgIsText); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "bool",
		Name:  "msg_is_text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgContent); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "msg_content",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgMediaType); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "msg_media_type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgMediaContent); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "msg_media_content",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgActive); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "bool",
		Name:  "msg_active",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgIsIndividual); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "bool",
		Name:  "msg_is_individual",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mcm.MsgRecvOrSent); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "msg_recv_or_sent",
		Value: string(buf),
	}
	return node, nil
}

func (mc *MstCustomer) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     mc.ID,
		Type:   "MstCustomer",
		Fields: make([]*Field, 20),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(mc.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustCode); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "cust_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "cust_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustAddress); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "cust_address",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustPlace); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "cust_place",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustState); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "cust_state",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustPin); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "cust_pin",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustContactPerson); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "cust_contact_person",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustPhone); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "cust_phone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustEmail); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "cust_email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustMobile); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "cust_mobile",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustURL); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "string",
		Name:  "cust_url",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustBanner1); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "string",
		Name:  "cust_banner_1",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustBanner2); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "string",
		Name:  "cust_banner_2",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustLogoURL); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "string",
		Name:  "cust_logo_url",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustIsActive); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "customtypes.IsActive",
		Name:  "cust_is_active",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustStatus); err != nil {
		return nil, err
	}
	node.Fields[17] = &Field{
		Type:  "string",
		Name:  "cust_status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustNumInst); err != nil {
		return nil, err
	}
	node.Fields[18] = &Field{
		Type:  "int",
		Name:  "cust_num_inst",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mc.CustTimeZone); err != nil {
		return nil, err
	}
	node.Fields[19] = &Field{
		Type:  "time.Time",
		Name:  "cust_time_zone",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "MstInst",
		Name: "Cust2Inst",
	}
	err = mc.QueryCust2Inst().
		Select(mstinst.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (mi *MstInst) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     mi.ID,
		Type:   "MstInst",
		Fields: make([]*Field, 21),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(mi.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstCode); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "inst_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "inst_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstShortName); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "inst_short_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstAddress); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "inst_address",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstPlace); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "inst_place",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstState); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "inst_state",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstPin); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "inst_pin",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstContactPerson); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "inst_contact_person",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstPhone); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "inst_phone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstEmail); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "inst_email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstMobile); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "string",
		Name:  "inst_mobile",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstURL); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "string",
		Name:  "inst_url",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstBanner1); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "string",
		Name:  "inst_banner_1",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstBanner2); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "string",
		Name:  "inst_banner_2",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstLogoURL); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "string",
		Name:  "inst_logo_url",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstIsActive); err != nil {
		return nil, err
	}
	node.Fields[17] = &Field{
		Type:  "customtypes.IsActive",
		Name:  "inst_is_active",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstStatus); err != nil {
		return nil, err
	}
	node.Fields[18] = &Field{
		Type:  "string",
		Name:  "inst_status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.InstTimeZone); err != nil {
		return nil, err
	}
	node.Fields[19] = &Field{
		Type:  "time.Time",
		Name:  "inst_time_zone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(mi.CustomerID); err != nil {
		return nil, err
	}
	node.Fields[20] = &Field{
		Type:  "int",
		Name:  "customer_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "MstCustomer",
		Name: "InstfromCust",
	}
	err = mi.QueryInstfromCust().
		Select(mstcustomer.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ms *MstStudent) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ms.ID,
		Type:   "MstStudent",
		Fields: make([]*Field, 9),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(ms.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdFirstName); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "std_first_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdMiddleName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "std_middle_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdLastName); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "std_last_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdStudying); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "bool",
		Name:  "std_studying",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdStatus); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "customtypes.StdStatus",
		Name:  "std_status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdSex); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "customtypes.Sex",
		Name:  "std_sex",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ms.StdRegNo); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "std_reg_no",
		Value: string(buf),
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case authparent.Table:
		n, err := c.AuthParent.Query().
			Where(authparent.ID(id)).
			CollectFields(ctx, "AuthParent").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case authstaff.Table:
		n, err := c.AuthStaff.Query().
			Where(authstaff.ID(id)).
			CollectFields(ctx, "AuthStaff").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case msgchannelmessage.Table:
		n, err := c.MsgChannelMessage.Query().
			Where(msgchannelmessage.ID(id)).
			CollectFields(ctx, "MsgChannelMessage").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case mstcustomer.Table:
		n, err := c.MstCustomer.Query().
			Where(mstcustomer.ID(id)).
			CollectFields(ctx, "MstCustomer").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case mstinst.Table:
		n, err := c.MstInst.Query().
			Where(mstinst.ID(id)).
			CollectFields(ctx, "MstInst").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case mststudent.Table:
		n, err := c.MstStudent.Query().
			Where(mststudent.ID(id)).
			CollectFields(ctx, "MstStudent").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case authparent.Table:
		nodes, err := c.AuthParent.Query().
			Where(authparent.IDIn(ids...)).
			CollectFields(ctx, "AuthParent").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case authstaff.Table:
		nodes, err := c.AuthStaff.Query().
			Where(authstaff.IDIn(ids...)).
			CollectFields(ctx, "AuthStaff").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msgchannelmessage.Table:
		nodes, err := c.MsgChannelMessage.Query().
			Where(msgchannelmessage.IDIn(ids...)).
			CollectFields(ctx, "MsgChannelMessage").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case mstcustomer.Table:
		nodes, err := c.MstCustomer.Query().
			Where(mstcustomer.IDIn(ids...)).
			CollectFields(ctx, "MstCustomer").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case mstinst.Table:
		nodes, err := c.MstInst.Query().
			Where(mstinst.IDIn(ids...)).
			CollectFields(ctx, "MstInst").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case mststudent.Table:
		nodes, err := c.MstStudent.Query().
			Where(mststudent.IDIn(ids...)).
			CollectFields(ctx, "MstStudent").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}

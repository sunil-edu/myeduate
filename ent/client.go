// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"myeduate/ent/migrate"

	"myeduate/ent/authparent"
	"myeduate/ent/authstaff"
	"myeduate/ent/mstcustomer"
	"myeduate/ent/mstinst"
	"myeduate/ent/mststudent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AuthParent is the client for interacting with the AuthParent builders.
	AuthParent *AuthParentClient
	// AuthStaff is the client for interacting with the AuthStaff builders.
	AuthStaff *AuthStaffClient
	// MstCustomer is the client for interacting with the MstCustomer builders.
	MstCustomer *MstCustomerClient
	// MstInst is the client for interacting with the MstInst builders.
	MstInst *MstInstClient
	// MstStudent is the client for interacting with the MstStudent builders.
	MstStudent *MstStudentClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AuthParent = NewAuthParentClient(c.config)
	c.AuthStaff = NewAuthStaffClient(c.config)
	c.MstCustomer = NewMstCustomerClient(c.config)
	c.MstInst = NewMstInstClient(c.config)
	c.MstStudent = NewMstStudentClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		AuthParent:  NewAuthParentClient(cfg),
		AuthStaff:   NewAuthStaffClient(cfg),
		MstCustomer: NewMstCustomerClient(cfg),
		MstInst:     NewMstInstClient(cfg),
		MstStudent:  NewMstStudentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		AuthParent:  NewAuthParentClient(cfg),
		AuthStaff:   NewAuthStaffClient(cfg),
		MstCustomer: NewMstCustomerClient(cfg),
		MstInst:     NewMstInstClient(cfg),
		MstStudent:  NewMstStudentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AuthParent.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AuthParent.Use(hooks...)
	c.AuthStaff.Use(hooks...)
	c.MstCustomer.Use(hooks...)
	c.MstInst.Use(hooks...)
	c.MstStudent.Use(hooks...)
}

// AuthParentClient is a client for the AuthParent schema.
type AuthParentClient struct {
	config
}

// NewAuthParentClient returns a client for the AuthParent from the given config.
func NewAuthParentClient(c config) *AuthParentClient {
	return &AuthParentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authparent.Hooks(f(g(h())))`.
func (c *AuthParentClient) Use(hooks ...Hook) {
	c.hooks.AuthParent = append(c.hooks.AuthParent, hooks...)
}

// Create returns a create builder for AuthParent.
func (c *AuthParentClient) Create() *AuthParentCreate {
	mutation := newAuthParentMutation(c.config, OpCreate)
	return &AuthParentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthParent entities.
func (c *AuthParentClient) CreateBulk(builders ...*AuthParentCreate) *AuthParentCreateBulk {
	return &AuthParentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthParent.
func (c *AuthParentClient) Update() *AuthParentUpdate {
	mutation := newAuthParentMutation(c.config, OpUpdate)
	return &AuthParentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthParentClient) UpdateOne(ap *AuthParent) *AuthParentUpdateOne {
	mutation := newAuthParentMutation(c.config, OpUpdateOne, withAuthParent(ap))
	return &AuthParentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthParentClient) UpdateOneID(id int) *AuthParentUpdateOne {
	mutation := newAuthParentMutation(c.config, OpUpdateOne, withAuthParentID(id))
	return &AuthParentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthParent.
func (c *AuthParentClient) Delete() *AuthParentDelete {
	mutation := newAuthParentMutation(c.config, OpDelete)
	return &AuthParentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuthParentClient) DeleteOne(ap *AuthParent) *AuthParentDeleteOne {
	return c.DeleteOneID(ap.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuthParentClient) DeleteOneID(id int) *AuthParentDeleteOne {
	builder := c.Delete().Where(authparent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthParentDeleteOne{builder}
}

// Query returns a query builder for AuthParent.
func (c *AuthParentClient) Query() *AuthParentQuery {
	return &AuthParentQuery{
		config: c.config,
	}
}

// Get returns a AuthParent entity by its id.
func (c *AuthParentClient) Get(ctx context.Context, id int) (*AuthParent, error) {
	return c.Query().Where(authparent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthParentClient) GetX(ctx context.Context, id int) *AuthParent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthParentClient) Hooks() []Hook {
	return c.hooks.AuthParent
}

// AuthStaffClient is a client for the AuthStaff schema.
type AuthStaffClient struct {
	config
}

// NewAuthStaffClient returns a client for the AuthStaff from the given config.
func NewAuthStaffClient(c config) *AuthStaffClient {
	return &AuthStaffClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authstaff.Hooks(f(g(h())))`.
func (c *AuthStaffClient) Use(hooks ...Hook) {
	c.hooks.AuthStaff = append(c.hooks.AuthStaff, hooks...)
}

// Create returns a create builder for AuthStaff.
func (c *AuthStaffClient) Create() *AuthStaffCreate {
	mutation := newAuthStaffMutation(c.config, OpCreate)
	return &AuthStaffCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthStaff entities.
func (c *AuthStaffClient) CreateBulk(builders ...*AuthStaffCreate) *AuthStaffCreateBulk {
	return &AuthStaffCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthStaff.
func (c *AuthStaffClient) Update() *AuthStaffUpdate {
	mutation := newAuthStaffMutation(c.config, OpUpdate)
	return &AuthStaffUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthStaffClient) UpdateOne(as *AuthStaff) *AuthStaffUpdateOne {
	mutation := newAuthStaffMutation(c.config, OpUpdateOne, withAuthStaff(as))
	return &AuthStaffUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthStaffClient) UpdateOneID(id int) *AuthStaffUpdateOne {
	mutation := newAuthStaffMutation(c.config, OpUpdateOne, withAuthStaffID(id))
	return &AuthStaffUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthStaff.
func (c *AuthStaffClient) Delete() *AuthStaffDelete {
	mutation := newAuthStaffMutation(c.config, OpDelete)
	return &AuthStaffDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuthStaffClient) DeleteOne(as *AuthStaff) *AuthStaffDeleteOne {
	return c.DeleteOneID(as.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuthStaffClient) DeleteOneID(id int) *AuthStaffDeleteOne {
	builder := c.Delete().Where(authstaff.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthStaffDeleteOne{builder}
}

// Query returns a query builder for AuthStaff.
func (c *AuthStaffClient) Query() *AuthStaffQuery {
	return &AuthStaffQuery{
		config: c.config,
	}
}

// Get returns a AuthStaff entity by its id.
func (c *AuthStaffClient) Get(ctx context.Context, id int) (*AuthStaff, error) {
	return c.Query().Where(authstaff.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthStaffClient) GetX(ctx context.Context, id int) *AuthStaff {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthStaffClient) Hooks() []Hook {
	return c.hooks.AuthStaff
}

// MstCustomerClient is a client for the MstCustomer schema.
type MstCustomerClient struct {
	config
}

// NewMstCustomerClient returns a client for the MstCustomer from the given config.
func NewMstCustomerClient(c config) *MstCustomerClient {
	return &MstCustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mstcustomer.Hooks(f(g(h())))`.
func (c *MstCustomerClient) Use(hooks ...Hook) {
	c.hooks.MstCustomer = append(c.hooks.MstCustomer, hooks...)
}

// Create returns a create builder for MstCustomer.
func (c *MstCustomerClient) Create() *MstCustomerCreate {
	mutation := newMstCustomerMutation(c.config, OpCreate)
	return &MstCustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MstCustomer entities.
func (c *MstCustomerClient) CreateBulk(builders ...*MstCustomerCreate) *MstCustomerCreateBulk {
	return &MstCustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MstCustomer.
func (c *MstCustomerClient) Update() *MstCustomerUpdate {
	mutation := newMstCustomerMutation(c.config, OpUpdate)
	return &MstCustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MstCustomerClient) UpdateOne(mc *MstCustomer) *MstCustomerUpdateOne {
	mutation := newMstCustomerMutation(c.config, OpUpdateOne, withMstCustomer(mc))
	return &MstCustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MstCustomerClient) UpdateOneID(id int) *MstCustomerUpdateOne {
	mutation := newMstCustomerMutation(c.config, OpUpdateOne, withMstCustomerID(id))
	return &MstCustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MstCustomer.
func (c *MstCustomerClient) Delete() *MstCustomerDelete {
	mutation := newMstCustomerMutation(c.config, OpDelete)
	return &MstCustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MstCustomerClient) DeleteOne(mc *MstCustomer) *MstCustomerDeleteOne {
	return c.DeleteOneID(mc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MstCustomerClient) DeleteOneID(id int) *MstCustomerDeleteOne {
	builder := c.Delete().Where(mstcustomer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MstCustomerDeleteOne{builder}
}

// Query returns a query builder for MstCustomer.
func (c *MstCustomerClient) Query() *MstCustomerQuery {
	return &MstCustomerQuery{
		config: c.config,
	}
}

// Get returns a MstCustomer entity by its id.
func (c *MstCustomerClient) Get(ctx context.Context, id int) (*MstCustomer, error) {
	return c.Query().Where(mstcustomer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MstCustomerClient) GetX(ctx context.Context, id int) *MstCustomer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCust2Inst queries the Cust2Inst edge of a MstCustomer.
func (c *MstCustomerClient) QueryCust2Inst(mc *MstCustomer) *MstInstQuery {
	query := &MstInstQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mstcustomer.Table, mstcustomer.FieldID, id),
			sqlgraph.To(mstinst.Table, mstinst.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mstcustomer.Cust2InstTable, mstcustomer.Cust2InstColumn),
		)
		fromV = sqlgraph.Neighbors(mc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MstCustomerClient) Hooks() []Hook {
	return c.hooks.MstCustomer
}

// MstInstClient is a client for the MstInst schema.
type MstInstClient struct {
	config
}

// NewMstInstClient returns a client for the MstInst from the given config.
func NewMstInstClient(c config) *MstInstClient {
	return &MstInstClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mstinst.Hooks(f(g(h())))`.
func (c *MstInstClient) Use(hooks ...Hook) {
	c.hooks.MstInst = append(c.hooks.MstInst, hooks...)
}

// Create returns a create builder for MstInst.
func (c *MstInstClient) Create() *MstInstCreate {
	mutation := newMstInstMutation(c.config, OpCreate)
	return &MstInstCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MstInst entities.
func (c *MstInstClient) CreateBulk(builders ...*MstInstCreate) *MstInstCreateBulk {
	return &MstInstCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MstInst.
func (c *MstInstClient) Update() *MstInstUpdate {
	mutation := newMstInstMutation(c.config, OpUpdate)
	return &MstInstUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MstInstClient) UpdateOne(mi *MstInst) *MstInstUpdateOne {
	mutation := newMstInstMutation(c.config, OpUpdateOne, withMstInst(mi))
	return &MstInstUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MstInstClient) UpdateOneID(id int) *MstInstUpdateOne {
	mutation := newMstInstMutation(c.config, OpUpdateOne, withMstInstID(id))
	return &MstInstUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MstInst.
func (c *MstInstClient) Delete() *MstInstDelete {
	mutation := newMstInstMutation(c.config, OpDelete)
	return &MstInstDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MstInstClient) DeleteOne(mi *MstInst) *MstInstDeleteOne {
	return c.DeleteOneID(mi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MstInstClient) DeleteOneID(id int) *MstInstDeleteOne {
	builder := c.Delete().Where(mstinst.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MstInstDeleteOne{builder}
}

// Query returns a query builder for MstInst.
func (c *MstInstClient) Query() *MstInstQuery {
	return &MstInstQuery{
		config: c.config,
	}
}

// Get returns a MstInst entity by its id.
func (c *MstInstClient) Get(ctx context.Context, id int) (*MstInst, error) {
	return c.Query().Where(mstinst.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MstInstClient) GetX(ctx context.Context, id int) *MstInst {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstfromCust queries the InstfromCust edge of a MstInst.
func (c *MstInstClient) QueryInstfromCust(mi *MstInst) *MstCustomerQuery {
	query := &MstCustomerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(mstinst.Table, mstinst.FieldID, id),
			sqlgraph.To(mstcustomer.Table, mstcustomer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mstinst.InstfromCustTable, mstinst.InstfromCustColumn),
		)
		fromV = sqlgraph.Neighbors(mi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MstInstClient) Hooks() []Hook {
	return c.hooks.MstInst
}

// MstStudentClient is a client for the MstStudent schema.
type MstStudentClient struct {
	config
}

// NewMstStudentClient returns a client for the MstStudent from the given config.
func NewMstStudentClient(c config) *MstStudentClient {
	return &MstStudentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mststudent.Hooks(f(g(h())))`.
func (c *MstStudentClient) Use(hooks ...Hook) {
	c.hooks.MstStudent = append(c.hooks.MstStudent, hooks...)
}

// Create returns a create builder for MstStudent.
func (c *MstStudentClient) Create() *MstStudentCreate {
	mutation := newMstStudentMutation(c.config, OpCreate)
	return &MstStudentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MstStudent entities.
func (c *MstStudentClient) CreateBulk(builders ...*MstStudentCreate) *MstStudentCreateBulk {
	return &MstStudentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MstStudent.
func (c *MstStudentClient) Update() *MstStudentUpdate {
	mutation := newMstStudentMutation(c.config, OpUpdate)
	return &MstStudentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MstStudentClient) UpdateOne(ms *MstStudent) *MstStudentUpdateOne {
	mutation := newMstStudentMutation(c.config, OpUpdateOne, withMstStudent(ms))
	return &MstStudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MstStudentClient) UpdateOneID(id int) *MstStudentUpdateOne {
	mutation := newMstStudentMutation(c.config, OpUpdateOne, withMstStudentID(id))
	return &MstStudentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MstStudent.
func (c *MstStudentClient) Delete() *MstStudentDelete {
	mutation := newMstStudentMutation(c.config, OpDelete)
	return &MstStudentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MstStudentClient) DeleteOne(ms *MstStudent) *MstStudentDeleteOne {
	return c.DeleteOneID(ms.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MstStudentClient) DeleteOneID(id int) *MstStudentDeleteOne {
	builder := c.Delete().Where(mststudent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MstStudentDeleteOne{builder}
}

// Query returns a query builder for MstStudent.
func (c *MstStudentClient) Query() *MstStudentQuery {
	return &MstStudentQuery{
		config: c.config,
	}
}

// Get returns a MstStudent entity by its id.
func (c *MstStudentClient) Get(ctx context.Context, id int) (*MstStudent, error) {
	return c.Query().Where(mststudent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MstStudentClient) GetX(ctx context.Context, id int) *MstStudent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MstStudentClient) Hooks() []Hook {
	return c.hooks.MstStudent
}

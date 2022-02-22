// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"myeduate/ent/migrate"
	"myeduate/ent/schema/pulid"

	"myeduate/ent/mstcustomer"
	"myeduate/ent/mstinst"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// MstCustomer is the client for interacting with the MstCustomer builders.
	MstCustomer *MstCustomerClient
	// MstInst is the client for interacting with the MstInst builders.
	MstInst *MstInstClient
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
	c.MstCustomer = NewMstCustomerClient(c.config)
	c.MstInst = NewMstInstClient(c.config)
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
		MstCustomer: NewMstCustomerClient(cfg),
		MstInst:     NewMstInstClient(cfg),
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
		MstCustomer: NewMstCustomerClient(cfg),
		MstInst:     NewMstInstClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		MstCustomer.
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
	c.MstCustomer.Use(hooks...)
	c.MstInst.Use(hooks...)
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
func (c *MstCustomerClient) UpdateOneID(id pulid.ID) *MstCustomerUpdateOne {
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
func (c *MstCustomerClient) DeleteOneID(id pulid.ID) *MstCustomerDeleteOne {
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
func (c *MstCustomerClient) Get(ctx context.Context, id pulid.ID) (*MstCustomer, error) {
	return c.Query().Where(mstcustomer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MstCustomerClient) GetX(ctx context.Context, id pulid.ID) *MstCustomer {
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
func (c *MstInstClient) UpdateOneID(id pulid.ID) *MstInstUpdateOne {
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
func (c *MstInstClient) DeleteOneID(id pulid.ID) *MstInstDeleteOne {
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
func (c *MstInstClient) Get(ctx context.Context, id pulid.ID) (*MstInst, error) {
	return c.Query().Where(mstinst.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MstInstClient) GetX(ctx context.Context, id pulid.ID) *MstInst {
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

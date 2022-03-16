// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"myeduate/ent/mstcustomer"
	"myeduate/ent/mstinst"
	"myeduate/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MstInstQuery is the builder for querying MstInst entities.
type MstInstQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MstInst
	// eager-loading edges.
	withInstfromCust *MstCustomerQuery
	modifiers        []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MstInstQuery builder.
func (miq *MstInstQuery) Where(ps ...predicate.MstInst) *MstInstQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit adds a limit step to the query.
func (miq *MstInstQuery) Limit(limit int) *MstInstQuery {
	miq.limit = &limit
	return miq
}

// Offset adds an offset step to the query.
func (miq *MstInstQuery) Offset(offset int) *MstInstQuery {
	miq.offset = &offset
	return miq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (miq *MstInstQuery) Unique(unique bool) *MstInstQuery {
	miq.unique = &unique
	return miq
}

// Order adds an order step to the query.
func (miq *MstInstQuery) Order(o ...OrderFunc) *MstInstQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// QueryInstfromCust chains the current query on the "InstfromCust" edge.
func (miq *MstInstQuery) QueryInstfromCust() *MstCustomerQuery {
	query := &MstCustomerQuery{config: miq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mstinst.Table, mstinst.FieldID, selector),
			sqlgraph.To(mstcustomer.Table, mstcustomer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mstinst.InstfromCustTable, mstinst.InstfromCustColumn),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MstInst entity from the query.
// Returns a *NotFoundError when no MstInst was found.
func (miq *MstInstQuery) First(ctx context.Context) (*MstInst, error) {
	nodes, err := miq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mstinst.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MstInstQuery) FirstX(ctx context.Context) *MstInst {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MstInst ID from the query.
// Returns a *NotFoundError when no MstInst ID was found.
func (miq *MstInstQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mstinst.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MstInstQuery) FirstIDX(ctx context.Context) int {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MstInst entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MstInst entity is found.
// Returns a *NotFoundError when no MstInst entities are found.
func (miq *MstInstQuery) Only(ctx context.Context) (*MstInst, error) {
	nodes, err := miq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mstinst.Label}
	default:
		return nil, &NotSingularError{mstinst.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MstInstQuery) OnlyX(ctx context.Context) *MstInst {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MstInst ID in the query.
// Returns a *NotSingularError when more than one MstInst ID is found.
// Returns a *NotFoundError when no entities are found.
func (miq *MstInstQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = &NotSingularError{mstinst.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MstInstQuery) OnlyIDX(ctx context.Context) int {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MstInsts.
func (miq *MstInstQuery) All(ctx context.Context) ([]*MstInst, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return miq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (miq *MstInstQuery) AllX(ctx context.Context) []*MstInst {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MstInst IDs.
func (miq *MstInstQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := miq.Select(mstinst.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MstInstQuery) IDsX(ctx context.Context) []int {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MstInstQuery) Count(ctx context.Context) (int, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return miq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MstInstQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MstInstQuery) Exist(ctx context.Context) (bool, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return miq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MstInstQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MstInstQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MstInstQuery) Clone() *MstInstQuery {
	if miq == nil {
		return nil
	}
	return &MstInstQuery{
		config:           miq.config,
		limit:            miq.limit,
		offset:           miq.offset,
		order:            append([]OrderFunc{}, miq.order...),
		predicates:       append([]predicate.MstInst{}, miq.predicates...),
		withInstfromCust: miq.withInstfromCust.Clone(),
		// clone intermediate query.
		sql:    miq.sql.Clone(),
		path:   miq.path,
		unique: miq.unique,
	}
}

// WithInstfromCust tells the query-builder to eager-load the nodes that are connected to
// the "InstfromCust" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MstInstQuery) WithInstfromCust(opts ...func(*MstCustomerQuery)) *MstInstQuery {
	query := &MstCustomerQuery{config: miq.config}
	for _, opt := range opts {
		opt(query)
	}
	miq.withInstfromCust = query
	return miq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MstInst.Query().
//		GroupBy(mstinst.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (miq *MstInstQuery) GroupBy(field string, fields ...string) *MstInstGroupBy {
	group := &MstInstGroupBy{config: miq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return miq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.MstInst.Query().
//		Select(mstinst.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (miq *MstInstQuery) Select(fields ...string) *MstInstSelect {
	miq.fields = append(miq.fields, fields...)
	return &MstInstSelect{MstInstQuery: miq}
}

func (miq *MstInstQuery) prepareQuery(ctx context.Context) error {
	for _, f := range miq.fields {
		if !mstinst.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MstInstQuery) sqlAll(ctx context.Context) ([]*MstInst, error) {
	var (
		nodes       = []*MstInst{}
		_spec       = miq.querySpec()
		loadedTypes = [1]bool{
			miq.withInstfromCust != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MstInst{config: miq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(miq.modifiers) > 0 {
		_spec.Modifiers = miq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := miq.withInstfromCust; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*MstInst)
		for i := range nodes {
			fk := nodes[i].CustomerID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(mstcustomer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.InstfromCust = n
			}
		}
	}

	return nodes, nil
}

func (miq *MstInstQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	if len(miq.modifiers) > 0 {
		_spec.Modifiers = miq.modifiers
	}
	_spec.Node.Columns = miq.fields
	if len(miq.fields) > 0 {
		_spec.Unique = miq.unique != nil && *miq.unique
	}
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MstInstQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := miq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (miq *MstInstQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mstinst.Table,
			Columns: mstinst.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mstinst.FieldID,
			},
		},
		From:   miq.sql,
		Unique: true,
	}
	if unique := miq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := miq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mstinst.FieldID)
		for i := range fields {
			if fields[i] != mstinst.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (miq *MstInstQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(mstinst.Table)
	columns := miq.fields
	if len(columns) == 0 {
		columns = mstinst.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if miq.unique != nil && *miq.unique {
		selector.Distinct()
	}
	for _, m := range miq.modifiers {
		m(selector)
	}
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector)
	}
	if offset := miq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (miq *MstInstQuery) Modify(modifiers ...func(s *sql.Selector)) *MstInstSelect {
	miq.modifiers = append(miq.modifiers, modifiers...)
	return miq.Select()
}

// MstInstGroupBy is the group-by builder for MstInst entities.
type MstInstGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MstInstGroupBy) Aggregate(fns ...AggregateFunc) *MstInstGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the group-by query and scans the result into the given value.
func (migb *MstInstGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := migb.path(ctx)
	if err != nil {
		return err
	}
	migb.sql = query
	return migb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (migb *MstInstGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := migb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MstInstGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (migb *MstInstGroupBy) StringsX(ctx context.Context) []string {
	v, err := migb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = migb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (migb *MstInstGroupBy) StringX(ctx context.Context) string {
	v, err := migb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MstInstGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (migb *MstInstGroupBy) IntsX(ctx context.Context) []int {
	v, err := migb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = migb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (migb *MstInstGroupBy) IntX(ctx context.Context) int {
	v, err := migb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MstInstGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (migb *MstInstGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := migb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = migb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (migb *MstInstGroupBy) Float64X(ctx context.Context) float64 {
	v, err := migb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MstInstGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (migb *MstInstGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := migb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (migb *MstInstGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = migb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (migb *MstInstGroupBy) BoolX(ctx context.Context) bool {
	v, err := migb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (migb *MstInstGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range migb.fields {
		if !mstinst.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := migb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (migb *MstInstGroupBy) sqlQuery() *sql.Selector {
	selector := migb.sql.Select()
	aggregation := make([]string, 0, len(migb.fns))
	for _, fn := range migb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(migb.fields)+len(migb.fns))
		for _, f := range migb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(migb.fields...)...)
}

// MstInstSelect is the builder for selecting fields of MstInst entities.
type MstInstSelect struct {
	*MstInstQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mis *MstInstSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mis.prepareQuery(ctx); err != nil {
		return err
	}
	mis.sql = mis.MstInstQuery.sqlQuery(ctx)
	return mis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mis *MstInstSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MstInstSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mis *MstInstSelect) StringsX(ctx context.Context) []string {
	v, err := mis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mis *MstInstSelect) StringX(ctx context.Context) string {
	v, err := mis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MstInstSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mis *MstInstSelect) IntsX(ctx context.Context) []int {
	v, err := mis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mis *MstInstSelect) IntX(ctx context.Context) int {
	v, err := mis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MstInstSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mis *MstInstSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mis *MstInstSelect) Float64X(ctx context.Context) float64 {
	v, err := mis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MstInstSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mis *MstInstSelect) BoolsX(ctx context.Context) []bool {
	v, err := mis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mis *MstInstSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mstinst.Label}
	default:
		err = fmt.Errorf("ent: MstInstSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mis *MstInstSelect) BoolX(ctx context.Context) bool {
	v, err := mis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mis *MstInstSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mis.sql.Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mis *MstInstSelect) Modify(modifiers ...func(s *sql.Selector)) *MstInstSelect {
	mis.modifiers = append(mis.modifiers, modifiers...)
	return mis
}

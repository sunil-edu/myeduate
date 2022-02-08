// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myeduate/ent/customtypes"
	"myeduate/ent/mstcustomer"
	"myeduate/ent/mstinst"
	"myeduate/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MstCustomerUpdate is the builder for updating MstCustomer entities.
type MstCustomerUpdate struct {
	config
	hooks    []Hook
	mutation *MstCustomerMutation
}

// Where appends a list predicates to the MstCustomerUpdate builder.
func (mcu *MstCustomerUpdate) Where(ps ...predicate.MstCustomer) *MstCustomerUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetUpdatedAt sets the "updated_at" field.
func (mcu *MstCustomerUpdate) SetUpdatedAt(t time.Time) *MstCustomerUpdate {
	mcu.mutation.SetUpdatedAt(t)
	return mcu
}

// SetCustCode sets the "cust_code" field.
func (mcu *MstCustomerUpdate) SetCustCode(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustCode(s)
	return mcu
}

// SetCustName sets the "cust_name" field.
func (mcu *MstCustomerUpdate) SetCustName(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustName(s)
	return mcu
}

// SetCustAddress sets the "cust_address" field.
func (mcu *MstCustomerUpdate) SetCustAddress(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustAddress(s)
	return mcu
}

// SetCustPlace sets the "cust_place" field.
func (mcu *MstCustomerUpdate) SetCustPlace(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustPlace(s)
	return mcu
}

// SetCustState sets the "cust_state" field.
func (mcu *MstCustomerUpdate) SetCustState(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustState(s)
	return mcu
}

// SetCustPin sets the "cust_pin" field.
func (mcu *MstCustomerUpdate) SetCustPin(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustPin(s)
	return mcu
}

// SetCustContactPerson sets the "cust_contact_person" field.
func (mcu *MstCustomerUpdate) SetCustContactPerson(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustContactPerson(s)
	return mcu
}

// SetCustPhone sets the "cust_phone" field.
func (mcu *MstCustomerUpdate) SetCustPhone(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustPhone(s)
	return mcu
}

// SetCustEmail sets the "cust_email" field.
func (mcu *MstCustomerUpdate) SetCustEmail(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustEmail(s)
	return mcu
}

// SetCustMobile sets the "cust_mobile" field.
func (mcu *MstCustomerUpdate) SetCustMobile(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustMobile(s)
	return mcu
}

// SetCustURL sets the "cust_url" field.
func (mcu *MstCustomerUpdate) SetCustURL(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustURL(s)
	return mcu
}

// SetCustBanner1 sets the "cust_banner_1" field.
func (mcu *MstCustomerUpdate) SetCustBanner1(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustBanner1(s)
	return mcu
}

// SetCustBanner2 sets the "cust_banner_2" field.
func (mcu *MstCustomerUpdate) SetCustBanner2(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustBanner2(s)
	return mcu
}

// SetCustLogoURL sets the "cust_logo_url" field.
func (mcu *MstCustomerUpdate) SetCustLogoURL(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustLogoURL(s)
	return mcu
}

// SetCustIsActive sets the "cust_is_active" field.
func (mcu *MstCustomerUpdate) SetCustIsActive(ca customtypes.IsActive) *MstCustomerUpdate {
	mcu.mutation.SetCustIsActive(ca)
	return mcu
}

// SetNillableCustIsActive sets the "cust_is_active" field if the given value is not nil.
func (mcu *MstCustomerUpdate) SetNillableCustIsActive(ca *customtypes.IsActive) *MstCustomerUpdate {
	if ca != nil {
		mcu.SetCustIsActive(*ca)
	}
	return mcu
}

// SetCustStatus sets the "cust_status" field.
func (mcu *MstCustomerUpdate) SetCustStatus(s string) *MstCustomerUpdate {
	mcu.mutation.SetCustStatus(s)
	return mcu
}

// SetCustNumInst sets the "cust_num_inst" field.
func (mcu *MstCustomerUpdate) SetCustNumInst(i int) *MstCustomerUpdate {
	mcu.mutation.ResetCustNumInst()
	mcu.mutation.SetCustNumInst(i)
	return mcu
}

// SetNillableCustNumInst sets the "cust_num_inst" field if the given value is not nil.
func (mcu *MstCustomerUpdate) SetNillableCustNumInst(i *int) *MstCustomerUpdate {
	if i != nil {
		mcu.SetCustNumInst(*i)
	}
	return mcu
}

// AddCustNumInst adds i to the "cust_num_inst" field.
func (mcu *MstCustomerUpdate) AddCustNumInst(i int) *MstCustomerUpdate {
	mcu.mutation.AddCustNumInst(i)
	return mcu
}

// SetCustTimeZone sets the "cust_time_zone" field.
func (mcu *MstCustomerUpdate) SetCustTimeZone(t time.Time) *MstCustomerUpdate {
	mcu.mutation.SetCustTimeZone(t)
	return mcu
}

// AddCust2InstIDs adds the "Cust2Inst" edge to the MstInst entity by IDs.
func (mcu *MstCustomerUpdate) AddCust2InstIDs(ids ...uuid.UUID) *MstCustomerUpdate {
	mcu.mutation.AddCust2InstIDs(ids...)
	return mcu
}

// AddCust2Inst adds the "Cust2Inst" edges to the MstInst entity.
func (mcu *MstCustomerUpdate) AddCust2Inst(m ...*MstInst) *MstCustomerUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.AddCust2InstIDs(ids...)
}

// Mutation returns the MstCustomerMutation object of the builder.
func (mcu *MstCustomerUpdate) Mutation() *MstCustomerMutation {
	return mcu.mutation
}

// ClearCust2Inst clears all "Cust2Inst" edges to the MstInst entity.
func (mcu *MstCustomerUpdate) ClearCust2Inst() *MstCustomerUpdate {
	mcu.mutation.ClearCust2Inst()
	return mcu
}

// RemoveCust2InstIDs removes the "Cust2Inst" edge to MstInst entities by IDs.
func (mcu *MstCustomerUpdate) RemoveCust2InstIDs(ids ...uuid.UUID) *MstCustomerUpdate {
	mcu.mutation.RemoveCust2InstIDs(ids...)
	return mcu
}

// RemoveCust2Inst removes "Cust2Inst" edges to MstInst entities.
func (mcu *MstCustomerUpdate) RemoveCust2Inst(m ...*MstInst) *MstCustomerUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcu.RemoveCust2InstIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MstCustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mcu.defaults()
	if len(mcu.hooks) == 0 {
		if err = mcu.check(); err != nil {
			return 0, err
		}
		affected, err = mcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MstCustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mcu.check(); err != nil {
				return 0, err
			}
			mcu.mutation = mutation
			affected, err = mcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mcu.hooks) - 1; i >= 0; i-- {
			if mcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MstCustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MstCustomerUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MstCustomerUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcu *MstCustomerUpdate) defaults() {
	if _, ok := mcu.mutation.UpdatedAt(); !ok {
		v := mstcustomer.UpdateDefaultUpdatedAt()
		mcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcu *MstCustomerUpdate) check() error {
	if v, ok := mcu.mutation.CustCode(); ok {
		if err := mstcustomer.CustCodeValidator(v); err != nil {
			return &ValidationError{Name: "cust_code", err: fmt.Errorf(`ent: validator failed for field "MstCustomer.cust_code": %w`, err)}
		}
	}
	if v, ok := mcu.mutation.CustName(); ok {
		if err := mstcustomer.CustNameValidator(v); err != nil {
			return &ValidationError{Name: "cust_name", err: fmt.Errorf(`ent: validator failed for field "MstCustomer.cust_name": %w`, err)}
		}
	}
	if v, ok := mcu.mutation.CustIsActive(); ok {
		if err := mstcustomer.CustIsActiveValidator(v); err != nil {
			return &ValidationError{Name: "cust_is_active", err: fmt.Errorf(`ent: validator failed for field "MstCustomer.cust_is_active": %w`, err)}
		}
	}
	return nil
}

func (mcu *MstCustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mstcustomer.Table,
			Columns: mstcustomer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mstcustomer.FieldID,
			},
		},
	}
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstcustomer.FieldUpdatedAt,
		})
	}
	if value, ok := mcu.mutation.CustCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustCode,
		})
	}
	if value, ok := mcu.mutation.CustName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustName,
		})
	}
	if value, ok := mcu.mutation.CustAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustAddress,
		})
	}
	if value, ok := mcu.mutation.CustPlace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustPlace,
		})
	}
	if value, ok := mcu.mutation.CustState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustState,
		})
	}
	if value, ok := mcu.mutation.CustPin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustPin,
		})
	}
	if value, ok := mcu.mutation.CustContactPerson(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustContactPerson,
		})
	}
	if value, ok := mcu.mutation.CustPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustPhone,
		})
	}
	if value, ok := mcu.mutation.CustEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustEmail,
		})
	}
	if value, ok := mcu.mutation.CustMobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustMobile,
		})
	}
	if value, ok := mcu.mutation.CustURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustURL,
		})
	}
	if value, ok := mcu.mutation.CustBanner1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustBanner1,
		})
	}
	if value, ok := mcu.mutation.CustBanner2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustBanner2,
		})
	}
	if value, ok := mcu.mutation.CustLogoURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustLogoURL,
		})
	}
	if value, ok := mcu.mutation.CustIsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mstcustomer.FieldCustIsActive,
		})
	}
	if value, ok := mcu.mutation.CustStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustStatus,
		})
	}
	if value, ok := mcu.mutation.CustNumInst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mstcustomer.FieldCustNumInst,
		})
	}
	if value, ok := mcu.mutation.AddedCustNumInst(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mstcustomer.FieldCustNumInst,
		})
	}
	if value, ok := mcu.mutation.CustTimeZone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstcustomer.FieldCustTimeZone,
		})
	}
	if mcu.mutation.Cust2InstCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mstcustomer.Cust2InstTable,
			Columns: []string{mstcustomer.Cust2InstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mstinst.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.RemovedCust2InstIDs(); len(nodes) > 0 && !mcu.mutation.Cust2InstCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mstcustomer.Cust2InstTable,
			Columns: []string{mstcustomer.Cust2InstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mstinst.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcu.mutation.Cust2InstIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mstcustomer.Cust2InstTable,
			Columns: []string{mstcustomer.Cust2InstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mstinst.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mstcustomer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MstCustomerUpdateOne is the builder for updating a single MstCustomer entity.
type MstCustomerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MstCustomerMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (mcuo *MstCustomerUpdateOne) SetUpdatedAt(t time.Time) *MstCustomerUpdateOne {
	mcuo.mutation.SetUpdatedAt(t)
	return mcuo
}

// SetCustCode sets the "cust_code" field.
func (mcuo *MstCustomerUpdateOne) SetCustCode(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustCode(s)
	return mcuo
}

// SetCustName sets the "cust_name" field.
func (mcuo *MstCustomerUpdateOne) SetCustName(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustName(s)
	return mcuo
}

// SetCustAddress sets the "cust_address" field.
func (mcuo *MstCustomerUpdateOne) SetCustAddress(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustAddress(s)
	return mcuo
}

// SetCustPlace sets the "cust_place" field.
func (mcuo *MstCustomerUpdateOne) SetCustPlace(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustPlace(s)
	return mcuo
}

// SetCustState sets the "cust_state" field.
func (mcuo *MstCustomerUpdateOne) SetCustState(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustState(s)
	return mcuo
}

// SetCustPin sets the "cust_pin" field.
func (mcuo *MstCustomerUpdateOne) SetCustPin(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustPin(s)
	return mcuo
}

// SetCustContactPerson sets the "cust_contact_person" field.
func (mcuo *MstCustomerUpdateOne) SetCustContactPerson(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustContactPerson(s)
	return mcuo
}

// SetCustPhone sets the "cust_phone" field.
func (mcuo *MstCustomerUpdateOne) SetCustPhone(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustPhone(s)
	return mcuo
}

// SetCustEmail sets the "cust_email" field.
func (mcuo *MstCustomerUpdateOne) SetCustEmail(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustEmail(s)
	return mcuo
}

// SetCustMobile sets the "cust_mobile" field.
func (mcuo *MstCustomerUpdateOne) SetCustMobile(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustMobile(s)
	return mcuo
}

// SetCustURL sets the "cust_url" field.
func (mcuo *MstCustomerUpdateOne) SetCustURL(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustURL(s)
	return mcuo
}

// SetCustBanner1 sets the "cust_banner_1" field.
func (mcuo *MstCustomerUpdateOne) SetCustBanner1(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustBanner1(s)
	return mcuo
}

// SetCustBanner2 sets the "cust_banner_2" field.
func (mcuo *MstCustomerUpdateOne) SetCustBanner2(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustBanner2(s)
	return mcuo
}

// SetCustLogoURL sets the "cust_logo_url" field.
func (mcuo *MstCustomerUpdateOne) SetCustLogoURL(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustLogoURL(s)
	return mcuo
}

// SetCustIsActive sets the "cust_is_active" field.
func (mcuo *MstCustomerUpdateOne) SetCustIsActive(ca customtypes.IsActive) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustIsActive(ca)
	return mcuo
}

// SetNillableCustIsActive sets the "cust_is_active" field if the given value is not nil.
func (mcuo *MstCustomerUpdateOne) SetNillableCustIsActive(ca *customtypes.IsActive) *MstCustomerUpdateOne {
	if ca != nil {
		mcuo.SetCustIsActive(*ca)
	}
	return mcuo
}

// SetCustStatus sets the "cust_status" field.
func (mcuo *MstCustomerUpdateOne) SetCustStatus(s string) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustStatus(s)
	return mcuo
}

// SetCustNumInst sets the "cust_num_inst" field.
func (mcuo *MstCustomerUpdateOne) SetCustNumInst(i int) *MstCustomerUpdateOne {
	mcuo.mutation.ResetCustNumInst()
	mcuo.mutation.SetCustNumInst(i)
	return mcuo
}

// SetNillableCustNumInst sets the "cust_num_inst" field if the given value is not nil.
func (mcuo *MstCustomerUpdateOne) SetNillableCustNumInst(i *int) *MstCustomerUpdateOne {
	if i != nil {
		mcuo.SetCustNumInst(*i)
	}
	return mcuo
}

// AddCustNumInst adds i to the "cust_num_inst" field.
func (mcuo *MstCustomerUpdateOne) AddCustNumInst(i int) *MstCustomerUpdateOne {
	mcuo.mutation.AddCustNumInst(i)
	return mcuo
}

// SetCustTimeZone sets the "cust_time_zone" field.
func (mcuo *MstCustomerUpdateOne) SetCustTimeZone(t time.Time) *MstCustomerUpdateOne {
	mcuo.mutation.SetCustTimeZone(t)
	return mcuo
}

// AddCust2InstIDs adds the "Cust2Inst" edge to the MstInst entity by IDs.
func (mcuo *MstCustomerUpdateOne) AddCust2InstIDs(ids ...uuid.UUID) *MstCustomerUpdateOne {
	mcuo.mutation.AddCust2InstIDs(ids...)
	return mcuo
}

// AddCust2Inst adds the "Cust2Inst" edges to the MstInst entity.
func (mcuo *MstCustomerUpdateOne) AddCust2Inst(m ...*MstInst) *MstCustomerUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.AddCust2InstIDs(ids...)
}

// Mutation returns the MstCustomerMutation object of the builder.
func (mcuo *MstCustomerUpdateOne) Mutation() *MstCustomerMutation {
	return mcuo.mutation
}

// ClearCust2Inst clears all "Cust2Inst" edges to the MstInst entity.
func (mcuo *MstCustomerUpdateOne) ClearCust2Inst() *MstCustomerUpdateOne {
	mcuo.mutation.ClearCust2Inst()
	return mcuo
}

// RemoveCust2InstIDs removes the "Cust2Inst" edge to MstInst entities by IDs.
func (mcuo *MstCustomerUpdateOne) RemoveCust2InstIDs(ids ...uuid.UUID) *MstCustomerUpdateOne {
	mcuo.mutation.RemoveCust2InstIDs(ids...)
	return mcuo
}

// RemoveCust2Inst removes "Cust2Inst" edges to MstInst entities.
func (mcuo *MstCustomerUpdateOne) RemoveCust2Inst(m ...*MstInst) *MstCustomerUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mcuo.RemoveCust2InstIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MstCustomerUpdateOne) Select(field string, fields ...string) *MstCustomerUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MstCustomer entity.
func (mcuo *MstCustomerUpdateOne) Save(ctx context.Context) (*MstCustomer, error) {
	var (
		err  error
		node *MstCustomer
	)
	mcuo.defaults()
	if len(mcuo.hooks) == 0 {
		if err = mcuo.check(); err != nil {
			return nil, err
		}
		node, err = mcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MstCustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mcuo.check(); err != nil {
				return nil, err
			}
			mcuo.mutation = mutation
			node, err = mcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mcuo.hooks) - 1; i >= 0; i-- {
			if mcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MstCustomerUpdateOne) SaveX(ctx context.Context) *MstCustomer {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MstCustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MstCustomerUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcuo *MstCustomerUpdateOne) defaults() {
	if _, ok := mcuo.mutation.UpdatedAt(); !ok {
		v := mstcustomer.UpdateDefaultUpdatedAt()
		mcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcuo *MstCustomerUpdateOne) check() error {
	if v, ok := mcuo.mutation.CustCode(); ok {
		if err := mstcustomer.CustCodeValidator(v); err != nil {
			return &ValidationError{Name: "cust_code", err: fmt.Errorf(`ent: validator failed for field "MstCustomer.cust_code": %w`, err)}
		}
	}
	if v, ok := mcuo.mutation.CustName(); ok {
		if err := mstcustomer.CustNameValidator(v); err != nil {
			return &ValidationError{Name: "cust_name", err: fmt.Errorf(`ent: validator failed for field "MstCustomer.cust_name": %w`, err)}
		}
	}
	if v, ok := mcuo.mutation.CustIsActive(); ok {
		if err := mstcustomer.CustIsActiveValidator(v); err != nil {
			return &ValidationError{Name: "cust_is_active", err: fmt.Errorf(`ent: validator failed for field "MstCustomer.cust_is_active": %w`, err)}
		}
	}
	return nil
}

func (mcuo *MstCustomerUpdateOne) sqlSave(ctx context.Context) (_node *MstCustomer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mstcustomer.Table,
			Columns: mstcustomer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mstcustomer.FieldID,
			},
		},
	}
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MstCustomer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mstcustomer.FieldID)
		for _, f := range fields {
			if !mstcustomer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mstcustomer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstcustomer.FieldUpdatedAt,
		})
	}
	if value, ok := mcuo.mutation.CustCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustCode,
		})
	}
	if value, ok := mcuo.mutation.CustName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustName,
		})
	}
	if value, ok := mcuo.mutation.CustAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustAddress,
		})
	}
	if value, ok := mcuo.mutation.CustPlace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustPlace,
		})
	}
	if value, ok := mcuo.mutation.CustState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustState,
		})
	}
	if value, ok := mcuo.mutation.CustPin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustPin,
		})
	}
	if value, ok := mcuo.mutation.CustContactPerson(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustContactPerson,
		})
	}
	if value, ok := mcuo.mutation.CustPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustPhone,
		})
	}
	if value, ok := mcuo.mutation.CustEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustEmail,
		})
	}
	if value, ok := mcuo.mutation.CustMobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustMobile,
		})
	}
	if value, ok := mcuo.mutation.CustURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustURL,
		})
	}
	if value, ok := mcuo.mutation.CustBanner1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustBanner1,
		})
	}
	if value, ok := mcuo.mutation.CustBanner2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustBanner2,
		})
	}
	if value, ok := mcuo.mutation.CustLogoURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustLogoURL,
		})
	}
	if value, ok := mcuo.mutation.CustIsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mstcustomer.FieldCustIsActive,
		})
	}
	if value, ok := mcuo.mutation.CustStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstcustomer.FieldCustStatus,
		})
	}
	if value, ok := mcuo.mutation.CustNumInst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mstcustomer.FieldCustNumInst,
		})
	}
	if value, ok := mcuo.mutation.AddedCustNumInst(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mstcustomer.FieldCustNumInst,
		})
	}
	if value, ok := mcuo.mutation.CustTimeZone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstcustomer.FieldCustTimeZone,
		})
	}
	if mcuo.mutation.Cust2InstCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mstcustomer.Cust2InstTable,
			Columns: []string{mstcustomer.Cust2InstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mstinst.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.RemovedCust2InstIDs(); len(nodes) > 0 && !mcuo.mutation.Cust2InstCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mstcustomer.Cust2InstTable,
			Columns: []string{mstcustomer.Cust2InstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mstinst.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mcuo.mutation.Cust2InstIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mstcustomer.Cust2InstTable,
			Columns: []string{mstcustomer.Cust2InstColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mstinst.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MstCustomer{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mstcustomer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

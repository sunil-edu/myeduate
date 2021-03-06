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
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MstInstCreate is the builder for creating a MstInst entity.
type MstInstCreate struct {
	config
	mutation *MstInstMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mic *MstInstCreate) SetCreatedAt(t time.Time) *MstInstCreate {
	mic.mutation.SetCreatedAt(t)
	return mic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mic *MstInstCreate) SetNillableCreatedAt(t *time.Time) *MstInstCreate {
	if t != nil {
		mic.SetCreatedAt(*t)
	}
	return mic
}

// SetUpdatedAt sets the "updated_at" field.
func (mic *MstInstCreate) SetUpdatedAt(t time.Time) *MstInstCreate {
	mic.mutation.SetUpdatedAt(t)
	return mic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mic *MstInstCreate) SetNillableUpdatedAt(t *time.Time) *MstInstCreate {
	if t != nil {
		mic.SetUpdatedAt(*t)
	}
	return mic
}

// SetInstCode sets the "inst_code" field.
func (mic *MstInstCreate) SetInstCode(s string) *MstInstCreate {
	mic.mutation.SetInstCode(s)
	return mic
}

// SetInstName sets the "inst_name" field.
func (mic *MstInstCreate) SetInstName(s string) *MstInstCreate {
	mic.mutation.SetInstName(s)
	return mic
}

// SetInstShortName sets the "inst_short_name" field.
func (mic *MstInstCreate) SetInstShortName(s string) *MstInstCreate {
	mic.mutation.SetInstShortName(s)
	return mic
}

// SetInstAddress sets the "inst_address" field.
func (mic *MstInstCreate) SetInstAddress(s string) *MstInstCreate {
	mic.mutation.SetInstAddress(s)
	return mic
}

// SetInstPlace sets the "inst_place" field.
func (mic *MstInstCreate) SetInstPlace(s string) *MstInstCreate {
	mic.mutation.SetInstPlace(s)
	return mic
}

// SetInstState sets the "inst_state" field.
func (mic *MstInstCreate) SetInstState(s string) *MstInstCreate {
	mic.mutation.SetInstState(s)
	return mic
}

// SetInstPin sets the "inst_pin" field.
func (mic *MstInstCreate) SetInstPin(s string) *MstInstCreate {
	mic.mutation.SetInstPin(s)
	return mic
}

// SetInstContactPerson sets the "inst_contact_person" field.
func (mic *MstInstCreate) SetInstContactPerson(s string) *MstInstCreate {
	mic.mutation.SetInstContactPerson(s)
	return mic
}

// SetInstPhone sets the "inst_phone" field.
func (mic *MstInstCreate) SetInstPhone(s string) *MstInstCreate {
	mic.mutation.SetInstPhone(s)
	return mic
}

// SetInstEmail sets the "inst_email" field.
func (mic *MstInstCreate) SetInstEmail(s string) *MstInstCreate {
	mic.mutation.SetInstEmail(s)
	return mic
}

// SetInstMobile sets the "inst_mobile" field.
func (mic *MstInstCreate) SetInstMobile(s string) *MstInstCreate {
	mic.mutation.SetInstMobile(s)
	return mic
}

// SetInstURL sets the "inst_url" field.
func (mic *MstInstCreate) SetInstURL(s string) *MstInstCreate {
	mic.mutation.SetInstURL(s)
	return mic
}

// SetInstBanner1 sets the "inst_banner_1" field.
func (mic *MstInstCreate) SetInstBanner1(s string) *MstInstCreate {
	mic.mutation.SetInstBanner1(s)
	return mic
}

// SetInstBanner2 sets the "inst_banner_2" field.
func (mic *MstInstCreate) SetInstBanner2(s string) *MstInstCreate {
	mic.mutation.SetInstBanner2(s)
	return mic
}

// SetInstLogoURL sets the "inst_logo_url" field.
func (mic *MstInstCreate) SetInstLogoURL(s string) *MstInstCreate {
	mic.mutation.SetInstLogoURL(s)
	return mic
}

// SetInstIsActive sets the "inst_is_active" field.
func (mic *MstInstCreate) SetInstIsActive(ca customtypes.IsActive) *MstInstCreate {
	mic.mutation.SetInstIsActive(ca)
	return mic
}

// SetNillableInstIsActive sets the "inst_is_active" field if the given value is not nil.
func (mic *MstInstCreate) SetNillableInstIsActive(ca *customtypes.IsActive) *MstInstCreate {
	if ca != nil {
		mic.SetInstIsActive(*ca)
	}
	return mic
}

// SetInstStatus sets the "inst_status" field.
func (mic *MstInstCreate) SetInstStatus(s string) *MstInstCreate {
	mic.mutation.SetInstStatus(s)
	return mic
}

// SetInstTimeZone sets the "inst_time_zone" field.
func (mic *MstInstCreate) SetInstTimeZone(t time.Time) *MstInstCreate {
	mic.mutation.SetInstTimeZone(t)
	return mic
}

// SetCustomerID sets the "customer_id" field.
func (mic *MstInstCreate) SetCustomerID(i int) *MstInstCreate {
	mic.mutation.SetCustomerID(i)
	return mic
}

// SetInstfromCustID sets the "InstfromCust" edge to the MstCustomer entity by ID.
func (mic *MstInstCreate) SetInstfromCustID(id int) *MstInstCreate {
	mic.mutation.SetInstfromCustID(id)
	return mic
}

// SetInstfromCust sets the "InstfromCust" edge to the MstCustomer entity.
func (mic *MstInstCreate) SetInstfromCust(m *MstCustomer) *MstInstCreate {
	return mic.SetInstfromCustID(m.ID)
}

// Mutation returns the MstInstMutation object of the builder.
func (mic *MstInstCreate) Mutation() *MstInstMutation {
	return mic.mutation
}

// Save creates the MstInst in the database.
func (mic *MstInstCreate) Save(ctx context.Context) (*MstInst, error) {
	var (
		err  error
		node *MstInst
	)
	mic.defaults()
	if len(mic.hooks) == 0 {
		if err = mic.check(); err != nil {
			return nil, err
		}
		node, err = mic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MstInstMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mic.check(); err != nil {
				return nil, err
			}
			mic.mutation = mutation
			if node, err = mic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mic.hooks) - 1; i >= 0; i-- {
			if mic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MstInstCreate) SaveX(ctx context.Context) *MstInst {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mic *MstInstCreate) Exec(ctx context.Context) error {
	_, err := mic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mic *MstInstCreate) ExecX(ctx context.Context) {
	if err := mic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mic *MstInstCreate) defaults() {
	if _, ok := mic.mutation.CreatedAt(); !ok {
		v := mstinst.DefaultCreatedAt()
		mic.mutation.SetCreatedAt(v)
	}
	if _, ok := mic.mutation.UpdatedAt(); !ok {
		v := mstinst.DefaultUpdatedAt()
		mic.mutation.SetUpdatedAt(v)
	}
	if _, ok := mic.mutation.InstIsActive(); !ok {
		v := mstinst.DefaultInstIsActive
		mic.mutation.SetInstIsActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mic *MstInstCreate) check() error {
	if _, ok := mic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MstInst.created_at"`)}
	}
	if _, ok := mic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MstInst.updated_at"`)}
	}
	if _, ok := mic.mutation.InstCode(); !ok {
		return &ValidationError{Name: "inst_code", err: errors.New(`ent: missing required field "MstInst.inst_code"`)}
	}
	if _, ok := mic.mutation.InstName(); !ok {
		return &ValidationError{Name: "inst_name", err: errors.New(`ent: missing required field "MstInst.inst_name"`)}
	}
	if v, ok := mic.mutation.InstName(); ok {
		if err := mstinst.InstNameValidator(v); err != nil {
			return &ValidationError{Name: "inst_name", err: fmt.Errorf(`ent: validator failed for field "MstInst.inst_name": %w`, err)}
		}
	}
	if _, ok := mic.mutation.InstShortName(); !ok {
		return &ValidationError{Name: "inst_short_name", err: errors.New(`ent: missing required field "MstInst.inst_short_name"`)}
	}
	if _, ok := mic.mutation.InstAddress(); !ok {
		return &ValidationError{Name: "inst_address", err: errors.New(`ent: missing required field "MstInst.inst_address"`)}
	}
	if _, ok := mic.mutation.InstPlace(); !ok {
		return &ValidationError{Name: "inst_place", err: errors.New(`ent: missing required field "MstInst.inst_place"`)}
	}
	if _, ok := mic.mutation.InstState(); !ok {
		return &ValidationError{Name: "inst_state", err: errors.New(`ent: missing required field "MstInst.inst_state"`)}
	}
	if _, ok := mic.mutation.InstPin(); !ok {
		return &ValidationError{Name: "inst_pin", err: errors.New(`ent: missing required field "MstInst.inst_pin"`)}
	}
	if _, ok := mic.mutation.InstContactPerson(); !ok {
		return &ValidationError{Name: "inst_contact_person", err: errors.New(`ent: missing required field "MstInst.inst_contact_person"`)}
	}
	if _, ok := mic.mutation.InstPhone(); !ok {
		return &ValidationError{Name: "inst_phone", err: errors.New(`ent: missing required field "MstInst.inst_phone"`)}
	}
	if _, ok := mic.mutation.InstEmail(); !ok {
		return &ValidationError{Name: "inst_email", err: errors.New(`ent: missing required field "MstInst.inst_email"`)}
	}
	if _, ok := mic.mutation.InstMobile(); !ok {
		return &ValidationError{Name: "inst_mobile", err: errors.New(`ent: missing required field "MstInst.inst_mobile"`)}
	}
	if _, ok := mic.mutation.InstURL(); !ok {
		return &ValidationError{Name: "inst_url", err: errors.New(`ent: missing required field "MstInst.inst_url"`)}
	}
	if _, ok := mic.mutation.InstBanner1(); !ok {
		return &ValidationError{Name: "inst_banner_1", err: errors.New(`ent: missing required field "MstInst.inst_banner_1"`)}
	}
	if _, ok := mic.mutation.InstBanner2(); !ok {
		return &ValidationError{Name: "inst_banner_2", err: errors.New(`ent: missing required field "MstInst.inst_banner_2"`)}
	}
	if _, ok := mic.mutation.InstLogoURL(); !ok {
		return &ValidationError{Name: "inst_logo_url", err: errors.New(`ent: missing required field "MstInst.inst_logo_url"`)}
	}
	if _, ok := mic.mutation.InstIsActive(); !ok {
		return &ValidationError{Name: "inst_is_active", err: errors.New(`ent: missing required field "MstInst.inst_is_active"`)}
	}
	if v, ok := mic.mutation.InstIsActive(); ok {
		if err := mstinst.InstIsActiveValidator(v); err != nil {
			return &ValidationError{Name: "inst_is_active", err: fmt.Errorf(`ent: validator failed for field "MstInst.inst_is_active": %w`, err)}
		}
	}
	if _, ok := mic.mutation.InstStatus(); !ok {
		return &ValidationError{Name: "inst_status", err: errors.New(`ent: missing required field "MstInst.inst_status"`)}
	}
	if _, ok := mic.mutation.InstTimeZone(); !ok {
		return &ValidationError{Name: "inst_time_zone", err: errors.New(`ent: missing required field "MstInst.inst_time_zone"`)}
	}
	if _, ok := mic.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`ent: missing required field "MstInst.customer_id"`)}
	}
	if _, ok := mic.mutation.InstfromCustID(); !ok {
		return &ValidationError{Name: "InstfromCust", err: errors.New(`ent: missing required edge "MstInst.InstfromCust"`)}
	}
	return nil
}

func (mic *MstInstCreate) sqlSave(ctx context.Context) (*MstInst, error) {
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mic *MstInstCreate) createSpec() (*MstInst, *sqlgraph.CreateSpec) {
	var (
		_node = &MstInst{config: mic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mstinst.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mstinst.FieldID,
			},
		}
	)
	if value, ok := mic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstinst.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstinst.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mic.mutation.InstCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstCode,
		})
		_node.InstCode = value
	}
	if value, ok := mic.mutation.InstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstName,
		})
		_node.InstName = value
	}
	if value, ok := mic.mutation.InstShortName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstShortName,
		})
		_node.InstShortName = value
	}
	if value, ok := mic.mutation.InstAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstAddress,
		})
		_node.InstAddress = value
	}
	if value, ok := mic.mutation.InstPlace(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstPlace,
		})
		_node.InstPlace = value
	}
	if value, ok := mic.mutation.InstState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstState,
		})
		_node.InstState = value
	}
	if value, ok := mic.mutation.InstPin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstPin,
		})
		_node.InstPin = value
	}
	if value, ok := mic.mutation.InstContactPerson(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstContactPerson,
		})
		_node.InstContactPerson = value
	}
	if value, ok := mic.mutation.InstPhone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstPhone,
		})
		_node.InstPhone = value
	}
	if value, ok := mic.mutation.InstEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstEmail,
		})
		_node.InstEmail = value
	}
	if value, ok := mic.mutation.InstMobile(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstMobile,
		})
		_node.InstMobile = value
	}
	if value, ok := mic.mutation.InstURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstURL,
		})
		_node.InstURL = value
	}
	if value, ok := mic.mutation.InstBanner1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstBanner1,
		})
		_node.InstBanner1 = value
	}
	if value, ok := mic.mutation.InstBanner2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstBanner2,
		})
		_node.InstBanner2 = value
	}
	if value, ok := mic.mutation.InstLogoURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstLogoURL,
		})
		_node.InstLogoURL = value
	}
	if value, ok := mic.mutation.InstIsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mstinst.FieldInstIsActive,
		})
		_node.InstIsActive = value
	}
	if value, ok := mic.mutation.InstStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mstinst.FieldInstStatus,
		})
		_node.InstStatus = value
	}
	if value, ok := mic.mutation.InstTimeZone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mstinst.FieldInstTimeZone,
		})
		_node.InstTimeZone = value
	}
	if nodes := mic.mutation.InstfromCustIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mstinst.InstfromCustTable,
			Columns: []string{mstinst.InstfromCustColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mstcustomer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CustomerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MstInstCreateBulk is the builder for creating many MstInst entities in bulk.
type MstInstCreateBulk struct {
	config
	builders []*MstInstCreate
}

// Save creates the MstInst entities in the database.
func (micb *MstInstCreateBulk) Save(ctx context.Context) ([]*MstInst, error) {
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MstInst, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MstInstMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MstInstCreateBulk) SaveX(ctx context.Context) []*MstInst {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (micb *MstInstCreateBulk) Exec(ctx context.Context) error {
	_, err := micb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (micb *MstInstCreateBulk) ExecX(ctx context.Context) {
	if err := micb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (mc *MstCustomerQuery) CollectFields(ctx context.Context, satisfies ...string) *MstCustomerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		mc = mc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return mc
}

func (mc *MstCustomerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *MstCustomerQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "Cust2Inst":
			mc = mc.WithCust2Inst(func(query *MstInstQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return mc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (mi *MstInstQuery) CollectFields(ctx context.Context, satisfies ...string) *MstInstQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		mi = mi.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return mi
}

func (mi *MstInstQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *MstInstQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "InstfromCust":
			mi = mi.WithInstfromCust(func(query *MstCustomerQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return mi
}

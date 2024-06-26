// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/hedwig-phan/assignment-3/ent/passwordresetrequest"
	"gitlab.com/hedwig-phan/assignment-3/ent/predicate"
)

// PasswordResetRequestDelete is the builder for deleting a PasswordResetRequest entity.
type PasswordResetRequestDelete struct {
	config
	hooks    []Hook
	mutation *PasswordResetRequestMutation
}

// Where appends a list predicates to the PasswordResetRequestDelete builder.
func (prrd *PasswordResetRequestDelete) Where(ps ...predicate.PasswordResetRequest) *PasswordResetRequestDelete {
	prrd.mutation.Where(ps...)
	return prrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prrd *PasswordResetRequestDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, prrd.sqlExec, prrd.mutation, prrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prrd *PasswordResetRequestDelete) ExecX(ctx context.Context) int {
	n, err := prrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prrd *PasswordResetRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(passwordresetrequest.Table, sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt))
	if ps := prrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prrd.mutation.done = true
	return affected, err
}

// PasswordResetRequestDeleteOne is the builder for deleting a single PasswordResetRequest entity.
type PasswordResetRequestDeleteOne struct {
	prrd *PasswordResetRequestDelete
}

// Where appends a list predicates to the PasswordResetRequestDelete builder.
func (prrdo *PasswordResetRequestDeleteOne) Where(ps ...predicate.PasswordResetRequest) *PasswordResetRequestDeleteOne {
	prrdo.prrd.mutation.Where(ps...)
	return prrdo
}

// Exec executes the deletion query.
func (prrdo *PasswordResetRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := prrdo.prrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{passwordresetrequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prrdo *PasswordResetRequestDeleteOne) ExecX(ctx context.Context) {
	if err := prrdo.Exec(ctx); err != nil {
		panic(err)
	}
}

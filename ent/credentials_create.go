// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/hedwig-phan/assignment-3/ent/credentials"
	"gitlab.com/hedwig-phan/assignment-3/ent/user"
)

// CredentialsCreate is the builder for creating a Credentials entity.
type CredentialsCreate struct {
	config
	mutation *CredentialsMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (cc *CredentialsCreate) SetUserID(i int) *CredentialsCreate {
	cc.mutation.SetUserID(i)
	return cc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cc *CredentialsCreate) SetNillableUserID(i *int) *CredentialsCreate {
	if i != nil {
		cc.SetUserID(*i)
	}
	return cc
}

// SetHashedPassword sets the "hashed_password" field.
func (cc *CredentialsCreate) SetHashedPassword(s string) *CredentialsCreate {
	cc.mutation.SetHashedPassword(s)
	return cc
}

// SetLastLogin sets the "last_login" field.
func (cc *CredentialsCreate) SetLastLogin(t time.Time) *CredentialsCreate {
	cc.mutation.SetLastLogin(t)
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CredentialsCreate) SetUser(u *User) *CredentialsCreate {
	return cc.SetUserID(u.ID)
}

// Mutation returns the CredentialsMutation object of the builder.
func (cc *CredentialsCreate) Mutation() *CredentialsMutation {
	return cc.mutation
}

// Save creates the Credentials in the database.
func (cc *CredentialsCreate) Save(ctx context.Context) (*Credentials, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CredentialsCreate) SaveX(ctx context.Context) *Credentials {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CredentialsCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CredentialsCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CredentialsCreate) check() error {
	if _, ok := cc.mutation.HashedPassword(); !ok {
		return &ValidationError{Name: "hashed_password", err: errors.New(`ent: missing required field "Credentials.hashed_password"`)}
	}
	if _, ok := cc.mutation.LastLogin(); !ok {
		return &ValidationError{Name: "last_login", err: errors.New(`ent: missing required field "Credentials.last_login"`)}
	}
	return nil
}

func (cc *CredentialsCreate) sqlSave(ctx context.Context) (*Credentials, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CredentialsCreate) createSpec() (*Credentials, *sqlgraph.CreateSpec) {
	var (
		_node = &Credentials{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(credentials.Table, sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.HashedPassword(); ok {
		_spec.SetField(credentials.FieldHashedPassword, field.TypeString, value)
		_node.HashedPassword = value
	}
	if value, ok := cc.mutation.LastLogin(); ok {
		_spec.SetField(credentials.FieldLastLogin, field.TypeTime, value)
		_node.LastLogin = value
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credentials.UserTable,
			Columns: []string{credentials.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CredentialsCreateBulk is the builder for creating many Credentials entities in bulk.
type CredentialsCreateBulk struct {
	config
	err      error
	builders []*CredentialsCreate
}

// Save creates the Credentials entities in the database.
func (ccb *CredentialsCreateBulk) Save(ctx context.Context) ([]*Credentials, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Credentials, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CredentialsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CredentialsCreateBulk) SaveX(ctx context.Context) []*Credentials {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CredentialsCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CredentialsCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

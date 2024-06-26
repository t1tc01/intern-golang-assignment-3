// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/hedwig-phan/assignment-3/ent/geometry"
	"gitlab.com/hedwig-phan/assignment-3/ent/location"
)

// LocationCreate is the builder for creating a Location entity.
type LocationCreate struct {
	config
	mutation *LocationMutation
	hooks    []Hook
}

// SetPlace sets the "place" field.
func (lc *LocationCreate) SetPlace(s string) *LocationCreate {
	lc.mutation.SetPlace(s)
	return lc
}

// SetNillablePlace sets the "place" field if the given value is not nil.
func (lc *LocationCreate) SetNillablePlace(s *string) *LocationCreate {
	if s != nil {
		lc.SetPlace(*s)
	}
	return lc
}

// SetAddress sets the "address" field.
func (lc *LocationCreate) SetAddress(s string) *LocationCreate {
	lc.mutation.SetAddress(s)
	return lc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (lc *LocationCreate) SetNillableAddress(s *string) *LocationCreate {
	if s != nil {
		lc.SetAddress(*s)
	}
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LocationCreate) SetCreatedAt(t time.Time) *LocationCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LocationCreate) SetUpdatedAt(t time.Time) *LocationCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetDeletedAt sets the "deleted_at" field.
func (lc *LocationCreate) SetDeletedAt(t time.Time) *LocationCreate {
	lc.mutation.SetDeletedAt(t)
	return lc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lc *LocationCreate) SetNillableDeletedAt(t *time.Time) *LocationCreate {
	if t != nil {
		lc.SetDeletedAt(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LocationCreate) SetID(i int) *LocationCreate {
	lc.mutation.SetID(i)
	return lc
}

// AddGeometryIDs adds the "geometries" edge to the Geometry entity by IDs.
func (lc *LocationCreate) AddGeometryIDs(ids ...int) *LocationCreate {
	lc.mutation.AddGeometryIDs(ids...)
	return lc
}

// AddGeometries adds the "geometries" edges to the Geometry entity.
func (lc *LocationCreate) AddGeometries(g ...*Geometry) *LocationCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return lc.AddGeometryIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lc *LocationCreate) Mutation() *LocationMutation {
	return lc.mutation
}

// Save creates the Location in the database.
func (lc *LocationCreate) Save(ctx context.Context) (*Location, error) {
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LocationCreate) SaveX(ctx context.Context) *Location {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LocationCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LocationCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LocationCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Location.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Location.updated_at"`)}
	}
	return nil
}

func (lc *LocationCreate) sqlSave(ctx context.Context) (*Location, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LocationCreate) createSpec() (*Location, *sqlgraph.CreateSpec) {
	var (
		_node = &Location{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(location.Table, sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Place(); ok {
		_spec.SetField(location.FieldPlace, field.TypeString, value)
		_node.Place = value
	}
	if value, ok := lc.mutation.Address(); ok {
		_spec.SetField(location.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(location.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.DeletedAt(); ok {
		_spec.SetField(location.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := lc.mutation.GeometriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.GeometriesTable,
			Columns: []string{location.GeometriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LocationCreateBulk is the builder for creating many Location entities in bulk.
type LocationCreateBulk struct {
	config
	err      error
	builders []*LocationCreate
}

// Save creates the Location entities in the database.
func (lcb *LocationCreateBulk) Save(ctx context.Context) ([]*Location, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Location, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LocationCreateBulk) SaveX(ctx context.Context) []*Location {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LocationCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LocationCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/hedwig-phan/assignment-3/ent/earthquake"
	"gitlab.com/hedwig-phan/assignment-3/ent/predicate"
	"gitlab.com/hedwig-phan/assignment-3/ent/source"
	"gitlab.com/hedwig-phan/assignment-3/ent/sourceearthquake"
)

// SourceEarthquakeUpdate is the builder for updating SourceEarthquake entities.
type SourceEarthquakeUpdate struct {
	config
	hooks    []Hook
	mutation *SourceEarthquakeMutation
}

// Where appends a list predicates to the SourceEarthquakeUpdate builder.
func (seu *SourceEarthquakeUpdate) Where(ps ...predicate.SourceEarthquake) *SourceEarthquakeUpdate {
	seu.mutation.Where(ps...)
	return seu
}

// SetSID sets the "s_id" field.
func (seu *SourceEarthquakeUpdate) SetSID(i int) *SourceEarthquakeUpdate {
	seu.mutation.SetSID(i)
	return seu
}

// SetNillableSID sets the "s_id" field if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableSID(i *int) *SourceEarthquakeUpdate {
	if i != nil {
		seu.SetSID(*i)
	}
	return seu
}

// ClearSID clears the value of the "s_id" field.
func (seu *SourceEarthquakeUpdate) ClearSID() *SourceEarthquakeUpdate {
	seu.mutation.ClearSID()
	return seu
}

// SetEqID sets the "eq_id" field.
func (seu *SourceEarthquakeUpdate) SetEqID(i int) *SourceEarthquakeUpdate {
	seu.mutation.SetEqID(i)
	return seu
}

// SetNillableEqID sets the "eq_id" field if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableEqID(i *int) *SourceEarthquakeUpdate {
	if i != nil {
		seu.SetEqID(*i)
	}
	return seu
}

// ClearEqID clears the value of the "eq_id" field.
func (seu *SourceEarthquakeUpdate) ClearEqID() *SourceEarthquakeUpdate {
	seu.mutation.ClearEqID()
	return seu
}

// SetCreatedAt sets the "created_at" field.
func (seu *SourceEarthquakeUpdate) SetCreatedAt(t time.Time) *SourceEarthquakeUpdate {
	seu.mutation.SetCreatedAt(t)
	return seu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableCreatedAt(t *time.Time) *SourceEarthquakeUpdate {
	if t != nil {
		seu.SetCreatedAt(*t)
	}
	return seu
}

// SetUpdatedAt sets the "updated_at" field.
func (seu *SourceEarthquakeUpdate) SetUpdatedAt(t time.Time) *SourceEarthquakeUpdate {
	seu.mutation.SetUpdatedAt(t)
	return seu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableUpdatedAt(t *time.Time) *SourceEarthquakeUpdate {
	if t != nil {
		seu.SetUpdatedAt(*t)
	}
	return seu
}

// SetDeletedAt sets the "deleted_at" field.
func (seu *SourceEarthquakeUpdate) SetDeletedAt(t time.Time) *SourceEarthquakeUpdate {
	seu.mutation.SetDeletedAt(t)
	return seu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableDeletedAt(t *time.Time) *SourceEarthquakeUpdate {
	if t != nil {
		seu.SetDeletedAt(*t)
	}
	return seu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (seu *SourceEarthquakeUpdate) ClearDeletedAt() *SourceEarthquakeUpdate {
	seu.mutation.ClearDeletedAt()
	return seu
}

// SetEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID.
func (seu *SourceEarthquakeUpdate) SetEarthquakeID(id int) *SourceEarthquakeUpdate {
	seu.mutation.SetEarthquakeID(id)
	return seu
}

// SetNillableEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableEarthquakeID(id *int) *SourceEarthquakeUpdate {
	if id != nil {
		seu = seu.SetEarthquakeID(*id)
	}
	return seu
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (seu *SourceEarthquakeUpdate) SetEarthquake(e *Earthquake) *SourceEarthquakeUpdate {
	return seu.SetEarthquakeID(e.ID)
}

// SetSourceID sets the "source" edge to the Source entity by ID.
func (seu *SourceEarthquakeUpdate) SetSourceID(id int) *SourceEarthquakeUpdate {
	seu.mutation.SetSourceID(id)
	return seu
}

// SetNillableSourceID sets the "source" edge to the Source entity by ID if the given value is not nil.
func (seu *SourceEarthquakeUpdate) SetNillableSourceID(id *int) *SourceEarthquakeUpdate {
	if id != nil {
		seu = seu.SetSourceID(*id)
	}
	return seu
}

// SetSource sets the "source" edge to the Source entity.
func (seu *SourceEarthquakeUpdate) SetSource(s *Source) *SourceEarthquakeUpdate {
	return seu.SetSourceID(s.ID)
}

// Mutation returns the SourceEarthquakeMutation object of the builder.
func (seu *SourceEarthquakeUpdate) Mutation() *SourceEarthquakeMutation {
	return seu.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (seu *SourceEarthquakeUpdate) ClearEarthquake() *SourceEarthquakeUpdate {
	seu.mutation.ClearEarthquake()
	return seu
}

// ClearSource clears the "source" edge to the Source entity.
func (seu *SourceEarthquakeUpdate) ClearSource() *SourceEarthquakeUpdate {
	seu.mutation.ClearSource()
	return seu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seu *SourceEarthquakeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, seu.sqlSave, seu.mutation, seu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seu *SourceEarthquakeUpdate) SaveX(ctx context.Context) int {
	affected, err := seu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seu *SourceEarthquakeUpdate) Exec(ctx context.Context) error {
	_, err := seu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seu *SourceEarthquakeUpdate) ExecX(ctx context.Context) {
	if err := seu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seu *SourceEarthquakeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sourceearthquake.Table, sourceearthquake.Columns, sqlgraph.NewFieldSpec(sourceearthquake.FieldID, field.TypeInt))
	if ps := seu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seu.mutation.CreatedAt(); ok {
		_spec.SetField(sourceearthquake.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := seu.mutation.UpdatedAt(); ok {
		_spec.SetField(sourceearthquake.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := seu.mutation.DeletedAt(); ok {
		_spec.SetField(sourceearthquake.FieldDeletedAt, field.TypeTime, value)
	}
	if seu.mutation.DeletedAtCleared() {
		_spec.ClearField(sourceearthquake.FieldDeletedAt, field.TypeTime)
	}
	if seu.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.EarthquakeTable,
			Columns: []string{sourceearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.EarthquakeTable,
			Columns: []string{sourceearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seu.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.SourceTable,
			Columns: []string{sourceearthquake.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.SourceTable,
			Columns: []string{sourceearthquake.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sourceearthquake.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	seu.mutation.done = true
	return n, nil
}

// SourceEarthquakeUpdateOne is the builder for updating a single SourceEarthquake entity.
type SourceEarthquakeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SourceEarthquakeMutation
}

// SetSID sets the "s_id" field.
func (seuo *SourceEarthquakeUpdateOne) SetSID(i int) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetSID(i)
	return seuo
}

// SetNillableSID sets the "s_id" field if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableSID(i *int) *SourceEarthquakeUpdateOne {
	if i != nil {
		seuo.SetSID(*i)
	}
	return seuo
}

// ClearSID clears the value of the "s_id" field.
func (seuo *SourceEarthquakeUpdateOne) ClearSID() *SourceEarthquakeUpdateOne {
	seuo.mutation.ClearSID()
	return seuo
}

// SetEqID sets the "eq_id" field.
func (seuo *SourceEarthquakeUpdateOne) SetEqID(i int) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetEqID(i)
	return seuo
}

// SetNillableEqID sets the "eq_id" field if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableEqID(i *int) *SourceEarthquakeUpdateOne {
	if i != nil {
		seuo.SetEqID(*i)
	}
	return seuo
}

// ClearEqID clears the value of the "eq_id" field.
func (seuo *SourceEarthquakeUpdateOne) ClearEqID() *SourceEarthquakeUpdateOne {
	seuo.mutation.ClearEqID()
	return seuo
}

// SetCreatedAt sets the "created_at" field.
func (seuo *SourceEarthquakeUpdateOne) SetCreatedAt(t time.Time) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetCreatedAt(t)
	return seuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableCreatedAt(t *time.Time) *SourceEarthquakeUpdateOne {
	if t != nil {
		seuo.SetCreatedAt(*t)
	}
	return seuo
}

// SetUpdatedAt sets the "updated_at" field.
func (seuo *SourceEarthquakeUpdateOne) SetUpdatedAt(t time.Time) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetUpdatedAt(t)
	return seuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableUpdatedAt(t *time.Time) *SourceEarthquakeUpdateOne {
	if t != nil {
		seuo.SetUpdatedAt(*t)
	}
	return seuo
}

// SetDeletedAt sets the "deleted_at" field.
func (seuo *SourceEarthquakeUpdateOne) SetDeletedAt(t time.Time) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetDeletedAt(t)
	return seuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableDeletedAt(t *time.Time) *SourceEarthquakeUpdateOne {
	if t != nil {
		seuo.SetDeletedAt(*t)
	}
	return seuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (seuo *SourceEarthquakeUpdateOne) ClearDeletedAt() *SourceEarthquakeUpdateOne {
	seuo.mutation.ClearDeletedAt()
	return seuo
}

// SetEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID.
func (seuo *SourceEarthquakeUpdateOne) SetEarthquakeID(id int) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetEarthquakeID(id)
	return seuo
}

// SetNillableEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableEarthquakeID(id *int) *SourceEarthquakeUpdateOne {
	if id != nil {
		seuo = seuo.SetEarthquakeID(*id)
	}
	return seuo
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (seuo *SourceEarthquakeUpdateOne) SetEarthquake(e *Earthquake) *SourceEarthquakeUpdateOne {
	return seuo.SetEarthquakeID(e.ID)
}

// SetSourceID sets the "source" edge to the Source entity by ID.
func (seuo *SourceEarthquakeUpdateOne) SetSourceID(id int) *SourceEarthquakeUpdateOne {
	seuo.mutation.SetSourceID(id)
	return seuo
}

// SetNillableSourceID sets the "source" edge to the Source entity by ID if the given value is not nil.
func (seuo *SourceEarthquakeUpdateOne) SetNillableSourceID(id *int) *SourceEarthquakeUpdateOne {
	if id != nil {
		seuo = seuo.SetSourceID(*id)
	}
	return seuo
}

// SetSource sets the "source" edge to the Source entity.
func (seuo *SourceEarthquakeUpdateOne) SetSource(s *Source) *SourceEarthquakeUpdateOne {
	return seuo.SetSourceID(s.ID)
}

// Mutation returns the SourceEarthquakeMutation object of the builder.
func (seuo *SourceEarthquakeUpdateOne) Mutation() *SourceEarthquakeMutation {
	return seuo.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (seuo *SourceEarthquakeUpdateOne) ClearEarthquake() *SourceEarthquakeUpdateOne {
	seuo.mutation.ClearEarthquake()
	return seuo
}

// ClearSource clears the "source" edge to the Source entity.
func (seuo *SourceEarthquakeUpdateOne) ClearSource() *SourceEarthquakeUpdateOne {
	seuo.mutation.ClearSource()
	return seuo
}

// Where appends a list predicates to the SourceEarthquakeUpdate builder.
func (seuo *SourceEarthquakeUpdateOne) Where(ps ...predicate.SourceEarthquake) *SourceEarthquakeUpdateOne {
	seuo.mutation.Where(ps...)
	return seuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seuo *SourceEarthquakeUpdateOne) Select(field string, fields ...string) *SourceEarthquakeUpdateOne {
	seuo.fields = append([]string{field}, fields...)
	return seuo
}

// Save executes the query and returns the updated SourceEarthquake entity.
func (seuo *SourceEarthquakeUpdateOne) Save(ctx context.Context) (*SourceEarthquake, error) {
	return withHooks(ctx, seuo.sqlSave, seuo.mutation, seuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seuo *SourceEarthquakeUpdateOne) SaveX(ctx context.Context) *SourceEarthquake {
	node, err := seuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seuo *SourceEarthquakeUpdateOne) Exec(ctx context.Context) error {
	_, err := seuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seuo *SourceEarthquakeUpdateOne) ExecX(ctx context.Context) {
	if err := seuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seuo *SourceEarthquakeUpdateOne) sqlSave(ctx context.Context) (_node *SourceEarthquake, err error) {
	_spec := sqlgraph.NewUpdateSpec(sourceearthquake.Table, sourceearthquake.Columns, sqlgraph.NewFieldSpec(sourceearthquake.FieldID, field.TypeInt))
	id, ok := seuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SourceEarthquake.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sourceearthquake.FieldID)
		for _, f := range fields {
			if !sourceearthquake.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sourceearthquake.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seuo.mutation.CreatedAt(); ok {
		_spec.SetField(sourceearthquake.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := seuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sourceearthquake.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := seuo.mutation.DeletedAt(); ok {
		_spec.SetField(sourceearthquake.FieldDeletedAt, field.TypeTime, value)
	}
	if seuo.mutation.DeletedAtCleared() {
		_spec.ClearField(sourceearthquake.FieldDeletedAt, field.TypeTime)
	}
	if seuo.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.EarthquakeTable,
			Columns: []string{sourceearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.EarthquakeTable,
			Columns: []string{sourceearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seuo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.SourceTable,
			Columns: []string{sourceearthquake.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sourceearthquake.SourceTable,
			Columns: []string{sourceearthquake.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SourceEarthquake{config: seuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sourceearthquake.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	seuo.mutation.done = true
	return _node, nil
}
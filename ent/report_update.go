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
	"gitlab.com/hedwig-phan/assignment-3/ent/report"
)

// ReportUpdate is the builder for updating Report entities.
type ReportUpdate struct {
	config
	hooks    []Hook
	mutation *ReportMutation
}

// Where appends a list predicates to the ReportUpdate builder.
func (ru *ReportUpdate) Where(ps ...predicate.Report) *ReportUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetFelt sets the "felt" field.
func (ru *ReportUpdate) SetFelt(i int32) *ReportUpdate {
	ru.mutation.ResetFelt()
	ru.mutation.SetFelt(i)
	return ru
}

// SetNillableFelt sets the "felt" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableFelt(i *int32) *ReportUpdate {
	if i != nil {
		ru.SetFelt(*i)
	}
	return ru
}

// AddFelt adds i to the "felt" field.
func (ru *ReportUpdate) AddFelt(i int32) *ReportUpdate {
	ru.mutation.AddFelt(i)
	return ru
}

// ClearFelt clears the value of the "felt" field.
func (ru *ReportUpdate) ClearFelt() *ReportUpdate {
	ru.mutation.ClearFelt()
	return ru
}

// SetCdi sets the "cdi" field.
func (ru *ReportUpdate) SetCdi(f float64) *ReportUpdate {
	ru.mutation.ResetCdi()
	ru.mutation.SetCdi(f)
	return ru
}

// SetNillableCdi sets the "cdi" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableCdi(f *float64) *ReportUpdate {
	if f != nil {
		ru.SetCdi(*f)
	}
	return ru
}

// AddCdi adds f to the "cdi" field.
func (ru *ReportUpdate) AddCdi(f float64) *ReportUpdate {
	ru.mutation.AddCdi(f)
	return ru
}

// ClearCdi clears the value of the "cdi" field.
func (ru *ReportUpdate) ClearCdi() *ReportUpdate {
	ru.mutation.ClearCdi()
	return ru
}

// SetMmi sets the "mmi" field.
func (ru *ReportUpdate) SetMmi(f float64) *ReportUpdate {
	ru.mutation.ResetMmi()
	ru.mutation.SetMmi(f)
	return ru
}

// SetNillableMmi sets the "mmi" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableMmi(f *float64) *ReportUpdate {
	if f != nil {
		ru.SetMmi(*f)
	}
	return ru
}

// AddMmi adds f to the "mmi" field.
func (ru *ReportUpdate) AddMmi(f float64) *ReportUpdate {
	ru.mutation.AddMmi(f)
	return ru
}

// ClearMmi clears the value of the "mmi" field.
func (ru *ReportUpdate) ClearMmi() *ReportUpdate {
	ru.mutation.ClearMmi()
	return ru
}

// SetAlert sets the "alert" field.
func (ru *ReportUpdate) SetAlert(s string) *ReportUpdate {
	ru.mutation.SetAlert(s)
	return ru
}

// SetNillableAlert sets the "alert" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableAlert(s *string) *ReportUpdate {
	if s != nil {
		ru.SetAlert(*s)
	}
	return ru
}

// ClearAlert clears the value of the "alert" field.
func (ru *ReportUpdate) ClearAlert() *ReportUpdate {
	ru.mutation.ClearAlert()
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *ReportUpdate) SetCreatedAt(t time.Time) *ReportUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableCreatedAt(t *time.Time) *ReportUpdate {
	if t != nil {
		ru.SetCreatedAt(*t)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReportUpdate) SetUpdatedAt(t time.Time) *ReportUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableUpdatedAt(t *time.Time) *ReportUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *ReportUpdate) SetDeletedAt(t time.Time) *ReportUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *ReportUpdate) SetNillableDeletedAt(t *time.Time) *ReportUpdate {
	if t != nil {
		ru.SetDeletedAt(*t)
	}
	return ru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ru *ReportUpdate) ClearDeletedAt() *ReportUpdate {
	ru.mutation.ClearDeletedAt()
	return ru
}

// AddEarthquakeIDs adds the "earthquakes" edge to the Earthquake entity by IDs.
func (ru *ReportUpdate) AddEarthquakeIDs(ids ...int) *ReportUpdate {
	ru.mutation.AddEarthquakeIDs(ids...)
	return ru
}

// AddEarthquakes adds the "earthquakes" edges to the Earthquake entity.
func (ru *ReportUpdate) AddEarthquakes(e ...*Earthquake) *ReportUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.AddEarthquakeIDs(ids...)
}

// Mutation returns the ReportMutation object of the builder.
func (ru *ReportUpdate) Mutation() *ReportMutation {
	return ru.mutation
}

// ClearEarthquakes clears all "earthquakes" edges to the Earthquake entity.
func (ru *ReportUpdate) ClearEarthquakes() *ReportUpdate {
	ru.mutation.ClearEarthquakes()
	return ru
}

// RemoveEarthquakeIDs removes the "earthquakes" edge to Earthquake entities by IDs.
func (ru *ReportUpdate) RemoveEarthquakeIDs(ids ...int) *ReportUpdate {
	ru.mutation.RemoveEarthquakeIDs(ids...)
	return ru
}

// RemoveEarthquakes removes "earthquakes" edges to Earthquake entities.
func (ru *ReportUpdate) RemoveEarthquakes(e ...*Earthquake) *ReportUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.RemoveEarthquakeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReportUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReportUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReportUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReportUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(report.Table, report.Columns, sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Felt(); ok {
		_spec.SetField(report.FieldFelt, field.TypeInt32, value)
	}
	if value, ok := ru.mutation.AddedFelt(); ok {
		_spec.AddField(report.FieldFelt, field.TypeInt32, value)
	}
	if ru.mutation.FeltCleared() {
		_spec.ClearField(report.FieldFelt, field.TypeInt32)
	}
	if value, ok := ru.mutation.Cdi(); ok {
		_spec.SetField(report.FieldCdi, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedCdi(); ok {
		_spec.AddField(report.FieldCdi, field.TypeFloat64, value)
	}
	if ru.mutation.CdiCleared() {
		_spec.ClearField(report.FieldCdi, field.TypeFloat64)
	}
	if value, ok := ru.mutation.Mmi(); ok {
		_spec.SetField(report.FieldMmi, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedMmi(); ok {
		_spec.AddField(report.FieldMmi, field.TypeFloat64, value)
	}
	if ru.mutation.MmiCleared() {
		_spec.ClearField(report.FieldMmi, field.TypeFloat64)
	}
	if value, ok := ru.mutation.Alert(); ok {
		_spec.SetField(report.FieldAlert, field.TypeString, value)
	}
	if ru.mutation.AlertCleared() {
		_spec.ClearField(report.FieldAlert, field.TypeString)
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.SetField(report.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(report.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(report.FieldDeletedAt, field.TypeTime, value)
	}
	if ru.mutation.DeletedAtCleared() {
		_spec.ClearField(report.FieldDeletedAt, field.TypeTime)
	}
	if ru.mutation.EarthquakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedEarthquakesIDs(); len(nodes) > 0 && !ru.mutation.EarthquakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.EarthquakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{report.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReportUpdateOne is the builder for updating a single Report entity.
type ReportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReportMutation
}

// SetFelt sets the "felt" field.
func (ruo *ReportUpdateOne) SetFelt(i int32) *ReportUpdateOne {
	ruo.mutation.ResetFelt()
	ruo.mutation.SetFelt(i)
	return ruo
}

// SetNillableFelt sets the "felt" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableFelt(i *int32) *ReportUpdateOne {
	if i != nil {
		ruo.SetFelt(*i)
	}
	return ruo
}

// AddFelt adds i to the "felt" field.
func (ruo *ReportUpdateOne) AddFelt(i int32) *ReportUpdateOne {
	ruo.mutation.AddFelt(i)
	return ruo
}

// ClearFelt clears the value of the "felt" field.
func (ruo *ReportUpdateOne) ClearFelt() *ReportUpdateOne {
	ruo.mutation.ClearFelt()
	return ruo
}

// SetCdi sets the "cdi" field.
func (ruo *ReportUpdateOne) SetCdi(f float64) *ReportUpdateOne {
	ruo.mutation.ResetCdi()
	ruo.mutation.SetCdi(f)
	return ruo
}

// SetNillableCdi sets the "cdi" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableCdi(f *float64) *ReportUpdateOne {
	if f != nil {
		ruo.SetCdi(*f)
	}
	return ruo
}

// AddCdi adds f to the "cdi" field.
func (ruo *ReportUpdateOne) AddCdi(f float64) *ReportUpdateOne {
	ruo.mutation.AddCdi(f)
	return ruo
}

// ClearCdi clears the value of the "cdi" field.
func (ruo *ReportUpdateOne) ClearCdi() *ReportUpdateOne {
	ruo.mutation.ClearCdi()
	return ruo
}

// SetMmi sets the "mmi" field.
func (ruo *ReportUpdateOne) SetMmi(f float64) *ReportUpdateOne {
	ruo.mutation.ResetMmi()
	ruo.mutation.SetMmi(f)
	return ruo
}

// SetNillableMmi sets the "mmi" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableMmi(f *float64) *ReportUpdateOne {
	if f != nil {
		ruo.SetMmi(*f)
	}
	return ruo
}

// AddMmi adds f to the "mmi" field.
func (ruo *ReportUpdateOne) AddMmi(f float64) *ReportUpdateOne {
	ruo.mutation.AddMmi(f)
	return ruo
}

// ClearMmi clears the value of the "mmi" field.
func (ruo *ReportUpdateOne) ClearMmi() *ReportUpdateOne {
	ruo.mutation.ClearMmi()
	return ruo
}

// SetAlert sets the "alert" field.
func (ruo *ReportUpdateOne) SetAlert(s string) *ReportUpdateOne {
	ruo.mutation.SetAlert(s)
	return ruo
}

// SetNillableAlert sets the "alert" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableAlert(s *string) *ReportUpdateOne {
	if s != nil {
		ruo.SetAlert(*s)
	}
	return ruo
}

// ClearAlert clears the value of the "alert" field.
func (ruo *ReportUpdateOne) ClearAlert() *ReportUpdateOne {
	ruo.mutation.ClearAlert()
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *ReportUpdateOne) SetCreatedAt(t time.Time) *ReportUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableCreatedAt(t *time.Time) *ReportUpdateOne {
	if t != nil {
		ruo.SetCreatedAt(*t)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReportUpdateOne) SetUpdatedAt(t time.Time) *ReportUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableUpdatedAt(t *time.Time) *ReportUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *ReportUpdateOne) SetDeletedAt(t time.Time) *ReportUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *ReportUpdateOne) SetNillableDeletedAt(t *time.Time) *ReportUpdateOne {
	if t != nil {
		ruo.SetDeletedAt(*t)
	}
	return ruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruo *ReportUpdateOne) ClearDeletedAt() *ReportUpdateOne {
	ruo.mutation.ClearDeletedAt()
	return ruo
}

// AddEarthquakeIDs adds the "earthquakes" edge to the Earthquake entity by IDs.
func (ruo *ReportUpdateOne) AddEarthquakeIDs(ids ...int) *ReportUpdateOne {
	ruo.mutation.AddEarthquakeIDs(ids...)
	return ruo
}

// AddEarthquakes adds the "earthquakes" edges to the Earthquake entity.
func (ruo *ReportUpdateOne) AddEarthquakes(e ...*Earthquake) *ReportUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.AddEarthquakeIDs(ids...)
}

// Mutation returns the ReportMutation object of the builder.
func (ruo *ReportUpdateOne) Mutation() *ReportMutation {
	return ruo.mutation
}

// ClearEarthquakes clears all "earthquakes" edges to the Earthquake entity.
func (ruo *ReportUpdateOne) ClearEarthquakes() *ReportUpdateOne {
	ruo.mutation.ClearEarthquakes()
	return ruo
}

// RemoveEarthquakeIDs removes the "earthquakes" edge to Earthquake entities by IDs.
func (ruo *ReportUpdateOne) RemoveEarthquakeIDs(ids ...int) *ReportUpdateOne {
	ruo.mutation.RemoveEarthquakeIDs(ids...)
	return ruo
}

// RemoveEarthquakes removes "earthquakes" edges to Earthquake entities.
func (ruo *ReportUpdateOne) RemoveEarthquakes(e ...*Earthquake) *ReportUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.RemoveEarthquakeIDs(ids...)
}

// Where appends a list predicates to the ReportUpdate builder.
func (ruo *ReportUpdateOne) Where(ps ...predicate.Report) *ReportUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReportUpdateOne) Select(field string, fields ...string) *ReportUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Report entity.
func (ruo *ReportUpdateOne) Save(ctx context.Context) (*Report, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReportUpdateOne) SaveX(ctx context.Context) *Report {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReportUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReportUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReportUpdateOne) sqlSave(ctx context.Context) (_node *Report, err error) {
	_spec := sqlgraph.NewUpdateSpec(report.Table, report.Columns, sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Report.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, report.FieldID)
		for _, f := range fields {
			if !report.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != report.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Felt(); ok {
		_spec.SetField(report.FieldFelt, field.TypeInt32, value)
	}
	if value, ok := ruo.mutation.AddedFelt(); ok {
		_spec.AddField(report.FieldFelt, field.TypeInt32, value)
	}
	if ruo.mutation.FeltCleared() {
		_spec.ClearField(report.FieldFelt, field.TypeInt32)
	}
	if value, ok := ruo.mutation.Cdi(); ok {
		_spec.SetField(report.FieldCdi, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedCdi(); ok {
		_spec.AddField(report.FieldCdi, field.TypeFloat64, value)
	}
	if ruo.mutation.CdiCleared() {
		_spec.ClearField(report.FieldCdi, field.TypeFloat64)
	}
	if value, ok := ruo.mutation.Mmi(); ok {
		_spec.SetField(report.FieldMmi, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedMmi(); ok {
		_spec.AddField(report.FieldMmi, field.TypeFloat64, value)
	}
	if ruo.mutation.MmiCleared() {
		_spec.ClearField(report.FieldMmi, field.TypeFloat64)
	}
	if value, ok := ruo.mutation.Alert(); ok {
		_spec.SetField(report.FieldAlert, field.TypeString, value)
	}
	if ruo.mutation.AlertCleared() {
		_spec.ClearField(report.FieldAlert, field.TypeString)
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.SetField(report.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(report.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(report.FieldDeletedAt, field.TypeTime, value)
	}
	if ruo.mutation.DeletedAtCleared() {
		_spec.ClearField(report.FieldDeletedAt, field.TypeTime)
	}
	if ruo.mutation.EarthquakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedEarthquakesIDs(); len(nodes) > 0 && !ruo.mutation.EarthquakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.EarthquakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
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
	_node = &Report{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{report.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}

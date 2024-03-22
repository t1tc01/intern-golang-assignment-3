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
	"gitlab.com/hedwig-phan/assignment-3/ent/credentials"
	"gitlab.com/hedwig-phan/assignment-3/ent/passwordresetrequest"
	"gitlab.com/hedwig-phan/assignment-3/ent/predicate"
	"gitlab.com/hedwig-phan/assignment-3/ent/session"
	"gitlab.com/hedwig-phan/assignment-3/ent/token"
	"gitlab.com/hedwig-phan/assignment-3/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(i int) *UserUpdate {
	uu.mutation.ResetRole()
	uu.mutation.SetRole(i)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(i *int) *UserUpdate {
	if i != nil {
		uu.SetRole(*i)
	}
	return uu
}

// AddRole adds i to the "role" field.
func (uu *UserUpdate) AddRole(i int) *UserUpdate {
	uu.mutation.AddRole(i)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// AddSessionIDs adds the "session" edge to the Session entity by IDs.
func (uu *UserUpdate) AddSessionIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSessionIDs(ids...)
	return uu
}

// AddSession adds the "session" edges to the Session entity.
func (uu *UserUpdate) AddSession(s ...*Session) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSessionIDs(ids...)
}

// AddPasswordResetRequestIDs adds the "password_reset_request" edge to the PasswordResetRequest entity by IDs.
func (uu *UserUpdate) AddPasswordResetRequestIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPasswordResetRequestIDs(ids...)
	return uu
}

// AddPasswordResetRequest adds the "password_reset_request" edges to the PasswordResetRequest entity.
func (uu *UserUpdate) AddPasswordResetRequest(p ...*PasswordResetRequest) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPasswordResetRequestIDs(ids...)
}

// AddCredentialIDs adds the "credentials" edge to the Credentials entity by IDs.
func (uu *UserUpdate) AddCredentialIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCredentialIDs(ids...)
	return uu
}

// AddCredentials adds the "credentials" edges to the Credentials entity.
func (uu *UserUpdate) AddCredentials(c ...*Credentials) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCredentialIDs(ids...)
}

// AddTokenIDs adds the "token" edge to the Token entity by IDs.
func (uu *UserUpdate) AddTokenIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTokenIDs(ids...)
	return uu
}

// AddToken adds the "token" edges to the Token entity.
func (uu *UserUpdate) AddToken(t ...*Token) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSession clears all "session" edges to the Session entity.
func (uu *UserUpdate) ClearSession() *UserUpdate {
	uu.mutation.ClearSession()
	return uu
}

// RemoveSessionIDs removes the "session" edge to Session entities by IDs.
func (uu *UserUpdate) RemoveSessionIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSessionIDs(ids...)
	return uu
}

// RemoveSession removes "session" edges to Session entities.
func (uu *UserUpdate) RemoveSession(s ...*Session) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSessionIDs(ids...)
}

// ClearPasswordResetRequest clears all "password_reset_request" edges to the PasswordResetRequest entity.
func (uu *UserUpdate) ClearPasswordResetRequest() *UserUpdate {
	uu.mutation.ClearPasswordResetRequest()
	return uu
}

// RemovePasswordResetRequestIDs removes the "password_reset_request" edge to PasswordResetRequest entities by IDs.
func (uu *UserUpdate) RemovePasswordResetRequestIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePasswordResetRequestIDs(ids...)
	return uu
}

// RemovePasswordResetRequest removes "password_reset_request" edges to PasswordResetRequest entities.
func (uu *UserUpdate) RemovePasswordResetRequest(p ...*PasswordResetRequest) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePasswordResetRequestIDs(ids...)
}

// ClearCredentials clears all "credentials" edges to the Credentials entity.
func (uu *UserUpdate) ClearCredentials() *UserUpdate {
	uu.mutation.ClearCredentials()
	return uu
}

// RemoveCredentialIDs removes the "credentials" edge to Credentials entities by IDs.
func (uu *UserUpdate) RemoveCredentialIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCredentialIDs(ids...)
	return uu
}

// RemoveCredentials removes "credentials" edges to Credentials entities.
func (uu *UserUpdate) RemoveCredentials(c ...*Credentials) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCredentialIDs(ids...)
}

// ClearToken clears all "token" edges to the Token entity.
func (uu *UserUpdate) ClearToken() *UserUpdate {
	uu.mutation.ClearToken()
	return uu
}

// RemoveTokenIDs removes the "token" edge to Token entities by IDs.
func (uu *UserUpdate) RemoveTokenIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTokenIDs(ids...)
	return uu
}

// RemoveToken removes "token" edges to Token entities.
func (uu *UserUpdate) RemoveToken(t ...*Token) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if uu.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionTable,
			Columns: []string{user.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSessionIDs(); len(nodes) > 0 && !uu.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionTable,
			Columns: []string{user.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionTable,
			Columns: []string{user.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PasswordResetRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasswordResetRequestTable,
			Columns: []string{user.PasswordResetRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPasswordResetRequestIDs(); len(nodes) > 0 && !uu.mutation.PasswordResetRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasswordResetRequestTable,
			Columns: []string{user.PasswordResetRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PasswordResetRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasswordResetRequestTable,
			Columns: []string{user.PasswordResetRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCredentialsIDs(); len(nodes) > 0 && !uu.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTokenIDs(); len(nodes) > 0 && !uu.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(i int) *UserUpdateOne {
	uuo.mutation.ResetRole()
	uuo.mutation.SetRole(i)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetRole(*i)
	}
	return uuo
}

// AddRole adds i to the "role" field.
func (uuo *UserUpdateOne) AddRole(i int) *UserUpdateOne {
	uuo.mutation.AddRole(i)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// AddSessionIDs adds the "session" edge to the Session entity by IDs.
func (uuo *UserUpdateOne) AddSessionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSessionIDs(ids...)
	return uuo
}

// AddSession adds the "session" edges to the Session entity.
func (uuo *UserUpdateOne) AddSession(s ...*Session) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSessionIDs(ids...)
}

// AddPasswordResetRequestIDs adds the "password_reset_request" edge to the PasswordResetRequest entity by IDs.
func (uuo *UserUpdateOne) AddPasswordResetRequestIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPasswordResetRequestIDs(ids...)
	return uuo
}

// AddPasswordResetRequest adds the "password_reset_request" edges to the PasswordResetRequest entity.
func (uuo *UserUpdateOne) AddPasswordResetRequest(p ...*PasswordResetRequest) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPasswordResetRequestIDs(ids...)
}

// AddCredentialIDs adds the "credentials" edge to the Credentials entity by IDs.
func (uuo *UserUpdateOne) AddCredentialIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCredentialIDs(ids...)
	return uuo
}

// AddCredentials adds the "credentials" edges to the Credentials entity.
func (uuo *UserUpdateOne) AddCredentials(c ...*Credentials) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCredentialIDs(ids...)
}

// AddTokenIDs adds the "token" edge to the Token entity by IDs.
func (uuo *UserUpdateOne) AddTokenIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTokenIDs(ids...)
	return uuo
}

// AddToken adds the "token" edges to the Token entity.
func (uuo *UserUpdateOne) AddToken(t ...*Token) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSession clears all "session" edges to the Session entity.
func (uuo *UserUpdateOne) ClearSession() *UserUpdateOne {
	uuo.mutation.ClearSession()
	return uuo
}

// RemoveSessionIDs removes the "session" edge to Session entities by IDs.
func (uuo *UserUpdateOne) RemoveSessionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSessionIDs(ids...)
	return uuo
}

// RemoveSession removes "session" edges to Session entities.
func (uuo *UserUpdateOne) RemoveSession(s ...*Session) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSessionIDs(ids...)
}

// ClearPasswordResetRequest clears all "password_reset_request" edges to the PasswordResetRequest entity.
func (uuo *UserUpdateOne) ClearPasswordResetRequest() *UserUpdateOne {
	uuo.mutation.ClearPasswordResetRequest()
	return uuo
}

// RemovePasswordResetRequestIDs removes the "password_reset_request" edge to PasswordResetRequest entities by IDs.
func (uuo *UserUpdateOne) RemovePasswordResetRequestIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePasswordResetRequestIDs(ids...)
	return uuo
}

// RemovePasswordResetRequest removes "password_reset_request" edges to PasswordResetRequest entities.
func (uuo *UserUpdateOne) RemovePasswordResetRequest(p ...*PasswordResetRequest) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePasswordResetRequestIDs(ids...)
}

// ClearCredentials clears all "credentials" edges to the Credentials entity.
func (uuo *UserUpdateOne) ClearCredentials() *UserUpdateOne {
	uuo.mutation.ClearCredentials()
	return uuo
}

// RemoveCredentialIDs removes the "credentials" edge to Credentials entities by IDs.
func (uuo *UserUpdateOne) RemoveCredentialIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCredentialIDs(ids...)
	return uuo
}

// RemoveCredentials removes "credentials" edges to Credentials entities.
func (uuo *UserUpdateOne) RemoveCredentials(c ...*Credentials) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCredentialIDs(ids...)
}

// ClearToken clears all "token" edges to the Token entity.
func (uuo *UserUpdateOne) ClearToken() *UserUpdateOne {
	uuo.mutation.ClearToken()
	return uuo
}

// RemoveTokenIDs removes the "token" edge to Token entities by IDs.
func (uuo *UserUpdateOne) RemoveTokenIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTokenIDs(ids...)
	return uuo
}

// RemoveToken removes "token" edges to Token entities.
func (uuo *UserUpdateOne) RemoveToken(t ...*Token) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTokenIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if uuo.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionTable,
			Columns: []string{user.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSessionIDs(); len(nodes) > 0 && !uuo.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionTable,
			Columns: []string{user.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionTable,
			Columns: []string{user.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PasswordResetRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasswordResetRequestTable,
			Columns: []string{user.PasswordResetRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPasswordResetRequestIDs(); len(nodes) > 0 && !uuo.mutation.PasswordResetRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasswordResetRequestTable,
			Columns: []string{user.PasswordResetRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PasswordResetRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PasswordResetRequestTable,
			Columns: []string{user.PasswordResetRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passwordresetrequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCredentialsIDs(); len(nodes) > 0 && !uuo.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credentials.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTokenIDs(); len(nodes) > 0 && !uuo.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}

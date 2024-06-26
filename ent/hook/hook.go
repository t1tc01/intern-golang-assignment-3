// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"gitlab.com/hedwig-phan/assignment-3/ent"
)

// The ApireqFunc type is an adapter to allow the use of ordinary
// function as Apireq mutator.
type ApireqFunc func(context.Context, *ent.ApireqMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ApireqFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ApireqMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ApireqMutation", m)
}

// The CredentialsFunc type is an adapter to allow the use of ordinary
// function as Credentials mutator.
type CredentialsFunc func(context.Context, *ent.CredentialsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CredentialsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CredentialsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CredentialsMutation", m)
}

// The EarthquakeFunc type is an adapter to allow the use of ordinary
// function as Earthquake mutator.
type EarthquakeFunc func(context.Context, *ent.EarthquakeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EarthquakeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EarthquakeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EarthquakeMutation", m)
}

// The FeatureTypeFunc type is an adapter to allow the use of ordinary
// function as FeatureType mutator.
type FeatureTypeFunc func(context.Context, *ent.FeatureTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeatureTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeatureTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeatureTypeMutation", m)
}

// The FtypeEarthquakeFunc type is an adapter to allow the use of ordinary
// function as FtypeEarthquake mutator.
type FtypeEarthquakeFunc func(context.Context, *ent.FtypeEarthquakeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FtypeEarthquakeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FtypeEarthquakeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FtypeEarthquakeMutation", m)
}

// The GeometryFunc type is an adapter to allow the use of ordinary
// function as Geometry mutator.
type GeometryFunc func(context.Context, *ent.GeometryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GeometryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GeometryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GeometryMutation", m)
}

// The LocationFunc type is an adapter to allow the use of ordinary
// function as Location mutator.
type LocationFunc func(context.Context, *ent.LocationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LocationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LocationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LocationMutation", m)
}

// The PasswordResetRequestFunc type is an adapter to allow the use of ordinary
// function as PasswordResetRequest mutator.
type PasswordResetRequestFunc func(context.Context, *ent.PasswordResetRequestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PasswordResetRequestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PasswordResetRequestMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PasswordResetRequestMutation", m)
}

// The ReportFunc type is an adapter to allow the use of ordinary
// function as Report mutator.
type ReportFunc func(context.Context, *ent.ReportMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReportFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReportMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReportMutation", m)
}

// The SessionFunc type is an adapter to allow the use of ordinary
// function as Session mutator.
type SessionFunc func(context.Context, *ent.SessionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SessionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SessionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SessionMutation", m)
}

// The SourceFunc type is an adapter to allow the use of ordinary
// function as Source mutator.
type SourceFunc func(context.Context, *ent.SourceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SourceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SourceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SourceMutation", m)
}

// The SourceEarthquakeFunc type is an adapter to allow the use of ordinary
// function as SourceEarthquake mutator.
type SourceEarthquakeFunc func(context.Context, *ent.SourceEarthquakeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SourceEarthquakeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SourceEarthquakeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SourceEarthquakeMutation", m)
}

// The TokenFunc type is an adapter to allow the use of ordinary
// function as Token mutator.
type TokenFunc func(context.Context, *ent.TokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TokenMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}

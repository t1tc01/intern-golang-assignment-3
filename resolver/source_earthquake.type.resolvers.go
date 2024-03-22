package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/graph"
)

// Earthquake is the resolver for the earthquake field.
func (r *sourceEarthquakeResolver) Earthquake(ctx context.Context, obj *ent.SourceEarthquake) (*ent.Earthquake, error) {
	panic(fmt.Errorf("not implemented: Earthquake - earthquake"))
}

// Source is the resolver for the source field.
func (r *sourceEarthquakeResolver) Source(ctx context.Context, obj *ent.SourceEarthquake) (*ent.Source, error) {
	panic(fmt.Errorf("not implemented: Source - source"))
}

// SourceEarthquake returns graph.SourceEarthquakeResolver implementation.
func (r *Resolver) SourceEarthquake() graph.SourceEarthquakeResolver {
	return &sourceEarthquakeResolver{r}
}

type sourceEarthquakeResolver struct{ *Resolver }
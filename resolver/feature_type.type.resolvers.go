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

// FtypeEarthquakes is the resolver for the ftypeEarthquakes field.
func (r *featureTypeResolver) FtypeEarthquakes(ctx context.Context, obj *ent.FeatureType) ([]*ent.FtypeEarthquake, error) {
	panic(fmt.Errorf("not implemented: FtypeEarthquakes - ftypeEarthquakes"))
}

// FeatureType returns graph.FeatureTypeResolver implementation.
func (r *Resolver) FeatureType() graph.FeatureTypeResolver { return &featureTypeResolver{r} }

type featureTypeResolver struct{ *Resolver }

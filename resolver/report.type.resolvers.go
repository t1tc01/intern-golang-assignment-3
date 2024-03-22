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

// Earthquakes is the resolver for the earthquakes field.
func (r *reportResolver) Earthquakes(ctx context.Context, obj *ent.Report) ([]*ent.Earthquake, error) {
	panic(fmt.Errorf("not implemented: Earthquakes - earthquakes"))
}

// Report returns graph.ReportResolver implementation.
func (r *Resolver) Report() graph.ReportResolver { return &reportResolver{r} }

type reportResolver struct{ *Resolver }

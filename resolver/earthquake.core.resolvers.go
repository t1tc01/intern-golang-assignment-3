package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/graph"
	"gitlab.com/hedwig-phan/assignment-3/internal/earthquake"
	"gitlab.com/hedwig-phan/assignment-3/middleware"
	"gitlab.com/hedwig-phan/assignment-3/model"
)

// CreateEarthquake is the resolver for the createEarthquake field.
func (r *mutationResolver) CreateEarthquake(ctx context.Context, input model.CreateEarthquake) (string, error) {
	panic(fmt.Errorf("not implemented: CreateEarthquake - createEarthquake"))
}

// Earthquakes is the resolver for the earthquakes field.
func (r *queryResolver) Earthquakes(ctx context.Context) ([]*ent.Earthquake, error) {

	//Authen
	user := middleware.ForContext(ctx)
	if user == nil {
		var tmp []*ent.Earthquake
		return tmp, fmt.Errorf("access denied")
	}

	return earthquake.QuerryAllEathquake(ctx)
}

// ListEarthquakes is the resolver for the listEarthquakes field.
func (r *queryResolver) ListEarthquakes(ctx context.Context, filter *model.EarthquakeFilter, pagination *model.PaginationInput) ([]*ent.Earthquake, error) {

	//
	user := middleware.ForContext(ctx)
	if user == nil {
		var tmp []*ent.Earthquake
		return tmp, fmt.Errorf("access denied")
	}

	limit := 1
	offset := 10

	var mag float64 = 0
	place := ""
	eq_type := ""

	if pagination != nil {
		if pagination.Limit != nil {
			limit = *pagination.Limit
		}
		if pagination.Offset != nil {
			offset = *pagination.Offset
		}
	}

	if filter != nil {
		if filter.Magnitude != nil {
			mag = *filter.Magnitude
		}
		if filter.Place != nil {
			place = *filter.Place
		}

		if filter.EarthquakeType != nil {
			eq_type = *filter.EarthquakeType
		}
	}

	return earthquake.QuerryEarthquakeMultiFiltered(ctx, limit, offset, mag, place, eq_type)
}

// FilterEarthquakesByID is the resolver for the filterEarthquakesByID field.
func (r *queryResolver) FilterEarthquakesByID(ctx context.Context, input int) ([]*ent.Earthquake, error) {

	//
	user := middleware.ForContext(ctx)
	if user == nil {
		var tmp []*ent.Earthquake
		return tmp, fmt.Errorf("access denied")
	}
	return earthquake.QuerryEarthquakeFilteredByID(ctx, input)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

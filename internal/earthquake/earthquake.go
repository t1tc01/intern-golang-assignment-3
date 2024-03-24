package earthquake

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/ent"
	"gitlab.com/hedwig-phan/assignment-3/ent/earthquake"
	"gitlab.com/hedwig-phan/assignment-3/ent/location"
)

// querry
func QuerryAllEathquake(c context.Context) ([]*ent.Earthquake, error) {
	eqs, err := db.Client.Earthquake.Query().All(c)

	if err != nil {
		return nil, err
	}

	return eqs, nil
}

func QuerryEarthquakeMultiFiltered(c context.Context, limit int, offset int,
	magnitude float64, place string, earthquakeType string) ([]*ent.Earthquake, error) {

	query := db.Client.Earthquake.Query()

	if magnitude > 0 {
		query.Where(earthquake.MagGT(magnitude))
	}
	if place != "" {
		query.QueryGeometry().QueryLocation().Where(location.Place(place))
	}
	if earthquakeType != "" {
		query.Where(earthquake.EqType(earthquakeType))
	}

	// Pagination
	eqs := query.
		Offset(offset).
		Limit(limit).
		AllX(c)

	return eqs, nil
}

func QuerryEarthquakeMultiFilteredByDay(c context.Context, limit int, offset int,
	magnitude float64, place string, earthquakeType string, day int) ([]*ent.Earthquake, error) {

	dayAgo := time.Now().AddDate(0, 0, -day)

	query := db.Client.Earthquake.Query().Where(sql.FieldGT(earthquake.FieldCreatedAt, dayAgo))

	if magnitude > 0 {
		query.Where(earthquake.MagGT(magnitude))
	}
	if place != "" {
		query.QueryGeometry().QueryLocation().Where(location.Place(place))
	}
	if earthquakeType != "" {
		query.Where(earthquake.EqType(earthquakeType))
	}

	// Pagination
	eqs := query.
		Offset(offset).
		Limit(limit).
		AllX(c)

	return eqs, nil
}

func QuerryEarthquakeFilteredByID(c context.Context, id int) ([]*ent.Earthquake, error) {
	eqs, err := db.Client.Earthquake.Query().
		Where(earthquake.ID(id)).
		All(c)

	if err != nil {
		return nil, err
	}

	return eqs, nil
}

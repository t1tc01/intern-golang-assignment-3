// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.com/hedwig-phan/assignment-3/ent/geometry"
	"gitlab.com/hedwig-phan/assignment-3/ent/location"
)

// Geometry is the model entity for the Geometry schema.
type Geometry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LocationID holds the value of the "location_id" field.
	LocationID int `json:"location_id,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Depth holds the value of the "depth" field.
	Depth float64 `json:"depth,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GeometryQuery when eager-loading is set.
	Edges        GeometryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GeometryEdges holds the relations/edges for other nodes in the graph.
type GeometryEdges struct {
	// Earthquakes holds the value of the earthquakes edge.
	Earthquakes []*Earthquake `json:"earthquakes,omitempty"`
	// Location holds the value of the location edge.
	Location *Location `json:"location,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EarthquakesOrErr returns the Earthquakes value or an error if the edge
// was not loaded in eager-loading.
func (e GeometryEdges) EarthquakesOrErr() ([]*Earthquake, error) {
	if e.loadedTypes[0] {
		return e.Earthquakes, nil
	}
	return nil, &NotLoadedError{edge: "earthquakes"}
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GeometryEdges) LocationOrErr() (*Location, error) {
	if e.Location != nil {
		return e.Location, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: location.Label}
	}
	return nil, &NotLoadedError{edge: "location"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Geometry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case geometry.FieldLongitude, geometry.FieldLatitude, geometry.FieldDepth:
			values[i] = new(sql.NullFloat64)
		case geometry.FieldID, geometry.FieldLocationID:
			values[i] = new(sql.NullInt64)
		case geometry.FieldCreatedAt, geometry.FieldUpdatedAt, geometry.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Geometry fields.
func (ge *Geometry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case geometry.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ge.ID = int(value.Int64)
		case geometry.FieldLocationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value.Valid {
				ge.LocationID = int(value.Int64)
			}
		case geometry.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				ge.Longitude = value.Float64
			}
		case geometry.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				ge.Latitude = value.Float64
			}
		case geometry.FieldDepth:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field depth", values[i])
			} else if value.Valid {
				ge.Depth = value.Float64
			}
		case geometry.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ge.CreatedAt = value.Time
			}
		case geometry.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ge.UpdatedAt = value.Time
			}
		case geometry.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ge.DeletedAt = value.Time
			}
		default:
			ge.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Geometry.
// This includes values selected through modifiers, order, etc.
func (ge *Geometry) Value(name string) (ent.Value, error) {
	return ge.selectValues.Get(name)
}

// QueryEarthquakes queries the "earthquakes" edge of the Geometry entity.
func (ge *Geometry) QueryEarthquakes() *EarthquakeQuery {
	return NewGeometryClient(ge.config).QueryEarthquakes(ge)
}

// QueryLocation queries the "location" edge of the Geometry entity.
func (ge *Geometry) QueryLocation() *LocationQuery {
	return NewGeometryClient(ge.config).QueryLocation(ge)
}

// Update returns a builder for updating this Geometry.
// Note that you need to call Geometry.Unwrap() before calling this method if this Geometry
// was returned from a transaction, and the transaction was committed or rolled back.
func (ge *Geometry) Update() *GeometryUpdateOne {
	return NewGeometryClient(ge.config).UpdateOne(ge)
}

// Unwrap unwraps the Geometry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ge *Geometry) Unwrap() *Geometry {
	_tx, ok := ge.config.driver.(*txDriver)
	if !ok {
		panic("ent: Geometry is not a transactional entity")
	}
	ge.config.driver = _tx.drv
	return ge
}

// String implements the fmt.Stringer.
func (ge *Geometry) String() string {
	var builder strings.Builder
	builder.WriteString("Geometry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ge.ID))
	builder.WriteString("location_id=")
	builder.WriteString(fmt.Sprintf("%v", ge.LocationID))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", ge.Longitude))
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", ge.Latitude))
	builder.WriteString(", ")
	builder.WriteString("depth=")
	builder.WriteString(fmt.Sprintf("%v", ge.Depth))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ge.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ge.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ge.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Geometries is a parsable slice of Geometry.
type Geometries []*Geometry
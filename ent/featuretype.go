// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.com/hedwig-phan/assignment-3/ent/featuretype"
)

// FeatureType is the model entity for the FeatureType schema.
type FeatureType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FeatType holds the value of the "feat_type" field.
	FeatType string `json:"feat_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FeatureTypeQuery when eager-loading is set.
	Edges        FeatureTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FeatureTypeEdges holds the relations/edges for other nodes in the graph.
type FeatureTypeEdges struct {
	// FtypeEarthquakes holds the value of the ftype_earthquakes edge.
	FtypeEarthquakes []*FtypeEarthquake `json:"ftype_earthquakes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FtypeEarthquakesOrErr returns the FtypeEarthquakes value or an error if the edge
// was not loaded in eager-loading.
func (e FeatureTypeEdges) FtypeEarthquakesOrErr() ([]*FtypeEarthquake, error) {
	if e.loadedTypes[0] {
		return e.FtypeEarthquakes, nil
	}
	return nil, &NotLoadedError{edge: "ftype_earthquakes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeatureType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case featuretype.FieldID:
			values[i] = new(sql.NullInt64)
		case featuretype.FieldFeatType:
			values[i] = new(sql.NullString)
		case featuretype.FieldCreatedAt, featuretype.FieldUpdatedAt, featuretype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeatureType fields.
func (ft *FeatureType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case featuretype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ft.ID = int(value.Int64)
		case featuretype.FieldFeatType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feat_type", values[i])
			} else if value.Valid {
				ft.FeatType = value.String
			}
		case featuretype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ft.CreatedAt = value.Time
			}
		case featuretype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ft.UpdatedAt = value.Time
			}
		case featuretype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ft.DeletedAt = value.Time
			}
		default:
			ft.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeatureType.
// This includes values selected through modifiers, order, etc.
func (ft *FeatureType) Value(name string) (ent.Value, error) {
	return ft.selectValues.Get(name)
}

// QueryFtypeEarthquakes queries the "ftype_earthquakes" edge of the FeatureType entity.
func (ft *FeatureType) QueryFtypeEarthquakes() *FtypeEarthquakeQuery {
	return NewFeatureTypeClient(ft.config).QueryFtypeEarthquakes(ft)
}

// Update returns a builder for updating this FeatureType.
// Note that you need to call FeatureType.Unwrap() before calling this method if this FeatureType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FeatureType) Update() *FeatureTypeUpdateOne {
	return NewFeatureTypeClient(ft.config).UpdateOne(ft)
}

// Unwrap unwraps the FeatureType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FeatureType) Unwrap() *FeatureType {
	_tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeatureType is not a transactional entity")
	}
	ft.config.driver = _tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FeatureType) String() string {
	var builder strings.Builder
	builder.WriteString("FeatureType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ft.ID))
	builder.WriteString("feat_type=")
	builder.WriteString(ft.FeatType)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ft.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ft.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ft.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FeatureTypes is a parsable slice of FeatureType.
type FeatureTypes []*FeatureType

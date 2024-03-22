// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.com/hedwig-phan/assignment-3/ent/earthquake"
	"gitlab.com/hedwig-phan/assignment-3/ent/featuretype"
	"gitlab.com/hedwig-phan/assignment-3/ent/ftypeearthquake"
)

// FtypeEarthquake is the model entity for the FtypeEarthquake schema.
type FtypeEarthquake struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FtID holds the value of the "ft_id" field.
	FtID int `json:"ft_id,omitempty"`
	// EqID holds the value of the "eq_id" field.
	EqID int `json:"eq_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FtypeEarthquakeQuery when eager-loading is set.
	Edges        FtypeEarthquakeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FtypeEarthquakeEdges holds the relations/edges for other nodes in the graph.
type FtypeEarthquakeEdges struct {
	// Earthquake holds the value of the earthquake edge.
	Earthquake *Earthquake `json:"earthquake,omitempty"`
	// FeatureType holds the value of the feature_type edge.
	FeatureType *FeatureType `json:"feature_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EarthquakeOrErr returns the Earthquake value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FtypeEarthquakeEdges) EarthquakeOrErr() (*Earthquake, error) {
	if e.Earthquake != nil {
		return e.Earthquake, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: earthquake.Label}
	}
	return nil, &NotLoadedError{edge: "earthquake"}
}

// FeatureTypeOrErr returns the FeatureType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FtypeEarthquakeEdges) FeatureTypeOrErr() (*FeatureType, error) {
	if e.FeatureType != nil {
		return e.FeatureType, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: featuretype.Label}
	}
	return nil, &NotLoadedError{edge: "feature_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FtypeEarthquake) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ftypeearthquake.FieldID, ftypeearthquake.FieldFtID, ftypeearthquake.FieldEqID:
			values[i] = new(sql.NullInt64)
		case ftypeearthquake.FieldCreatedAt, ftypeearthquake.FieldUpdatedAt, ftypeearthquake.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FtypeEarthquake fields.
func (fe *FtypeEarthquake) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ftypeearthquake.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fe.ID = int(value.Int64)
		case ftypeearthquake.FieldFtID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ft_id", values[i])
			} else if value.Valid {
				fe.FtID = int(value.Int64)
			}
		case ftypeearthquake.FieldEqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field eq_id", values[i])
			} else if value.Valid {
				fe.EqID = int(value.Int64)
			}
		case ftypeearthquake.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fe.CreatedAt = value.Time
			}
		case ftypeearthquake.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fe.UpdatedAt = value.Time
			}
		case ftypeearthquake.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fe.DeletedAt = value.Time
			}
		default:
			fe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FtypeEarthquake.
// This includes values selected through modifiers, order, etc.
func (fe *FtypeEarthquake) Value(name string) (ent.Value, error) {
	return fe.selectValues.Get(name)
}

// QueryEarthquake queries the "earthquake" edge of the FtypeEarthquake entity.
func (fe *FtypeEarthquake) QueryEarthquake() *EarthquakeQuery {
	return NewFtypeEarthquakeClient(fe.config).QueryEarthquake(fe)
}

// QueryFeatureType queries the "feature_type" edge of the FtypeEarthquake entity.
func (fe *FtypeEarthquake) QueryFeatureType() *FeatureTypeQuery {
	return NewFtypeEarthquakeClient(fe.config).QueryFeatureType(fe)
}

// Update returns a builder for updating this FtypeEarthquake.
// Note that you need to call FtypeEarthquake.Unwrap() before calling this method if this FtypeEarthquake
// was returned from a transaction, and the transaction was committed or rolled back.
func (fe *FtypeEarthquake) Update() *FtypeEarthquakeUpdateOne {
	return NewFtypeEarthquakeClient(fe.config).UpdateOne(fe)
}

// Unwrap unwraps the FtypeEarthquake entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fe *FtypeEarthquake) Unwrap() *FtypeEarthquake {
	_tx, ok := fe.config.driver.(*txDriver)
	if !ok {
		panic("ent: FtypeEarthquake is not a transactional entity")
	}
	fe.config.driver = _tx.drv
	return fe
}

// String implements the fmt.Stringer.
func (fe *FtypeEarthquake) String() string {
	var builder strings.Builder
	builder.WriteString("FtypeEarthquake(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fe.ID))
	builder.WriteString("ft_id=")
	builder.WriteString(fmt.Sprintf("%v", fe.FtID))
	builder.WriteString(", ")
	builder.WriteString("eq_id=")
	builder.WriteString(fmt.Sprintf("%v", fe.EqID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fe.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FtypeEarthquakes is a parsable slice of FtypeEarthquake.
type FtypeEarthquakes []*FtypeEarthquake

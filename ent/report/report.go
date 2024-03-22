// Code generated by ent, DO NOT EDIT.

package report

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the report type in the database.
	Label = "report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFelt holds the string denoting the felt field in the database.
	FieldFelt = "felt"
	// FieldCdi holds the string denoting the cdi field in the database.
	FieldCdi = "cdi"
	// FieldMmi holds the string denoting the mmi field in the database.
	FieldMmi = "mmi"
	// FieldAlert holds the string denoting the alert field in the database.
	FieldAlert = "alert"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeEarthquakes holds the string denoting the earthquakes edge name in mutations.
	EdgeEarthquakes = "earthquakes"
	// Table holds the table name of the report in the database.
	Table = "report"
	// EarthquakesTable is the table that holds the earthquakes relation/edge.
	EarthquakesTable = "earthquake"
	// EarthquakesInverseTable is the table name for the Earthquake entity.
	// It exists in this package in order to avoid circular dependency with the "earthquake" package.
	EarthquakesInverseTable = "earthquake"
	// EarthquakesColumn is the table column denoting the earthquakes relation/edge.
	EarthquakesColumn = "report_id"
)

// Columns holds all SQL columns for report fields.
var Columns = []string{
	FieldID,
	FieldFelt,
	FieldCdi,
	FieldMmi,
	FieldAlert,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Report queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFelt orders the results by the felt field.
func ByFelt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFelt, opts...).ToFunc()
}

// ByCdi orders the results by the cdi field.
func ByCdi(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCdi, opts...).ToFunc()
}

// ByMmi orders the results by the mmi field.
func ByMmi(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMmi, opts...).ToFunc()
}

// ByAlert orders the results by the alert field.
func ByAlert(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlert, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByEarthquakesCount orders the results by earthquakes count.
func ByEarthquakesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEarthquakesStep(), opts...)
	}
}

// ByEarthquakes orders the results by earthquakes terms.
func ByEarthquakes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEarthquakesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEarthquakesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EarthquakesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EarthquakesTable, EarthquakesColumn),
	)
}

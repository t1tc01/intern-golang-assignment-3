// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApireqColumns holds the columns for the "apireq" table.
	ApireqColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "req_time", Type: field.TypeTime},
		{Name: "req_param", Type: field.TypeJSON, Nullable: true},
		{Name: "req_body", Type: field.TypeJSON, Nullable: true},
		{Name: "req_headers", Type: field.TypeJSON, Nullable: true},
		{Name: "req_metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "resp_time", Type: field.TypeTime},
		{Name: "resp_status", Type: field.TypeInt, Nullable: true},
		{Name: "resp_body", Type: field.TypeJSON, Nullable: true},
		{Name: "resp_headers", Type: field.TypeJSON, Nullable: true},
		{Name: "resp_metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// ApireqTable holds the schema information for the "apireq" table.
	ApireqTable = &schema.Table{
		Name:       "apireq",
		Columns:    ApireqColumns,
		PrimaryKey: []*schema.Column{ApireqColumns[0]},
	}
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hashed_password", Type: field.TypeString},
		{Name: "last_login", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:       "credentials",
		Columns:    CredentialsColumns,
		PrimaryKey: []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credentials_user_credentials",
				Columns:    []*schema.Column{CredentialsColumns[3]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EarthquakeColumns holds the columns for the "earthquake" table.
	EarthquakeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "mag", Type: field.TypeFloat64},
		{Name: "time", Type: field.TypeTime},
		{Name: "updated_time", Type: field.TypeTime, Nullable: true},
		{Name: "tz", Type: field.TypeInt32, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "detail", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "tsunami", Type: field.TypeInt32, Nullable: true},
		{Name: "sig", Type: field.TypeInt32, Nullable: true},
		{Name: "net", Type: field.TypeString, Nullable: true},
		{Name: "code", Type: field.TypeString, Nullable: true},
		{Name: "nst", Type: field.TypeInt32, Nullable: true},
		{Name: "dmin", Type: field.TypeFloat64, Nullable: true},
		{Name: "rms", Type: field.TypeFloat64, Nullable: true},
		{Name: "gap", Type: field.TypeFloat64, Nullable: true},
		{Name: "mag_type", Type: field.TypeString, Nullable: true},
		{Name: "eq_type", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "geo_id", Type: field.TypeInt, Nullable: true},
		{Name: "report_id", Type: field.TypeInt, Nullable: true},
	}
	// EarthquakeTable holds the schema information for the "earthquake" table.
	EarthquakeTable = &schema.Table{
		Name:       "earthquake",
		Columns:    EarthquakeColumns,
		PrimaryKey: []*schema.Column{EarthquakeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "earthquake_geometry_earthquakes",
				Columns:    []*schema.Column{EarthquakeColumns[21]},
				RefColumns: []*schema.Column{GeometryColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "earthquake_report_earthquakes",
				Columns:    []*schema.Column{EarthquakeColumns[22]},
				RefColumns: []*schema.Column{ReportColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FeatureTypeColumns holds the columns for the "feature_type" table.
	FeatureTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "feat_type", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// FeatureTypeTable holds the schema information for the "feature_type" table.
	FeatureTypeTable = &schema.Table{
		Name:       "feature_type",
		Columns:    FeatureTypeColumns,
		PrimaryKey: []*schema.Column{FeatureTypeColumns[0]},
	}
	// FtypeEarthquakeColumns holds the columns for the "ftype_earthquake" table.
	FtypeEarthquakeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "eq_id", Type: field.TypeInt, Nullable: true},
		{Name: "ft_id", Type: field.TypeInt, Nullable: true},
	}
	// FtypeEarthquakeTable holds the schema information for the "ftype_earthquake" table.
	FtypeEarthquakeTable = &schema.Table{
		Name:       "ftype_earthquake",
		Columns:    FtypeEarthquakeColumns,
		PrimaryKey: []*schema.Column{FtypeEarthquakeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ftype_earthquake_earthquake_ftype_earthquakes",
				Columns:    []*schema.Column{FtypeEarthquakeColumns[4]},
				RefColumns: []*schema.Column{EarthquakeColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ftype_earthquake_feature_type_ftype_earthquakes",
				Columns:    []*schema.Column{FtypeEarthquakeColumns[5]},
				RefColumns: []*schema.Column{FeatureTypeColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GeometryColumns holds the columns for the "geometry" table.
	GeometryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "depth", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "location_id", Type: field.TypeInt, Nullable: true},
	}
	// GeometryTable holds the schema information for the "geometry" table.
	GeometryTable = &schema.Table{
		Name:       "geometry",
		Columns:    GeometryColumns,
		PrimaryKey: []*schema.Column{GeometryColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "geometry_location_geometries",
				Columns:    []*schema.Column{GeometryColumns[7]},
				RefColumns: []*schema.Column{LocationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LocationColumns holds the columns for the "location" table.
	LocationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "place", Type: field.TypeString, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// LocationTable holds the schema information for the "location" table.
	LocationTable = &schema.Table{
		Name:       "location",
		Columns:    LocationColumns,
		PrimaryKey: []*schema.Column{LocationColumns[0]},
	}
	// PasswordResetRequestColumns holds the columns for the "password_reset_request" table.
	PasswordResetRequestColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token_value", Type: field.TypeString},
		{Name: "expiration_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// PasswordResetRequestTable holds the schema information for the "password_reset_request" table.
	PasswordResetRequestTable = &schema.Table{
		Name:       "password_reset_request",
		Columns:    PasswordResetRequestColumns,
		PrimaryKey: []*schema.Column{PasswordResetRequestColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "password_reset_request_user_password_reset_request",
				Columns:    []*schema.Column{PasswordResetRequestColumns[3]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReportColumns holds the columns for the "report" table.
	ReportColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "felt", Type: field.TypeInt32, Nullable: true},
		{Name: "cdi", Type: field.TypeFloat64, Nullable: true},
		{Name: "mmi", Type: field.TypeFloat64, Nullable: true},
		{Name: "alert", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// ReportTable holds the schema information for the "report" table.
	ReportTable = &schema.Table{
		Name:       "report",
		Columns:    ReportColumns,
		PrimaryKey: []*schema.Column{ReportColumns[0]},
	}
	// SessionColumns holds the columns for the "session" table.
	SessionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "login_time", Type: field.TypeTime},
		{Name: "last_activity", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// SessionTable holds the schema information for the "session" table.
	SessionTable = &schema.Table{
		Name:       "session",
		Columns:    SessionColumns,
		PrimaryKey: []*schema.Column{SessionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "session_user_session",
				Columns:    []*schema.Column{SessionColumns[3]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SourceColumns holds the columns for the "source" table.
	SourceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SourceTable holds the schema information for the "source" table.
	SourceTable = &schema.Table{
		Name:       "source",
		Columns:    SourceColumns,
		PrimaryKey: []*schema.Column{SourceColumns[0]},
	}
	// SourceEarthquakeColumns holds the columns for the "source_earthquake" table.
	SourceEarthquakeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "eq_id", Type: field.TypeInt, Nullable: true},
		{Name: "s_id", Type: field.TypeInt, Nullable: true},
	}
	// SourceEarthquakeTable holds the schema information for the "source_earthquake" table.
	SourceEarthquakeTable = &schema.Table{
		Name:       "source_earthquake",
		Columns:    SourceEarthquakeColumns,
		PrimaryKey: []*schema.Column{SourceEarthquakeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_earthquake_earthquake_source_earthquakes",
				Columns:    []*schema.Column{SourceEarthquakeColumns[4]},
				RefColumns: []*schema.Column{EarthquakeColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "source_earthquake_source_source_earthquakes",
				Columns:    []*schema.Column{SourceEarthquakeColumns[5]},
				RefColumns: []*schema.Column{SourceColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TokenColumns holds the columns for the "token" table.
	TokenColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token_value", Type: field.TypeString},
		{Name: "expiration_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// TokenTable holds the schema information for the "token" table.
	TokenTable = &schema.Table{
		Name:       "token",
		Columns:    TokenColumns,
		PrimaryKey: []*schema.Column{TokenColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "token_user_token",
				Columns:    []*schema.Column{TokenColumns[3]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApireqTable,
		CredentialsTable,
		EarthquakeTable,
		FeatureTypeTable,
		FtypeEarthquakeTable,
		GeometryTable,
		LocationTable,
		PasswordResetRequestTable,
		ReportTable,
		SessionTable,
		SourceTable,
		SourceEarthquakeTable,
		TokenTable,
		UserTable,
	}
)

func init() {
	ApireqTable.Annotation = &entsql.Annotation{
		Table: "apireq",
	}
	CredentialsTable.ForeignKeys[0].RefTable = UserTable
	CredentialsTable.Annotation = &entsql.Annotation{
		Table: "credentials",
	}
	EarthquakeTable.ForeignKeys[0].RefTable = GeometryTable
	EarthquakeTable.ForeignKeys[1].RefTable = ReportTable
	EarthquakeTable.Annotation = &entsql.Annotation{
		Table: "earthquake",
	}
	FeatureTypeTable.Annotation = &entsql.Annotation{
		Table: "feature_type",
	}
	FtypeEarthquakeTable.ForeignKeys[0].RefTable = EarthquakeTable
	FtypeEarthquakeTable.ForeignKeys[1].RefTable = FeatureTypeTable
	FtypeEarthquakeTable.Annotation = &entsql.Annotation{
		Table: "ftype_earthquake",
	}
	GeometryTable.ForeignKeys[0].RefTable = LocationTable
	GeometryTable.Annotation = &entsql.Annotation{
		Table: "geometry",
	}
	LocationTable.Annotation = &entsql.Annotation{
		Table: "location",
	}
	PasswordResetRequestTable.ForeignKeys[0].RefTable = UserTable
	PasswordResetRequestTable.Annotation = &entsql.Annotation{
		Table: "password_reset_request",
	}
	ReportTable.Annotation = &entsql.Annotation{
		Table: "report",
	}
	SessionTable.ForeignKeys[0].RefTable = UserTable
	SessionTable.Annotation = &entsql.Annotation{
		Table: "session",
	}
	SourceTable.Annotation = &entsql.Annotation{
		Table: "source",
	}
	SourceEarthquakeTable.ForeignKeys[0].RefTable = EarthquakeTable
	SourceEarthquakeTable.ForeignKeys[1].RefTable = SourceTable
	SourceEarthquakeTable.Annotation = &entsql.Annotation{
		Table: "source_earthquake",
	}
	TokenTable.ForeignKeys[0].RefTable = UserTable
	TokenTable.Annotation = &entsql.Annotation{
		Table: "token",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}

package reflection

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"k8s.io/klog/v2"
)

// Schema contains information about the tag types in a BMDB. It also contains an
// active connection to the BMDB, allowing retrieval of data based on the
// detected schema.
//
// It also contains an embedded connection to the CockroachDB database backing
// this BMDB which is then used to retrieve data described by this schema.
type Schema struct {
	// TagTypes is the list of tag types extracted from the BMDB.
	TagTypes []TagType
	// Version is the go-migrate schema version of the BMDB this schema was extracted
	// from. By convention, it is a stringified base-10 number representing the number
	// of seconds since UNIX epoch of when the migration version was created, but
	// this is not guaranteed.
	Version string

	db *sql.DB
}

// TagType describes the type of a BMDB Tag. Each tag in turn corresponds to a
// CockroachDB database.
type TagType struct {
	// NativeName is the name of the table that holds tags of this type.
	NativeName string
	// Fields are the types of fields contained in this tag type.
	Fields []TagFieldType
}

// Name returns the canonical name of this tag type. For example, a table named
// machine_agent_started will have a canonical name AgentStarted.
func (r *TagType) Name() string {
	tableSuffix := strings.TrimPrefix(r.NativeName, "machine_")
	parts := strings.Split(tableSuffix, "_")
	// Capitalize some known acronyms.
	for i, p := range parts {
		parts[i] = strings.ReplaceAll(p, "os", "OS")
	}
	return strcase.ToCamel(strings.Join(parts, "_"))
}

// TagFieldType is the type of a field within a BMDB Tag. Each tag field in turn
// corresponds to a column inside its Tag table.
type TagFieldType struct {
	// NativeName is the name of the column that holds this field type. It is also
	// the canonical name of the field type.
	NativeName string
	// NativeType is the CockroachDB type name of this field.
	NativeType string
	// NativeUDTName is the CockroachDB user-defined-type name of this field. This is
	// only valid if NativeType is 'USER-DEFINED'.
	NativeUDTName string
}

// HumanType returns a human-readable representation of the field's type. This is
// not well-defined, and should be used only informatively.
func (r *TagFieldType) HumanType() string {
	switch r.NativeType {
	case "USER-DEFINED":
		return r.NativeUDTName
	case "timestamp with time zone":
		return "timestamp"
	case "bytea":
		return "bytes"
	default:
		return r.NativeType
	}
}

// Reflect builds a runtime BMDB schema from a raw SQL connection to the BMDB
// database. You're probably looking for bmdb.Connection.Reflect.
func Reflect(ctx context.Context, db *sql.DB) (*Schema, error) {
	// Get all tables in the currently connected to database.
	rows, err := db.QueryContext(ctx, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_catalog = current_database()
          AND table_schema = 'public'
          AND table_name LIKE 'machine\_%'
	`)
	if err != nil {
		return nil, fmt.Errorf("could not query table names: %w", err)
	}
	defer rows.Close()

	// Collect all table names for further processing.
	var tableNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("table name scan failed: %w", err)
		}
		tableNames = append(tableNames, name)
	}

	// Start processing each table into a TagType.
	tags := make([]TagType, 0, len(tableNames))
	for _, tagName := range tableNames {
		// Get all columns of the table.
		rows, err := db.QueryContext(ctx, `
			SELECT column_name, data_type, udt_name
			FROM information_schema.columns
		    WHERE table_catalog = current_database()
			  AND table_schema = 'public'
		      AND table_name = $1
		`, tagName)
		if err != nil {
			return nil, fmt.Errorf("could not query columns: %w", err)
		}

		tag := TagType{
			NativeName: tagName,
		}

		// Build field types from columns.
		foundMachineID := false
		for rows.Next() {
			var column_name, data_type, udt_name string
			if err := rows.Scan(&column_name, &data_type, &udt_name); err != nil {
				rows.Close()
				return nil, fmt.Errorf("column scan failed: %w", err)
			}
			if column_name == "machine_id" {
				foundMachineID = true
				continue
			}
			tag.Fields = append(tag.Fields, TagFieldType{
				NativeName:    column_name,
				NativeType:    data_type,
				NativeUDTName: udt_name,
			})
		}

		// Make sure there's a machine_id key in the table, then remove it.
		if !foundMachineID {
			klog.Warningf("Table %q has no machine_id column, skipping", tag.NativeName)
			continue
		}

		tags = append(tags, tag)
	}

	// Retrieve version information from go-migrate's schema_migrations table.
	var version string
	var dirty bool
	if err := db.QueryRowContext(ctx, "SELECT version, dirty FROM schema_migrations").Scan(&version, &dirty); err != nil {
		return nil, fmt.Errorf("could not select schema version: %w", err)
	}
	if dirty {
		version += " DIRTY!!!"
	}

	return &Schema{
		TagTypes: tags,
		Version:  version,

		db: db,
	}, nil
}

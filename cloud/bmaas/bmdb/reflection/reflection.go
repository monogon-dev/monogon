// Package reflection implements facilities to retrieve information about the
// implemented Tags and their types from a plain CockroachDB SQL connection,
// bypassing the queries/types defined in models. Then, the retrieved Schema can
// be used to retrieve information about machines.
//
// This is designed to be used in debugging facilities to allow arbitrary machine
// introspection. It must _not_ be used in the user path, as the schema
// extraction functionality is implemented best-effort.
package reflection

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GetMachinesOpts influences the behaviour of GetMachines.
type GetMachinesOpts struct {
	// FilterMachine, if set, will only retrieve information about the machine with
	// the given UUID. In case the given machine UUID does not exist in the database,
	// an empty result will be returned and _no_ error will be set.
	FilterMachine *uuid.UUID
	// Strict enables strict consistency. This is not recommended for use when
	// retrieving all machines, as such queries will compete against all currently
	// running operations. When not enabled, the retrieval will be executed AS OF
	// SYSTEM TIME follower_timestamp(), meaning the data might be a few seconds out
	// of date. Regardless of the option, the returned machine data will be
	// internally consistent, even across machines - but when not enabled the data
	// might be stale.
	Strict bool
	// ExpiredBackoffs enables the retrieval of information about all machine
	// backoffs, including expired backoff. Note that expired backoffs might be
	// garbage collected in the future, and their long-term storage is not
	// guaranteed.
	ExpiredBackoffs bool
}

// GetMachines retrieves all available BMDB data about one or more machines. The
// Schema's embedded SQL connection is used to performed the retrieval.
//
// Options can be specified to influenced the exact operation performed. By
// default (with a zeroed structure or nil pointer), all machines with active
// backoffs are retrieved with weak consistency. See GetMachineOpts to influence
// this behaviour.
func (r *Schema) GetMachines(ctx context.Context, opts *GetMachinesOpts) (*Reflected[[]*Machine], error) {
	if opts == nil {
		opts = &GetMachinesOpts{}
	}

	// We're about to build a pretty big SELECT query with a ton of joins.
	//
	// First, we join against work_backoff and work to get information about active
	// work and backoffs on the machines we're retrieving.
	//
	// Second, we join against all the tags that are declared in the schema.

	// These are the colums we'll SELECT <...> FROM
	columns := []string{
		"machines.machine_id",
		"machines.machine_created_at",
		"work_backoff.process",
		"work_backoff.cause",
		"work_backoff.until",
		"work.process",
		"work.session_id",
		// ... tag columns will come after this.
	}
	// These are tha args we'll pass to the query.
	var args []any

	// Start building joins. First, against work_backoff and work.
	backoffFilter := " AND work_backoff.until > now()"
	if opts.ExpiredBackoffs {
		backoffFilter = ""
	}
	joins := []string{
		"LEFT JOIN work_backoff ON machines.machine_id = work_backoff.machine_id" + backoffFilter,
		"LEFT JOIN work ON machines.machine_id = work.machine_id",
	}

	// Then, against tags. Also populate columns as we go along.
	for _, tagType := range r.TagTypes {
		joins = append(joins, fmt.Sprintf("LEFT JOIN %s ON machines.machine_id = %s.machine_id", tagType.NativeName, tagType.NativeName))
		columns = append(columns, fmt.Sprintf("%s.machine_id", tagType.NativeName))
		for _, fieldType := range tagType.Fields {
			columns = append(columns, fmt.Sprintf("%s.%s", tagType.NativeName, fieldType.NativeName))
		}
	}

	// Finalize query.
	q := []string{
		"SELECT",
		strings.Join(columns, ", "),
		"FROM machines",
	}
	q = append(q, joins...)
	if !opts.Strict {
		q = append(q, "AS OF SYSTEM TIME follower_read_timestamp()")
	}
	if opts.FilterMachine != nil {
		q = append(q, "WHERE machines.machine_id = $1")
		args = append(args, *opts.FilterMachine)
	}

	rows, err := r.db.QueryContext(ctx, strings.Join(q, "\n"), args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	// Okay, we can start scanning the result rows.
	//
	// As this is a complex join, we need to merge some rows together and discard
	// some NULLs. We do merging/deduplication using machine_id values for the
	// machine data, and abuse UNIQUE constraints in the work_backoff/work tables to
	// deduplicate these.
	//
	// The alternative would be to rewrite this query to use array_agg, and we might
	// do that at some point. This is only really a problem if we
	// have _a lot_ of active work/backoffs (as that effectively duplicates all
	// machine/tag data), which isn't the case yet. But we should keep an eye out for
	// this.

	var machines []*Machine
	for rows.Next() {

		// We need to scan this row back into columns. For constant columns we'll just
		// create the data here and refer to it later.
		var dests []any

		// Add non-tag always-retrieved constants.
		var mid uuid.UUID
		var machineCreated time.Time
		var workSession uuid.NullUUID
		var backoffProcess, backoffCause, workProcess sql.NullString
		var backoffUntil sql.NullTime

		dests = append(dests, &mid, &machineCreated, &backoffProcess, &backoffCause, &backoffUntil, &workProcess, &workSession)

		// For dynamic data, we need to keep a reference to a list of columns that are
		// part of tags, and then refer to them later. We can't just refer back to dests
		// as the types are erased into `any`. scannedTags is that data storage.
		type scannedTag struct {
			ty     *TagType
			id     uuid.NullUUID
			fields []*TagField
		}
		var scannedTags []*scannedTag
		for _, tagType := range r.TagTypes {
			tagType := tagType
			st := scannedTag{
				ty: &tagType,
			}
			scannedTags = append(scannedTags, &st)
			dests = append(dests, &st.id)
			for _, fieldType := range tagType.Fields {
				fieldType := fieldType
				field := TagField{
					Type: &fieldType,
				}
				dests = append(dests, &field)
				st.fields = append(st.fields, &field)

			}
		}

		if err := rows.Scan(dests...); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		// Now comes the merging/deduplication.

		// First, check if we are processing a new machine. If so, create a new
		// Machine. Otherwise, pick up the previous one.
		var machine *Machine
		if len(machines) == 0 || machines[len(machines)-1].ID.String() != mid.String() {
			// New machine or no machine yet.
			machine = &Machine{
				ID:       mid,
				Created:  machineCreated,
				Tags:     make(map[string]*Tag),
				Backoffs: make(map[string]Backoff),
				Work:     make(map[string]Work),
			}

			// Collect tags into machine.
			for _, st := range scannedTags {
				if !st.id.Valid {
					continue
				}
				var fields []TagField
				for _, f := range st.fields {
					fields = append(fields, *f)
				}
				machine.Tags[st.ty.Name()] = &Tag{
					Type:   st.ty,
					Fields: fields,
				}
			}
			machines = append(machines, machine)
		} else {
			// Continue previous machine.
			machine = machines[len(machines)-1]
		}

		// Do we have a backoff? Upsert it to the machine. This works because there's a
		// UNIQUE(machine_id, process) constraint on the work_backoff table, and we're
		// effectively rebuilding that keyspace here by indexing first by machine then by
		// process.
		if backoffCause.Valid && backoffProcess.Valid && backoffUntil.Valid {
			process := backoffProcess.String
			machine.Backoffs[process] = Backoff{
				Cause:   backoffCause.String,
				Process: process,
				Until:   backoffUntil.Time,
			}
		}

		// Do we have an active work item? Upsert it to the machine. Same UNIQUE
		// constraint abuse happening here.
		if workProcess.Valid && workSession.Valid {
			process := workProcess.String
			machine.Work[process] = Work{
				SessionID: workSession.UUID,
				Process:   process,
			}
		}
	}

	return &Reflected[[]*Machine]{
		Data:  machines,
		Query: strings.Join(q, " "),
	}, nil
}

// Reflected wraps data retrieved by reflection (T) with metadata about the
// retrieval.
type Reflected[T any] struct {
	Data T
	// Effective SQL query performed on the database.
	Query string
}

// Machine retrieved from BMDB.
type Machine struct {
	ID      uuid.UUID
	Created time.Time

	// Tags on this machine, keyed by Tag type name (canonical, not native).
	Tags map[string]*Tag

	// Backoffs on this machine, keyed by process name. By default these are only
	// active backoffs, unless ExpiredBackoffs was set on GetMachineOptions.
	Backoffs map[string]Backoff

	// Work active on this machine, keyed by process name.
	Work map[string]Work
}

// ActiveBackoffs retrieves a copy of a Machine's active backoffs. Note: the
// expiration check is performed according tu current system time, so it might
// not be consistent with the data snapshot retrieved from the database.
func (r *Machine) ActiveBackoffs() []*Backoff {
	var res []*Backoff
	for _, bo := range r.Backoffs {
		bo := bo
		if !bo.Active() {
			continue
		}
		res = append(res, &bo)
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Process < res[j].Process })
	return res
}

// ExpiredBackoffs retrieves a copy of a Machine's expired backoffs. Note: the
// expiration check is performed according tu current system time, so it might
// not be consistent with the data snapshot retrieved from the database.
func (r *Machine) ExpiredBackoffs() []*Backoff {
	var res []*Backoff
	for _, bo := range r.Backoffs {
		bo := bo
		if bo.Active() {
			continue
		}
		res = append(res, &bo)
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Process < res[j].Process })
	return res
}

// Tag value set on a Machine.
type Tag struct {
	// Type describing this tag.
	Type *TagType
	// Field data contained in this tag, sorted alphabetically by name.
	Fields []TagField
}

// Field is a shorthand for returning a TagField by its name.
func (r *Tag) Field(name string) *TagField {
	for _, f := range r.Fields {
		if f.Type.NativeName == name {
			return &f
		}
	}
	return nil
}

// TagField value which is part of a Tag set on a Machine.
type TagField struct {
	// Type describing this field.
	Type *TagFieldType

	text  *string
	bytes *[]byte
	time  *time.Time
}

// HumanValue returns a human-readable (best effort) representation of the field
// value.
func (r *TagField) HumanValue() string {
	switch {
	case r.text != nil:
		return *r.text
	case r.bytes != nil:
		return hex.EncodeToString(*r.bytes)
	case r.time != nil:
		return r.time.String()
	default:
		return "<unknown>"
	}
}

// Backoff on a Machine.
type Backoff struct {
	// Process which established Backoff.
	Process string
	// Time when Backoff expires.
	Until time.Time
	// Cause for the Backoff as emitted by worker.
	Cause string
}

// Active returns whether this Backoff is _currently_ active per the _local_ time.
func (r Backoff) Active() bool {
	return time.Now().Before(r.Until)
}

// Work being actively performed on a Machine.
type Work struct {
	// SessionID of the worker performing this Work.
	SessionID uuid.UUID
	// Process name of this Work.
	Process string
}

// Scan implements sql.Scanner for direct scanning of query results into a
// reflected tag value. This method is not meant to by used outside the
// reflection package.
func (r *TagField) Scan(src any) error {
	if src == nil {
		return nil
	}

	switch r.Type.NativeType {
	case "text":
		src2, ok := src.(string)
		if !ok {
			return fmt.Errorf("SQL type %q, but got %+v", r.Type.NativeType, src)
		}
		r.text = &src2
	case "bytea":
		src2, ok := src.([]byte)
		if !ok {
			return fmt.Errorf("SQL type %q, but got %+v", r.Type.NativeType, src)
		}
		// Copy the bytes, as they are otherwise going to be reused by the pq library.
		copied := make([]byte, len(src2))
		copy(copied[:], src2)
		r.bytes = &copied
	case "USER-DEFINED":
		switch r.Type.NativeUDTName {
		case "provider":
			src2, ok := src.([]byte)
			if !ok {
				return fmt.Errorf("SQL type %q, but got %+v", r.Type.NativeType, src)
			}
			src3 := string(src2)
			r.text = &src3
		}
	case "timestamp with time zone":
		src2, ok := src.(time.Time)
		if !ok {
			return fmt.Errorf("SQL type %q, but got %+v", r.Type.NativeType, src)
		}
		r.time = &src2
	default:
		return fmt.Errorf("unimplemented SQL type %q", r.Type.NativeType)
	}

	return nil
}

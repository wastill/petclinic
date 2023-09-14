// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wastill/petclinic/internal/data/ent/role"
	"github.com/wastill/petclinic/internal/data/ent/user"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges        RoleEdges `json:"edges"`
	user_roles   *int
	selectValues sql.SelectValues
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// BelongUser holds the value of the belong_user edge.
	BelongUser *User `json:"belong_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BelongUserOrErr returns the BelongUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleEdges) BelongUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.BelongUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.BelongUser, nil
	}
	return nil, &NotLoadedError{edge: "belong_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			values[i] = new(sql.NullInt64)
		case role.FieldUsername, role.FieldRole:
			values[i] = new(sql.NullString)
		case role.ForeignKeys[0]: // user_roles
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case role.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				r.Username = value.String
			}
		case role.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				r.Role = value.String
			}
		case role.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_roles", value)
			} else if value.Valid {
				r.user_roles = new(int)
				*r.user_roles = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role.
// This includes values selected through modifiers, order, etc.
func (r *Role) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryBelongUser queries the "belong_user" edge of the Role entity.
func (r *Role) QueryBelongUser() *UserQuery {
	return NewRoleClient(r.config).QueryBelongUser(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("username=")
	builder.WriteString(r.Username)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(r.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
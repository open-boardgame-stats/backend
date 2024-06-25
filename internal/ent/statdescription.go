// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
)

// StatDescription is the model entity for the StatDescription schema.
type StatDescription struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type stat.StatType `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata string `json:"metadata,omitempty"`
	// OrderNumber holds the value of the "order_number" field.
	OrderNumber int `json:"order_number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatDescriptionQuery when eager-loading is set.
	Edges        StatDescriptionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StatDescriptionEdges holds the relations/edges for other nodes in the graph.
type StatDescriptionEdges struct {
	// GameVersion holds the value of the game_version edge.
	GameVersion []*GameVersion `json:"game_version,omitempty"`
	// Stats holds the value of the stats edge.
	Stats []*Statistic `json:"stats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool

	namedGameVersion map[string][]*GameVersion
	namedStats       map[string][]*Statistic
}

// GameVersionOrErr returns the GameVersion value or an error if the edge
// was not loaded in eager-loading.
func (e StatDescriptionEdges) GameVersionOrErr() ([]*GameVersion, error) {
	if e.loadedTypes[0] {
		return e.GameVersion, nil
	}
	return nil, &NotLoadedError{edge: "game_version"}
}

// StatsOrErr returns the Stats value or an error if the edge
// was not loaded in eager-loading.
func (e StatDescriptionEdges) StatsOrErr() ([]*Statistic, error) {
	if e.loadedTypes[1] {
		return e.Stats, nil
	}
	return nil, &NotLoadedError{edge: "stats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StatDescription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case statdescription.FieldID:
			values[i] = new(guidgql.GUID)
		case statdescription.FieldOrderNumber:
			values[i] = new(sql.NullInt64)
		case statdescription.FieldName, statdescription.FieldDescription, statdescription.FieldMetadata:
			values[i] = new(sql.NullString)
		case statdescription.FieldType:
			values[i] = new(stat.StatType)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StatDescription fields.
func (sd *StatDescription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case statdescription.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sd.ID = *value
			}
		case statdescription.FieldType:
			if value, ok := values[i].(*stat.StatType); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value != nil {
				sd.Type = *value
			}
		case statdescription.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sd.Name = value.String
			}
		case statdescription.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				sd.Description = value.String
			}
		case statdescription.FieldMetadata:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value.Valid {
				sd.Metadata = value.String
			}
		case statdescription.FieldOrderNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_number", values[i])
			} else if value.Valid {
				sd.OrderNumber = int(value.Int64)
			}
		default:
			sd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StatDescription.
// This includes values selected through modifiers, order, etc.
func (sd *StatDescription) Value(name string) (ent.Value, error) {
	return sd.selectValues.Get(name)
}

// QueryGameVersion queries the "game_version" edge of the StatDescription entity.
func (sd *StatDescription) QueryGameVersion() *GameVersionQuery {
	return NewStatDescriptionClient(sd.config).QueryGameVersion(sd)
}

// QueryStats queries the "stats" edge of the StatDescription entity.
func (sd *StatDescription) QueryStats() *StatisticQuery {
	return NewStatDescriptionClient(sd.config).QueryStats(sd)
}

// Update returns a builder for updating this StatDescription.
// Note that you need to call StatDescription.Unwrap() before calling this method if this StatDescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (sd *StatDescription) Update() *StatDescriptionUpdateOne {
	return NewStatDescriptionClient(sd.config).UpdateOne(sd)
}

// Unwrap unwraps the StatDescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sd *StatDescription) Unwrap() *StatDescription {
	_tx, ok := sd.config.driver.(*txDriver)
	if !ok {
		panic("ent: StatDescription is not a transactional entity")
	}
	sd.config.driver = _tx.drv
	return sd
}

// String implements the fmt.Stringer.
func (sd *StatDescription) String() string {
	var builder strings.Builder
	builder.WriteString("StatDescription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sd.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", sd.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sd.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(sd.Description)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(sd.Metadata)
	builder.WriteString(", ")
	builder.WriteString("order_number=")
	builder.WriteString(fmt.Sprintf("%v", sd.OrderNumber))
	builder.WriteByte(')')
	return builder.String()
}

// NamedGameVersion returns the GameVersion named value or an error if the edge was not
// loaded in eager-loading with this name.
func (sd *StatDescription) NamedGameVersion(name string) ([]*GameVersion, error) {
	if sd.Edges.namedGameVersion == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := sd.Edges.namedGameVersion[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (sd *StatDescription) appendNamedGameVersion(name string, edges ...*GameVersion) {
	if sd.Edges.namedGameVersion == nil {
		sd.Edges.namedGameVersion = make(map[string][]*GameVersion)
	}
	if len(edges) == 0 {
		sd.Edges.namedGameVersion[name] = []*GameVersion{}
	} else {
		sd.Edges.namedGameVersion[name] = append(sd.Edges.namedGameVersion[name], edges...)
	}
}

// NamedStats returns the Stats named value or an error if the edge was not
// loaded in eager-loading with this name.
func (sd *StatDescription) NamedStats(name string) ([]*Statistic, error) {
	if sd.Edges.namedStats == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := sd.Edges.namedStats[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (sd *StatDescription) appendNamedStats(name string, edges ...*Statistic) {
	if sd.Edges.namedStats == nil {
		sd.Edges.namedStats = make(map[string][]*Statistic)
	}
	if len(edges) == 0 {
		sd.Edges.namedStats[name] = []*Statistic{}
	} else {
		sd.Edges.namedStats[name] = append(sd.Edges.namedStats[name], edges...)
	}
}

// StatDescriptions is a parsable slice of StatDescription.
type StatDescriptions []*StatDescription

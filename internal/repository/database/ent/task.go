// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"samsamoohooh-go-api/internal/repository/database/ent/group"
	"samsamoohooh-go-api/internal/repository/database/ent/task"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline time.Time `json:"deadline,omitempty"`
	// Range holds the value of the "range" field.
	Range int `json:"range,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges        TaskEdges `json:"edges"`
	group_tasks  *int
	selectValues sql.SelectValues
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Topics holds the value of the topics edge.
	Topics []*Topic `json:"topics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// TopicsOrErr returns the Topics value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TopicsOrErr() ([]*Topic, error) {
	if e.loadedTypes[1] {
		return e.Topics, nil
	}
	return nil, &NotLoadedError{edge: "topics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldID, task.FieldRange:
			values[i] = new(sql.NullInt64)
		case task.FieldCreatedAt, task.FieldUpdatedAt, task.FieldDeleteTime, task.FieldDeadline:
			values[i] = new(sql.NullTime)
		case task.ForeignKeys[0]: // group_tasks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case task.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				t.DeleteTime = value.Time
			}
		case task.FieldDeadline:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				t.Deadline = value.Time
			}
		case task.FieldRange:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field range", values[i])
			} else if value.Valid {
				t.Range = int(value.Int64)
			}
		case task.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_tasks", value)
			} else if value.Valid {
				t.group_tasks = new(int)
				*t.group_tasks = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Task.
// This includes values selected through modifiers, order, etc.
func (t *Task) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the Task entity.
func (t *Task) QueryGroup() *GroupQuery {
	return NewTaskClient(t.config).QueryGroup(t)
}

// QueryTopics queries the "topics" edge of the Task entity.
func (t *Task) QueryTopics() *TopicQuery {
	return NewTaskClient(t.config).QueryTopics(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return NewTaskClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_time=")
	builder.WriteString(t.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deadline=")
	builder.WriteString(t.Deadline.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("range=")
	builder.WriteString(fmt.Sprintf("%v", t.Range))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

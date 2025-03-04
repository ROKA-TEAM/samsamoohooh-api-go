// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent/group"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent/predicate"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent/task"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TaskUpdate) SetUpdatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetDeadline sets the "deadline" field.
func (tu *TaskUpdate) SetDeadline(t time.Time) *TaskUpdate {
	tu.mutation.SetDeadline(t)
	return tu
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDeadline(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDeadline(*t)
	}
	return tu
}

// SetRange sets the "range" field.
func (tu *TaskUpdate) SetRange(i int) *TaskUpdate {
	tu.mutation.ResetRange()
	tu.mutation.SetRange(i)
	return tu
}

// SetNillableRange sets the "range" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRange(i *int) *TaskUpdate {
	if i != nil {
		tu.SetRange(*i)
	}
	return tu
}

// AddRange adds i to the "range" field.
func (tu *TaskUpdate) AddRange(i int) *TaskUpdate {
	tu.mutation.AddRange(i)
	return tu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (tu *TaskUpdate) SetGroupID(id int) *TaskUpdate {
	tu.mutation.SetGroupID(id)
	return tu
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableGroupID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetGroupID(*id)
	}
	return tu
}

// SetGroup sets the "group" edge to the Group entity.
func (tu *TaskUpdate) SetGroup(g *Group) *TaskUpdate {
	return tu.SetGroupID(g.ID)
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (tu *TaskUpdate) AddTopicIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddTopicIDs(ids...)
	return tu
}

// AddTopics adds the "topics" edges to the Topic entity.
func (tu *TaskUpdate) AddTopics(t ...*Topic) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTopicIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (tu *TaskUpdate) ClearGroup() *TaskUpdate {
	tu.mutation.ClearGroup()
	return tu
}

// ClearTopics clears all "topics" edges to the Topic entity.
func (tu *TaskUpdate) ClearTopics() *TaskUpdate {
	tu.mutation.ClearTopics()
	return tu
}

// RemoveTopicIDs removes the "topics" edge to Topic entities by IDs.
func (tu *TaskUpdate) RemoveTopicIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveTopicIDs(ids...)
	return tu
}

// RemoveTopics removes "topics" edges to Topic entities.
func (tu *TaskUpdate) RemoveTopics(t ...*Topic) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTopicIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := task.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Deadline(); ok {
		_spec.SetField(task.FieldDeadline, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Range(); ok {
		_spec.SetField(task.FieldRange, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedRange(); ok {
		_spec.AddField(task.FieldRange, field.TypeInt, value)
	}
	if tu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.GroupTable,
			Columns: []string{task.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.GroupTable,
			Columns: []string{task.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TopicsTable,
			Columns: []string{task.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTopicsIDs(); len(nodes) > 0 && !tu.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TopicsTable,
			Columns: []string{task.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TopicsTable,
			Columns: []string{task.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TaskUpdateOne) SetUpdatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetDeadline sets the "deadline" field.
func (tuo *TaskUpdateOne) SetDeadline(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDeadline(t)
	return tuo
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDeadline(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDeadline(*t)
	}
	return tuo
}

// SetRange sets the "range" field.
func (tuo *TaskUpdateOne) SetRange(i int) *TaskUpdateOne {
	tuo.mutation.ResetRange()
	tuo.mutation.SetRange(i)
	return tuo
}

// SetNillableRange sets the "range" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRange(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetRange(*i)
	}
	return tuo
}

// AddRange adds i to the "range" field.
func (tuo *TaskUpdateOne) AddRange(i int) *TaskUpdateOne {
	tuo.mutation.AddRange(i)
	return tuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (tuo *TaskUpdateOne) SetGroupID(id int) *TaskUpdateOne {
	tuo.mutation.SetGroupID(id)
	return tuo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableGroupID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetGroupID(*id)
	}
	return tuo
}

// SetGroup sets the "group" edge to the Group entity.
func (tuo *TaskUpdateOne) SetGroup(g *Group) *TaskUpdateOne {
	return tuo.SetGroupID(g.ID)
}

// AddTopicIDs adds the "topics" edge to the Topic entity by IDs.
func (tuo *TaskUpdateOne) AddTopicIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddTopicIDs(ids...)
	return tuo
}

// AddTopics adds the "topics" edges to the Topic entity.
func (tuo *TaskUpdateOne) AddTopics(t ...*Topic) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTopicIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (tuo *TaskUpdateOne) ClearGroup() *TaskUpdateOne {
	tuo.mutation.ClearGroup()
	return tuo
}

// ClearTopics clears all "topics" edges to the Topic entity.
func (tuo *TaskUpdateOne) ClearTopics() *TaskUpdateOne {
	tuo.mutation.ClearTopics()
	return tuo
}

// RemoveTopicIDs removes the "topics" edge to Topic entities by IDs.
func (tuo *TaskUpdateOne) RemoveTopicIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveTopicIDs(ids...)
	return tuo
}

// RemoveTopics removes "topics" edges to Topic entities.
func (tuo *TaskUpdateOne) RemoveTopics(t ...*Topic) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTopicIDs(ids...)
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := task.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Deadline(); ok {
		_spec.SetField(task.FieldDeadline, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Range(); ok {
		_spec.SetField(task.FieldRange, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedRange(); ok {
		_spec.AddField(task.FieldRange, field.TypeInt, value)
	}
	if tuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.GroupTable,
			Columns: []string{task.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.GroupTable,
			Columns: []string{task.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TopicsTable,
			Columns: []string{task.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTopicsIDs(); len(nodes) > 0 && !tuo.mutation.TopicsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TopicsTable,
			Columns: []string{task.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TopicsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.TopicsTable,
			Columns: []string{task.TopicsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}

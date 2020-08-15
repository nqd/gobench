// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/ent/eventlog"
	"github.com/gobench-io/gobench/ent/group"
	"github.com/gobench-io/gobench/ent/predicate"
)

// ApplicationUpdate is the builder for updating Application entities.
type ApplicationUpdate struct {
	config
	hooks      []Hook
	mutation   *ApplicationMutation
	predicates []predicate.Application
}

// Where adds a new predicate for the builder.
func (au *ApplicationUpdate) Where(ps ...predicate.Application) *ApplicationUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetName sets the name field.
func (au *ApplicationUpdate) SetName(s string) *ApplicationUpdate {
	au.mutation.SetName(s)
	return au
}

// SetStatus sets the status field.
func (au *ApplicationUpdate) SetStatus(s string) *ApplicationUpdate {
	au.mutation.SetStatus(s)
	return au
}

// SetCreatedAt sets the created_at field.
func (au *ApplicationUpdate) SetCreatedAt(t time.Time) *ApplicationUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableCreatedAt(t *time.Time) *ApplicationUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the updated_at field.
func (au *ApplicationUpdate) SetUpdatedAt(t time.Time) *ApplicationUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetScenario sets the scenario field.
func (au *ApplicationUpdate) SetScenario(s string) *ApplicationUpdate {
	au.mutation.SetScenario(s)
	return au
}

// AddGroupIDs adds the groups edge to Group by ids.
func (au *ApplicationUpdate) AddGroupIDs(ids ...int) *ApplicationUpdate {
	au.mutation.AddGroupIDs(ids...)
	return au
}

// AddGroups adds the groups edges to Group.
func (au *ApplicationUpdate) AddGroups(g ...*Group) *ApplicationUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.AddGroupIDs(ids...)
}

// AddEventLogIDs adds the eventLogs edge to EventLog by ids.
func (au *ApplicationUpdate) AddEventLogIDs(ids ...int) *ApplicationUpdate {
	au.mutation.AddEventLogIDs(ids...)
	return au
}

// AddEventLogs adds the eventLogs edges to EventLog.
func (au *ApplicationUpdate) AddEventLogs(e ...*EventLog) *ApplicationUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddEventLogIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (au *ApplicationUpdate) RemoveGroupIDs(ids ...int) *ApplicationUpdate {
	au.mutation.RemoveGroupIDs(ids...)
	return au
}

// RemoveGroups removes groups edges to Group.
func (au *ApplicationUpdate) RemoveGroups(g ...*Group) *ApplicationUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.RemoveGroupIDs(ids...)
}

// RemoveEventLogIDs removes the eventLogs edge to EventLog by ids.
func (au *ApplicationUpdate) RemoveEventLogIDs(ids ...int) *ApplicationUpdate {
	au.mutation.RemoveEventLogIDs(ids...)
	return au
}

// RemoveEventLogs removes eventLogs edges to EventLog.
func (au *ApplicationUpdate) RemoveEventLogs(e ...*EventLog) *ApplicationUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveEventLogIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *ApplicationUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := application.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApplicationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApplicationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ApplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   application.Table,
			Columns: application.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: application.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldName,
		})
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldStatus,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.Scenario(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldScenario,
		})
	}
	if nodes := au.mutation.RemovedGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.GroupsTable,
			Columns: []string{application.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.GroupsTable,
			Columns: []string{application.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := au.mutation.RemovedEventLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.EventLogsTable,
			Columns: []string{application.EventLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eventlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EventLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.EventLogsTable,
			Columns: []string{application.EventLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eventlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ApplicationUpdateOne is the builder for updating a single Application entity.
type ApplicationUpdateOne struct {
	config
	hooks    []Hook
	mutation *ApplicationMutation
}

// SetName sets the name field.
func (auo *ApplicationUpdateOne) SetName(s string) *ApplicationUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetStatus sets the status field.
func (auo *ApplicationUpdateOne) SetStatus(s string) *ApplicationUpdateOne {
	auo.mutation.SetStatus(s)
	return auo
}

// SetCreatedAt sets the created_at field.
func (auo *ApplicationUpdateOne) SetCreatedAt(t time.Time) *ApplicationUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableCreatedAt(t *time.Time) *ApplicationUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the updated_at field.
func (auo *ApplicationUpdateOne) SetUpdatedAt(t time.Time) *ApplicationUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetScenario sets the scenario field.
func (auo *ApplicationUpdateOne) SetScenario(s string) *ApplicationUpdateOne {
	auo.mutation.SetScenario(s)
	return auo
}

// AddGroupIDs adds the groups edge to Group by ids.
func (auo *ApplicationUpdateOne) AddGroupIDs(ids ...int) *ApplicationUpdateOne {
	auo.mutation.AddGroupIDs(ids...)
	return auo
}

// AddGroups adds the groups edges to Group.
func (auo *ApplicationUpdateOne) AddGroups(g ...*Group) *ApplicationUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.AddGroupIDs(ids...)
}

// AddEventLogIDs adds the eventLogs edge to EventLog by ids.
func (auo *ApplicationUpdateOne) AddEventLogIDs(ids ...int) *ApplicationUpdateOne {
	auo.mutation.AddEventLogIDs(ids...)
	return auo
}

// AddEventLogs adds the eventLogs edges to EventLog.
func (auo *ApplicationUpdateOne) AddEventLogs(e ...*EventLog) *ApplicationUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddEventLogIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (auo *ApplicationUpdateOne) RemoveGroupIDs(ids ...int) *ApplicationUpdateOne {
	auo.mutation.RemoveGroupIDs(ids...)
	return auo
}

// RemoveGroups removes groups edges to Group.
func (auo *ApplicationUpdateOne) RemoveGroups(g ...*Group) *ApplicationUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.RemoveGroupIDs(ids...)
}

// RemoveEventLogIDs removes the eventLogs edge to EventLog by ids.
func (auo *ApplicationUpdateOne) RemoveEventLogIDs(ids ...int) *ApplicationUpdateOne {
	auo.mutation.RemoveEventLogIDs(ids...)
	return auo
}

// RemoveEventLogs removes eventLogs edges to EventLog.
func (auo *ApplicationUpdateOne) RemoveEventLogs(e ...*EventLog) *ApplicationUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveEventLogIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *ApplicationUpdateOne) Save(ctx context.Context) (*Application, error) {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := application.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}

	var (
		err  error
		node *Application
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApplicationUpdateOne) SaveX(ctx context.Context) *Application {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *ApplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApplicationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ApplicationUpdateOne) sqlSave(ctx context.Context) (a *Application, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   application.Table,
			Columns: application.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: application.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Application.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldName,
		})
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldStatus,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.Scenario(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldScenario,
		})
	}
	if nodes := auo.mutation.RemovedGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.GroupsTable,
			Columns: []string{application.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.GroupsTable,
			Columns: []string{application.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := auo.mutation.RemovedEventLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.EventLogsTable,
			Columns: []string{application.EventLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eventlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EventLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.EventLogsTable,
			Columns: []string{application.EventLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: eventlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Application{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/chatroom"
	"server/ent/predicate"
	"server/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatRoomUpdate is the builder for updating ChatRoom entities.
type ChatRoomUpdate struct {
	config
	hooks    []Hook
	mutation *ChatRoomMutation
}

// Where appends a list predicates to the ChatRoomUpdate builder.
func (cru *ChatRoomUpdate) Where(ps ...predicate.ChatRoom) *ChatRoomUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetCreatedAt sets the "created_at" field.
func (cru *ChatRoomUpdate) SetCreatedAt(t time.Time) *ChatRoomUpdate {
	cru.mutation.SetCreatedAt(t)
	return cru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cru *ChatRoomUpdate) SetNillableCreatedAt(t *time.Time) *ChatRoomUpdate {
	if t != nil {
		cru.SetCreatedAt(*t)
	}
	return cru
}

// SetUpdatedAt sets the "updated_at" field.
func (cru *ChatRoomUpdate) SetUpdatedAt(t time.Time) *ChatRoomUpdate {
	cru.mutation.SetUpdatedAt(t)
	return cru
}

// SetName sets the "name" field.
func (cru *ChatRoomUpdate) SetName(s string) *ChatRoomUpdate {
	cru.mutation.SetName(s)
	return cru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cru *ChatRoomUpdate) SetNillableName(s *string) *ChatRoomUpdate {
	if s != nil {
		cru.SetName(*s)
	}
	return cru
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cru *ChatRoomUpdate) AddUserIDs(ids ...int) *ChatRoomUpdate {
	cru.mutation.AddUserIDs(ids...)
	return cru
}

// AddUsers adds the "users" edges to the User entity.
func (cru *ChatRoomUpdate) AddUsers(u ...*User) *ChatRoomUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cru.AddUserIDs(ids...)
}

// Mutation returns the ChatRoomMutation object of the builder.
func (cru *ChatRoomUpdate) Mutation() *ChatRoomMutation {
	return cru.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cru *ChatRoomUpdate) ClearUsers() *ChatRoomUpdate {
	cru.mutation.ClearUsers()
	return cru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cru *ChatRoomUpdate) RemoveUserIDs(ids ...int) *ChatRoomUpdate {
	cru.mutation.RemoveUserIDs(ids...)
	return cru
}

// RemoveUsers removes "users" edges to User entities.
func (cru *ChatRoomUpdate) RemoveUsers(u ...*User) *ChatRoomUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *ChatRoomUpdate) Save(ctx context.Context) (int, error) {
	cru.defaults()
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *ChatRoomUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *ChatRoomUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *ChatRoomUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cru *ChatRoomUpdate) defaults() {
	if _, ok := cru.mutation.UpdatedAt(); !ok {
		v := chatroom.UpdateDefaultUpdatedAt()
		cru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *ChatRoomUpdate) check() error {
	if v, ok := cru.mutation.Name(); ok {
		if err := chatroom.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ChatRoom.name": %w`, err)}
		}
	}
	return nil
}

func (cru *ChatRoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chatroom.Table, chatroom.Columns, sqlgraph.NewFieldSpec(chatroom.FieldID, field.TypeInt))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.CreatedAt(); ok {
		_spec.SetField(chatroom.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cru.mutation.UpdatedAt(); ok {
		_spec.SetField(chatroom.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cru.mutation.Name(); ok {
		_spec.SetField(chatroom.FieldName, field.TypeString, value)
	}
	if cru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.UsersTable,
			Columns: []string{chatroom.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.UsersTable,
			Columns: []string{chatroom.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.UsersTable,
			Columns: []string{chatroom.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatroom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// ChatRoomUpdateOne is the builder for updating a single ChatRoom entity.
type ChatRoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatRoomMutation
}

// SetCreatedAt sets the "created_at" field.
func (cruo *ChatRoomUpdateOne) SetCreatedAt(t time.Time) *ChatRoomUpdateOne {
	cruo.mutation.SetCreatedAt(t)
	return cruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cruo *ChatRoomUpdateOne) SetNillableCreatedAt(t *time.Time) *ChatRoomUpdateOne {
	if t != nil {
		cruo.SetCreatedAt(*t)
	}
	return cruo
}

// SetUpdatedAt sets the "updated_at" field.
func (cruo *ChatRoomUpdateOne) SetUpdatedAt(t time.Time) *ChatRoomUpdateOne {
	cruo.mutation.SetUpdatedAt(t)
	return cruo
}

// SetName sets the "name" field.
func (cruo *ChatRoomUpdateOne) SetName(s string) *ChatRoomUpdateOne {
	cruo.mutation.SetName(s)
	return cruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cruo *ChatRoomUpdateOne) SetNillableName(s *string) *ChatRoomUpdateOne {
	if s != nil {
		cruo.SetName(*s)
	}
	return cruo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cruo *ChatRoomUpdateOne) AddUserIDs(ids ...int) *ChatRoomUpdateOne {
	cruo.mutation.AddUserIDs(ids...)
	return cruo
}

// AddUsers adds the "users" edges to the User entity.
func (cruo *ChatRoomUpdateOne) AddUsers(u ...*User) *ChatRoomUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cruo.AddUserIDs(ids...)
}

// Mutation returns the ChatRoomMutation object of the builder.
func (cruo *ChatRoomUpdateOne) Mutation() *ChatRoomMutation {
	return cruo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cruo *ChatRoomUpdateOne) ClearUsers() *ChatRoomUpdateOne {
	cruo.mutation.ClearUsers()
	return cruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cruo *ChatRoomUpdateOne) RemoveUserIDs(ids ...int) *ChatRoomUpdateOne {
	cruo.mutation.RemoveUserIDs(ids...)
	return cruo
}

// RemoveUsers removes "users" edges to User entities.
func (cruo *ChatRoomUpdateOne) RemoveUsers(u ...*User) *ChatRoomUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cruo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the ChatRoomUpdate builder.
func (cruo *ChatRoomUpdateOne) Where(ps ...predicate.ChatRoom) *ChatRoomUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *ChatRoomUpdateOne) Select(field string, fields ...string) *ChatRoomUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated ChatRoom entity.
func (cruo *ChatRoomUpdateOne) Save(ctx context.Context) (*ChatRoom, error) {
	cruo.defaults()
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *ChatRoomUpdateOne) SaveX(ctx context.Context) *ChatRoom {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *ChatRoomUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *ChatRoomUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cruo *ChatRoomUpdateOne) defaults() {
	if _, ok := cruo.mutation.UpdatedAt(); !ok {
		v := chatroom.UpdateDefaultUpdatedAt()
		cruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *ChatRoomUpdateOne) check() error {
	if v, ok := cruo.mutation.Name(); ok {
		if err := chatroom.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ChatRoom.name": %w`, err)}
		}
	}
	return nil
}

func (cruo *ChatRoomUpdateOne) sqlSave(ctx context.Context) (_node *ChatRoom, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chatroom.Table, chatroom.Columns, sqlgraph.NewFieldSpec(chatroom.FieldID, field.TypeInt))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChatRoom.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatroom.FieldID)
		for _, f := range fields {
			if !chatroom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chatroom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.CreatedAt(); ok {
		_spec.SetField(chatroom.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.UpdatedAt(); ok {
		_spec.SetField(chatroom.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.Name(); ok {
		_spec.SetField(chatroom.FieldName, field.TypeString, value)
	}
	if cruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.UsersTable,
			Columns: []string{chatroom.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.UsersTable,
			Columns: []string{chatroom.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.UsersTable,
			Columns: []string{chatroom.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ChatRoom{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatroom.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}

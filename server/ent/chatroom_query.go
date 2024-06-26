// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"server/ent/chatroom"
	"server/ent/predicate"
	"server/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChatRoomQuery is the builder for querying ChatRoom entities.
type ChatRoomQuery struct {
	config
	ctx        *QueryContext
	order      []chatroom.OrderOption
	inters     []Interceptor
	predicates []predicate.ChatRoom
	withUsers  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChatRoomQuery builder.
func (crq *ChatRoomQuery) Where(ps ...predicate.ChatRoom) *ChatRoomQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit the number of records to be returned by this query.
func (crq *ChatRoomQuery) Limit(limit int) *ChatRoomQuery {
	crq.ctx.Limit = &limit
	return crq
}

// Offset to start from.
func (crq *ChatRoomQuery) Offset(offset int) *ChatRoomQuery {
	crq.ctx.Offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *ChatRoomQuery) Unique(unique bool) *ChatRoomQuery {
	crq.ctx.Unique = &unique
	return crq
}

// Order specifies how the records should be ordered.
func (crq *ChatRoomQuery) Order(o ...chatroom.OrderOption) *ChatRoomQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QueryUsers chains the current query on the "users" edge.
func (crq *ChatRoomQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: crq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chatroom.Table, chatroom.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chatroom.UsersTable, chatroom.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ChatRoom entity from the query.
// Returns a *NotFoundError when no ChatRoom was found.
func (crq *ChatRoomQuery) First(ctx context.Context) (*ChatRoom, error) {
	nodes, err := crq.Limit(1).All(setContextOp(ctx, crq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chatroom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *ChatRoomQuery) FirstX(ctx context.Context) *ChatRoom {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ChatRoom ID from the query.
// Returns a *NotFoundError when no ChatRoom ID was found.
func (crq *ChatRoomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(setContextOp(ctx, crq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chatroom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *ChatRoomQuery) FirstIDX(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ChatRoom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ChatRoom entity is found.
// Returns a *NotFoundError when no ChatRoom entities are found.
func (crq *ChatRoomQuery) Only(ctx context.Context) (*ChatRoom, error) {
	nodes, err := crq.Limit(2).All(setContextOp(ctx, crq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chatroom.Label}
	default:
		return nil, &NotSingularError{chatroom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *ChatRoomQuery) OnlyX(ctx context.Context) *ChatRoom {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ChatRoom ID in the query.
// Returns a *NotSingularError when more than one ChatRoom ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *ChatRoomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(setContextOp(ctx, crq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chatroom.Label}
	default:
		err = &NotSingularError{chatroom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *ChatRoomQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ChatRooms.
func (crq *ChatRoomQuery) All(ctx context.Context) ([]*ChatRoom, error) {
	ctx = setContextOp(ctx, crq.ctx, "All")
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ChatRoom, *ChatRoomQuery]()
	return withInterceptors[[]*ChatRoom](ctx, crq, qr, crq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crq *ChatRoomQuery) AllX(ctx context.Context) []*ChatRoom {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ChatRoom IDs.
func (crq *ChatRoomQuery) IDs(ctx context.Context) (ids []int, err error) {
	if crq.ctx.Unique == nil && crq.path != nil {
		crq.Unique(true)
	}
	ctx = setContextOp(ctx, crq.ctx, "IDs")
	if err = crq.Select(chatroom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *ChatRoomQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *ChatRoomQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crq.ctx, "Count")
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crq, querierCount[*ChatRoomQuery](), crq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crq *ChatRoomQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *ChatRoomQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crq.ctx, "Exist")
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *ChatRoomQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChatRoomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *ChatRoomQuery) Clone() *ChatRoomQuery {
	if crq == nil {
		return nil
	}
	return &ChatRoomQuery{
		config:     crq.config,
		ctx:        crq.ctx.Clone(),
		order:      append([]chatroom.OrderOption{}, crq.order...),
		inters:     append([]Interceptor{}, crq.inters...),
		predicates: append([]predicate.ChatRoom{}, crq.predicates...),
		withUsers:  crq.withUsers.Clone(),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (crq *ChatRoomQuery) WithUsers(opts ...func(*UserQuery)) *ChatRoomQuery {
	query := (&UserClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crq.withUsers = query
	return crq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ChatRoom.Query().
//		GroupBy(chatroom.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *ChatRoomQuery) GroupBy(field string, fields ...string) *ChatRoomGroupBy {
	crq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChatRoomGroupBy{build: crq}
	grbuild.flds = &crq.ctx.Fields
	grbuild.label = chatroom.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ChatRoom.Query().
//		Select(chatroom.FieldCreatedAt).
//		Scan(ctx, &v)
func (crq *ChatRoomQuery) Select(fields ...string) *ChatRoomSelect {
	crq.ctx.Fields = append(crq.ctx.Fields, fields...)
	sbuild := &ChatRoomSelect{ChatRoomQuery: crq}
	sbuild.label = chatroom.Label
	sbuild.flds, sbuild.scan = &crq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChatRoomSelect configured with the given aggregations.
func (crq *ChatRoomQuery) Aggregate(fns ...AggregateFunc) *ChatRoomSelect {
	return crq.Select().Aggregate(fns...)
}

func (crq *ChatRoomQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crq); err != nil {
				return err
			}
		}
	}
	for _, f := range crq.ctx.Fields {
		if !chatroom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *ChatRoomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ChatRoom, error) {
	var (
		nodes       = []*ChatRoom{}
		_spec       = crq.querySpec()
		loadedTypes = [1]bool{
			crq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ChatRoom).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ChatRoom{config: crq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := crq.withUsers; query != nil {
		if err := crq.loadUsers(ctx, query, nodes,
			func(n *ChatRoom) { n.Edges.Users = []*User{} },
			func(n *ChatRoom, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (crq *ChatRoomQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*ChatRoom, init func(*ChatRoom), assign func(*ChatRoom, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ChatRoom)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(chatroom.UsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.chat_room_users
		if fk == nil {
			return fmt.Errorf(`foreign-key "chat_room_users" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "chat_room_users" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (crq *ChatRoomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	_spec.Node.Columns = crq.ctx.Fields
	if len(crq.ctx.Fields) > 0 {
		_spec.Unique = crq.ctx.Unique != nil && *crq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *ChatRoomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chatroom.Table, chatroom.Columns, sqlgraph.NewFieldSpec(chatroom.FieldID, field.TypeInt))
	_spec.From = crq.sql
	if unique := crq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if crq.path != nil {
		_spec.Unique = true
	}
	if fields := crq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatroom.FieldID)
		for i := range fields {
			if fields[i] != chatroom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *ChatRoomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(chatroom.Table)
	columns := crq.ctx.Fields
	if len(columns) == 0 {
		columns = chatroom.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.ctx.Unique != nil && *crq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChatRoomGroupBy is the group-by builder for ChatRoom entities.
type ChatRoomGroupBy struct {
	selector
	build *ChatRoomQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *ChatRoomGroupBy) Aggregate(fns ...AggregateFunc) *ChatRoomGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the selector query and scans the result into the given value.
func (crgb *ChatRoomGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crgb.build.ctx, "GroupBy")
	if err := crgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatRoomQuery, *ChatRoomGroupBy](ctx, crgb.build, crgb, crgb.build.inters, v)
}

func (crgb *ChatRoomGroupBy) sqlScan(ctx context.Context, root *ChatRoomQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crgb.flds)+len(crgb.fns))
		for _, f := range *crgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChatRoomSelect is the builder for selecting fields of ChatRoom entities.
type ChatRoomSelect struct {
	*ChatRoomQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crs *ChatRoomSelect) Aggregate(fns ...AggregateFunc) *ChatRoomSelect {
	crs.fns = append(crs.fns, fns...)
	return crs
}

// Scan applies the selector query and scans the result into the given value.
func (crs *ChatRoomSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crs.ctx, "Select")
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChatRoomQuery, *ChatRoomSelect](ctx, crs.ChatRoomQuery, crs, crs.inters, v)
}

func (crs *ChatRoomSelect) sqlScan(ctx context.Context, root *ChatRoomQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crs.fns))
	for _, fn := range crs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

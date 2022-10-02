// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/image"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ImageQuery is the builder for querying Image entities.
type ImageQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	predicates         []predicate.Image
	withUserProfilePic *UserQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImageQuery builder.
func (iq *ImageQuery) Where(ps ...predicate.Image) *ImageQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *ImageQuery) Limit(limit int) *ImageQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *ImageQuery) Offset(offset int) *ImageQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ImageQuery) Unique(unique bool) *ImageQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *ImageQuery) Order(o ...OrderFunc) *ImageQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryUserProfilePic chains the current query on the "user_profile_pic" edge.
func (iq *ImageQuery) QueryUserProfilePic() *UserQuery {
	query := &UserQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, image.UserProfilePicTable, image.UserProfilePicColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Image entity from the query.
// Returns a *NotFoundError when no Image was found.
func (iq *ImageQuery) First(ctx context.Context) (*Image, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{image.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ImageQuery) FirstX(ctx context.Context) *Image {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Image ID from the query.
// Returns a *NotFoundError when no Image ID was found.
func (iq *ImageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{image.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ImageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Image entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Image entity is found.
// Returns a *NotFoundError when no Image entities are found.
func (iq *ImageQuery) Only(ctx context.Context) (*Image, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{image.Label}
	default:
		return nil, &NotSingularError{image.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ImageQuery) OnlyX(ctx context.Context) *Image {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Image ID in the query.
// Returns a *NotSingularError when more than one Image ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ImageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{image.Label}
	default:
		err = &NotSingularError{image.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ImageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Images.
func (iq *ImageQuery) All(ctx context.Context) ([]*Image, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *ImageQuery) AllX(ctx context.Context) []*Image {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Image IDs.
func (iq *ImageQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iq.Select(image.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ImageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ImageQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ImageQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ImageQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ImageQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ImageQuery) Clone() *ImageQuery {
	if iq == nil {
		return nil
	}
	return &ImageQuery{
		config:             iq.config,
		limit:              iq.limit,
		offset:             iq.offset,
		order:              append([]OrderFunc{}, iq.order...),
		predicates:         append([]predicate.Image{}, iq.predicates...),
		withUserProfilePic: iq.withUserProfilePic.Clone(),
		// clone intermediate query.
		sql:    iq.sql.Clone(),
		path:   iq.path,
		unique: iq.unique,
	}
}

// WithUserProfilePic tells the query-builder to eager-load the nodes that are connected to
// the "user_profile_pic" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImageQuery) WithUserProfilePic(opts ...func(*UserQuery)) *ImageQuery {
	query := &UserQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withUserProfilePic = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URLPath string `json:"url_path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Image.Query().
//		GroupBy(image.FieldURLPath).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *ImageQuery) GroupBy(field string, fields ...string) *ImageGroupBy {
	grbuild := &ImageGroupBy{config: iq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	grbuild.label = image.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URLPath string `json:"url_path,omitempty"`
//	}
//
//	client.Image.Query().
//		Select(image.FieldURLPath).
//		Scan(ctx, &v)
func (iq *ImageQuery) Select(fields ...string) *ImageSelect {
	iq.fields = append(iq.fields, fields...)
	selbuild := &ImageSelect{ImageQuery: iq}
	selbuild.label = image.Label
	selbuild.flds, selbuild.scan = &iq.fields, selbuild.Scan
	return selbuild
}

func (iq *ImageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !image.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *ImageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Image, error) {
	var (
		nodes       = []*Image{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withUserProfilePic != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, image.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Image).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Image{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withUserProfilePic; query != nil {
		if err := iq.loadUserProfilePic(ctx, query, nodes,
			func(n *Image) { n.Edges.UserProfilePic = []*User{} },
			func(n *Image, e *User) { n.Edges.UserProfilePic = append(n.Edges.UserProfilePic, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *ImageQuery) loadUserProfilePic(ctx context.Context, query *UserQuery, nodes []*Image, init func(*Image), assign func(*Image, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Image)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(image.UserProfilePicColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_profile_pic
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_profile_pic" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_profile_pic" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *ImageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.fields
	if len(iq.fields) > 0 {
		_spec.Unique = iq.unique != nil && *iq.unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *ImageQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (iq *ImageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: image.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for i := range fields {
			if fields[i] != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *ImageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(image.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = image.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.unique != nil && *iq.unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImageGroupBy is the group-by builder for Image entities.
type ImageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ImageGroupBy) Aggregate(fns ...AggregateFunc) *ImageGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *ImageGroupBy) Scan(ctx context.Context, v any) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

func (igb *ImageGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range igb.fields {
		if !image.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *ImageGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// ImageSelect is the builder for selecting fields of Image entities.
type ImageSelect struct {
	*ImageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *ImageSelect) Scan(ctx context.Context, v any) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.ImageQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

func (is *ImageSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

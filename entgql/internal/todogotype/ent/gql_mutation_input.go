// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/schema/schematype"
	"entgo.io/contrib/entgql/internal/todogotype/ent/category"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/bigintgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/todo"
)

// CreateCategoryInput represents a mutation input for creating categories.
type CreateCategoryInput struct {
	Text     string
	Status   category.Status
	Config   *schematype.CategoryConfig
	Duration *time.Duration
	Count    *uint64
	Strings  *[]string
	TodoIDs  []string
}

// Mutate applies the CreateCategoryInput on the CategoryMutation builder.
func (i *CreateCategoryInput) Mutate(m *CategoryMutation) {
	m.SetText(i.Text)
	m.SetStatus(i.Status)
	m.SetConfig(i.Config)
	if v := i.Duration; v != nil {
		m.SetDuration(*v)
	}
	if v := i.Count; v != nil {
		m.SetCount(*v)
	}
	if v := i.Strings; v != nil {
		m.SetStrings(*v)
	}
	if v := i.TodoIDs; len(v) > 0 {
		m.AddTodoIDs(v...)
	}
}

// SetInput applies the change-set in the CreateCategoryInput on the CategoryCreate builder.
func (c *CategoryCreate) SetInput(i CreateCategoryInput) *CategoryCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCategoryInput represents a mutation input for updating categories.
type UpdateCategoryInput struct {
	Text          *string
	Status        *category.Status
	ClearConfig   bool
	Config        *schematype.CategoryConfig
	ClearDuration bool
	Duration      *time.Duration
	ClearCount    bool
	Count         *uint64
	ClearStrings  bool
	Strings       *[]string
	AppendStrings *[]string
	AddTodoIDs    []string
	RemoveTodoIDs []string
}

// Mutate applies the UpdateCategoryInput on the CategoryMutation builder.
func (i *UpdateCategoryInput) Mutate(m *CategoryMutation) {
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearConfig {
		m.ClearConfig()
	}
	m.SetConfig(i.Config)
	if i.ClearDuration {
		m.ClearDuration()
	}
	if v := i.Duration; v != nil {
		m.SetDuration(*v)
	}
	if i.ClearCount {
		m.ClearCount()
	}
	if v := i.Count; v != nil {
		m.SetCount(*v)
	}
	if i.ClearStrings {
		m.ClearStrings()
	}
	if v := i.Strings; v != nil {
		m.SetStrings(*v)
	}
	if i.AppendStrings != nil {
		m.AppendStrings(*i.Strings)
	}
	if v := i.AddTodoIDs; len(v) > 0 {
		m.AddTodoIDs(v...)
	}
	if v := i.RemoveTodoIDs; len(v) > 0 {
		m.RemoveTodoIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateCategoryInput on the CategoryUpdate builder.
func (c *CategoryUpdate) SetInput(i UpdateCategoryInput) *CategoryUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCategoryInput on the CategoryUpdateOne builder.
func (c *CategoryUpdateOne) SetInput(i UpdateCategoryInput) *CategoryUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Status     todo.Status
	Priority   *int
	Text       string
	ParentID   *string
	ChildIDs   []string
	CategoryID *bigintgql.BigInt
	SecretID   *string
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetStatus(i.Status)
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	m.SetText(i.Text)
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
	if v := i.SecretID; v != nil {
		m.SetSecretID(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Status         *todo.Status
	Priority       *int
	Text           *string
	ClearParent    bool
	ParentID       *string
	AddChildIDs    []string
	RemoveChildIDs []string
	ClearSecret    bool
	SecretID       *string
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
	if i.ClearSecret {
		m.ClearSecret()
	}
	if v := i.SecretID; v != nil {
		m.SetSecretID(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

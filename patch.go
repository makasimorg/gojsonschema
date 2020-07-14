package gojsonschema

import (
	"github.com/xeipuuv/gojsonreference"
)

type JsonSchemaType interface {
	IsTyped() bool
	Add(etype string) error
	Contains(etype string) bool
	String() string
}

type SubSchema interface {
	Draft() *Draft
	ID() *gojsonreference.JsonReference
	Title() string
	Description() string
	Types() JsonSchemaType
	Property() string
	Ref() *gojsonreference.JsonReference
	RefSchema() SubSchema
	Parent() SubSchema
	ItemsChildren() []SubSchema
	ItemsChildrenIsSingleSchema() bool
	PropertiesChildren() []SubSchema
	Required() []string
	Const() string
	Enum() []string
	OneOf() []SubSchema
	AnyOf() []SubSchema
	AllOf() []SubSchema

	Extras() map[string]interface{}
	ExtraString(n string) (string, bool)
	ExtraBool(n string) (bool, bool)
}
func (s *subSchema) Draft() *Draft {
	return s.draft
}

func (s *subSchema) ID() *gojsonreference.JsonReference {
	return s.id
}

func (s *subSchema) Title() string{
	if s.title == nil {
		return ""
	}

	return *s.title
}

func (s *subSchema) Description() string{
	if s.description == nil {
		return ""
	}

	return *s.description
}

func (s *subSchema) Types() JsonSchemaType {
	return &s.types
}

func (s *subSchema) Property() string {
	return s.property
}

func (s *subSchema) Ref() *gojsonreference.JsonReference {
	return s.ref
}

func (s *subSchema) RefSchema() SubSchema {
	return s.refSchema
}

// todo add to interface

func (s *subSchema) Parent() SubSchema {
	return s.parent
}

func (s *subSchema) ItemsChildren() []SubSchema {
	items := make([]SubSchema, 0, len(s.itemsChildren))
	for _, i := range s.itemsChildren {
		items = append(items, i)
	}

	return items
}

func (s *subSchema) ItemsChildrenIsSingleSchema() bool {
	return s.itemsChildrenIsSingleSchema
}

func (s *subSchema) PropertiesChildren() []SubSchema {
	items := make([]SubSchema, 0, len(s.propertiesChildren))
	for _, i := range s.propertiesChildren {
		items = append(items, i)
	}

	return items
}

func (s *subSchema) Required() []string {
	return s.required
}

func (s *subSchema) Const() string {
	if s._const == nil {
		return ""
	}

	return *s._const
}

func (s *subSchema) Enum() []string {
	return s.enum
}

func (s *subSchema) OneOf() []SubSchema {
	items := make([]SubSchema, 0, len(s.oneOf))
	for _, i := range s.oneOf {
		items = append(items, i)
	}

	return items
}

func (s *subSchema) AnyOf() []SubSchema {
	items := make([]SubSchema, 0, len(s.anyOf))
	for _, i := range s.anyOf {
		items = append(items, i)
	}

	return items
}

func (s *subSchema) AllOf() []SubSchema {
	// log.Printf("s = %T: %+v", s, s)
	// log.Printf("s.allOf = %T: %+v", s.allOf, s.allOf)
	items := make([]SubSchema, 0, len(s.allOf))
	for _, i := range s.allOf {
		items = append(items, i)
	}

	return items
}

func (s *subSchema) Extras() map[string]interface{} {
	return s.extras
}

func (s *subSchema) ExtraString(n string) (string, bool) {
	v, ok := s.extras[n].(string)
	if !ok {
		return "", false
	}

	return v, true
}

func (s *subSchema) ExtraBool(n string) (bool, bool) {
	v, ok := s.extras[n].(bool)
	if !ok {
		return false, false
	}

	return v, true
}

func (s *Schema) SubSchema() SubSchema {
	return s.rootSchema
}
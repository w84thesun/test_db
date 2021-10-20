// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package sakila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type filmCategoryViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *filmCategoryViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("film_category").
func (v *filmCategoryViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *filmCategoryViewType) Columns() []string {
	return []string{
		"film_id",
		"category_id",
		"last_update",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *filmCategoryViewType) NewStruct() reform.Struct {
	return new(FilmCategory)
}

// FilmCategoryView represents film_category view or table in SQL database.
var FilmCategoryView = &filmCategoryViewType{
	s: parse.StructInfo{
		Type:    "FilmCategory",
		SQLName: "film_category",
		Fields: []parse.FieldInfo{
			{Name: "FilmID", Type: "int16", Column: "film_id"},
			{Name: "CategoryID", Type: "int8", Column: "category_id"},
			{Name: "LastUpdate", Type: "time.Time", Column: "last_update"},
		},
		PKFieldIndex: -1,
	},
	z: new(FilmCategory).Values(),
}

// String returns a string representation of this struct or record.
func (s FilmCategory) String() string {
	res := make([]string, 3)
	res[0] = "FilmID: " + reform.Inspect(s.FilmID, true)
	res[1] = "CategoryID: " + reform.Inspect(s.CategoryID, true)
	res[2] = "LastUpdate: " + reform.Inspect(s.LastUpdate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *FilmCategory) Values() []interface{} {
	return []interface{}{
		s.FilmID,
		s.CategoryID,
		s.LastUpdate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *FilmCategory) Pointers() []interface{} {
	return []interface{}{
		&s.FilmID,
		&s.CategoryID,
		&s.LastUpdate,
	}
}

// View returns View object for that struct.
func (s *FilmCategory) View() reform.View {
	return FilmCategoryView
}

// check interfaces
var (
	_ reform.View   = FilmCategoryView
	_ reform.Struct = (*FilmCategory)(nil)
	_ fmt.Stringer  = (*FilmCategory)(nil)
)

func init() {
	parse.AssertUpToDate(&FilmCategoryView.s, new(FilmCategory))
}

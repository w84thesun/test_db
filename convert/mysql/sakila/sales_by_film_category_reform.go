// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package sakila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type salesByFilmCategoryViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *salesByFilmCategoryViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("sales_by_film_category").
func (v *salesByFilmCategoryViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *salesByFilmCategoryViewType) Columns() []string {
	return []string{
		"category",
		"total_sales",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *salesByFilmCategoryViewType) NewStruct() reform.Struct {
	return new(SalesByFilmCategory)
}

// SalesByFilmCategoryView represents sales_by_film_category view or table in SQL database.
var SalesByFilmCategoryView = &salesByFilmCategoryViewType{
	s: parse.StructInfo{
		Type:    "SalesByFilmCategory",
		SQLName: "sales_by_film_category",
		Fields: []parse.FieldInfo{
			{Name: "Category", Type: "string", Column: "category"},
			{Name: "TotalSales", Type: "*string", Column: "total_sales"},
		},
		PKFieldIndex: -1,
	},
	z: new(SalesByFilmCategory).Values(),
}

// String returns a string representation of this struct or record.
func (s SalesByFilmCategory) String() string {
	res := make([]string, 2)
	res[0] = "Category: " + reform.Inspect(s.Category, true)
	res[1] = "TotalSales: " + reform.Inspect(s.TotalSales, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *SalesByFilmCategory) Values() []interface{} {
	return []interface{}{
		s.Category,
		s.TotalSales,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *SalesByFilmCategory) Pointers() []interface{} {
	return []interface{}{
		&s.Category,
		&s.TotalSales,
	}
}

// View returns View object for that struct.
func (s *SalesByFilmCategory) View() reform.View {
	return SalesByFilmCategoryView
}

// check interfaces
var (
	_ reform.View   = SalesByFilmCategoryView
	_ reform.Struct = (*SalesByFilmCategory)(nil)
	_ fmt.Stringer  = (*SalesByFilmCategory)(nil)
)

func init() {
	parse.AssertUpToDate(&SalesByFilmCategoryView.s, new(SalesByFilmCategory))
}

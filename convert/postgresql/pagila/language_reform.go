// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package pagila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type languageTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *languageTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("language").
func (v *languageTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *languageTableType) Columns() []string {
	return []string{
		"language_id",
		"name",
		"last_update",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *languageTableType) NewStruct() reform.Struct {
	return new(Language)
}

// NewRecord makes a new record for that table.
func (v *languageTableType) NewRecord() reform.Record {
	return new(Language)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *languageTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// LanguageTable represents language view or table in SQL database.
var LanguageTable = &languageTableType{
	s: parse.StructInfo{
		Type:    "Language",
		SQLName: "language",
		Fields: []parse.FieldInfo{
			{Name: "LanguageID", Type: "int32", Column: "language_id"},
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "LastUpdate", Type: "time.Time", Column: "last_update"},
		},
		PKFieldIndex: 0,
	},
	z: new(Language).Values(),
}

// String returns a string representation of this struct or record.
func (s Language) String() string {
	res := make([]string, 3)
	res[0] = "LanguageID: " + reform.Inspect(s.LanguageID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "LastUpdate: " + reform.Inspect(s.LastUpdate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Language) Values() []interface{} {
	return []interface{}{
		s.LanguageID,
		s.Name,
		s.LastUpdate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Language) Pointers() []interface{} {
	return []interface{}{
		&s.LanguageID,
		&s.Name,
		&s.LastUpdate,
	}
}

// View returns View object for that struct.
func (s *Language) View() reform.View {
	return LanguageTable
}

// Table returns Table object for that record.
func (s *Language) Table() reform.Table {
	return LanguageTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Language) PKValue() interface{} {
	return s.LanguageID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Language) PKPointer() interface{} {
	return &s.LanguageID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Language) HasPK() bool {
	return s.LanguageID != LanguageTable.z[LanguageTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.LanguageID = pk.
func (s *Language) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = LanguageTable
	_ reform.Struct = (*Language)(nil)
	_ reform.Table  = LanguageTable
	_ reform.Record = (*Language)(nil)
	_ fmt.Stringer  = (*Language)(nil)
)

func init() {
	parse.AssertUpToDate(&LanguageTable.s, new(Language))
}

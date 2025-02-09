// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package pagila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type actorInfoViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *actorInfoViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("actor_info").
func (v *actorInfoViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *actorInfoViewType) Columns() []string {
	return []string{
		"actor_id",
		"first_name",
		"last_name",
		"film_info",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *actorInfoViewType) NewStruct() reform.Struct {
	return new(ActorInfo)
}

// ActorInfoView represents actor_info view or table in SQL database.
var ActorInfoView = &actorInfoViewType{
	s: parse.StructInfo{
		Type:    "ActorInfo",
		SQLName: "actor_info",
		Fields: []parse.FieldInfo{
			{Name: "ActorID", Type: "*int32", Column: "actor_id"},
			{Name: "FirstName", Type: "*string", Column: "first_name"},
			{Name: "LastName", Type: "*string", Column: "last_name"},
			{Name: "FilmInfo", Type: "*string", Column: "film_info"},
		},
		PKFieldIndex: -1,
	},
	z: new(ActorInfo).Values(),
}

// String returns a string representation of this struct or record.
func (s ActorInfo) String() string {
	res := make([]string, 4)
	res[0] = "ActorID: " + reform.Inspect(s.ActorID, true)
	res[1] = "FirstName: " + reform.Inspect(s.FirstName, true)
	res[2] = "LastName: " + reform.Inspect(s.LastName, true)
	res[3] = "FilmInfo: " + reform.Inspect(s.FilmInfo, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ActorInfo) Values() []interface{} {
	return []interface{}{
		s.ActorID,
		s.FirstName,
		s.LastName,
		s.FilmInfo,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ActorInfo) Pointers() []interface{} {
	return []interface{}{
		&s.ActorID,
		&s.FirstName,
		&s.LastName,
		&s.FilmInfo,
	}
}

// View returns View object for that struct.
func (s *ActorInfo) View() reform.View {
	return ActorInfoView
}

// check interfaces
var (
	_ reform.View   = ActorInfoView
	_ reform.Struct = (*ActorInfo)(nil)
	_ fmt.Stringer  = (*ActorInfo)(nil)
)

func init() {
	parse.AssertUpToDate(&ActorInfoView.s, new(ActorInfo))
}

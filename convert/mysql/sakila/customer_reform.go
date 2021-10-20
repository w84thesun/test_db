// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package sakila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type customerTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *customerTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("customer").
func (v *customerTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *customerTableType) Columns() []string {
	return []string{
		"customer_id",
		"store_id",
		"first_name",
		"last_name",
		"email",
		"address_id",
		"active",
		"create_date",
		"last_update",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *customerTableType) NewStruct() reform.Struct {
	return new(Customer)
}

// NewRecord makes a new record for that table.
func (v *customerTableType) NewRecord() reform.Record {
	return new(Customer)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *customerTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// CustomerTable represents customer view or table in SQL database.
var CustomerTable = &customerTableType{
	s: parse.StructInfo{
		Type:    "Customer",
		SQLName: "customer",
		Fields: []parse.FieldInfo{
			{Name: "CustomerID", Type: "int16", Column: "customer_id"},
			{Name: "StoreID", Type: "int8", Column: "store_id"},
			{Name: "FirstName", Type: "string", Column: "first_name"},
			{Name: "LastName", Type: "string", Column: "last_name"},
			{Name: "Email", Type: "*string", Column: "email"},
			{Name: "AddressID", Type: "int16", Column: "address_id"},
			{Name: "Active", Type: "int8", Column: "active"},
			{Name: "CreateDate", Type: "time.Time", Column: "create_date"},
			{Name: "LastUpdate", Type: "*time.Time", Column: "last_update"},
		},
		PKFieldIndex: 0,
	},
	z: new(Customer).Values(),
}

// String returns a string representation of this struct or record.
func (s Customer) String() string {
	res := make([]string, 9)
	res[0] = "CustomerID: " + reform.Inspect(s.CustomerID, true)
	res[1] = "StoreID: " + reform.Inspect(s.StoreID, true)
	res[2] = "FirstName: " + reform.Inspect(s.FirstName, true)
	res[3] = "LastName: " + reform.Inspect(s.LastName, true)
	res[4] = "Email: " + reform.Inspect(s.Email, true)
	res[5] = "AddressID: " + reform.Inspect(s.AddressID, true)
	res[6] = "Active: " + reform.Inspect(s.Active, true)
	res[7] = "CreateDate: " + reform.Inspect(s.CreateDate, true)
	res[8] = "LastUpdate: " + reform.Inspect(s.LastUpdate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Customer) Values() []interface{} {
	return []interface{}{
		s.CustomerID,
		s.StoreID,
		s.FirstName,
		s.LastName,
		s.Email,
		s.AddressID,
		s.Active,
		s.CreateDate,
		s.LastUpdate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Customer) Pointers() []interface{} {
	return []interface{}{
		&s.CustomerID,
		&s.StoreID,
		&s.FirstName,
		&s.LastName,
		&s.Email,
		&s.AddressID,
		&s.Active,
		&s.CreateDate,
		&s.LastUpdate,
	}
}

// View returns View object for that struct.
func (s *Customer) View() reform.View {
	return CustomerTable
}

// Table returns Table object for that record.
func (s *Customer) Table() reform.Table {
	return CustomerTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Customer) PKValue() interface{} {
	return s.CustomerID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Customer) PKPointer() interface{} {
	return &s.CustomerID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Customer) HasPK() bool {
	return s.CustomerID != CustomerTable.z[CustomerTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.CustomerID = pk.
func (s *Customer) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = CustomerTable
	_ reform.Struct = (*Customer)(nil)
	_ reform.Table  = CustomerTable
	_ reform.Record = (*Customer)(nil)
	_ fmt.Stringer  = (*Customer)(nil)
)

func init() {
	parse.AssertUpToDate(&CustomerTable.s, new(Customer))
}

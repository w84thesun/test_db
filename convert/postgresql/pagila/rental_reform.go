// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package pagila

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type rentalTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *rentalTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("rental").
func (v *rentalTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *rentalTableType) Columns() []string {
	return []string{
		"rental_id",
		"rental_date",
		"inventory_id",
		"customer_id",
		"return_date",
		"staff_id",
		"last_update",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *rentalTableType) NewStruct() reform.Struct {
	return new(Rental)
}

// NewRecord makes a new record for that table.
func (v *rentalTableType) NewRecord() reform.Record {
	return new(Rental)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *rentalTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// RentalTable represents rental view or table in SQL database.
var RentalTable = &rentalTableType{
	s: parse.StructInfo{
		Type:    "Rental",
		SQLName: "rental",
		Fields: []parse.FieldInfo{
			{Name: "RentalID", Type: "int32", Column: "rental_id"},
			{Name: "RentalDate", Type: "time.Time", Column: "rental_date"},
			{Name: "InventoryID", Type: "int32", Column: "inventory_id"},
			{Name: "CustomerID", Type: "int32", Column: "customer_id"},
			{Name: "ReturnDate", Type: "*time.Time", Column: "return_date"},
			{Name: "StaffID", Type: "int32", Column: "staff_id"},
			{Name: "LastUpdate", Type: "time.Time", Column: "last_update"},
		},
		PKFieldIndex: 0,
	},
	z: new(Rental).Values(),
}

// String returns a string representation of this struct or record.
func (s Rental) String() string {
	res := make([]string, 7)
	res[0] = "RentalID: " + reform.Inspect(s.RentalID, true)
	res[1] = "RentalDate: " + reform.Inspect(s.RentalDate, true)
	res[2] = "InventoryID: " + reform.Inspect(s.InventoryID, true)
	res[3] = "CustomerID: " + reform.Inspect(s.CustomerID, true)
	res[4] = "ReturnDate: " + reform.Inspect(s.ReturnDate, true)
	res[5] = "StaffID: " + reform.Inspect(s.StaffID, true)
	res[6] = "LastUpdate: " + reform.Inspect(s.LastUpdate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Rental) Values() []interface{} {
	return []interface{}{
		s.RentalID,
		s.RentalDate,
		s.InventoryID,
		s.CustomerID,
		s.ReturnDate,
		s.StaffID,
		s.LastUpdate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Rental) Pointers() []interface{} {
	return []interface{}{
		&s.RentalID,
		&s.RentalDate,
		&s.InventoryID,
		&s.CustomerID,
		&s.ReturnDate,
		&s.StaffID,
		&s.LastUpdate,
	}
}

// View returns View object for that struct.
func (s *Rental) View() reform.View {
	return RentalTable
}

// Table returns Table object for that record.
func (s *Rental) Table() reform.Table {
	return RentalTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Rental) PKValue() interface{} {
	return s.RentalID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Rental) PKPointer() interface{} {
	return &s.RentalID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Rental) HasPK() bool {
	return s.RentalID != RentalTable.z[RentalTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.RentalID = pk.
func (s *Rental) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = RentalTable
	_ reform.Struct = (*Rental)(nil)
	_ reform.Table  = RentalTable
	_ reform.Record = (*Rental)(nil)
	_ fmt.Stringer  = (*Rental)(nil)
)

func init() {
	parse.AssertUpToDate(&RentalTable.s, new(Rental))
}

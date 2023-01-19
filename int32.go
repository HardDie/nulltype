package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int32 struct {
	Data sql.NullInt32
}

func NewInt32(data *int32) Int32 {
	if data == nil {
		return Int32{}
	}
	return Int32{
		Data: sql.NullInt32{
			Int32: *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *Int32) Valid() bool {
	return t.Data.Valid
}
func (t *Int32) Get() int32 {
	return t.Data.Int32
}
func (t *Int32) GetPtr() *int32 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Int32) Reset() {
	t.Data = sql.NullInt32{
		Int32: 0,
		Valid: false,
	}
}
func (t *Int32) Set(data int32) {
	t.Data = sql.NullInt32{
		Int32: data,
		Valid: true,
	}
}

// fmt

func (t *Int32) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Int32)
}

// SQL

func (t *Int32) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Int32) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Int32) UnmarshalJSON(data []byte) error {
	var value *int32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullInt32{
			Int32: 0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullInt32{
		Int32: *value,
		Valid: true,
	}
	return nil
}
func (t *Int32) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Int32)
}

package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int16 struct {
	Data sql.NullInt16
}

func NewInt16(data *int16) Int16 {
	if data == nil {
		return Int16{}
	}
	return Int16{
		Data: sql.NullInt16{
			Int16: *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *Int16) Valid() bool {
	return t.Data.Valid
}
func (t *Int16) Get() int16 {
	return t.Data.Int16
}
func (t *Int16) GetPtr() *int16 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Int16) Reset() {
	t.Data = sql.NullInt16{
		Int16: 0,
		Valid: false,
	}
}
func (t *Int16) Set(data int16) {
	t.Data = sql.NullInt16{
		Int16: data,
		Valid: true,
	}
}

// fmt

func (t *Int16) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Int16)
}

// SQL

func (t *Int16) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Int16) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Int16) UnmarshalJSON(data []byte) error {
	var value *int16
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullInt16{
			Int16: 0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullInt16{
		Int16: *value,
		Valid: true,
	}
	return nil
}
func (t *Int16) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Int16)
}

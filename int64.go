package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int64 struct {
	Data sql.NullInt64
}

func NewInt64(data *int64) Int64 {
	if data == nil {
		return Int64{}
	}
	return Int64{
		Data: sql.NullInt64{
			Int64: *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *Int64) Valid() bool {
	return t.Data.Valid
}
func (t *Int64) Get() int64 {
	return t.Data.Int64
}
func (t *Int64) GetPtr() *int64 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Int64) Reset() {
	t.Data = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
}
func (t *Int64) Set(data int64) {
	t.Data = sql.NullInt64{
		Int64: data,
		Valid: true,
	}
}

// fmt

func (t *Int64) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Int64)
}

// SQL

func (t *Int64) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Int64) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Int64) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullInt64{
		Int64: *value,
		Valid: true,
	}
	return nil
}
func (t *Int64) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Int64)
}

package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Float64 struct {
	Data sql.NullFloat64
}

func NewFloat64(data *float64) Float64 {
	if data == nil {
		return Float64{}
	}
	return Float64{
		Data: sql.NullFloat64{
			Float64: *data,
			Valid:   true,
		},
	}
}

// Methods for the user

func (t *Float64) Valid() bool {
	return t.Data.Valid
}
func (t *Float64) Get() float64 {
	return t.Data.Float64
}
func (t *Float64) GetPtr() *float64 {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Float64) Reset() {
	t.Data = sql.NullFloat64{
		Float64: 0,
		Valid:   false,
	}
}
func (t *Float64) Set(data float64) {
	t.Data = sql.NullFloat64{
		Float64: data,
		Valid:   true,
	}
}

// fmt

func (t *Float64) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Float64)
}

// SQL

func (t *Float64) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Float64) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Float64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullFloat64{
			Float64: 0,
			Valid:   false,
		}
		return nil
	}
	t.Data = sql.NullFloat64{
		Float64: *value,
		Valid:   true,
	}
	return nil
}
func (t *Float64) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Float64)
}

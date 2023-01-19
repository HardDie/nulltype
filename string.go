package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type String struct {
	Data sql.NullString
}

func NewString(data *string) String {
	if data == nil {
		return String{}
	}
	return String{
		Data: sql.NullString{
			String: *data,
			Valid:  true,
		},
	}
}

// Methods for the user

func (t *String) Valid() bool {
	return t.Data.Valid
}
func (t *String) Get() string {
	return t.Data.String
}
func (t *String) GetPtr() *string {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *String) Reset() {
	t.Data = sql.NullString{
		String: "",
		Valid:  false,
	}
}
func (t *String) Set(data string) {
	t.Data = sql.NullString{
		String: data,
		Valid:  true,
	}
}

// fmt

func (t *String) String() string {
	if !t.Valid() {
		return ""
	}
	return t.Data.String
}

// SQL

func (t *String) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *String) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *String) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullString{
			String: "",
			Valid:  false,
		}
		return nil
	}
	t.Data = sql.NullString{
		String: *value,
		Valid:  true,
	}
	return nil
}
func (t *String) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.String)
}

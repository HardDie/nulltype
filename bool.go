package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Bool struct {
	Data sql.NullBool
}

func NewBool(data *bool) Bool {
	if data == nil {
		return Bool{}
	}
	return Bool{
		Data: sql.NullBool{
			Bool:  *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *Bool) Valid() bool {
	return t.Data.Valid
}
func (t *Bool) Get() bool {
	return t.Data.Bool
}
func (t *Bool) GetPtr() *bool {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Bool) Reset() {
	t.Data = sql.NullBool{
		Bool:  false,
		Valid: false,
	}
}
func (t *Bool) Set(data bool) {
	t.Data = sql.NullBool{
		Bool:  data,
		Valid: true,
	}
}

// fmt

func (t *Bool) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Bool)
}

// SQL

func (t *Bool) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Bool) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Bool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullBool{
			Bool:  false,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullBool{
		Bool:  *value,
		Valid: true,
	}
	return nil
}
func (t *Bool) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Bool)
}

package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Byte struct {
	Data sql.NullByte
}

func NewByte(data *byte) Byte {
	if data == nil {
		return Byte{}
	}
	return Byte{
		Data: sql.NullByte{
			Byte:  *data,
			Valid: true,
		},
	}
}

// Methods for the user

func (t *Byte) Valid() bool {
	return t.Data.Valid
}
func (t *Byte) Get() byte {
	return t.Data.Byte
}
func (t *Byte) GetPtr() *byte {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Byte) Reset() {
	t.Data = sql.NullByte{
		Byte:  0,
		Valid: false,
	}
}
func (t *Byte) Set(data byte) {
	t.Data = sql.NullByte{
		Byte:  data,
		Valid: true,
	}
}

// fmt

func (t *Byte) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Byte)
}

// SQL

func (t *Byte) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Byte) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Byte) UnmarshalJSON(data []byte) error {
	var value *byte
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullByte{
			Byte:  0,
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullByte{
		Byte:  *value,
		Valid: true,
	}
	return nil
}
func (t *Byte) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Byte)
}

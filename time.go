package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Time struct {
	Data   sql.NullTime
	Format string
}

func NewTime(data *time.Time) Time {
	if data == nil {
		return Time{
			// Default time format. Copy from format.go file (t Time) String() method
			Format: "2006-01-02 15:04:05.999999999 -0700 MST",
		}
	}
	return Time{
		Data: sql.NullTime{
			Time:  *data,
			Valid: true,
		},
		// Default time format. Copy from format.go file (t Time) String() method
		Format: "2006-01-02 15:04:05.999999999 -0700 MST",
	}
}

func (t Time) SetFormat(format string) Time {
	t.Format = format
	return t
}

// Methods for the user

func (t *Time) Valid() bool {
	return t.Data.Valid
}
func (t *Time) Get() time.Time {
	return t.Data.Time
}
func (t *Time) GetPtr() *time.Time {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Time) Reset() {
	t.Data = sql.NullTime{
		Time:  time.Unix(0, 0),
		Valid: false,
	}
}
func (t *Time) Set(data time.Time) {
	t.Data = sql.NullTime{
		Time:  data,
		Valid: true,
	}
}
func (t *Time) ToString(format string) string {
	if !t.Valid() {
		return ""
	}
	return t.Get().Format(format)
}
func (t *Time) ToStringPtr(format string) *string {
	if !t.Valid() {
		return nil
	}
	val := t.Get().Format(format)
	return &val
}

// Methods for use with proto structures

func NewTimeFromTimestamppb(data *timestamppb.Timestamp) Time {
	if data == nil || !data.IsValid() {
		return NewTime(nil)
	}
	date := data.AsTime()
	return NewTime(&date)
}
func (t *Time) GetTimestamppb() *timestamppb.Timestamp {
	if !t.Valid() {
		return nil
	}
	return timestamppb.New(t.Get())
}

// fmt

func (t *Time) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Time.Format(t.Format))
}

// SQL

func (t *Time) Scan(value interface{}) error {
	return t.Data.Scan(value)
}
func (t *Time) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Time) UnmarshalJSON(data []byte) error {
	var value *time.Time
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		t.Data = sql.NullTime{
			Time:  time.Unix(0, 0),
			Valid: false,
		}
		return nil
	}
	t.Data = sql.NullTime{
		Time:  *value,
		Valid: true,
	}
	return nil
}
func (t *Time) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Time.Format(t.Format))
}

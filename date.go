package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Date struct {
	Data   sql.NullTime
	Format string
}

func NewDate(data *time.Time) Date {
	if data == nil {
		return Date{
			// Default time format. Copy from format.go file (t Time) String() method
			Format: "2006-01-02",
		}
	}
	return Date{
		Data: sql.NullTime{
			Time:  timeToDate(*data),
			Valid: true,
		},
		// Default time format. Copy from format.go file (t Time) String() method
		Format: "2006-01-02",
	}
}

func (t Date) SetFormat(format string) Date {
	t.Format = format
	return t
}

// Methods for the user

func (t *Date) Valid() bool {
	return t.Data.Valid
}
func (t *Date) Get() time.Time {
	return t.Data.Time
}
func (t *Date) GetPtr() *time.Time {
	if !t.Valid() {
		return nil
	}
	val := t.Get()
	return &val
}
func (t *Date) Reset() {
	t.Data = sql.NullTime{
		Time:  time.Unix(0, 0),
		Valid: false,
	}
}
func (t *Date) Set(data time.Time) {
	t.Data = sql.NullTime{
		Time:  timeToDate(data),
		Valid: true,
	}
}
func (t *Date) ToString(format string) string {
	if !t.Valid() {
		return ""
	}
	return t.Get().Format(format)
}
func (t *Date) ToStringPtr(format string) *string {
	if !t.Valid() {
		return nil
	}
	val := t.Get().Format(format)
	return &val
}

// Methods for use with proto structures

func NewDateFromTimestamppb(data *timestamppb.Timestamp) Date {
	if data == nil || !data.IsValid() {
		return NewDate(nil)
	}
	date := data.AsTime()
	return NewDate(&date)
}
func (t *Date) GetTimestamppb() *timestamppb.Timestamp {
	if !t.Valid() {
		return nil
	}
	return timestamppb.New(t.Get())
}

// fmt

func (t *Date) String() string {
	if !t.Valid() {
		return ""
	}
	return fmt.Sprint(t.Data.Time.Format(t.Format))
}

// SQL

func (t *Date) Scan(value interface{}) error {
	err := t.Data.Scan(value)
	if err != nil {
		return err
	}
	if t.Valid() {
		t.Data.Time = timeToDate(t.Data.Time)
	}
	return nil
}
func (t *Date) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.Data.Value()
}

// JSON

func (t *Date) UnmarshalJSON(data []byte) error {
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
		Time:  timeToDate(*value),
		Valid: true,
	}
	return nil
}
func (t *Date) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Data.Time.Format(t.Format))
}

// Utils

func timeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

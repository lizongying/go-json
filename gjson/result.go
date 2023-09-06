package gjson

import (
	"strconv"
	"strings"
)

type Result struct {
	Raw string
}

func NewResult(raw string) (r *Result) {
	raw = strings.TrimSpace(raw)
	r = &Result{
		Raw: raw,
	}
	return
}
func (r *Result) Int() (ii int) {
	if r.Raw == "" {
		return
	}
	ii, _ = strconv.Atoi(r.Raw)
	return
}
func (r *Result) Int8() (i8 int8) {
	if r.Raw == "" {
		return
	}
	i64, err := strconv.ParseInt(r.Raw, 10, 8)
	if err != nil {
		return
	}
	i8 = int8(i64)
	return
}
func (r *Result) Int16() (i16 int16) {
	if r.Raw == "" {
		return
	}
	i64, err := strconv.ParseInt(r.Raw, 10, 16)
	if err != nil {
		return
	}
	i16 = int16(i64)
	return
}
func (r *Result) Int32() (i32 int32) {
	if r.Raw == "" {
		return
	}
	i64, err := strconv.ParseInt(r.Raw, 10, 32)
	if err != nil {
		return
	}
	i32 = int32(i64)
	return
}
func (r *Result) Int64() (i64 int64) {
	if r.Raw == "" {
		return
	}
	i64, _ = strconv.ParseInt(r.Raw, 10, 64)
	return
}
func (r *Result) IntOr(i int) (ii int) {
	if r.Raw == "" {
		ii = i
		return
	}
	ii, _ = strconv.Atoi(r.Raw)
	return
}
func (r *Result) Int8Or(i int8) (i8 int8) {
	if r.Raw == "" {
		i8 = i
		return
	}
	i64, err := strconv.ParseInt(r.Raw, 10, 8)
	if err != nil {
		return
	}
	i8 = int8(i64)
	return
}
func (r *Result) Int16Or(i int16) (i16 int16) {
	if r.Raw == "" {
		i16 = i
		return
	}
	i64, err := strconv.ParseInt(r.Raw, 10, 16)
	if err != nil {
		return
	}
	i16 = int16(i64)
	return
}
func (r *Result) Int32Or(i int32) (i32 int32) {
	if r.Raw == "" {
		i32 = i
		return
	}
	i64, err := strconv.ParseInt(r.Raw, 10, 32)
	if err != nil {
		return
	}
	i32 = int32(i64)
	return
}
func (r *Result) Int64Or(i int64) (i64 int64) {
	if r.Raw == "" {
		i64 = i
		return
	}
	i64, _ = strconv.ParseInt(r.Raw, 10, 64)
	return
}
func (r *Result) Uint() (uu uint) {
	if r.Raw == "" {
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 64)
	if err != nil {
		return
	}
	uu = uint(u64)
	return
}
func (r *Result) Uint8() (u8 uint8) {
	if r.Raw == "" {
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 8)
	if err != nil {
		return
	}
	u8 = uint8(u64)
	return
}
func (r *Result) Uint16() (u16 uint16) {
	if r.Raw == "" {
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 16)
	if err != nil {
		return
	}
	u16 = uint16(u64)
	return
}
func (r *Result) Uint32() (u32 uint32) {
	if r.Raw == "" {
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 32)
	if err != nil {
		return
	}
	u32 = uint32(u64)
	return
}
func (r *Result) Uint64() (u64 uint64) {
	if r.Raw == "" {
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 64)
	if err != nil {
		return
	}
	return
}
func (r *Result) UintOr(u uint) (uu uint) {
	if r.Raw == "" {
		uu = u
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 64)
	if err != nil {
		return
	}
	uu = uint(u64)
	return
}
func (r *Result) Uint8Or(u uint8) (u8 uint8) {
	if r.Raw == "" {
		u8 = u
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 8)
	if err != nil {
		return
	}
	u8 = uint8(u64)
	return
}
func (r *Result) Uint16Or(u uint16) (u16 uint16) {
	if r.Raw == "" {
		u16 = u
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 16)
	if err != nil {
		return
	}
	u16 = uint16(u64)
	return
}
func (r *Result) Uint32Or(u uint32) (u32 uint32) {
	if r.Raw == "" {
		u32 = u
		return
	}
	u64, err := strconv.ParseUint(r.Raw, 10, 32)
	if err != nil {
		return
	}
	u32 = uint32(u64)
	return
}
func (r *Result) Uint64Or(u uint64) (u64 uint64) {
	if r.Raw == "" {
		u64 = u
		return
	}
	u64, _ = strconv.ParseUint(r.Raw, 10, 64)
	return
}
func (r *Result) Bool() (bb bool) {
	if r.Raw == "true" {
		bb = true
		return
	}
	if r.Raw == "True" {
		bb = true
		return
	}
	if r.Raw == "TRUE" {
		bb = true
		return
	}
	return
}
func (r *Result) boolOr(b bool) (bb bool) {
	if r.Raw == "true" {
		bb = true
		return
	}
	if r.Raw == "True" {
		bb = true
		return
	}
	if r.Raw == "TRUE" {
		bb = true
		return
	}
	if r.Raw == "" {
		bb = b
		return
	}
	return
}
func (r *Result) String() (ss string) {
	ss = r.Raw
	return
}
func (r *Result) StringOr(s string) (ss string) {
	if r.Raw == "" {
		ss = s
		return
	}
	return
}
func (r *Result) Float32() (f32 float32) {
	if r.Raw == "" {
		return
	}
	f64, err := strconv.ParseFloat(r.Raw, 32)
	if err != nil {
		return
	}
	f32 = float32(f64)
	return
}
func (r *Result) Float64() (f64 float64) {
	if r.Raw == "" {
		return
	}
	f64, _ = strconv.ParseFloat(r.Raw, 64)
	return
}
func (r *Result) Float32Or(f float32) (f32 float32) {
	if r.Raw == "" {
		f32 = f
		return
	}
	f64, err := strconv.ParseFloat(r.Raw, 32)
	if err != nil {
		return
	}
	f32 = float32(f64)
	return
}
func (r *Result) Float64Or(f float64) (f64 float64) {
	if r.Raw == "" {
		f64 = f
		return
	}
	f64, _ = strconv.ParseFloat(r.Raw, 64)
	return
}

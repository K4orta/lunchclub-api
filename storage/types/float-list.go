package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type FloatList []float64

func (p FloatList) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}
	m, _ := json.Marshal(p)
	replaceLeft := strings.Replace(string(m), "[", "{", 1)
	return strings.Replace(replaceLeft, "]", "}", 1), nil
}

func (p *FloatList) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || v.IsNil() {
		return nil
	}
	if data, ok := src.([]byte); ok {
		replaceLeft := strings.Replace(string(data), "{", `[`, 1)
		replaceRight := strings.Replace(replaceLeft, "}", `]`, 1)
		return json.Unmarshal([]byte(replaceRight), &p)
	}
	return fmt.Errorf("Could not not decode type %T -> %T", src, p)
}

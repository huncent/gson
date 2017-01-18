package gson

import (
	"encoding/json"

	"github.com/najeira/conv"
)

func Decode(s []byte) (*Value, error) {
	v := &Value{}
	err := v.UnmarshalJSON(s)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func DecodeString(s string) (*Value, error) {
	return Decode([]byte(s))
}

type Value struct {
	value interface{}
}

func (v Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(&v.value)
}

func (v *Value) UnmarshalJSON(b []byte) error {
	var obj interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}
	v.value = obj
	return nil
}

func (v Value) IsNull() bool {
	return v.value == nil
}

func (v Value) String() string {
	return conv.String(v.value)
}

func (v Value) Int() int64 {
	return conv.Int(v.value)
}

func (v Value) Float() float64 {
	return conv.Float(v.value)
}

func (v Value) Bool() bool {
	return conv.Bool(v.value)
}

func (v Value) Map() map[string]interface{} {
	switch d := v.value.(type) {
	case map[string]interface{}:
		return d
	}
	return nil
}

func (v Value) Array() []interface{} {
	switch d := v.value.(type) {
	case []interface{}:
		return d
	}
	return nil
}

func (v Value) Get(names ...string) Value {
	var cur interface{} = v.value
	for _, name := range names {
		mp, ok := cur.(map[string]interface{})
		if !ok {
			return Value{}
		}
		next, ok := mp[name]
		if !ok {
			return Value{}
		}
		cur = next
	}
	return Value{cur}
}

func (v Value) GetAt(index int) Value {
	arr, ok := v.value.([]interface{})
	if !ok {
		return Value{}
	} else if len(arr) <= index {
		return Value{}
	}
	return Value{arr[index]}
}

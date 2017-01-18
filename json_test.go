package gson

import (
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	s := `{"users": [{"id": 123, "name": "Alice"}]}`

	// decode
	v, err := DecodeString(s)
	if err != nil {
		t.Error(err)
		return
	} else if v == nil {
		t.Error("DecodeString() retuns nil")
		return
	}

	// v should be a Map
	if mp := v.Map(); mp == nil {
		t.Error("Map() retuns nil")
	} else if len(mp) != 1 {
		t.Error("invalid length")
	}

	// v should have users
	users := v.Get("users")
	if users.IsNull() {
		t.Error("Get() retuns null object")
		return
	} else if users.value == nil {
		t.Error("Get() retuns null object")
		return
	}

	// users should be Array
	if arr := users.Array(); arr == nil {
		t.Error("Array() retuns nil")
	} else if len(arr) != 1 {
		t.Error("invalid length")
	}

	// get a value
	user := users.GetAt(0)
	if user.IsNull() {
		t.Error("GetAt() retuns null object")
		return
	}

	// user should be a Map
	if userMap := user.Map(); err != nil {
		t.Error(err)
	} else if userMap == nil {
		t.Error("Map() retuns nil")
	}

	userId := user.Get("id").Int()
	if userId != 123 {
		t.Error("invalid id")
	}

	userName := user.Get("name").String()
	if userName != "Alice" {
		t.Error("invalid name")
	}
}

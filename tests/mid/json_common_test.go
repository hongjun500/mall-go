// @author hongjun500
// @date 2023/6/9 13:50
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package mid

import (
	"github.com/hongjun500/mall-go/pkg/convert"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonCommon(t *testing.T) {

	jsonString := `{"name":"hongjun500","age":18}`
	t.Log(jsonString)
	user := new(User)
	err := convert.JsonToStruct(jsonString, user)
	if err != nil {
		return
	}
	assert.NotEmpty(t, user)
	t.Log(user)

	jsonToMap, err := convert.JsonToMap(jsonString)
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, jsonToMap, "name")
	sliceStruct, err := convert.JsonToSliceStruct(jsonString, []User{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, sliceStruct)

	// jsonarr := `[{"name":"hongjun500","age":18},{"name":"hongjun502","age":25}]`
}

// @author hongjun500
// @date 2023/6/9 13:50
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package common

import (
	"testing"

	"github.com/hongjun500/mall-go/pkg/convert"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	user = User{
		Name: "hongjun500",
		Age:  18,
	}
	userMap = map[string]any{
		"name": "hongjun500",
		"age":  "18",
		/*1:      2,
		2.0:    user,*/
	}

	users = []User{
		user, user,
	}

	usersMap = []map[string]any{
		userMap, userMap,
		{
			"name": "hongjun502",
			"age":  "25",
		},
	}
)

func TestAnyToJson(t *testing.T) {
	userStr := convert.AnyToJson(user)
	assert.NotEmpty(t, userStr)
	mapStr := convert.AnyToJson(userMap)
	assert.NotEmpty(t, mapStr)
	usersStr := convert.AnyToJson(users)
	assert.NotEmpty(t, usersStr)
	usersMapStr := convert.AnyToJson(usersMap)
	assert.NotEmpty(t, usersMapStr)
}

func TestJsonToAny(t *testing.T) {
	user1 := new(User)
	err := convert.JsonToAny(convert.AnyToJson(user), user1)
	assert.Nil(t, err)
	assert.Equal(t, user.Name, user1.Name)
	var userMap1 map[string]any
	err = convert.JsonToAny(convert.AnyToJson(userMap), &userMap1)
	assert.Nil(t, err)
	assert.Equal(t, len(userMap), len(userMap1))
	var users1 []User
	err = convert.JsonToAny(convert.AnyToJson(users), &users1)
	assert.Nil(t, err)
	assert.Equal(t, len(users), len(users1))
	var usersMap1 []map[string]any
	err = convert.JsonToAny(convert.AnyToJson(usersMap), &usersMap1)
	assert.Nil(t, err)
	assert.Equal(t, len(usersMap), len(usersMap1))
}

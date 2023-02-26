package main

import (
	"fmt"
	"reflect"
)

type Favor struct {
	Name string
	ID   int64
}

type User struct {
	Name   string
	Age    int
	Favors []Favor
}

type IUserService interface {
	GetUser(userID int64) User
}

type UserService struct {
}

func (u *UserService) GetUser(userID int64) User {
	user := User{
		Name:   "defaultTestUser",
		Age:    -1,
		Favors: []Favor{},
	}
	fmt.Printf("GetUser exec: result = %#v\n", user)
	return user
}

func TestFunc(names string) {
	fmt.Printf("TestFunc exec: result = %v\n", names)
}

func reflectFuncAndMethod() {
	// 调用函数
	f := reflect.ValueOf(TestFunc)
	args := []reflect.Value{
		reflect.ValueOf("pbrong"),
	}
	f.Call(args)
	// 调用方法
	us := reflect.ValueOf(&UserService{})
	args = []reflect.Value{
		reflect.ValueOf(int64(123)),
	}
	results := us.MethodByName("GetUser").Call(args)
	for _, r := range results {
		fmt.Printf("GetUser resp: %#v", r.Interface())
	}

}

func reflectSetValue() {
	u := User{
		Name: "arong",
		Age:  23,
	}
	fmt.Printf("原始字段值:%#v\n", u)

	// 一定要取指针，不然无法赋值
	rval := reflect.ValueOf(&u).Elem()
	rvalNameField := rval.FieldByName("Name")
	// 字段是否可写入
	if rvalNameField.CanSet() {
		rvalNameField.Set(reflect.ValueOf("pbrong"))
	}

	fmt.Printf("改变已知字段值:%#v\n", u)
}

func rangeStruct() {
	u := User{
		Name: "arong",
		Age:  23,
		Favors: []Favor{
			{
				Name: "篮球",
				ID:   1,
			},
			{
				Name: "唱跳",
				ID:   2,
			},
			{
				Name: "RAP",
				ID:   3,
			},
		},
	}
	rval := reflect.ValueOf(&u)
	if rval.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}
	for i := 0; i < rval.NumField(); i++ {
		name := rval.Type().Field(i).Name
		val := rval.Field(i).Interface()
		fmt.Printf("字段%v值为%v\n", name, val)
	}
}

func reflectValue() {
	var (
		intType    int                    = 10
		stringType string                 = "hello world"
		mapType    map[string]interface{} = nil
	)
	var types []interface{}
	types = append(types, intType)
	types = append(types, stringType)
	types = append(types, mapType)
	// 类型判断
	for i, v := range types {
		rval := reflect.ValueOf(v)
		if rval.Kind() == reflect.Ptr {
			rval = rval.Elem()
		}
		fmt.Printf("(%v)的值%v\n", i, rval.Interface())
	}
}

func reflectType() {
	var (
		intType       int
		stringType    *string
		mapType       map[string]interface{}
		structType    *User
		interfaceType IUserService
	)
	var types []interface{}
	types = append(types, intType)
	types = append(types, stringType)
	types = append(types, mapType)
	types = append(types, structType)
	types = append(types, interfaceType)

	// 类型判断
	for i, v := range types {
		rtyp := reflect.TypeOf(v)
		if rtyp == nil {
			fmt.Printf("(%v)类型获取Type为nil，跳过\n", i)
			continue
		}
		rkind := rtyp.Kind()
		// 如果为指针类型，还原出真实类型
		if rtyp.Kind() == reflect.Ptr {
			fmt.Printf("(%v)指针类型: %v\n", i, rkind.String())
			rtyp = rtyp.Elem()
		}
		rkind = rtyp.Kind()
		fmt.Printf("(%v)类型%v\n", i, rkind.String())
	}
}

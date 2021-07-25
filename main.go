package main

import (
	"fmt"
	"reflect"
	"strings"
)

type A struct {
	B string
	C int
	D bool
	E struct{
		F string
		G bool
		H struct{
			I int
			J float32
			K struct{
				L string
			}
		}
	}
}

type StructGetType struct {
	Name string
	Type reflect.Type
	Value reflect.Value
}

func linearSearch(datalist []string, key string) int {
	count := 0
	for _, item := range datalist {
		if item == key {
			count++
		}
	}
	return count
}

func getVar(varstruct reflect.Value) []StructGetType {
	var listStructGetType []StructGetType
	elems := varstruct.Elem()
	for i := 0; i < elems.NumField(); i++ {
		varName := elems.Type().Field(i).Name
		varType := elems.Type().Field(i).Type
		varValue := elems.Field(i)
		//fmt.Printf("%v %v %v\n", varName,varType,varValue)
		structGetType := StructGetType{
			Name:  varName,
			Type:  varType,
			Value: varValue,
		}
		listStructGetType = append(listStructGetType, structGetType)
	}
	return listStructGetType
}

func getElemVars(structType StructGetType) []StructGetType {
	var listStructGetType []StructGetType
	for i := 0; i < structType.Value.NumField(); i++ {
		varName := structType.Value.Type().Field(i).Name
		varType := structType.Value.Type().Field(i).Type
		varValue := structType.Value.Field(i)
		//fmt.Printf("%v %v %v\n", varName,varType,varValue)
		structGetType := StructGetType{
			Name:  structType.Name + "." + varName,
			Type:  varType,
			Value: varValue,
		}
		listStructGetType = append(listStructGetType, structGetType)
	}
	return listStructGetType
}


func (a *A) GetValue(input string)  {
	inputSlice := strings.Split(input, ".")
	fmt.Println(inputSlice)

	//fmt.Println(getVar(a))

	for _, structType := range getVar(reflect.ValueOf(a)) {
		fmt.Println(structType.Name, structType.Type.String())

		if strings.HasPrefix(structType.Type.String(), "struct") {
			//splitStr := strings.Split(structType.Type.String(), " ")
			//fmt.Println(linearSearch(splitStr, "struct"))

			for i := 0; i < 3; i++ {
				fmt.Println(getElemVars(structType))
			}

			//elems := reflect.ValueOf(structType.Value).Elem()
			//
			//fmt.Println(elems)

			//fmt.Println(getVar(structType.Value.Elem()))
		}
	}
}

func main()  {
	a := A{
		B: "Hello",
		C: 0,
		D: false,
		E: struct {
			F string
			G bool
			H struct {
				I int
				J float32
				K struct{ L string }
			}
		}{F: "New", G: false, H: struct {
			I int
			J float32
			K struct{ L string }
		}{I: 0, J: 1, K: struct{ L string }{L: "str"}}},
	}

	//marshal, err := json.Marshal(a)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(marshal))

	//val := reflect.Indirect(reflect.ValueOf(a))
	//fmt.Println(val.Field(0).Type().Name())

	a.GetValue("E")
}
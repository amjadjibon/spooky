package main

import (
	"fmt"
	"reflect"
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

	M map[string]interface{}
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

var TypeMap = make(map[string]interface{})

func getStructPrintRec(structType StructGetType) {
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

		if structGetType.Type.Kind() != reflect.Struct {
			TypeMap[structGetType.Name] = structGetType.Value.Interface()
			fmt.Println(structGetType.Name, structGetType.Type.Kind(), structGetType.Value.Interface())
		}

		if structGetType.Type.Kind() == reflect.Struct {
			//fmt.Println(structGetType.Type.Kind())
			getStructPrintRec(structGetType)
		} else {

		}
	}
}

func getStructPrint(structType StructGetType) {
	for i := 0; i < structType.Value.NumField(); i++ {
		varName := structType.Value.Type().Field(i).Name
		varType := structType.Value.Type().Field(i).Type
		varValue := structType.Value.Field(i)
		//fmt.Printf("%v %v %v\n\n", varName,varType,varValue)
		structGetType := StructGetType{
			Name:  structType.Name + "." + varName,
			Type:  varType,
			Value: varValue,
		}
		fmt.Println(structType.Name ,structGetType.Type.String())
	}
}


func (a *A) GetValue(input string) interface{} {
	for _, structType := range getVar(reflect.ValueOf(a)) {
		if structType.Type.Kind() != reflect.Struct {
			TypeMap[structType.Name] = structType.Value.Interface()
			fmt.Println(structType.Name, structType.Type.Kind(), structType.Value.Interface())
		}
		if structType.Type.Kind() == reflect.Struct {
			getStructPrintRec(structType)
		}
	}

	return TypeMap[input]
}

func getMapStruct(structType StructGetType)  {
	for _, element := range structType.Value.MapKeys() {
		structGetType := StructGetType{
			Name:  structType.Name + "." + element.String(),
			Type:  structType.Value.MapIndex(element).Type(),
			Value: structType.Value.MapIndex(element),
		}

		fmt.Println(structGetType.Name, structGetType.Type.Kind(), structGetType.Value.Interface())
		TypeMap[structGetType.Name] = structGetType.Value.Interface()
		//
		////fmt.Println(structGetType.Name, structGetType.Type.Kind())
		//
		switch structGetType.Value.Interface().(type) {
		case map[string]interface{}:
			//fmt.Println(t)
			fmt.Println(reflect.TypeOf(structGetType.Value))

			//getMapStruct(structGetType)
		default:
			TypeMap[structGetType.Name] = structGetType.Value.Interface()
		}
	}
}


func GetValue2(a *A, input string) interface{} {
	for _, structType := range getVar(reflect.ValueOf(a)) {
		if structType.Type.Kind() != reflect.Struct && structType.Type.Kind() != reflect.Map {
			TypeMap[structType.Name] = structType.Value.Interface()
			fmt.Println(structType.Name, structType.Type.Kind(), structType.Value.Interface())
		}
		if structType.Type.Kind() == reflect.Struct {
			getStructPrintRec(structType)
		}

		if structType.Type.Kind() == reflect.Map {
			//getMapStruct(structType)
			fmt.Println(structType.Value.Interface())
		}

		//switch structType.Value.Interface().(type) {
		//case map[string]interface{}:
		//	getMapStruct(structType)
		//}
	}

	return TypeMap[input]
}

func main()  {
	a := A{
		B: "Hello",
		C: 2,
		D: false,
		E: struct {
			F string
			G bool
			H struct {
				I int
				J float32
				K struct{ L string }
			}
		}{F: "Go", G: false, H: struct {
			I int
			J float32
			K struct{ L string }
		}{I: 7, J: 1.5, K: struct{ L string }{L: "World"}}},
		M: map[string]interface{}{
			"key1" : "value1",
			"key2" : "value2",
			"key3" : "value3",

			"key4" : 10,
			"key5" : false,

			"key6" : map[string]interface{} {
				"key7" : "value7",
			},
		},
	}

	fmt.Println(GetValue2(&a,"M"))
}
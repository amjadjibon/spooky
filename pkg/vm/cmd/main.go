package main

import (
	"fmt"
	"github.com/amjadjibon/spooky/pkg/vm"
)



type Location struct {
	Lat  string
	Long string
}

type Payment struct {
	Amount string
	Date   string
}

type Data struct {
	Method   string
	Code     int
	Complete bool
	Payment *Payment
}

type InputModel struct {
	Id       int
	Name     string
	Age      int
	Location *Location
	Data     *Data
}


func main() {
	inputModel := InputModel{
		Id:       1,
		Name:     "Jibon",
		Age:      32,
		Location: &Location{
			Lat:  "12",
			Long: "13",
		},
		Data:     &Data{
			Method:   "Post",
			Code:     400,
			Complete: false,
			Payment:  &Payment{
				Amount: "100",
				Date:   "21-2-2020",
			},
		},
	}
	fmt.Println(vm.GetModelKeyValue(inputModel, "Id"))
}

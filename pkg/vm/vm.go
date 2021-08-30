package vm

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

var KeyValueMap = make(map[string]interface{})

func GetValues(inputModel interface{}, keyStruct string) {
	mapInputModel := make(map[string]interface{})
	inputModelBytes, err := json.Marshal(inputModel)
	if err != nil {
		return
	}
	err = json.Unmarshal(inputModelBytes, &mapInputModel)
	if err != nil {
		return
	}
	for key, val := range mapInputModel {
		if reflect.ValueOf(val).Type().Kind() == reflect.Map {
			GetValues(val, keyStruct+"."+key)
		} else {
			storeKeys := keyStruct + "." + key
			storeKeys = strings.TrimPrefix(storeKeys, ".")
			KeyValueMap[storeKeys] = val
		}
	}
}

func GetModelKeyValue(inputModel interface{}, key string) string {
	GetValues(inputModel, "")
	val, found := KeyValueMap[key]
	if !found {
		log.Warn("key not found")
		return ""
	}
	return fmt.Sprintf("%v", val)
}

//=========================vm===================

type InputValidator struct {
	InputModel map[string]interface{}
	Key        string
	Required   bool
	Regex      string
}

func (i *InputValidator) GetInputModel() map[string]interface{} {
	return i.InputModel
}

func (i *InputValidator) GetKey() string {
	return i.Key
}

func (i *InputValidator) GetRequired() bool {
	return i.Required
}

func (i *InputValidator) GetRegex() string {
	return i.Regex
}

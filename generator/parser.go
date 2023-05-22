package generator

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Schema struct {
	Name         string
	ApiFields    []ApiField    `json:"apiFields"`
	CustomFields []CustomField `json:"customFields"`
}

type ApiField struct {
	FieldId  string `json:"fieldId"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	Required bool   `json:"required"`
}

type CustomField struct {
}

func GenerateSchema(dname string, fname string) Schema {

	barray, _ := ioutil.ReadFile(dname + "/" + fname)

	var schema Schema
	schema.Name = strings.Replace(fname, ".json", "", -1)

	json.Unmarshal(barray, &schema)

	return schema

}

func GenerateSchemas(dname string) []Schema {

	files, _ := ioutil.ReadDir(dname)

	var schemas []Schema

	for _, file := range files {
		schemas = append(schemas, GenerateSchema(dname, file.Name()))
	}

	return schemas

}

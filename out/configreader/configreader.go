package goMysqlConnect

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type jsonobject struct {
	Object ObjectType
}

type ObjectType struct {
	Buffer_size int
	Databases   []DatabasesType
}

type DatabasesType struct {
	Host   string
	User   string
	Pass   string
	Type   string
	Name   string
	Tables []TablesType
}

type TablesType struct {
	Name     string
	Statment string
	Regex    string
	Types    []TypesType
}

type TypesType struct {
	Id    string
	Value string
}


func ReadConfig(file string) {
	var dbconn jsonobject

	configdata,_ := ioutil.ReadFile(file)
	json.Unmarshal(configdata,&dbconn)
	fmt.Print(string(dbconn.Object.Databases[0].Host))

}
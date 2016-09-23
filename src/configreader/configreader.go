package configreader

import (

	"io/ioutil"
	"encoding/json"
)
import "os/exec"

type Jsonobject struct {
	Object ObjectType
}

type ObjectType struct {
	Buffer_size int
	DumpExecutable string
	ClientExecutable string
	Databases   []DatabasesType
}

type DatabasesType struct {
	Host   string
	User   string
	Pass   string
	Type   string
	Name   string
	Port string
	DumpSchema []string
	DumpOption []string
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


func ReadConfig(file string) Jsonobject{
	var dbconn Jsonobject

	configdata,_ := ioutil.ReadFile(file)
	json.Unmarshal(configdata,&dbconn)
return dbconn

}
func SetMysqlExe(dbclient string) string{
	m,_:=exec.LookPath(dbclient)
return m


}
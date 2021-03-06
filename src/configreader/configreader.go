package configreader

import (

	"io/ioutil"
	"encoding/json"
)
import (
	"os/exec"
	"fmt"
	"os"
)

type Jsonobject struct {
	Object ObjectType
}

type ObjectType struct {
	Buffer_size int
	Executables [] OSExecutable
	Databases   [] DatabasesType
}
type OSExecutable struct {
	OS string
	Mysqldump string
	Myslclient string
	TargetDir string

}

type DatabasesType struct {
	Host   string
	User   string
	Pass   string
	Type   string
	Name   string
	Port string
	DumpSchema []string
	DumpOptions []string
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

func checkError(err error) {

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func ReadConfig(file string) Jsonobject{
	var dbconn Jsonobject

	configdata,err := ioutil.ReadFile(file)
	checkError(err)
	json.Unmarshal(configdata,&dbconn)
return dbconn

}
func SetMysqlExe(dbclient string) string{
	m,_:=exec.LookPath(dbclient)
return m


}
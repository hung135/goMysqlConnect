package main

import (
	"configreader"

	//"mysqldumper"
	"fmt"
	"mysqldumper"
	"runtime"
	"os"
)


func main() {




	var mysqlDumpPath string
	var TargetPath string

	argsWithProg := os.Args
	var jsonPath string
	if (len(argsWithProg)==1){
		jsonPath="../src/dbconn.json"
	} else{
		jsonPath=os.Args[1]
	}

	f := configreader.ReadConfig(jsonPath)
	if runtime.GOOS == "windows" {

		fmt.Print(f.Object.Executables[1].Mysqldump)
		mysqlDumpPath =f.Object.Executables[1].Mysqldump
		TargetPath=f.Object.Executables[1].TargetDir
	} else {
		fmt.Print(f.Object.Executables[0].Mysqldump)
		mysqlDumpPath =f.Object.Executables[0].Mysqldump
		TargetPath=f.Object.Executables[0].TargetDir
	}



	cmdName := configreader.SetMysqlExe(mysqlDumpPath)
	//fmt.Print(configreader.SetMysqlExe("mysqldump.exe"))

	mysqldumper.DumpFile(f,cmdName,TargetPath)



}

package main

import (
	"configreader"

	//"mysqldumper"
	"fmt"
	"mysqldumper"
	"runtime"
)


func main() {




	var mysqlDumpPath string
	var TargetPath string
	f := configreader.ReadConfig("../src/dbconn.json")
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
		fmt.Print(f.Object.Executables[1].Mysqldump)
		mysqlDumpPath =f.Object.Executables[1].Mysqldump
		TargetPath=f.Object.Executables[1].TargetDir
	} else {fmt.Println("Hello from Linux")
		fmt.Print(f.Object.Executables[0].Mysqldump)
		mysqlDumpPath =f.Object.Executables[0].Mysqldump
		TargetPath=f.Object.Executables[0].TargetDir
	}



	cmdName := configreader.SetMysqlExe(mysqlDumpPath)
	//fmt.Print(configreader.SetMysqlExe("mysqldump.exe"))

	mysqldumper.DumpFile(f,cmdName,TargetPath)



}

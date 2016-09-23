package main

import (
	"configreader"

	"mysqldumper"
)

func main() {
	f := configreader.ReadConfig("dbconn.json")
	//fmt.Print(f.Object.Databases[0].Host)
	dmpExecutable := configreader.SetMysqlExe(f.Object.DumpExecutable)
	//fmt.Print(configreader.SetMysqlExe("mysqldump.exe"))

	mysqldumper.DumpFile(f, dmpExecutable)


}

package mysqldumper

import (
	"configreader"
	"fmt"
	"os/exec"
	"os"
	"time"
)



type MySQL struct {
	// DB Host (e.g. 127.0.0.1)
	Host     string
	// DB Port (e.g. 3306)
	Port     string
	// DB Name
	DB       string
	// DB User
	User     string
	// DB Password
	Password string
	// Extra mysqldump options
	// e.g []string{"--extended-insert"}
	Options  []string
}

func (x MySQL) dumpOptions() []string {
	var options[] string
	//options := x.Options
	options = append(options, fmt.Sprintf(`-h%v`, x.Host))
	options = append(options, fmt.Sprintf(`-P%v`, x.Port))
	options = append(options, fmt.Sprintf(`-u%v`, x.User))
	options = append(options, fmt.Sprintf(`-p%v`, x.Password))

	lenOpt:=len(x.Options)
	for ii:=0;ii<lenOpt;ii++{
		options = append(options, x.Options[ii])
	}

	options = append(options, x.DB)
	//fmt.Println("\n",options ,"xxxxx \n")

	return options
}
func checkError(err error) {

	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(0)
	}
}
func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	checkError(err)
}



func DumpFile(f configreader.Jsonobject, cmdName string,TargetPath string) {


	num_schemas := len(f.Object.Databases[0].DumpSchema)
	num_db:=len(f.Object.Databases)

	//var x MySQL
	for ii := 0; ii < num_db; ii++ {
		for jj := 0; jj < num_schemas; jj++ {
			x := MySQL{Host:f.Object.Databases[ii].Host,
				User:f.Object.Databases[ii].User,
				Password:f.Object.Databases[ii].Pass,
				Port:f.Object.Databases[ii].Port,
				DB:f.Object.Databases[ii].DumpSchema[jj],
				Options:f.Object.Databases[ii].DumpOptions}
			ServerName:=f.Object.Databases[ii].Name
			FullPath:=TargetPath+"/"+ServerName

			t:=time.Now()

			filename := (f.Object.Databases[ii].DumpSchema[jj]+ t.Format("20060102")  +".sql")


			//dumpPath := fmt.Sprintf(`bu_%v_%v.sql`, x.DB, time.Now().Unix())
			dumpPath:=FullPath+"/"+filename

			os.MkdirAll(FullPath,os.ModePerm)
			deleteFile(dumpPath)
			options := append(x.dumpOptions(), fmt.Sprintf(`-r%v`, dumpPath))
			fmt.Println("Dumping",x.DB,"-->",options)
			 dmpout, err := exec.Command(cmdName, options...).Output()
			if err != nil {
				fmt.Print("error occured dumping schema:",x.DB, dmpout)
				deleteFile(dumpPath)
			}
		}
	}

}

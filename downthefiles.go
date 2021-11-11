package downfiles

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

//get files in windows download folder
func DownloadFiles() []fs.FileInfo {
	u, _ := os.UserHomeDir()
	//assume download folder is in user folder
	files, err := ioutil.ReadDir(string(u) + `\Downloads`)
	if err != nil {
		//assign download folder to D:\Downloads if not found in user folder
		fmt.Println(`D:\Downloads`)
		files, err = ioutil.ReadDir(`D:\Downloads`)
	} else {
		fmt.Println(string(u) + `\Downloads`)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(files)
	return files
}

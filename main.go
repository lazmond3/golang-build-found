package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func mylog(sts ...interface{}) {
	if os.Getenv("LOGLEVEL") == "DEBUG" {
		fmt.Println(sts...)
	}
}



func checkDir(dirPath string) {
	mylog("now is : ", dirPath, " moved to the location! ")
	os.Chdir(dirPath)
	listupDir, _ := ioutil.ReadDir(".")

	for _, f := range listupDir {
		mylog("  listup: f: ", absPath(f.Name()))
	}

	for _, filesIn := range listupDir {
		absP := absPath(filesIn.Name())

		if filesIn.IsDir() == false {
			continue
		}

		if filesIn.Name() == "build" {
			println("FOUND BUILD!: ", absP)
			continue
		}

		var wg = &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			checkDir(absP)
		}()
		wg.Wait()
	}
}


func main() {
	mylog("--- before checkdir")
	checkDir(".")
	mylog("--- after  checkdir")
}

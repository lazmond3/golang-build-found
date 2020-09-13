package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func mylog(sts ...interface{}) {
	if os.Getenv("LOGLEVEL") == "DEBUG" {
		fmt.Println(sts...)
	}
}

func checkDir(dirPath string) {
	mylog("now is : ", dirPath)
	listupDir, _ := ioutil.ReadDir(absPath(dirPath))

	for _, filesIn := range listupDir {
		absP := filepath.Join(dirPath, filesIn.Name())

		if filesIn.IsDir() == false {
			continue
		}

		if filesIn.Name() == "build" {
			println("FOUND BUILD!: ", absP)
			continue
		}

		wg := &sync.WaitGroup{}
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
	checkDir(absPath("."))
	mylog("--- after  checkdir")
}

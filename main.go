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
	// give dir name in dirPath

	mylog("now is : ", dirPath, " moved to the location! ")
	os.Chdir(dirPath)
	listupDir, _ := ioutil.ReadDir(dirPath)

	for _, f := range listupDir {
		mylog("           listup: f: ", absPath(f.Name()))
	}

	for _, filesIn := range listupDir {
		absP := absPath(filesIn.Name())

		if filesIn.IsDir() == false {
			//mylog(absP, " is not a directory.")
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
			checkDir(filesIn.Name())
		}()
		wg.Wait()
	}
}


func main() {

	mylog("--- before checkdir")
	checkDir(".")
	mylog("--- after  checkdir")

	//dir, _ := os.Getwd()
	//fpathList := []string{}
	//
	//fpathList = append(fpathList, dir)
	//for {
	//	if len(fpathList) == 0 {
	//		break
	//	}
	//	lastEl := fpathList[len(fpathList)-1]
	//	fpathList = fpathList[:len(fpathList)-1]
	//
	//
	//	filesInfos, _ := ioutil.ReadDir(lastEl)
	//
	//	for _, f := range filesInfos {
	//		mylog(f.Name(), f.IsDir(), fmt.Sprint(f.ModTime()), f.Size(), f.Sys())
	//		st, _ := filepath.Abs(f.Name())
	//		mylog("absPath: ", st)
	//
	//
	//		if f.IsDir() {
	//			absPath, _ := filepath.Abs(f.Name())
	//
	//			if f.Name() == "build" {
	//				//os.RemoveAll(absPath)
	//				println("found build! : ", absPath)
	//				continue
	//			}
	//
	//			fpathList = append(fpathList, absPath)
	//		}
	//	}
	//
	//}



}

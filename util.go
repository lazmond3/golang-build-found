package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func absPath(st string) string {
	p, _ := filepath.Abs(st)
	return p
}

func readDir(path string) []os.FileInfo {
	files, _ := ioutil.ReadDir(path)
	return files
}
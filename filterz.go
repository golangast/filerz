package filerz

import (
	"fmt"
	"os"
	"strings"
)

func Findext(s string) string {
	idx := strings.Index(s, ".")
	if idx != -1 {
		return s[idx:]
	}
	return s
}
func removefile(s string, file string) string {
	if strings.Contains(s, file) {
		idx := strings.Index(s, file)
		if idx != -1 {
			return s[:idx]
		}
		return s

	} else {
		idx := strings.Index(s, file)
		if idx != -1 {
			return s[idx:]
		}
		return s

	}

}

func SeparateDirFile(s string) (string, string) {

	idx := strings.Index(s, ".")
	if idx != -1 {
		slashes := strings.Count(s[:idx], "/")

		str := strings.Split(s[:idx], "/")

		file := str[slashes] + Findext(s)
		dir := removefile(s, file)

		return dir, file

	}
	return s, ""
}
func Makefolder(p string) {
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
	} else {
		fmt.Println("Directory " + p + " successfully created with sticky bits and full permissions")
	}
}

//make any file
func Makefile(p string) *os.File {
	file, err := os.Create(p)
	if IsError(err) {
		fmt.Println("error -", err, file)
	}
	return file
}

//error checker
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

package utils

import (
	"fmt"
	"os"
)

func JudgeTheFileExist(fp string) (file *os.File) {
	// judge the file exists or not
	exist, err := os.Stat(fp)
	if err != nil {
		fmt.Println(err)
	}
	if exist != nil {
		// if the file exists, remove it
		os.Remove(fp)
	}
	// create the file if it doesn't exist
	file, _ = os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0644)
	return file
}

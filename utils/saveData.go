package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// ServerStatusWatcher saves the server data to the file
func ServerStatusWatcher(data interface{}, fp string, mutex *sync.Mutex) {
	//file := JudgeTheFileExist(fp)
	for {
		time.Sleep(10)
		mutex.Lock()
		SaveRequest(data, fp)
		mutex.Unlock()
	}
}

// SaveRequest saves the request to the file
func SaveRequest(rq interface{}, fp string) {
	// save the request (a struct) to the file request.txt
	// the format is like this:
	// HandleType:0,TargetID:1,TargetStatus:2
	//file := JudgeTheFileExist(fp)
	jsonRq, err := json.Marshal(rq)
	if err != nil {
		fmt.Println(err)
	} else {
		// write the json to the data.txt
		file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println(err)
		}
		file.Write(jsonRq)
		file.WriteString("\n")
		file.Close()
	}
}

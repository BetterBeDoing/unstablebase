package main

import (
	"fmt"
	"time"

	"UnstableBase/appcomplement"
	"UnstableBase/config"
	"UnstableBase/request"
	"UnstableBase/server"
	"UnstableBase/utils"
)

func main() {
	go request.GenerateRequest(10)
	go server.RequestHandler(config.ServerMutex)
	//Start a goroutine to save the server data to the file
	go utils.ServerStatusWatcher(config.ServerData, config.ServerDataFile, config.ServerMutex)
	// try to handle the request
	go appcomplement.AppHandler()
	go appcomplement.Watcher()
	for {
		time.Sleep(10)
		fmt.Println(config.ServerData)
	}
}

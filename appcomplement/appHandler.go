package appcomplement

import (
	"fmt"
	"time"

	cfg "UnstableBase/config"
	"UnstableBase/utils"
)

func AppHandler() {
	// Handle the request
	for {
		select {
		case rq := <-cfg.AppChan:
			// check if the request has been handled
			if ok, err := checkRepeated(rq); ok {
				cfg.RequestMap[rq.TargetID] = rq
				cfg.RequestChannel <- rq
				utils.JudgeTheFileExist(cfg.AppRequestFile)
				utils.SaveRequest(rq, cfg.AppRequestFile)
			} else {
				// if the request has been handled, send the response to the client
				fmt.Println(err)
				// do nothing because repeated
			}
		default:
			time.Sleep(10)
		}
	}
}

func checkRepeated(rq cfg.Request) (bool, error) {
	if _, ok := cfg.RequestMap[rq.TargetID]; ok {
		return false, ErrRepeated
	}
	return true, nil
}

// Watcher watches the request finished or not
func Watcher() {
	for {
		time.Sleep(10)
		for key, rq := range cfg.RequestMap {
			if CheckRequest(rq) {
				// do something
				cfg.AppMutex.Lock()
				delete(cfg.RequestMap, key)
				cfg.AppMutex.Unlock()
			} else {
				// time expired
				if time.Now().Unix()-rq.StartTime > 50 {
					// do something
					cfg.AppMutex.Lock()
					delete(cfg.RequestMap, key)
					cfg.AppMutex.Unlock()
				}
			}
		}
		utils.JudgeTheFileExist(cfg.RequestMapFile)
		utils.SaveRequest(cfg.RequestMap, cfg.RequestMapFile)
	}
}

// CheckRequest checks if the request is finished
func CheckRequest(rq cfg.Request) bool {
	if cfg.ServerData[rq.TargetID] == rq.TargetStatus {
		return true
	} else {
		return false
	}
}

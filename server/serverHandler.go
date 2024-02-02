package server

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

	cfg "UnstableBase/config"
)

/*
Handle the Request
*/
func RequestHandler(mu *sync.Mutex) {
	for {
		r := <-cfg.RequestChannel
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		SuccessFlag := n.Int64() % 2
		SleepTime := n.Int64() % 2000
		fmt.Printf("HandleType %d,Id %s,SleepTime %d\n", r.HandleType, r.TargetID, SleepTime)
		if SuccessFlag == 1 {
			time.Sleep(time.Duration(SleepTime))
			fmt.Printf("HandleType %d,Id %s,Start\n", r.HandleType, r.TargetID)
			switch r.HandleType {
			case 0:
				CreateTarget(r, mu)
			case 1:
				UpdateTarget(r, mu)
			default:
				DeleteTarget(r, mu)
			}

			fmt.Printf("HandleType %d,Id %s,End\n", r.HandleType, r.TargetID)
		}
	}
}

func CreateTarget(r cfg.Request, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	cfg.ServerData[r.TargetID] = r.TargetStatus
}

func UpdateTarget(r cfg.Request, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	cfg.ServerData[r.TargetID] = r.TargetStatus
}

func DeleteTarget(r cfg.Request, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	delete(cfg.ServerData, r.TargetID)
}

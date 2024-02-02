package config

import (
	"sync"
	"time"
)

/*
HandleType means what kind of handle you want ,eg. CRUD
*/
type Request struct {
	HandleType   int64  `json:"HandleType"`
	TargetID     string `json:"TargetID"`
	TargetStatus string `json:"TargetStatus"`
	StartTime    int64  `json:"StartTime"`
}

/*
h means the type of handle;
status means the change
*/
func (r *Request) PackageTheRequest(h int64, id string, status string) {
	r.HandleType = h
	r.TargetID = id
	r.TargetStatus = status
	r.StartTime = time.Now().UnixNano()
}

var SumHandleType = int64(3)

// User request
var RequestChannel = make(chan Request, 1000)

// Handle channel
var AppChan = make(chan Request, 1000)
var ServerMutex = &sync.Mutex{}
var AppMutex = &sync.Mutex{}

// save the request to check if exists
var RequestMap = make(map[string]Request, 1000)

// Server data
var ServerData = make(map[string]string, 1000)

var (
	RequestFile    = "./data/request.json"
	ServerDataFile = "./data/serverdata.json"
	RequestMapFile = "./data/requestmap.json"
	AppRequestFile = "./data/apprequest.json"
)

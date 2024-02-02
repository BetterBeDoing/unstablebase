package request

import (
	"crypto/rand"
	"fmt"
	"math/big"

	cfg "UnstableBase/config"
	"UnstableBase/utils"
)

/*
Generate the request
*/
func GenerateRequest(sum int) {
	utils.JudgeTheFileExist(cfg.RequestFile)
	for i := 0; i < sum; i++ {
		var rq cfg.Request
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		HandleType := n.Int64() % 3
		TargetID := fmt.Sprintf("%d", n.Int64()%1000)
		TargetStatus := fmt.Sprintf("%d", n.Int64()%1000)
		rq.PackageTheRequest(HandleType, TargetID, TargetStatus)
		cfg.AppChan <- rq
		// save the request to the file
		utils.SaveRequest(rq, cfg.RequestFile)
	}
}

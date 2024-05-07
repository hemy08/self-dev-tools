package main

import (
	"time"

	log "codehub.huawei.com/hemyzhao/lib/hemyutils/logsdk"
	"codehub.huawei.com/hemyzhao/tools/efficiency/unixclient/client"
)

func main() {
	log.NewLogger().InitDefault().
		SetOutPutToConsole(true).
		SetOutPutToFile(false).
		SetLogSkipStep(5).
		SetAllowedLogLevel(log.LevelInfo)

	client.StartUnixDemoClient()
	time.Sleep(time.Second)
}

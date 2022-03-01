package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xelaj/mtproto/telegram"
)

var (
	appHash     = os.Getenv("APP_HASH")
	Token       = os.Getenv("TOKEN")
	appID, _    = strconv.Atoi(os.Getenv("APP_ID"))
	sessionFile = "session.json"
	publicKeys  = "tg_public_keys.pem"
)

func main() {
	c, err := telegram.NewClient(telegram.ClientConfig{
		SessionFile:     sessionFile,
		PublicKeysFile:  publicKeys,
		ServerHost:      "149.154.167.50:443",
		AppID:           appID,
		DeviceModel:     "Unknown",
		SystemVersion:   "linux/amd64",
		AppHash:         appHash,
		InitWarnChannel: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hi there!", c)
}

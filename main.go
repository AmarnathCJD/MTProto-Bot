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
		SessionFile:    sessionFile,
		PublicKeysFile: publicKeys,
		AppID:          appID,
		AppHash:        appHash,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hi there!", c)
}

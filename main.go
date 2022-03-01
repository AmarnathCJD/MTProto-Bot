package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	utils "github.com/xelaj/mtproto/examples/example_utils"
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
	appStorage := utils.PrepareAppStorageForExamples()
	sessionFile := filepath.Join(appStorage, "session.json")
	publicKeys := filepath.Join(appStorage, "tg_public_keys.pem")
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

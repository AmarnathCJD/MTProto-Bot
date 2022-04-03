package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xelaj/mtproto/telegram"
)

var (
	appHash  = os.Getenv("APP_HASH")
	Token    = os.Getenv("TOKEN")
	appID, _ = strconv.Atoi(os.Getenv("APP_ID"))
)

const (
	// from https://my.telegram.org/apps
	TgAppID      = 8950585              // integer value from "App api_id" field
	TgProdServer = "149.154.167.50:443" // string value from "Production configuration" field

	TgBotUserName = "aiko_robot" // username of the bot
)

var TgAppHash = "9b0f3eed1e952ed235eeae933b1daef6"
var TgBotToken = os.Getenv("TOKEN")

func main() {
        appStorage := utils.PrepareAppStorageForExamples()
	sessionFile := filepath.Join(appStorage, "session.json")
	publicKeys := filepath.Join(appStorage, "tg_public_keys.pem")

	c, err := telegram.NewClient(telegram.ClientConfig{
		SessionFile:    sessionFile,
		PublicKeysFile: publicKeys,
		AppID:          appID,
		AppHash:        appHash,
                ServerHost: "149.154.167.50:443",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(c)
}

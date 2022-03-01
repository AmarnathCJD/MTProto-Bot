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
	appHash  = os.Getenv("APP_HASH")
	Token    = os.Getenv("TOKEN")
	appID, _ = strconv.Atoi(os.Getenv("APP_ID"))
)

const (
	// from https://my.telegram.org/apps
	TgAppID      = 8950585              // integer value from "App api_id" field      // string value from "App api_hash" field
	TgTestServer = "149.154.167.40:443" // string value from "Test configuration" field
	TgProdServer = "149.154.167.50:443" // string value from "Production configuration" field

	TgBotUserName = "aiko_robot" // username of the bot
)

var TgAppHash = os.Getenv("APP_HASH")
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
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(c)
}

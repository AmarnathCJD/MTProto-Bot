package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/amarnathcjd/gogram/telegram"
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
	c, err := telegram.NewClient(telegram.ClientConfig{
		SessionFile:    "session.json",
		PublicKeysFile: "tg_public_keys.pem",
		AppID:          appID,
		AppHash:        appHash,
                ServerHost: "149.154.167.50:443",
	})
	if err != nil {
		fmt.Println(err)
	}
        _, errx := c.AuthImportBotAuthorization(
		1, // flags, it's reserved, must be set (don't mind how does it works, we don't know too)
		int32(appID),
		appHash,
		Token,
	)
        if errx != nil {
		fmt.Println(errx)
	}
	uname, err := c.ContactsResolveUsername("telegram")
	if err != nil {
		fmt.Println("ResolveUsername error:", err.Error())
		os.Exit(1)
	}
        fmt.Println(uname)
}

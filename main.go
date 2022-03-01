package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/xelaj/mtproto/telegram"
)

var (
	appHash     = os.Getenv("APP_HASH")
	Token       = os.Getenv("TOKEN")
	appID, _    = strconv.Atoi(os.Getenv("APP_ID"))
	sessionFile = filepath.Join("./", "session.json")
	publicKeys  = filepath.Join("./", "tg_public_keys.pem")
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
	_, err2 := c.AuthImportBotAuthorization(1, int32(appID), appHash, Token)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Hi there!")
}

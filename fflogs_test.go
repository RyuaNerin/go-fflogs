package fflogs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// https://ko.fflogs.com/reports/r2BjwQK7ng9kALt1
	testCode          = "r2BjwQK7ng9kALt1"
	testStart         = 2648515  // fights. 1 start
	testEnd           = 12935833 // fights.30 end
	testUserName      = "RyuaNerin"
	testGuildName     = "륜아린의 재활운동"
	testCharacterName = "륜아린"
	testServerName    = "Moogle"
	testServerRegion  = RegionKR
	testZone          = 38 // Eden's Promise
	testEncounter     = 73 // Cloud of Darkness
)

func newTestClient() *Client {
	godotenv.Load(".env")

	c, err := NewClient(&ClientOpt{
		ApiKey: os.Getenv("FFLOGS_V1_APIKEY"),
	})
	if err != nil {
		panic(err)
	}

	return c
}

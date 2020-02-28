package contentfabric

import (
	"fmt"
	"testing"
)

var testConfig Config = Config{
	configUrl:       ConfigContentFabricV3.configUrl,
	authStaticToken: "NOTAREALTOKEN",
}
var testHash = "hq__Kx352AFDyot4hsTGUzAhhkXwkUDSXJoUk7X7mxnAxXhagay6RPRtgE1um1yaJz12xPemgrkfta"

func TestStaticUrl(t *testing.T) {
	urlDash, urlHls, err := PlayoutStaticUrl(testConfig, testHash)
	if err != nil {
		t.Error(err)
	}
	if urlDash == "" || urlHls == "" {
		t.Error("URL empty")
	}
}

func ExampleStaticUrl() {
	urlDash, urlHls, err := PlayoutStaticUrl(testConfig, testHash)
	if err != nil {
		return
	}
	fmt.Println(urlDash)
	fmt.Println(urlHls)

	// Output:
	// https://main.net955305.contentfabric.io/q/hq__Kx352AFDyot4hsTGUzAhhkXwkUDSXJoUk7X7mxnAxXhagay6RPRtgE1um1yaJz12xPemgrkfta/rep/playout/default/dash-clear/dash.mpd?authorization=NOTAREALTOKEN
	// https://main.net955305.contentfabric.io/q/hq__Kx352AFDyot4hsTGUzAhhkXwkUDSXJoUk7X7mxnAxXhagay6RPRtgE1um1yaJz12xPemgrkfta/rep/playout/default/hls-clear/playlist.m3u8?authorization=NOTAREALTOKEN
}

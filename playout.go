package contentfabric

import "net/url"

const (
	defaultOffering string = "default"
	playoutPath     string = "/rep/playout/"
	dashClear       string = "/dash-clear/dash.mpd"
	hlsClear        string = "/hls-clear/playlist.m3u8"
	optionsPath     string = "options.json"
)

type Config struct {
	ConfigUrl       string
	AuthStaticToken string
}

var ConfigContentFabricV3 Config = Config{ConfigUrl: "https://main.net955305.contentfabric.io/config"}
var ConfigDemoContentFabricV3 Config = Config{ConfigUrl: "https://demov3.net955210.contentfabric.io/config"}
var ConfigContentFabric Config = ConfigContentFabricV3 // Latest version

func PlayoutStaticUrl(c Config, qihot string) (string, string, error) {
	url, err := c.staticUrl()
	urlDash := url + "/q/" + qihot + playoutPath + defaultOffering + dashClear + "?authorization=" + c.AuthStaticToken
	urlHls := url + "/q/" + qihot + playoutPath + defaultOffering + hlsClear + "?authorization=" + c.AuthStaticToken
	return urlDash, urlHls, err
}

func PlayoutOptionsUrl() {
	/* Not yet implemented */
}

func (c *Config) staticUrl() (string, error) {
	u, err := url.Parse(c.ConfigUrl)
	if err != nil {
		return "", err
	}
	u.Path = ""
	return u.String(), nil
}

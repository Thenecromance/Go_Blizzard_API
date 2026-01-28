package global

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	cookie_prefix = "cookie"
)

var (
	netEase_cookie = fmt.Sprintf("%s.netease", cookie_prefix)
	dd_cookie      = fmt.Sprintf("%s.DD", cookie_prefix)
)

func init() {
	viper.SetDefault(netEase_cookie, "") // storing cookie for wow.blizzard.cn
	viper.SetDefault(dd_cookie, "")      // storing cookie for NetEase DD using
}
func GetNetEaseCookie() string {
	return viper.GetString(netEase_cookie)
}

func GetDDCookie() string {
	return viper.GetString(dd_cookie)
}

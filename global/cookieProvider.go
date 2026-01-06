package global

import "github.com/spf13/viper"

func GetCookie() string {
	return viper.GetString("cookie")
}

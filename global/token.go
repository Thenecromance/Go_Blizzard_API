package global

import "github.com/spf13/viper"

func GetClientID() string {
	return viper.GetString("client_id")
}
func GetClientSecret() string {
	return viper.GetString("client_secret")
}

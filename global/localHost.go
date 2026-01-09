package global

import "github.com/spf13/viper"

var (
	localHost = "http://localhost"
)

func init() {
	viper.SetDefault("local_host", "http://localhost")
}
func GetLocalHost() string {
	return localHost
}

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("universal_config")
	viper.SetConfigType("yaml")

	//viper.AddConfigPath("/etc/unofficial_api/")
	//viper.AddConfigPath("$HOME/.unofficial_api")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		viper.Set("client_id", "")
		viper.Set("client_secret", "")
		viper.Set("cookie", "")
		err := viper.SafeWriteConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func main() {

	/*res, err := server.UpdateAllServerStatus()
	if err != nil {
		panic(err)
	}
	os.WriteFile("server_status.json", []byte(res), 0644)*/
}

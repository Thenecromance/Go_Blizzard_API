package Authentication

import (
	"context"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

var cli *http.Client

func init() {

	// Setup oauth2 config
	conf := &clientcredentials.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		TokenURL:     "https://us.battle.net/oauth/token",
	}
	cli = conf.Client(context.Background())

}

func Client() *http.Client {
	return cli
}

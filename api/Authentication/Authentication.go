package Authentication

import (
	"context"
	"net/http"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

var _cli *http.Client

func init() {
	viper.SetDefault("blizzard.id", "")
	viper.SetDefault("blizzard.secret", "")

}

func Client() *http.Client {
	if _cli == nil {
		// Setup oauth2 config
		conf := &clientcredentials.Config{
			ClientID:     viper.GetString("blizzard.id"),
			ClientSecret: viper.GetString("blizzard.secret"),
			TokenURL:     "https://us.battle.net/oauth/token",
		}
		_cli = conf.Client(context.Background())
	}
	return _cli
}

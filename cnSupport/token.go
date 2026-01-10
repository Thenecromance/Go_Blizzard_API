package cn

import (
	"fmt"

	"github.com/Thenecromance/Go_Blizzard_API/utils"
)

var (
	playerTokens *utils.Cache
)

func init() {
	playerTokens = utils.New(1 << 20)
}

func AddPlayerToken(name, realm, token string) {
	playerTokens.Add(fmt.Sprintf("%s:%s", name, realm), token)
}

func TryGetPlayerToken(name, realm string) (string, bool) {
	token, ok := playerTokens.Get(fmt.Sprintf("%s:%s", name, realm))
	return token.(string), ok
}

package command

import (
	"github.com/libertylocked/torchprint"
	"github.com/spf13/viper"
)

func newAPI() *torchprint.API {
	userid := viper.GetString("userid")
	token := viper.GetString("token")
	username := viper.GetString("username")
	password := viper.GetString("password")

	api := torchprint.NewAPI(userid)
	if len(username) != 0 && len(password) != 0 {
		return api.SetUserPass(username, password)
	}
	return api.SetToken(token)
}

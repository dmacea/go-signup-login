package useraccounts

import (
	"github.com/dmaceasistemas/go-backend-dev-login/helpers"
	"github.com/dmaceasistemas/go-backend-dev-login/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ? ", id).Update("balance", amount)
	defer db.Close()
}

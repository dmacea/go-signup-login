package migrations

import (
	"github.com/dmaceasistemas/go-backend-dev-login/helpers"
	"github.com/dmaceasistemas/go-backend-dev-login/interfaces"
)

/*type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}*/

func createAccounts() {
	db := helpers.ConnectDB()

	users := [2]interfaces.User{
		{Username: "Danny", Email: "dmacea.sistemas@gmail.com"},
		{Username: "Kenny", Email: "maceakenny@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := interfaces.Account{Type: "Daily Accoubt", Name: string(users[i].Username + "'s'" + " account"), Balance: uint(10000 * int(i+1)), UserId: user.ID}
		db.Create(&account)
	}
	defer db.Close()

}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User{}, &Account{})

	defer db.Close()

	createAccounts()
}

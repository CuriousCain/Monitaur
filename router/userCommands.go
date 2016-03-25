package router

import (
	"encoding/json"
	"github.com/curiouscain/monitaur/models"
	"net/http"
)

func (env *Env) NewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User

	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	user.HashPassword()
	env.Db.SubmitUser(&user)
}

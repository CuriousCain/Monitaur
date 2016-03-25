package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/curiouscain/monitaur/models"
	"github.com/curiouscain/monitaur/router"
	"net/http"
)

type Config struct {
	DBuser   string
	DBpass   string
	DBname   string
	CertPath string
	KeyPath  string
}

func main() {
	fmt.Println("Importing conf...")

	var conf Config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		panic(err)
	}

	fmt.Println("Server starting on port 443...")

	conStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=require", conf.DBuser, conf.DBpass, conf.DBname)

	db := models.NewDB(conStr)
	env := &router.Env{db}

	http.HandleFunc("/users/new", env.NewUser)
	http.ListenAndServeTLS(":443", conf.CertPath, conf.KeyPath, nil)
}

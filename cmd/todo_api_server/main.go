package main

import (
	"fmt"
	"github.com/us-learn-and-devops/todoapi/cmd/todo_api_server/handlers"
	"github.com/us-learn-and-devops/todoapi/configs"
	envcfg "github.com/us-learn-and-devops/todoapi/pkg/envconfig"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	cfgs := &configs.Settings{}
	err := envcfg.Unmarshal(cfgs)
	if err != nil {
		log.Fatalf("failed to get configs: %s", err)
	}

	var dbCreds handlers.DBCredentials

	dbUsername, err := ioutil.ReadFile(cfgs.DatabaseUserNameFilePath)
	if err != nil {
		log.Fatalf("failed to get DB username: %v", err)
	}
	dbCreds.Username = string(dbUsername)

	dbPswd, err := ioutil.ReadFile(cfgs.DatabasePswdFilePath)
	if err != nil {
		log.Fatalf("failed to get DB password: %v", err)
	}
	dbCreds.Password = string(dbPswd)

	tl, err:= handlers.NewTodoListHandler(cfgs, dbCreds)
	if err != nil {
		log.Fatalf("failed to create TodoListHandler: %v", err)
	}

	r := handlers.NewRouter(tl)

	host := fmt.Sprintf("%s:%s", cfgs.ServerAddr, cfgs.ServerPort)

	log.Printf("serving todo-api on %s\n", host)
	log.Fatal(http.ListenAndServe(host, r))
}

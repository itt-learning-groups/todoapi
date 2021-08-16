package main

import (
	"fmt"
	"github.com/us-learn-and-devops/todoapi/cmd/todo_api_server/handlers"
	"github.com/us-learn-and-devops/todoapi/configs"
	envcfg "github.com/us-learn-and-devops/todoapi/pkg/envconfig"
	"log"
	"net/http"
)

func main() {
	cfgs := &configs.Settings{}
	err := envcfg.Unmarshal(cfgs)
	if err != nil {
		log.Fatalf("Failed to get configs: %s", err)
	}

	tl, err:= handlers.NewTodoListHandler(cfgs)
	if err != nil {
		log.Fatalf("main.main failed to create TodoListHandler: %v", err)
	}

	r := handlers.NewRouter(tl)

	host := fmt.Sprintf("%s:%s", cfgs.ServerAddr, cfgs.ServerPort)

	log.Printf("serving todo-api on %s\n", host)
	log.Fatal(http.ListenAndServe(host, r))
}

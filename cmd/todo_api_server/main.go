package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/us-learn-and-devops/todoapi/cmd/todo_api_server/handlers"
	"github.com/us-learn-and-devops/todoapi/configs"
	envcfg "github.com/us-learn-and-devops/todoapi/pkg/envconfig"
)

func main() {
	cfgs := &configs.Settings{}
	err := envcfg.Unmarshal(cfgs)
	if err != nil {
		fmt.Printf("Failed to get configs: %s", err)
		os.Exit(1)
	}

	tl := handlers.NewTodoListHandler()
	r := handlers.NewRouter(tl)

	host := fmt.Sprintf("%s:%s", cfgs.ServerAddr, cfgs.ServerPort)

	fmt.Printf("serving todo-api on %s", host)
	log.Fatal(http.ListenAndServe(host, r))
}

package main

import (
	"github.com/unfavorablenode/thin_node/env"
	"github.com/unfavorablenode/thin_node/handlers"
)


func main() {
    env.RegisterEnv()
    Router := handlers.GetRouterWithRoutes()
    Server := handlers.GetServer(Router)
    handlers.StartServerWithGracefullShutdown(Server)
}

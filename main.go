package main

import (
	"github.com/unfavorablenode/thin_node/handlers"
)


func main() {
    Router := handlers.GetRouterWithRoutes()
    Server := handlers.GetServer(Router)
    handlers.StartServerWithGracefullShutdown(Server)
}

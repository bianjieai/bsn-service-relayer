package server

import (
	"fmt"
	"relayer/appchains"
)

// StartWebServer starts the web server with a ChainManager instance
func StartWebServer(
	chainManager appchains.AppChainHandlerI,
	port int,
) {
	srv := NewHTTPService(chainManager)

	err := srv.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
	}
}

package server

import (
	"fmt"
)

// StartWebServer starts the web server with a ChainManager instance
func StartWebServer(
	chainManager *ChainManager,
	marketManager *MarketManager,
	port int,
) {
	srv := NewHTTPService(chainManager, marketManager)

	err := srv.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
	}
}

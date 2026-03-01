package main

import (
	http_utils "Client-Server-API/client/http-utils"
	"context"
)

func main() {
	ctx := context.Background()
	http_utils.ChamarEndpoint(ctx)
}

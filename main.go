package main

import (
	"github.com/mjl776/sports-management-platform/internal/api"
)

func main() {

	server := api.NewAPIServer(":3000");
	server.Run();

}
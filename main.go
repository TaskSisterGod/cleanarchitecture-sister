package main

import (
	"cleanarchitecture-sister/external"
	"cleanarchitecture-sister/external/mysql"
)

func main() {
	defer mysql.CloseConn()
	external.Router.Run()
}

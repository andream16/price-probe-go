package main

import (
	"fmt"

	"../configuration"
	"../server"
)

func main() {

	conf := configuration.InitConfiguration()
	/*fmt.Println("Trying to enstablish a connection to a Cassandra Cluster . . .")
	s, err := cassandramanager.InitCassandraClient(conf)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Connection to Cassandra Cluster successfully enstabilshed!")
	}*/

	fmt.Println("Starting server . . .")
	server.InitServer(conf)

}

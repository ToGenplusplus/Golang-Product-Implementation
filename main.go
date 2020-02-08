package main
/**
Main Program
**/

import (
	"log"
	"net/http"
)

func main() {
	
	go uploadCounters()
	//dont need to issue sync.waitgroup because the routine will continue to run as long
	//as server stays alive

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/stats/", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

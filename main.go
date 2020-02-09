package main
/**
Main Program
**/

import (
	"log"
	"net/http"	
)


func main() {
	
	go uploadCounters()		//go routine to upload counters every 5 seconds
	go isAllowed()		//runs in the background used for global rate limiter

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/stats/", statsHandler)


	log.Fatal(http.ListenAndServe(":8080", nil))

}

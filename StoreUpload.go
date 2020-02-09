package main

import (
	
	"time"
	"os"
	"strconv"		//used for any conversion between string and ints
)

var timeaccess = time.Now()

var PreviousUploadTime = timeaccess.Format("2 Jan 2006 15:04:05") //keep track of last upload to data store

func uploadCounters() {
	for {

		time.Sleep(5 * time.Second)
			//attempt to open a file, that will be used to store the counters
		file, err := os.OpenFile("counterstore.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		//check to see if an error occured 
		if err != nil {
			panic(err)
		}

		GetCurrentTime := time.Now()
		//let the user know time and date the datastore was updated
		UploadTime := GetCurrentTime.Format("2 Jan 2006 15:04:05")

		file.WriteString("Prev Upload:" + PreviousUploadTime + "-----------Current: " + UploadTime+ "--------------\n\n")

		//set previous upload time, to the most recent upload time for next call
		PreviousUploadTime = UploadTime	

		//iterate through map contents, getting the key and value and write to store
		for key, value := range counterMap {

			file.WriteString( " Key: '" + key + "' Value : { views : " + strconv.Itoa(value.view) + " clicks : " + strconv.Itoa(value.click) + " }\n\n")
			//time.Sleep(time.Second)	//simulate large data set upload

		}
		
	}
	
}
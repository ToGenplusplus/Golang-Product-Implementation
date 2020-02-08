package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"		//func Contains(s, substr)
	"sync"
	"time"
	"strconv"		//used for any conversion between string and ints
)

type counters struct {
	sync.Mutex
	view  int
	click int
}

var (
	c = counters{}

	content = []string{"sports", "entertainment", "business", "education"}

	counterMap = make(map[string]counters)	//data structure to hold counters as values and data content as key.

)


func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	data := content[rand.Intn(len(content))]

	c.Lock()
	c.view++
	c.Unlock()

	countSupport(data)	//call the support function to update counters and populate counterMap

	err := processRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	if rand.Intn(100) < 50 {
		processClick(data)
	}

}

func processRequest(r *http.Request) error {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}

/**
Function will be used to support counters, by content and time selections.
takes in a string representing a data content as a paramter
**/

func countSupport(data string) {
	/*
		//i want this function to take in a data(representing content)
		// then for each data attach a time stamp to it, and that will generate a key string
		//for the map of string to counters created
		//everytime view handler is called, check the key and incrment
		the appropriate values for that key.
	*/
	timeaccess := time.Now()	//get current time

	data = data + " : " + timeaccess.Format("2006-01-02 15:04")		//format time to readable string and concatenate with content data

	//check if the key already exist in map, if it does
	if _, ok := counterMap[data]; ok {

		//call the Incr_view function on the key counter, and store the value in DatView
		DatView := Incr_view(counterMap[data])	

		//get the click value from the Key value
		DatClick := counterMap[data].click

		//to simiulate click calls
		if rand.Intn(100) < 50 {
			DatClick = Incr_click(counterMap[data])	//set DatClick to new click value, returned by Incr_click
		}

		//update Value for the Key
		counterMap[data] = counters{view : DatView, click : DatClick}
		
	}else {	//if key doesn't exist

		//initialize a new key value pair
		counterMap[data] = counters{ view : 1, click : 0}

	}
	
	
}

/*

This functions takes in a string representing either, a content or timestamp or both
representing portion of the key(or whole key) associated with a counter/counters
and prints out the counters if the key exist in map.

*/

func RecieveCounters(content string) {
	//check to make sure counters are availble 

	if len(counterMap) != 0{

		found := 0	//keep track of # of matches for specified content(key)

		for key, value := range counterMap {

			//check to see if the paramter passes in exist as part of a key or whole key in map
			if strings.Contains(key,content) {
				found++	//incrment found 
				fmt.Println(" Key: '" + key + "' Values : { views : " + strconv.Itoa(value.view) + " clicks : " + strconv.Itoa(value.click) + " }\n")
			}
		}
		//if found = 0 that means, no key was found for the paramter passed in

		if found == 0{
				fmt.Println("No counters for specified content")
		}

	}else{	//let the users know no counters are available yet
		fmt.Println("There are no counters yet")
	}

}

/**
Function takes in a counter as a paramter, and increments the view value of counter
returns the new view value
**/

func Incr_view(count counters) int {
	count.Lock()
	count.view++
	defer count.Unlock()	//Unlock counter mutex once view value is returned

	return count.view
}

func processClick(data string) error {

	c.Lock()
	c.click++
	c.Unlock()

	return nil
}

/**
Function takes in a counter as a paramter, and increments the click value of counter
returns the new click value. (will only be called by above countSuport function)
**/

func Incr_click(count counters) int {

	count.Lock()
	count.click++
	defer count.Unlock()	//Unlock counter mutex once click value is returned
	return count.click
}

func statsHandler(w http.ResponseWriter, r *http.Request) {

	if !isAllowed() {
		w.WriteHeader(429)
		http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
		return
	}
}


func isAllowed() bool {

	return true
}


/**
Function displays a formated output of the contents in map to the terminal
(implemented as a go routine)
**/
func printMapContents() {
	for {

		time.Sleep(10 * time.Second)	//run every 10 seconds

		i := 0	//intialize variable i will track # of contents in map

		fmt.Println("----------------------Map Contents-------------------------")
		//iterate through map contents, getting the key and value
		for key, value := range counterMap {
			i++	//incrment i for each row of map contents
			fmt.Println( strconv.Itoa(i) + " Key: '" + key + "' Values : { views : " + strconv.Itoa(value.view) + " clicks : " + strconv.Itoa(value.click) + " }\n")
		}

	}
		
}

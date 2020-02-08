package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
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

	contentMap = make(map[string]counters)	//data structure to hold counters as values and data content as key.
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	data := content[rand.Intn(len(content))]

	c.Lock()
	c.view++
	c.Unlock()

	countSupport(data)	//call the support function to update counters as needed

	err := processRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	if rand.Intn(100) < 50 {
		processClick(data)
	}

	printMapContents()	// to track current contents in map ( will remove once data store is implemented)

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
	if _, ok := contentMap[data]; ok {

		//call the Incr_view function on the key counter, and store the value in DatView
		DatView := Incr_view(contentMap[data])	

		//get the click value from the Key value
		DatClick := contentMap[data].click

		//to simiulate click calls
		if rand.Intn(100) < 50 {
			DatClick = Incr_click(contentMap[data])	//set DatClick to new click value, returned by Incr_click
		}

		//update Value for the Key
		contentMap[data] = counters{view : DatView, click : DatClick}
		
	}else {	//if key doesn't exist

		//initialize a new key value pair
		contentMap[data] = counters{ view : 1, click : 1}

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

var Limit = 10

var timer = time.NewTimer(time.Minute)

func isAllowed() bool {

	for Limit >= 0 {
		Limit--

		if Limit == 0 {
			return false
		}
	}
	fmt.Println(Limit)
	return true
}

func uploadCounters() error {
	return nil
}


/**
Function displays a formated output of the contents in map to the terminal
**/
func printMapContents() {
	i := 0	//intialize variable i will track # of contents in map

	fmt.Println("----------------------Map Contents-------------------------\n")
	//iterate through map contents, getting the key and value
	for key, value := range contentMap {
		i++	//incrment i for each row of map contents
		fmt.Println( strconv.Itoa(i) + " Key: '" + key + "' Values : { views : " + strconv.Itoa(value.view) + " clicks : " + strconv.Itoa(value.click) + " }\n")
	}

	
}
func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/stats/", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

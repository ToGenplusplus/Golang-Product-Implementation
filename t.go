package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type count struct {
	view  int
	click int
}

var (
	cc = count{}

	cont = []string{"sports", "entertainment", "business", "education"}

	prevdata = " "

	m = make(map[string]count)
)


func trial() {

	data := cont[rand.Intn(len(cont))]

	timeacces := time.Now()

	data = data + " : " + timeacces.Format("2006-01-02 15:04") //this will be our ke

	
	if _, ok := m[data]; ok {
		view1 := m[data].view
		click1 := m[data].click

		view1++
		click1++

		m[data] = count{view: view1, click: click1}

		fmt.Println("Key : " + data + " Values : { views : " + strconv.Itoa(view1) + " clicks : " + strconv.Itoa(click1) + " }\n")
	} else {
		fmt.Print(data + "-- ")
		fmt.Println("key not in dictionary, initializing key now\n")
		m[data] = count{view: 1, click: 1}

	}

}

/*
func main() {

	for i := 0; i <= 20; i++ {
		trial()
	}
}
*/

/**
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type counters struct {
	sync.Mutex
	view  int
	click int
}

type values struct {
	views  int
	clicks int
}

var (
	c = counters{}

	content = []string{"sports", "entertainment", "business", "education"}
)


	contVal[data] = &{views : c.view , clicks : c.click}
	data = prevdata


func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	data := content[rand.Intn(len(content))]

	err := processRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	processView() //call below method to increment view counter

	clickInit := randGenerate() //calls boolean function, return true if click was incremented

	timeacces := time.Now()

	data = data + " : " + timeacces.Format("2006-01-02 15:04") //this will be our key

	contVal := make(map[string]*values) //map string to values struct

	prevdata := "" //used to set the values struct appropriately

	if prevdata == data {
		contVal[data] = &values{views: c.click, clicks: c.view}
	} else {
		if clickInit {
			datViews := &contVal[data].views   //datViews stores the address of data views value
			datClicks := &contVal[data].clicks //datClicks stores the address of data clicks value

			*datViews += 1  //views always as to be incremented by one.
			*datClicks += 1 // if click was incremnetd in c.click, increment it for data clicks value
			contVal[data] = &values{views: *datViews, clicks: *datClicks}
		} else {
			datViews := &contVal[data].views   //datViews stores the address of data views value
			datClicks := &contVal[data].clicks //datClicks stores the address of data clicks value

			*datViews += 1 //views always as to be incremented by one.
			contVal[data] = &values{views: *datViews, clicks: *datClicks}
		}

	}

	prevdata = data // set previous recieved data key to recent data key

	//fmt.Fprintln(w, contVal[data].views)
	fmt.Fprintln(w, data)
	fmt.Fprintln(w, c.view)
	fmt.Fprintln(w, c.click)
}

/**
boolean funciton to check if processClick was called.

func randGenerate() bool {
	if rand.Intn(100) < 50 {
		processClick()
		return true
	}

	return false
}

func processView() error {
	c.Lock()
	c.view++
	c.Unlock()
	return nil
}

func processRequest(r *http.Request) error {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}

func processClick() error {
	c.Lock()
	c.click++
	c.Unlock()
	return nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if !isAllowed() {
		w.WriteHeader(429)
		return
	}
}

func isAllowed() bool {
	return true
}

func uploadCounters() error {
	return nil
}


func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/stats/", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
**/

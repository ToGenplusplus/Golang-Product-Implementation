package main 

import (
	"fmt"
	"time"
	"sync"
	
)


//struct to hold Rate Limiter information 
type Limiter struct {

	sync.Mutex
	NumberRequest int
	RequestConsumed int
	PerDuration Ticker

}

//ticker to track request duration
type Ticker struct {

	duration time.Duration
	ticker 	time.Ticker
}

//create a new Ticker struct with the time duartion passed into the function
func createTicker(period time.Duration) *Ticker {
    return &Ticker{period, *time.NewTicker(period)}
}

//function to reset a ticker
func (t * Ticker) resetTicker(){

	t.ticker = *time.NewTicker(t.duration)

}

var (

	tick = createTicker(time.Minute)	//create a new ticker with a minute duration
	RateLimit = Limiter{NumberRequest : 10 , RequestConsumed : 0 , PerDuration : *tick}	//issue a new rate limit
	DurationEllapsed = make (chan struct{})	//channel to check if request limit duration has elapsed
	ReqAllowed = true	//boolean variable to toggle flow or api request
)

/*
This function will uses the rate limit ands ticker to issue a rate limiter
*/
func isAllowed() error {

	go func() {
		for {

			select {
			
			case <-tick.ticker.C:
				
				if RateLimit.RequestConsumed >= RateLimit.NumberRequest {
					ReqAllowed = false		//not allowing any more request for remaing duration
					fmt.Println("limit has been reached, wait a few seconds")	//signal limit has been reached
				}

			case <- DurationEllapsed:	//if duration interval has passed
				RateLimit.RequestConsumed = 0	//reset request tokens
				ReqAllowed = true	//being to allow request again
				tick.resetTicker()	//reset ticker for new time.duration
			}
	
		}
		
	}()
	return nil
}

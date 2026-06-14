package routes

import "time"

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      time.Duration
}

type response struct {
	URL 			string
	CustomShort 	
	Expiry
	XRateRemaining
	XRateLimitReset
}

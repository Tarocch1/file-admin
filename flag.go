package main

import "flag"

var flagHost string

var flagPort string

var flagDir string

var flagAuth string

func initFlag() {
	flag.StringVar(&flagHost, "h", "0.0.0.0", "Host to listen.")
	flag.StringVar(&flagPort, "p", "8080", "Port to listen.")
	flag.StringVar(&flagDir, "d", ".", "Dir path to serve.")
	flag.StringVar(&flagAuth, "a", "", "<username:password> Basic auth user.")
}

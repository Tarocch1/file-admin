package common

import (
	"flag"
	"strings"
)

var FlagHost string

var FlagPort string

var FlagDir string

var FlagAuth string

var FlagHTTPSCert string

var FlagHTTPSKey string

var FlagVersion bool

func init() {
	flag.StringVar(&FlagHost, "h", "0.0.0.0", "Host to listen.")
	flag.StringVar(&FlagPort, "p", "3000", "Port to listen.")
	flag.StringVar(&FlagDir, "d", ".", "Dir path to serve.")
	flag.StringVar(&FlagAuth, "a", "", "<username:password> Basic auth user.")
	flag.StringVar(&FlagHTTPSCert, "https-cert", "", "Path to https cert.")
	flag.StringVar(&FlagHTTPSKey, "https-key", "", "Path to https key.")
	flag.BoolVar(&FlagVersion, "v", false, "Print version information.")
}

func ParseFlag() {
	flag.Parse()

	if FlagAuth != "" {
		index := strings.Index(FlagAuth, ":")
		if index != -1 {
			AuthUser = FlagAuth[:index]
			AuthPass = FlagAuth[index+1:]
		}
	}
}

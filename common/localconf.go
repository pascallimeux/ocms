package common

import (
	"time"
)

type Configuration struct {
	Logger          string
	HttpHyperledger string
	HttpHostUrl     string
	ChainCodePath   string
	EnrollID        string
	EnrollSecret    string
	LogFileName     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	HandlerTimeout  time.Duration
}

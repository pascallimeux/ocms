package common

import (
	"time"
)

type Configuration struct {
	Logger          string
	HttpHyperledger string
	HttpHostUrl     string
	ChainCodePath   string
	ChainCodeName   string
	EnrollID        string
	EnrollSecret    string
	LogFileName     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	HandlerTimeout  time.Duration
	ApplicationID   string
}

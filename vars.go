package fifochan

import (
	"sync"
)

var chans map[string]chan interface{}
var mutex sync.Mutex
var wg sync.WaitGroup
var Err chan error
var Done chan struct{}
var stopPubChan chan struct{}
var stopSubChan chan struct{}

// Copyright (c) 2015 HPE Software Inc. All rights reserved.
// Copyright (c) 2013 ActiveState Software Inc. All rights reserved.

package watch

import (
	"errors"
	"log"
	"os"
	"sync"
)

type Watcher interface {
	Add(string) error
	Remove(string) error
}

type Event struct {
	Name string
}

type Op = interface{}

type InotifyTracker struct {
	mux       sync.Mutex
	watcher   Watcher
	chans     map[string]chan Event
	done      map[string]chan bool
	watchNums map[string]int
	watch     chan *watchInfo
	remove    chan *watchInfo
	error     chan error
}

type watchInfo struct {
	op    Op
	fname string
}

func (this *watchInfo) isCreate() bool {
	return false
}

var (
	// globally shared InotifyTracker; ensures only one fsnotify.Watcher is used
	shared *InotifyTracker

	// these are used to ensure the shared InotifyTracker is run exactly once
	once  = sync.Once{}
	goRun = func() {
		shared = &InotifyTracker{
			mux:       sync.Mutex{},
			chans:     make(map[string]chan Event),
			done:      make(map[string]chan bool),
			watchNums: make(map[string]int),
			watch:     make(chan *watchInfo),
			remove:    make(chan *watchInfo),
			error:     make(chan error),
		}
		go shared.run()
	}

	logger = log.New(os.Stderr, "", log.LstdFlags)
)

// Watch signals the run goroutine to begin watching the input filename
func Watch(fname string) error {
	return errors.New("aix 不支持")
}

// Watch create signals the run goroutine to begin watching the input filename
// if call the WatchCreate function, don't call the Cleanup, call the RemoveWatchCreate
func WatchCreate(fname string) error {
	return errors.New("aix 不支持")
}

func watch(winfo *watchInfo) error {
	return nil
}

// RemoveWatch signals the run goroutine to remove the watch for the input filename
func RemoveWatch(fname string) error {
	return nil
}

// RemoveWatch create signals the run goroutine to remove the watch for the input filename
func RemoveWatchCreate(fname string) error {
	return nil
}

func remove(winfo *watchInfo) error {
	return nil
}

// Events returns a channel to which FileEvents corresponding to the input filename
// will be sent. This channel will be closed when removeWatch is called on this
// filename.
func Events(fname string) <-chan Event {
	shared.mux.Lock()
	defer shared.mux.Unlock()

	return shared.chans[fname]
}

// Cleanup removes the watch for the input filename if necessary.
func Cleanup(fname string) error {
	return RemoveWatch(fname)
}

// watchFlags calls fsnotify.WatchFlags for the input filename and flags, creating
// a new Watcher if the previous Watcher was closed.
func (shared *InotifyTracker) addWatch(winfo *watchInfo) error {
	return nil
}

// removeWatch calls fsnotify.RemoveWatch for the input filename and closes the
// corresponding events channel.
func (shared *InotifyTracker) removeWatch(winfo *watchInfo) error {
	return nil
}

// sendEvent sends the input event to the appropriate Tail.
func (shared *InotifyTracker) sendEvent(event Event) {
}

// run starts the goroutine in which the shared struct reads events from its
// Watcher's Event channel and sends the events to the appropriate Tail.
func (shared *InotifyTracker) run() {
}

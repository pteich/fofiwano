package fofiwano

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"strings"

	"github.com/rjeczalik/notify"
	"golang.org/x/net/context"
)

// WatcherNotify defines a notification type for a watcher
type WatcherNotify struct {
	Notify  string
	Event   string
	Options map[string]string
}

// Watcher defines a watcher configuration
// Target can be a folder or a file, add /... to a folder for recursive watching
// e.g. ./test/...
type Watcher struct {
	Target        string
	Notifications []WatcherNotify
}

// Watch starts the watcher with the given configuration
func Watch(watches []Watcher) {

	ctx, done := context.WithCancel(context.Background())

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		switch <-signalChannel {
		case os.Interrupt:
			done()
		case syscall.SIGTERM:
			done()
		}
	}()

	for _, watcher := range watches {

		stopfunc := func(watcher Watcher) func() {
			watcherEvents := make(chan notify.EventInfo, 2)
			go func() {
				for event := range watcherEvents {
					eventString := strings.ToLower(event.Event().String())
					for _, notification := range watcher.Notifications {
						if strings.ToLower(notification.Event) == "all" || strings.Contains(eventString, strings.ToLower(notification.Event)) {
							// TODO implement real notification
							log.Printf("Got event: %+v | %s | %+v", event.Event(), event.Path(), watcher)
						}
					}
				}
			}()

			log.Printf("Starting watcher for %s\n", watcher.Target)
			if err := notify.Watch(watcher.Target, watcherEvents, notify.All); err != nil {
				log.Fatal(err)
			}

			return func() {
				log.Printf("Stopping watcher for %s\n", watcher.Target)
				notify.Stop(watcherEvents)
			}
		}(watcher)

		defer stopfunc()
	}

	<-ctx.Done()

}

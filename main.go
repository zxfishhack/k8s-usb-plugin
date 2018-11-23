package main

import (
	"log"
	"syscall"

)

func main() {
	sigs := newOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	dp := NewUSBDevicePlugin()
	dp.Serve()
	
	for {
		select {
		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				log.Println("Received SIGHUP, restarting.")
			default:
				log.Printf("Received signal \"%v\", shutting down.", s)
				return
			}
		}
	}
}

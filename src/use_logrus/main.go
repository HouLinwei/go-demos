package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Debug("A walrus appears")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Warn("A walrus appears")
}

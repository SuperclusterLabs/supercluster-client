package supercluster

import (
	firebase "firebase.google.com/go"
)

type DB struct {
	instance *firebase.App
}

var db DB

package main

import "time"

// User : a structure of users
type User struct {
	ID         int
	Name       string
	Avatar     string
	Email      string
	JoinDate   time.Time
	LastOnline time.Time
}

package main

var buddyListPath = "config/buddy.list"
var GlobalBuddyList Buddylist

type Buddylist struct {
	Buddys []Buddy
}

type Buddy struct {
	Username string
	Pubkey   string
	Friend   bool
}

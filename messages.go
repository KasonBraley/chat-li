package main

type Message struct {
	From    User
	To      Room
	Type    string
	Content string
}

type DirectMessage struct {
	Users    []User
	Messages []Message
}

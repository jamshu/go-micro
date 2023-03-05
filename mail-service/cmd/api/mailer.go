package main

type Mail struct {

	Domain string 
	Host string
	Port int
	Username string
	Password string
	Encryption string
	FromAddress string
	FromName string

}

type Message struct {

	From string
	FromName string
	To string
	Subject string
	Attachemes []string
	Data any
	DataMap map[string]any
}
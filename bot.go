package main

// you need to create a file called secrets.go and put in your API_KEY in there for this to work

import (
	"time"
	"log"
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  API_KEY,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hello!\nThe possible / commands are:\n1. hello (greets the user)\n2. time (tells the UTC time)")
		fmt.Println("Got a valid request for /start")
		fmt.Printf("Sender's ID: %d", m.Sender.ID)
		fmt.Println("\n")
	})

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello " + m.Sender.FirstName + m.Sender.LastName)
		fmt.Println("Got a valid request for /hello")
		fmt.Printf("Sender's ID: %d", m.Sender.ID)
		fmt.Println("\n")
	})
	
	b.Handle("/time", func(m *tb.Message) {
		b.Send(m.Sender, "The UTC time is now: " + time.Now().UTC().String())
		fmt.Println("Got a valid request for /time")
		fmt.Printf("Sender's ID: %d", m.Sender.ID)
		fmt.Println("\n")
	})

	b.Start()
}

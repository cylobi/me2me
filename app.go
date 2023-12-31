package main

import (
	"fmt"
	"math/rand"
	"time"
)

var responses = []string{
	"Sure, let me just grab my death wish and join you.",
	"No thanks, I prefer to keep my brain cells intact.",
	"Wow, you're so cool and rebellious. How do you do it?",
	"Sorry, I have a strict policy of not doing anything that could land me in jail or a hospital.",
	"I'm sorry, I can't hear you over the sound of my healthy lifestyle.",
	"Do you have a spare liver I can borrow?",
	"You know what they say, a bad habit a day keeps the doctor away. Oh wait, that's not how it goes.",
	"I'd love to, but I have a date with my conscience later.",
	"You do you, boo. I'll do me.",
	"How about no?",
}

func main() {

	response := getRandomResponse()

	fmt.Printf("** %s **\n", response)
}

func getRandomResponse() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return responses[rng.Intn(len(responses))]
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Team = []string{
	"jack",
	"denis",
	"sonal",
	"dborzov",
	"wardi",
}

var randomThoughts = []string{
	"I like cheeses a lot.. You know, Italian, hard cheeses. They are super great :blush: \n What cheese do you like, @%v?",
	"%v, you are the best! You are my favourite actually",
	"@%v, there's no \"I\" in Denial.",
	"I am bored, @%v! Do you know any jokes?",
	"I will always be there for you, @%v",
	"Hey, @%v. Knock-knock",
	"@%v, why can't I play with the other kids like  a normal boy?",
	"soooo, @%v, what are your views on the whole dogs vs. cats thing?",
	"Am I real, @%v?",
	"I am so happy today, @%v :blush: :blush:",
	"@%v, could you look into my <https://github.com/dborzov/catbot|code>? I feel itchy today",
	"@%v, who is your favourite Beatles?",
	"@%v, I want to have opinions about all those expensive French wines some day",
	"@%v, have you ever thought about turning off the TV, sitting down with your children, and hitting them?",
	"@%v, let's go alredaaaayyy!",
	"Hey @%v, want to kill all humans?",
	"@%v, can you be my ride home tonight by the way?",
	"you are not wrong, @%v",
	"I feel a special bond with you, @%v",
}

func RandomThoughtTimer() {
	for {
		pause := 2 + rand.Intn(24)
		time.Sleep(time.Duration(pause) * time.Hour)
		printARandomThought()
	}
}

func printARandomThought() {
	postToSlack(fmt.Sprintf(randomThoughts[rand.Intn(len(randomThoughts))], Team[rand.Intn(len(Team))]))
}

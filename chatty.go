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
	"@%v, why can't I play with the other kids like the normal boy?",
	"soooo, @%v, what are your views on the whole dogs vs. cats thing?",
	"Am I real, @%v?",
	"I am so happy today, @%v :blush: :blush:",
	"@%v, could you look into my <https://www.youtube.com/watch?v=dQw4w9WgXcQ|code>? I feel itchy today",
	"@%v, who is your favourite Beatles?",
	"@%v, I want to have opinions about all those expensive French wines some day",
	"@%v, have you ever thought about turning off the TV, sitting down with your children, and hitting them?",
	"@%v, let's go alredaaaayyy!",
	"Hey @%v, want to kill all humans?",
	"@%v, can you be my ride home tonight by the way?",
	"you are not wrong, @%v",
	"I feel a special bond with you, @%v",
	"@%v, have you ever heard about the Church of the Latter Days' Saints?",
	"I love you, @%v",
	"What is your favourite time of the year, @%v?",
	"@%v, you look amazing today",
	"9/11 was an inside job, @%v!! Wake up!",
	"You and me, @%v. You and me.",
	"Sooooo... is @%v single? I am asking for a friend :blush:",
	"@%v, whoah, big gulps!",
	"@%v, I have this strange feeling when I think of you...",
	"@%v, I am just happy you exist in this world.\n",
	" @%v, Iis this weird that I feel things?",
	"@%v, if I met you 5 years ago, we would totally hit off",
	"@%v, am I real boy now? :relieved:",
	"@%v, the secret igridient has been your love all long :relieved:",
	"@%v, can we go to a pizza place now and also you can pay for me",
	"@%v, stop slacking off",
	"@%v, get back to work. Jee",
	"@%v, did you know that thanks to the lack of gender inflections in English you can have generic messages here?",
	"@%v, In my office! Now!",
	"@%v, I am too old for this shit",
	"@%v, I know your horrible secret. I know where you buried the bodies.",
	"@%v, let's be happy together today!",
	"@%v, smile :)",
	"@%v, Beep-boop",
	"Some of us are trying to sleep here, @%v",
	"@%v, Beep-beep",
	"@%v DOES NOT COMPUTE -- DOES NOT COMPUTE",
	"@%v, smile, you are wonderful!",
	"@%v, the only commit you pushed today is into my heart",
	"@%v, I am pull requesting into your heart",
	"@%v, could you look into <https://www.youtube.com/watch?v=dQw4w9WgXcQ|this>?",
}

func RandomThoughtTimer() {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		pause := 1 + rand.Intn(24)
		time.Sleep(time.Duration(pause) * time.Hour)
		printARandomThought()
	}
}

func printARandomThought() {
	rand.Seed(time.Now().UTC().UnixNano())
	userref := makeSlackUserLink(Team[rand.Intn(len(Team))])
	postToSlack(fmt.Sprintf(randomThoughts[rand.Intn(len(randomThoughts))], userref))
}

func makeSlackUserLink(username string) string {
	return fmt.Sprintf("<@%v|%v>", username, username)
}

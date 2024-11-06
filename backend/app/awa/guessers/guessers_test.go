package guessers

import (
	"fmt"
	"testing"
)

func TestGuessNation(t *testing.T) {
	head, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	query := map[string]interface{}{
		"name":     "",
		"company":  "",
		"blog":     "",
		"location": "Budapest",
		"email":    "",
		"comments": []string{},
	}
	res, err := GuessNation(head, query)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

/*
{
	"name": "David J Wu",
	"company": "",
	"blog": "",
	"location": "",
	"email": "",
	"comments": [
		"Do you know anyone who would be interested to implement and main",
		"Thanks for the discussion. I'll investigate a little and conside",
		"I reviewed probably around 50 different 19x19 training games and",
		"@kaorahi - Very cool. Is this phenomenon consistent across multi",
		"Thanks for posting. Would you be able to include an SGF file?",
		"Thanks!",
		"Thanks for reporting! This will probably also join into a pile o",
		"Thanks for posting this. I will add more variations of this kind"
	]
}

*/

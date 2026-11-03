package main

import "testing"

type testCase struct {
	input    string
	expected string
}

func checkAll(t *testing.T, tests []testCase) {
	for _, test := range tests {
		got := process(test.input)
		if got != test.expected {
			t.Errorf("\nentree  : %q\nobtenu  : %q\nattendu : %q", test.input, got, test.expected)
		}
	}
}

func TestHexBin(t *testing.T) {
	checkAll(t, []testCase{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"zz (hex) stays", "zz stays"},
		{"(hex) alone", "alone"},
	})
}

func TestCase(t *testing.T) {
	checkAll(t, []testCase{
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"hi there (up, 10)", "HI THERE"},
	})
}

func TestPunctuation(t *testing.T) {
	checkAll(t, []testCase{
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"What !? Really", "What!? Really"},
		{"wait : here ; there", "wait: here; there"},
	})
}

func TestQuotes(t *testing.T) {
	checkAll(t, []testCase{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"' one ' and ' two '", "'one' and 'two'"},
		{"I don't know", "I don't know"},
	})
}

func TestArticles(t *testing.T) {
	checkAll(t, []testCase{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"a untold story", "an untold story"},
		{"a hour", "an hour"},
		{"a book", "a book"},
		{"this is a", "this is a"},
	})
}

func TestSubjectExamples(t *testing.T) {
	checkAll(t, []testCase{
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
	})
}

func TestEdgeCases(t *testing.T) {
	checkAll(t, []testCase{
		{"", ""},
		{"hello   world", "hello world"},
		{"hello (up)\nworld (cap)", "HELLO\nWorld"},
	})
}

package main

import "testing"

func TestValidateWord(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "valid word", word: "apple", want: true},
		{name: "another valid word", word: "crane", want: true},
		{name: "uppercase CLI guess", word: "SLATE", want: true},
		{name: "mixed case guess", word: "Slate", want: true},
		{name: "unknown word", word: "zzzzz", want: false},
		{name: "partial word", word: "app", want: false},
		{name: "word too long", word: "applesauce", want: false},
		{name: "empty input", word: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateWord(tt.word)
			if err != nil {
				t.Fatalf("validateWord(%q) returned an unexpected error: %v", tt.word, err)
			}
			if got != tt.want {
				t.Errorf("validateWord(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

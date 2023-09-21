package main

import (
	"log"
	"os"
	"testing"
)

func Test_buddyStrings(t *testing.T) {
	t.Run("should return true with aaaaaaabc", func(t *testing.T) {
		if g := buddyStrings("aaaaaaabc", "aaaaaaacb"); g != true {
			t.Errorf("buddyStrings() = %v, want %v", g, true)
		}
	})

	t.Run("should return true with 2 equal runes", func(t *testing.T) {
		if g := buddyStrings("aa", "aa"); g != true {
			t.Errorf("buddyStrings() = %v, want %v", g, true)
		}
	})

	t.Run("should return true with aaab", func(t *testing.T) {
		if g := buddyStrings("aaab", "aaab"); g != true {
			t.Errorf("buddyStrings() = %v, want %v", g, true)
		}
	})

	t.Run("should return true with long inputs", func(t *testing.T) {
		testFile1, err := os.ReadFile("./testdata/buddystrings-longinput-1.txt")
		if err != nil {
			log.Fatal(err)
		}

		testFile2, err := os.ReadFile("./testdata/buddystrings-longinput-2.txt")
		if err != nil {
			log.Fatal(err)
		}

		if g := buddyStrings(string(testFile1), string(testFile2)); g != true {
			t.Errorf("buddyStrings() = %v, want %v", g, true)
		}
	})
}

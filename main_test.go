package main_test

import (
	src "go-reloaded/core"
	"testing"
)

func TestSample1(t *testing.T) {
	correct := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
	filecontent := src.FileOpen("sample1.txt")
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)
	if filecontent != correct {
		t.Errorf("Program Output: %v \n", filecontent)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func TestSample2(t *testing.T) {
	correct := "Simply add 66 and 2 and you will see the result is 68."
	filecontent := src.FileOpen("sample2.txt")
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)
	if filecontent != correct {
		t.Errorf("Program Output: %v \n", filecontent)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func TestSample3(t *testing.T) {
	correct := "There is no greater agony than bearing an untold story inside you."
	filecontent := src.FileOpen("sample3.txt")
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)
	if filecontent != correct {
		t.Errorf("Program Output: %v \n", filecontent)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func TestSample4(t *testing.T) {
	correct := "Punctuation tests are... kinda boring, don't you think!?"
	filecontent := src.FileOpen("sample4.txt")
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)
	if filecontent != correct {
		t.Errorf("Program Output: %v \n", filecontent)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func TestSample5(t *testing.T) {
	correct := "As Elton John said this: 's I am the most well-known homosexual in the world'. I am exactly how they describe me: 'awesome'"
	filecontent := src.FileOpen("sample5.txt")
	filecontent = src.Converting(filecontent)
	filecontent = src.Letters(filecontent)
	filecontent = src.Punc(filecontent)
	filecontent = src.Quote(filecontent)
	filecontent = src.Vowel(filecontent)
	if filecontent != correct {
		t.Errorf("Program Output: %v \n", filecontent)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

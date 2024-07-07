package persiangenderdetection

import (
	"testing"
)

func TestClearName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  عــــلی  ", "علی"},
		{"بیــ🥲ــتا", "بیتا"},
		{"محــ🌚ــمد", "محمد"},
		{"۱۲۳مهناز۴۵۶", "مهناز"},
		{"مهدي-1980", "مهدی"},
		{"", ""},
	}

	for _, test := range tests {
		result := clearName(test.input)
		if result != test.expected {
			t.Errorf("clearName(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestGetGender(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"آرمان", "MALE"},
		{"بابك آذر مهر", "MALE"},
		{"بی بی سلطان", "FEMALE"},
		{"شیدا علیزاده", "FEMALE"},
		{"ممد رضا", "MALE"},
		{"پانته‌آ عَبّاسی", "FEMALE"},
		{"دکتر ندا حسینی", "FEMALE"}, // Name with prefix
		{"سید امیر موسوی", "MALE"},   // Name with prefix
		{"جناب آقای X", "MALE"},      // Name with prefix
		{"سرکار خانم Y", "FEMALE"},   // Name with prefix
		{"", "UNKNOWN"},              // Empty name
		{"0", "UNKNOWN"},             // Name not in dataset
	}

	for _, test := range tests {
		result := GetGender(test.name)
		if result != test.expected {
			t.Errorf("GetGender(%s) = %s; expected %s", test.name, result, test.expected)
		}
	}
}

func TestIsUnwantedRune(t *testing.T) {
	tests := []struct {
		r        rune
		expected bool
	}{
		{'!', true},       // Punctuation
		{'آ', false},      // Valid character
		{'ـ', false},      // Specific unwanted rune
		{'1', true},       // Digit
		{'\u200c', false}, // Zero width non-joiner
	}

	for _, test := range tests {
		result := isUnwantedRune(test.r)
		if result != test.expected {
			t.Errorf("isUnwantedRune(%c) = %t; expected %t", test.r, result, test.expected)
		}
	}
}

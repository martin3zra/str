package str_test

import (
	"github.com/martin3zra/str"
	"regexp"
	"testing"
)

func TestCamel(t *testing.T) {

	expected := "MustChange"
	got := str.Camel("must_change")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestSnake(t *testing.T) {
	expected := "must_change"
	got := str.Snake("MustChange")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestTrim(t *testing.T) {
	expected := "without_space"
	got := str.Trim(" without_space ")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestEquals(t *testing.T) {
	expected := "one-text"
	got := str.Trim("one-text")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}

	got = str.Trim("another-text")

	if got == expected {
		t.Errorf("got %s, want %s", got, expected)
	}

}

func TestAddslashes(t *testing.T) {
	expected := "one\\'text"
	got := str.Addslashes("one'text")

	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}

}

func TestUUID(t *testing.T) {
	regx := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	got := str.UUID()
	if !regx.MatchString(got) {
		t.Errorf("got %s, want %s", got, "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	}
}

func TestRandom(t *testing.T) {
	expected := "some-value"
	got := str.Random(10)
	if got == expected {
		t.Errorf("got %s, want %s", got, expected)
	}

	expected = got
	got = str.Random(10)
	if got == expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

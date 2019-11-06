package freemail

import "testing"

var testDomains = map[string]bool{
	"bademail.com":    true,
	"bad.email.co.uk": true,
}

func TestFind(t *testing.T) {
	cases := map[string]bool{
		"bademail.com":        true,
		"goodemail.com":       false,
		"sub.bademail.com":    true,
		"bad.email.co.uk":     true,
		"sub.bad.email.co.uk": true,
		"email.co.uk":         false,
	}
	for domain, expected := range cases {
		actual := find(domain, testDomains)
		if actual != expected {
			t.Errorf("For %s, expected %v was %v", domain, expected, actual)
		}
	}
}

func TestIsFree(t *testing.T) {
	if !IsFree("gmail.com") {
		t.Errorf("expected 'gmail.com' to be free")
	}
	if IsFree("google.com") {
		t.Errorf("expected 'google.com' not to be free")
	}
}

func TestIsDisposable(t *testing.T) {
	if !IsDisposable("mailinator.com") {
		t.Errorf("expected 'mailinator.com' to be disposable")
	}
	if IsDisposable("google.com") {
		t.Errorf("expected 'google.com' not to be disposable")
	}
}

func TestIsSpammy(t *testing.T) {
	if !IsSpammy("aribeth.ru") {
		t.Errorf("expected 'aribeth.ru' to be spammy")
	}
	if IsSpammy("google.com") {
		t.Errorf("expected 'google.com' not to be disposable")
	}
}

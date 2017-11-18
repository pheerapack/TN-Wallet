package middleware

import "testing"

type fakePin struct {
	pin int
}

func (p fakePin) Pin() int {
	return p.pin
}

func Test10002Pin1234Correct(t *testing.T) {
	pass := false
	validation := Validation{P: (fakePin{1234})}
	pass = validation.validatePin(10002, 1234)

	if pass != true {
		t.Error("Error")
	}
}

func Test10002Pin4321Incorrect(t *testing.T) {
	pass := false
	validation := Validation{P: (fakePin{4321})}
	pass = validation.validatePin(10002, 1234)

	if pass == true {
		t.Error("Error")
	}
}

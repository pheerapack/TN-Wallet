package middleware

type Validation struct {
	P
}

func (v Validation) getPin(wallid int) int {
	return v.Pin()
}

// interface
type P interface {
	Pin() int
}

type fakePin struct {
	pin int
}

func (p fakePin) Pin() int {
	return p.pin
}

func (v Validation) validatePin(wallid int, pinid int) bool {
	pinFromDb:=v.getPin(wallid)
	if pinFromDb==pinid {
		return true
	}
	return false
}

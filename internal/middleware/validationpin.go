package middleware

type Validation struct {
	P
}

func (v Validation) getPin(wallid int) int {
	return v.Pin()
}

type P interface {
	Pin() int
}

func (v Validation) validatePin(wallid int, pinid int) bool {
	pinFromDb := v.getPin(wallid)
	if pinFromDb == pinid {
		return true
	}
	return false
}

package middleware

type merchant struct {
	Wallid string
	Balance float32
}

func (m *merchant)Increase(bal float32) {
	m.Balance+=bal
}
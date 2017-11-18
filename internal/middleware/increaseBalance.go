package middleware

type merchant struct {
	Wallid string `json:"wallid"`
	Balance float32 `json:"balance"`
}

func (m *merchant)Increase(bal float32) {
	m.Balance+=bal
}

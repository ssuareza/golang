package wise

// Balance represents the balance of a Wise account
type Balance struct {
	ID       int64  `json:"id"`
	Currency string `json:"currency"`
	Amount   Amount `json:"amount"`
}

// Balances represents a slice of Balance
type Balances []Balance

// Amount represents the amount of a currency
type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

// Rate represents the exchange rate between two currencies
type Rate struct {
	Rate   float64 `json:"rate"`
	Source string  `json:"source"`
	Target string  `json:"target"`
}

// Rates represents a slice of Rate
type Rates []Rate

// Transaction represents a transaction
type Transaction struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Amount string `json:"primaryAmount"`
}

// Transactions represents a slice of Transaction
type Transactions struct {
	Activities []Transaction `json:"activities"`
}

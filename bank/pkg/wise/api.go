package wise

// transaction represents a transaction
type transaction struct {
	ID     string `json:"id"`
	Amount string `json:"primaryAmount"`
	Date   string `json:"createdOn"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Type   string `json:"type"`
}

// transactions represents a slice of Transaction
type transactions struct {
	Activities []transaction `json:"activities"`
}

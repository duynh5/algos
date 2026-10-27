package orderapi

// Product represents the small catalog entity used by the first lab.
type Product struct {
	ID   string
	SKU  string
	Name string
}

// Store is the deliberately minimal in-memory store that learners
// extend while working through the backend curriculum.
type Store struct {
	products map[string]Product
}

// NewStore creates an empty store with initialized writable state.
func NewStore() *Store {
	return &Store{products: make(map[string]Product)}
}

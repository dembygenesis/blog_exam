package store

// Store has our app methods
type Store interface {
	AdminStore
}

// AdminStore contains our
type AdminStore interface {
	Create() error
	Read(id int) error
}
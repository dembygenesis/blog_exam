package logic

import "fmt"

func (h *Handler) Create() error {

	return nil
}

func (h *Handler) Read(id int) error {
	fmt.Println("I READ")
	h.store.Read(3)
	return nil
}
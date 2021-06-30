package load

import "log"

type TransactionsTypes struct {
	ID   int
	Type string
}

func (t TransactionsTypes) Delete() error {
	panic("implement me")
}

func (t TransactionsTypes) Exists() error {
	panic("implement me")
}

func (t TransactionsTypes) Insert() error {
	log.Println(t.ID)
	log.Println(t.Type)
	return nil
}

func (t TransactionsTypes) Update() error {
	panic("implement me")
}

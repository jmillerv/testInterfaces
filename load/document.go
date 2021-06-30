package load

import "log"

type DocumentTypes struct {
	ID int
	Type string
}

func (d *DocumentTypes) Insert() error {
	log.Println(d.ID)
	log.Println(d.Type)
	log.Println("I'M DOCUMENT")
	return nil

}

func (d *DocumentTypes) Delete() error {
	 panic("implement me")
}

func (d *DocumentTypes) Update() error {
	 panic("implement me")
}

func (d *DocumentTypes) Exists() error {
	 panic("implement me")
}
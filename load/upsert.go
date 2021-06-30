package load

type Upsert interface {
	Delete() error
	Exists() error
	Insert() error
	Update() error
}





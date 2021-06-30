package load

import "testing"

func TestProcess(t *testing.T)  {
	var upserts = make([]Upsert, 0)
	u1 := &TransactionsTypes{
		ID:   1,
		Type: "foo",
	}
	upserts = append(upserts, u1)
	u2 := &TransactionsTypes{
		ID:   2,
		Type: "bar",
	}
	upserts = append(upserts, u2)

	u3 := &DocumentTypes {
		ID: 3,
		Type: "baz",
	}
	upserts = append(upserts, u3)

	for _, u := range upserts {
		_ = u.Insert()
	}
}

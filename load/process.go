package load

func Process(records []Upsert) error {
	for _, r := range records {
		err := r.Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

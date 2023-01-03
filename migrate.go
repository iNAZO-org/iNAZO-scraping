package main

func migrate() error {
	err := db.AutoMigrate(&GradeDistribution{})
	if err != nil {
		return err
	}

	return nil
}

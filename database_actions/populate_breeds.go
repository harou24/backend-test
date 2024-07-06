package database_actions

import (
	"database/sql"
	"encoding/csv"
	"os"
)

func LoadBreedsFromCSV(db *sql.DB, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // skip header
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		_, err = db.Exec(`INSERT INTO breeds (species, pet_size, name, average_male_adult_weight, average_female_adult_weight) VALUES (?, ?, ?, ?, ?)`,
			record[1], record[2], record[3], record[4], record[5])
		if err != nil {
			return err
		}
	}
	return nil
}

package operations

import (
	"database/sql"
)

func ListCountries(db *sql.DB) ([]Country, error) {

	rows, err := list(db, "countries")
	if err != nil {
		return nil, err
	}

	var countries []Country

	for rows.Next() {
		var country Country
		err := rows.Scan(&country.Id, &country.Name, &country.Abbreviation, &country.CountryCode)
		if err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}
	rows.Close()

	return countries, nil

}

func list(db *sql.DB, table string) (*sql.Rows, error) {

	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rows, err
}

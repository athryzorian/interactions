package operations

import (
	"database/sql"

	"github.com/athryzorian/interactions/backend/dal/datatypes"
)

func ListCountries(db *sql.DB) ([]datatypes.Country, error) {

	rows, err := list(db, "countries")
	if err != nil {
		return nil, err
	}

	var countries []datatypes.Country

	for rows.Next() {
		var country datatypes.Country
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

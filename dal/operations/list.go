package operations

import (
	"database/sql"
	"log"

	"github.com/athryzorian/interactions/dal/datatypes"
)

func ListCountries(db *sql.DB) ([]datatypes.Country, error) {

	rows, err := db.Query("SELECT * FROM " + "countries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Println("Fetched rows from countries table")

	var countries []datatypes.Country

	for rows.Next() {
		var country datatypes.Country
		err := rows.Scan(&country.Id, &country.Name, &country.Abbreviation, &country.CountryCode)
		if err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	log.Println("No of countries fetched:", len(countries))

	return countries, nil

}

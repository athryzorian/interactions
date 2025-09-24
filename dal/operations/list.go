package operations

import (
	"database/sql"
	"log"

	"github.com/athryzorian/interactions/dal/datatypes"
)

func ListCountries(db *sql.DB) ([]datatypes.Country, error) {

	queryStatement := "SELECT * FROM " + "countries"
	rows, err := db.Query(queryStatement)
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

func ListStates(db *sql.DB, queryFilter string) ([]datatypes.State, error) {

	queryStatement := "SELECT * FROM " + "states"
	queryStatement += " WHERE " + queryFilter

	rows, err := db.Query(queryStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Println("Fetched rows from states table")

	var states []datatypes.State

	for rows.Next() {
		var state datatypes.State
		err := rows.Scan(&state.Id, &state.Name, &state.Abbreviation, &state.ParentCountry)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}

	log.Println("No of states fetched:", len(states))

	return states, nil

}

func ListCities(db *sql.DB, queryFilter string) ([]datatypes.City, error) {

	queryStatement := "SELECT * FROM " + "cities"
	queryStatement += " WHERE " + queryFilter

	rows, err := db.Query(queryStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Println("Fetched rows from cities table")

	var cities []datatypes.City

	for rows.Next() {
		var city datatypes.City
		err := rows.Scan(&city.Id, &city.Name, &city.ParentState)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	log.Println("No of cities fetched:", len(cities))

	return cities, nil

}

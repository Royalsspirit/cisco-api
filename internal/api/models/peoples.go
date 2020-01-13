package models

import (
	"context"
	"database/sql"
	"fmt"
)

type vehicle struct {
	ID    string
	Name  string
	Model string
}
type specie struct {
	ID             string
	Name           string
	Classification string
}

// People a main field to return as responses
type People struct {
	Name      string `validate:"max=50,min=2"`
	Height    string `validate:"max=50,min=2"`
	Mass      string `validate:"max=50,min=2"`
	HairColor string `validate:"max=50,min=2"`
	SkinColor string `validate:"max=50,min=2"`
	EyeColor  string `validate:"max=50,min=2"`
	BirthYear string `validate:"max=50,min=2"`
	Gender    string `validate:"max=50,min=2"`
	Homeworld string `validate:"min=1"`
	Films     string
	Vehicles  []vehicle
	Species   []specie
	Starships string
	Created   string `validate:"max=50,min=2"`
	URL       string `validate:"min=1"`
	ID        string `json:"omitempty"`
}

// AllPeoples return people item with some details
func AllPeoples(db *sql.DB) ([]People, error) {
	peoples, err := peoples(db)

	if err != nil {
		return nil, err
	}

	for i := range peoples {
		currentValue := &peoples[i]

		vehicles, err := peopleVehicles(db, currentValue.ID)

		if err != nil {
			return nil, err
		}

		species, err := peopleSpecies(db, currentValue.ID)

		if err != nil {
			return nil, err
		}

		if vehicles != nil {
			currentValue.Vehicles = vehicles
		}

		if species != nil {
			currentValue.Species = species
		}

	}

	return peoples, nil
}

func peoples(db *sql.DB) ([]People, error) {
	var people People

	sqlPeople := `SELECT 
		name, 
		height, 
		COALESCE(mass,''), 
		hair_color, 
		skin_color, 
		eye_color, 
		birth_year, 
		gender, 
		homeworld, 
		COALESCE(films,''), 
		COALESCE(starships, ''),
		COALESCE(created, ''),
		url, 
		COALESCE(id,'') 
	FROM 
		people 
	`
	rows, err := db.Query(sqlPeople)

	if err != nil {
		return nil, err
	}

	var peopleResponse []People

	for rows.Next() {
		err := rows.Scan(&people.Name,
			&people.Height,
			&people.Mass,
			&people.HairColor,
			&people.SkinColor,
			&people.EyeColor,
			&people.BirthYear,
			&people.Gender,
			&people.Homeworld,
			&people.Films,
			&people.Starships,
			&people.Created,
			&people.URL,
			&people.ID)

		if err != nil {
			return nil, err
		}

		peopleResponse = append(peopleResponse, people)
	}

	return peopleResponse, nil
}

func updatePeople(db *sql.DB, value People, id string) (int64, error) {
	sqlPeople := `UPDATE people
	SET name = $2, 
	height = $3, 
	mass = $4, 
	hair_color = $5, 
	skin_color = $6, 
	eye_color = $7, 
	birth_year = $8, 
	gender = $9, 
	homeworld = $10, 
	films = $11, 
	starships = $12,
	created = $13,
	url = $14
where 
	id = $1`
	result, err := db.ExecContext(context.Background(),
		sqlPeople,
		id,
		value.Name,
		value.Height,
		value.Mass,
		value.HairColor,
		value.SkinColor,
		value.EyeColor,
		value.BirthYear,
		value.Gender,
		value.Homeworld,
		value.Films,
		value.Starships,
		value.Created,
		value.URL)

	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if rows > 0 {
		return rows, nil
	}
	return 0, fmt.Errorf("None fields effected")

}

func peopleVehicles(db *sql.DB, people string) ([]vehicle, error) {
	var vehicleItem vehicle

	sqlPeopleVehicles := `
	select 
		v.name, 
		v.model,
		v.id 
	FROM people_vehicles p 
	INNER JOIN vehicles v ON p.vehicles = v.id 
	WHERE p.people = $1`

	rowsPv, errPv := db.Query(sqlPeopleVehicles, people)

	if errPv != nil {
		return nil, errPv
	}
	var vehicles []vehicle

	for rowsPv.Next() {
		err := rowsPv.Scan(&vehicleItem.Name,
			&vehicleItem.Model, &vehicleItem.ID)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicleItem)
	}

	return vehicles, nil
}

func deletePeopleVehicles(db *sql.DB, ID string) (int64, error) {
	sqlPeople := `DELETE FROM people_vehicles WHERE people = $1`
	result, err := db.ExecContext(context.Background(), sqlPeople, ID)

	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err == nil {
		return rows, nil
	}

	return 0, fmt.Errorf("None fields effected")
}

func insertPeopleVehicles(db *sql.DB, vehicleID string, ID string) (int64, error) {
	sqlPeople := `INSERT INTO people_vehicles VALUES (?, ?)`
	result, err := db.Prepare(sqlPeople)
	if err != nil {
		return 0, err
	}
	result.Exec(ID, vehicleID)

	return 1, nil

}

func peopleSpecies(db *sql.DB, people string) ([]specie, error) {
	var specieItem specie
	sqlPeopleSpecies := `
		select 
			v.name, 
			v.classification, 
			v.id
		FROM people_species p 
		INNER JOIN species v ON p.species = v.id 
		WHERE p.people = $1`

	rowsPv, errPv := db.Query(sqlPeopleSpecies, people)

	if errPv != nil {
		return nil, errPv
	}

	var species []specie

	for rowsPv.Next() {
		err := rowsPv.Scan(&specieItem.Name,
			&specieItem.Classification, &specieItem.ID)

		if err != nil {
			return nil, err
		}

		species = append(species, specieItem)
	}

	return species, nil
}

func deletePeopleSpecies(db *sql.DB, ID string) (int64, error) {
	sqlPeople := `DELETE FROM people_species WHERE people = $1`
	result, err := db.ExecContext(context.Background(), sqlPeople, ID)

	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err == nil {
		return rows, nil
	}

	return 0, fmt.Errorf("None fields effected")
}

func insertPeopleSpecies(db *sql.DB, vehicleID string, ID string) (int64, error) {
	sqlPeople := `INSERT INTO people_species VALUES (?, ?)`
	result, err := db.Prepare(sqlPeople)
	if err != nil {
		return 0, err
	}
	result.Exec(ID, vehicleID)

	return 1, nil

}

// PeopleUpdateHandler update people
func PeopleUpdateHandler(db *sql.DB, value People, id string) error {

	_, peopleErr := updatePeople(db, value, id)

	if peopleErr != nil {
		return peopleErr
	}
	_, vehicleDelErr := deletePeopleVehicles(db, id)

	if vehicleDelErr != nil {
		return vehicleDelErr
	}

	_, speciesDelErr := deletePeopleSpecies(db, id)

	if speciesDelErr != nil {
		return speciesDelErr
	}

	for _, vehicle := range value.Vehicles {

		_, vehicleInsertErr := insertPeopleVehicles(db, vehicle.ID, id)

		if vehicleInsertErr != nil {
			return vehicleInsertErr
		}
	}

	for _, specy := range value.Species {

		_, speciesInsertErr := insertPeopleSpecies(db, specy.ID, id)

		if speciesInsertErr != nil {
			return speciesInsertErr
		}
	}

	return nil
}

//InsertPeople InsertPeople
func InsertPeople(db *sql.DB, value People) error {
	sqlPeople := `INSERT INTO people(name, height, mass, hair_color, skin_color, eye_color, birth_year, gender, Homeworld, films, starships, created, url) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	result, err := db.ExecContext(context.Background(),
		sqlPeople,
		value.Name,
		value.Height,
		value.Mass,
		value.HairColor,
		value.SkinColor,
		value.EyeColor,
		value.BirthYear,
		value.Gender,
		value.Homeworld,
		value.Films,
		value.Starships,
		value.Created,
		value.URL)

	if err != nil {
		return err
	}
	id, err := result.RowsAffected()

	if err == nil {
		for _, vehicle := range value.Vehicles {

			_, vehicleInsertErr := insertPeopleVehicles(db, vehicle.ID, string(id))

			if vehicleInsertErr != nil {
				return vehicleInsertErr
			}
		}

		for _, specy := range value.Species {

			_, speciesInsertErr := insertPeopleSpecies(db, specy.ID, string(id))

			if speciesInsertErr != nil {
				return speciesInsertErr
			}
		}
	}

	return err

}

//DeletePeople InsertPeople
func DeletePeople(db *sql.DB, id string) error {
	sqlPeople := `DELETE FROM people WHERE id = $1`

	result, err := db.ExecContext(context.Background(), sqlPeople, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		_, vehicleDelErr := deletePeopleVehicles(db, id)
		if vehicleDelErr != nil {
			return vehicleDelErr
		}

		_, speciesDelErr := deletePeopleSpecies(db, id)
		if speciesDelErr != nil {
			return speciesDelErr
		}
	}

	return nil
}

package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	sights "github.com/Safwood/go-server"
	"github.com/jmoiron/sqlx"
)

type ParkPostgres struct {
	db *sqlx.DB
}

func NewParkPostgres(db *sqlx.DB) *ParkPostgres {
	return &ParkPostgres{db}
}

func (r *ParkPostgres) CreatePark(userId int, park sights.Park) (int, error)  {
    jsonCoords, err := json.Marshal(park.Coords)
   if err != nil {
       log.Fatalf("Error converting to JSON: %v", err)
   }

	// начало транзакций
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	parkQuery := fmt.Sprintf("INSERT INTO %s (title, description, coords) VALUES ($1, $2, $3) RETURNING id", parksTable)
	row := tx.QueryRow(parkQuery, park.Title, park.Description, jsonCoords)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersParksQuery := fmt.Sprintf("INSERT INTO %s (user_id, park_id) VALUES ($1, $2)", usersParksTable)
	if _, err := tx.Exec(usersParksQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
	// конец транзакций
}

func (r *ParkPostgres) GetAllParks(userId int) ([]sights.Park, error) {
	var result []sights.Park
	query := fmt.Sprintf(`SELECT pl.id, pl.title, pl.description, pl.address, pl.coords FROM %s pl INNER JOIN %s up on pl.id = up.park_id WHERE up.user_id = $1`, parksTable, usersParksTable)
    
	err := r.db.Select(&result, query, userId)
	log.Print(result)

	return result, err
}

func (r *ParkPostgres) GetParkById(userId int, parkId int) (sights.GetParkOutput, error) {
	var park sights.GetParkOutput
	
	query := fmt.Sprintf(`SELECT pl.id, pl.title, pl.address, pl.description, pl.Coords FROM %s pl INNER JOIN %s up on pl.id = up.park_id WHERE up.user_id = $1 AND up.park_id = $2`, parksTable, usersParksTable)
    rows := r.db.QueryRow(query, userId, parkId)
	var jsonData []byte
	err := rows.Scan(&park.Id, &park.Title, &park.Description, &park.Address, &jsonData)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	
    json.Unmarshal(jsonData, &park.Coords)
	log.Println(&park.Coords)

	return park, err
}

func (r *ParkPostgres) DeletePark(userId, itemId int) (error)  {
	query := fmt.Sprintf("DELETE FROM %s pl USING %s up, %s ul WHERE pl.id = $1 AND up.park_id = pl.id AND up.user_id = $2", parksTable, usersParksTable, usersTable)
	_, err := r.db.Exec(query, itemId, userId)

	return err
}

func (r *ParkPostgres) UpdatePark(userId, parkId int, input sights.UpdateParkInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Coords != nil {
		jsonCoords, err := json.Marshal(input.Coords)
		if err != nil {
			log.Fatalf("Error converting to JSON: %v", err)
		}
		setValues = append(setValues, fmt.Sprintf("coords=$%d", argId))
		args = append(args, jsonCoords)
		argId++
	}

	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, *input.Address)
		argId++
	}

	
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s pl SET %s FROM %s up	WHERE pl.id = up.park_id AND up.user_id = $%d AND pl.id = $%d`,
		parksTable, setQuery, usersParksTable, argId, argId+1)
	args = append(args, userId, parkId)

	_, err := r.db.Exec(query, args...)
	return err
}
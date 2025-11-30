package data

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)



type Perfume struct {
	Id				string     			`json:"id"`
	HouseId   		*string   			`json:"house_id"`
	Name      		string   			`json:"name"`
	Year	  		*int      			`json:"year"`
	Concentration   *string      		`json:"concentration"`
	Description   	*string      		`json:"description"`
	ImageURL  		*string      		`json:"image_url"`
	Gender			*string      		`json:"gender"`
	Category		*string      		`json:"category"`
	Longevity		*string      		`json:"longevity"`
}

type PerfumeRepository struct {
	db *pgxpool.Pool
}

func NewPerfumeRepository(db *pgxpool.Pool) *PerfumeRepository {
	return &PerfumeRepository{
		db: db,
	}
}

func (r *PerfumeRepository) GetAllPerfumes(ctx context.Context, filters map[string]string) ([]Perfume, error) {

	var perfumes []Perfume

	query := `
		SELECT house_id, name, year, concentration, description, image_url, gender, category, longevity
		FROM perfumes
	`

	var args []interface {}
	var conditions []string
	argPos := 1

	if filters["q"] != "" {
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR description ILIKE $%d)", argPos, argPos+1))
		searchTerm := "%" + filters["q"] + "%"
		args = append(args, searchTerm, searchTerm)
		argPos += 2
	}
	if filters["year"] != "" {
		conditions = append(conditions, fmt.Sprintf("year = $%d", argPos))
		args = append(args, filters["year"])
		argPos++
	}
	if filters["concentration"] != "" {
		conditions = append(conditions, fmt.Sprintf("concentration = $%d", argPos))
		args = append(args, filters["concentration"])
		argPos++
	}
	if filters["gender"] != "" {
		conditions = append(conditions, fmt.Sprintf("gender = $%d", argPos))
		args = append(args, filters["gender"])
		argPos++
	}
	if filters["category"] != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", argPos))
		args = append(args, filters["category"])
		argPos++
	}
	if filters["longevity"] != "" {
		conditions = append(conditions, fmt.Sprintf("longevity ILIKE $%d", argPos))
		args = append(args, filters["longevity"])
		argPos++
	}


	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		fmt.Println("Error querying perfumes:", err)
		return perfumes, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Perfume
		err := rows.Scan(
			&p.HouseId, &p.Name, &p.Year, &p.Concentration, &p.Description,
			&p.ImageURL, &p.Gender, &p.Category, &p.Longevity,
		)
		if err != nil {
			fmt.Println("Error scanning perfume:", err)
			continue
		}
		perfumes = append(perfumes, p)
	}

	return perfumes, nil
	

}
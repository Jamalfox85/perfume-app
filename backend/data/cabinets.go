package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)



type Cabinet struct {
	Shelf   		[]Perfume 			`json:"shelf"`
	LayeringSets	[]LayeringSet 		`json:"layering_sets"`
	Collections		[]Collection 		`json:"collections"`
	Wishlist  		[]Perfume 			`json:"wishlist"`
}

type LayeringSetItem struct {
	ItemId   string			`json:"item_id"`
	Name     string			`json:"name"`
	ItemType string			`json:"item_type"`
}

type LayeringSet struct {
	Id    	string				`json:"id"`
	Name  	string				`json:"name"`
	Items 	[]LayeringSetItem	`json:"items"`
}

type CabinetRepository struct {
	db *pgxpool.Pool
}

func NewCabinetRepository(db *pgxpool.Pool) *CabinetRepository {
	return &CabinetRepository{
		db: db,
	}
}

func (r *CabinetRepository) GetProfileCabinet(ctx context.Context, userId string) (Cabinet, error) {

	shelfPerfumes, err := r.getShelfPerfumes(ctx, userId)
	if err != nil {
		return Cabinet{}, err
	}

	layeringSets, err := r.getLayeringSets(ctx, userId)
	if err != nil {
		return Cabinet{}, err
	}

	collections, err := r.getUserCollections(ctx, userId)
	if err != nil {
		return Cabinet{}, err
	}

	wishlistPerfumes, err := r.getWishlistPerfumes(ctx, userId)
	if err != nil {
		return Cabinet{}, err
	}

	cabinet := Cabinet{
		Shelf: shelfPerfumes,
		LayeringSets: layeringSets,
		Collections: collections,
		Wishlist: wishlistPerfumes,
	}

	return cabinet, nil

	
}

func (r *CabinetRepository) getShelfPerfumes(ctx context.Context, userID string) ([]Perfume, error) {
	query := `
		SELECT p.id, p.name, p.house_id, p.year, p.concentration
		FROM perfumes p
		JOIN user_perfumes up ON p.id = up.perfume_id
		WHERE up.user_id = $1
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		fmt.Println("Error querying shelf perfumes:", err)
		return nil, err
	}
	defer rows.Close()

	var perfumes []Perfume
	for rows.Next() {
		var perfume Perfume
		if err := rows.Scan(
			&perfume.Id,
			&perfume.Name,
			&perfume.HouseId,
			&perfume.Year,
			&perfume.Concentration,
		); err != nil {
			return nil, err
		}
		perfumes = append(perfumes, perfume)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return perfumes, nil
}


func (r *CabinetRepository) getLayeringSets(ctx context.Context, userID string) ([]LayeringSet, error) {
	// First, get all sets for the user
	setQuery := `
		SELECT id, name
		FROM layering_sets
		WHERE user_id = $1
	`
	rows, err := r.db.Query(ctx, setQuery, userID)
	if err != nil {
		fmt.Println("Error querying layering sets:", err)
		return nil, err
	}
	defer rows.Close()

	var sets []LayeringSet
	for rows.Next() {
		var set LayeringSet
		if err := rows.Scan(&set.Id, &set.Name); err != nil {
			fmt.Println("Error scanning layering set:", err)
			return nil, err
		}

		// Now get the items for this set, filtering to perfumes only
		itemQuery := `
			SELECT lsi.item_id, p.name, it.name as item_type
			FROM layering_set_items lsi
			JOIN item_types it ON lsi.item_type_id = it.id
			JOIN perfumes p ON lsi.item_id = p.id::uuid
			WHERE lsi.layering_set_id = $1
		`

		itemRows, err := r.db.Query(ctx, itemQuery, set.Id)
		if err != nil {
			fmt.Println("Error querying layering set items:", err)
			return nil, err
		}

		var items []LayeringSetItem
		for itemRows.Next() {
			var item LayeringSetItem
			if err := itemRows.Scan(&item.ItemId, &item.Name, &item.ItemType); err != nil {
				itemRows.Close()
				fmt.Println("Error scanning layering set item:", err)
				return nil, err
			}
			items = append(items, item)
		}
		itemRows.Close()

		set.Items = items
		sets = append(sets, set)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error after iterating layering sets:", err)
		return nil, err
	}

	return sets, nil
}

func (r *CabinetRepository) getUserCollections(ctx context.Context, userID string) ([]Collection, error) {
	// 1. Fetch all collections for the user
	setQuery := `
		SELECT id, name
		FROM user_collections
		WHERE user_id = $1
	`
	rows, err := r.db.Query(ctx, setQuery, userID)
	if err != nil {
		fmt.Println("Error querying user collections:", err)
		return nil, err
	}
	defer rows.Close()

	var collections []Collection
	for rows.Next() {
		var collection Collection
		if err := rows.Scan(&collection.Id, &collection.Name); err != nil {
			fmt.Println("Error scanning collection:", err)
			return nil, err
		}

		// 2. Fetch perfumes in this collection
		itemQuery := `
			SELECT p.id, p.name, p.house_id, p.year, p.concentration
			FROM user_collection_items uci
			JOIN perfumes p ON uci.perfume_id = p.id
			WHERE uci.collection_id = $1
		`
		itemRows, err := r.db.Query(ctx, itemQuery, collection.Id)
		if err != nil {
			fmt.Println("Error querying collection items:", err)
			return nil, err
		}

		var perfumes []Perfume
		for itemRows.Next() {
			var perfume Perfume
			if err := itemRows.Scan(
				&perfume.Id,
				&perfume.Name,
				&perfume.HouseId,
				&perfume.Year,
				&perfume.Concentration,
			); err != nil {
				itemRows.Close()
				fmt.Println("Error scanning collection perfume:", err)
				return nil, err
			}
			perfumes = append(perfumes, perfume)
		}
		itemRows.Close()

		collection.Perfumes = perfumes
		fmt.Println("Fetched collection:", collection.Name, "with", len(perfumes), "perfumes")
		collections = append(collections, collection)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error after iterating collections:", err)
		return nil, err
	}

	return collections, nil
}


func (r *CabinetRepository) getWishlistPerfumes(ctx context.Context, userID string) ([]Perfume, error) {
	query := `
		SELECT p.id, p.name, p.house_id, p.year, p.concentration
		FROM perfumes p
		JOIN user_wishlist_perfumes uwp ON p.id = uwp.perfume_id
		WHERE uwp.user_id = $1
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		fmt.Println("Error querying wishlist perfumes:", err)
		return nil, err
	}
	defer rows.Close()

	var perfumes []Perfume
	for rows.Next() {
		var perfume Perfume
		if err := rows.Scan(
			&perfume.Id,
			&perfume.Name,
			&perfume.HouseId,
			&perfume.Year,
			&perfume.Concentration,
		); err != nil {
			fmt.Println("Error scanning wishlist perfume:", err)
			return nil, err
		}
		perfumes = append(perfumes, perfume)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error after iterating wishlist perfumes:", err)
		return nil, err
	}

	return perfumes, nil
}

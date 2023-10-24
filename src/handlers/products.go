package handlers

import (
	"log"
	"project-a/types"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

var appcontext *types.ApplicationContext

const query_get_all_products_basic = "select id, name, price, rating from product"

func CreateProductHandler(ctx *types.ApplicationContext) {
	appcontext = ctx
}

func GetProductsBasicHandler(ctx *fiber.Ctx) error {

	rows, err := appcontext.DB.Query(query_get_all_products_basic)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}

	result := []types.ProductBasic{}

	for rows.Next() {
		var id string
		var name string
		var rating float32
		price := decimal.Decimal{}

		if err := rows.Scan(&id, &name, &price, &rating); err != nil {
			log.Fatal(err)
		}

		prod := types.ProductBasic{
			ID:     id,
			Name:   name,
			Price:  price,
			Rating: rating,
		}

		result = append(result, prod)
	}

	return ctx.JSON(result)
}

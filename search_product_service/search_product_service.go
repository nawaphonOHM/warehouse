package search_product_service

import (
	"fmt"
	"uuid"

	"github.com/nawaphonOHM/warehouse/db"
	"github.com/nawaphonOHM/warehouse/transaction"
)

func SearchProduct() {

	database := &db.DB{}

	for {
		fmt.Println("Do you want to export products? (yes/no)")

		var choice string

		_, _ = fmt.Scanf("%v", &choice)

		if choice == "no" {

			return
		}

		var productId uuid.UUID

		_, _ = fmt.Scanf("%v", &productId)

		found := false

		transaction.Transaction(func() {
			found, _ = database.SearchProductId(productId)
		})

		if found {
			fmt.Printf("ProductId: %v found", productId)
		} else {
			fmt.Printf("ProductId: %v not found", productId)
		}
	}

}

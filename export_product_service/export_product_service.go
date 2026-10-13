package export_product_service

import (
	"fmt"
	"time"
	"uuid"

	"github.com/nawaphonOHM/warehouse/db"
	"github.com/nawaphonOHM/warehouse/storage"
	"github.com/nawaphonOHM/warehouse/transaction"
)

func ExportProduct() {

	database := &db.DB{}
	space := new(storage.Storages)

	for {

		fmt.Println("Do you want to update product's export code? (yes/no)")

		var choice string

		_, _ = fmt.Scanf("%v", &choice)

		if choice == "no" {

			break
		}

		var productId uuid.UUID
		var exportCode uuid.UUID

		_, _ = fmt.Scanf("%v", &productId)
		_, _ = fmt.Scanf("%v", &exportCode)

		transaction.Transaction(func() {

			found, location := database.SearchProductId(productId)

			if !found {
				fmt.Printf("product id: %v, reason: not found", productId)
				return
			}

			if match := database.ProductIdAndExportCodeMatch(productId, exportCode); !match {
				fmt.Printf("product id: %v, reason: export code does not match", productId)
				return
			}

			message := fmt.Sprintf(
				"Record: At %v, ProductId %v has been exported from the warehouse in position between (%v, %v, %v) and (%v, %v, %v)",
				time.Now(),
				productId,
				location.Start().X(),
				location.Start().Y(),
				location.Start().Z(),
				location.End().X(),
				location.End().Y(),
				location.End().Z(),
			)
			space.PopProduct(storage.NewStorage(location.End(), location.Start(), productId))
			database.RemoveProduct(productId)
			database.RemoveProductAndLocationTable(productId)

			database.SaveOnProductTransactionTable(db.NewProductTransactionRow(db.Success, message))

		})

	}

	return
}

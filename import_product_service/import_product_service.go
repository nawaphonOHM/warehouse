package import_product_service

import (
	"fmt"
	"time"

	"github.com/nawaphonOHM/warehouse/db"
	"github.com/nawaphonOHM/warehouse/product"
	"github.com/nawaphonOHM/warehouse/storage"
	"github.com/nawaphonOHM/warehouse/transaction"
)

func ImportProduct() {

	database := &db.DB{}
	space := new(storage.Storages)

	for {

		fmt.Println("Do you want to import products? (yes/no)")

		var choice string

		_, _ = fmt.Scanf("%v", &choice)

		if choice == "no" {

			break
		}

		var product_ *product.Product

		_, _ = fmt.Scanf("%v", &product_)

		transaction.Transaction(func() {

			if yes := database.HasSufficientSpace(product_.Width(), product_.Length(), product_.Height()); !yes {

				message := fmt.Sprintf("product id: %v, reason: insufficient space", product_.ProductId())

				database.SaveOnProductTransactionTable(db.NewProductTransactionRow(db.Error, message))

				return
			}

			start, end := database.AllocateSpace(product_.Width(), product_.Length(), product_.Height())
			space.PutProduct(storage.NewStorage(end, start, product_.ProductId()))

			database.SaveOnProductAndLocationTable(db.NewProductAndLocationRow(product_.ProductId(), start, end))

			message := fmt.Sprintf(
				"Record: At %v, ProductId %v has been imported to the warehouse in position between (%v, %v, %v) and (%v, %v, %v) and will be marked exported code %v",
				time.Now(),
				product_.ProductId(),
				start.X(),
				start.Y(),
				start.Z(),
				end.X(),
				end.Y(),
				end.Z(),
				database.SaveExportCode(product_.ProductId(), "South"),
			)
			database.SaveOnProductTransactionTable(db.NewProductTransactionRow(db.Success, message))

		})

	}

	return
}

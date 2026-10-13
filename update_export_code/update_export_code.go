package update_export_code

import (
	"fmt"
	"time"
	"uuid"

	"github.com/nawaphonOHM/warehouse/db"
	"github.com/nawaphonOHM/warehouse/transaction"
)

func UpdateExportCode() {

	database := &db.DB{}

	for {

		var productId uuid.UUID
		var exportCode uuid.UUID

		_, _ = fmt.Scanf("%v", &productId)
		_, _ = fmt.Scanf("%v", &exportCode)

		transaction.Transaction(func() {

			found, _ := database.SearchProductId(productId)

			if !found {
				fmt.Printf("product id: %v, reason: not found", productId)
				return
			}

			database.UpdateProductIdSExportCode(productId, exportCode)

			message := fmt.Sprintf(
				"Record: At %v, ProductId %v has been updated with export code %v",
				time.Now(),
				productId,
				exportCode,
			)

			database.SaveOnProductTransactionTable(db.NewProductTransactionRow(db.Success, message))
		})

	}
}

package warehouse

import (
	"fmt"
	"os"

	"github.com/nawaphonOHM/warehouse/export_product_service"
	"github.com/nawaphonOHM/warehouse/import_product_service"
	"github.com/nawaphonOHM/warehouse/search_product_service"
	"github.com/nawaphonOHM/warehouse/update_export_code"
)

func FrontDesk() {

	for {
		fmt.Println("Do you want to use our warehouse?")

		var choice string

		_, _ = fmt.Scanf("%s", &choice)

		if choice == "no" {
			os.Exit(0)
		}

		fmt.Println("What service do you want to use?")

		_, _ = fmt.Scanf("%s", &choice)

		if choice == "importProduct" {
			import_product_service.ImportProduct()
			continue
		}

		if choice == "exportProduct" {
			export_product_service.ExportProduct()
			continue
		}

		if choice == "searchProduct" {
			search_product_service.SearchProduct()
			continue
		}

		if choice == "updateProductSExportCode" {
			update_export_code.UpdateExportCode()
			continue
		}

	}
}

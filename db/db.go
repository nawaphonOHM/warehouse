package db

import (
	"uuid"

	"github.com/nawaphonOHM/warehouse/position"
)

type DB struct {
	productAndLocation ProductAndLocationTable

	productTransaction ProductTransactionTable

	exportCode ExportCodeTable
}

func (d *DB) SaveExportCode(productId uuid.UUID, detail string) uuid.UUID {
	new_ := NewExportCodeRow(productId, detail)

	d.exportCode = append(d.exportCode, new_)

	return new_.exportCode
}

func (d *DB) SearchProductId(productId uuid.UUID) (found bool, location *Location) {

	for _, row := range d.productAndLocation {
		if row.ProductId() == productId {
			return true, row.GetLocation()
		}
	}

	return false, nil
}

func (d *DB) ProductIdAndExportCodeMatch(productId uuid.UUID, exportCode uuid.UUID) bool {

	for _, row := range d.exportCode {
		if row.ProductID() == productId && row.ExportCode() == exportCode {
			return true
		}
	}

	return false
}

func (d *DB) UpdateProductIdSExportCode(productId uuid.UUID, newExportCode uuid.UUID) {

	index := -1

	for iindex, row := range d.exportCode {
		if row.ProductID() == productId {
			index = iindex
			break
		}
	}

	d.exportCode[index].SetExportCode(newExportCode)
}

func (d *DB) HasSufficientSpace(width float64, length float64, heigh float64) bool {

	return d.productAndLocation.CanFindConsecutiveSpace(width, length, heigh)
}

func (d *DB) SaveOnProductTransactionTable(new *ProductTransactionRow) {

	d.productTransaction = append(d.productTransaction, new)

}

func (d *DB) SaveOnProductAndLocationTable(new *ProductAndLocationRow) {

	d.productAndLocation = append(d.productAndLocation, new)

}

func (d *DB) RemoveProductAndLocationTable(pl uuid.UUID) {
	index := -1

	for i, row := range d.productAndLocation {
		if row.ProductId() == pl {
			index = i
			break
		}
	}

	d.productAndLocation = append(d.productAndLocation[:index], d.productAndLocation[index+1:]...)
}

func (d *DB) RemoveProduct(productId uuid.UUID) {
	index := -1

	for i, row := range d.productAndLocation {
		if row.productId == productId {
			index = i
			break
		}
	}

	d.productAndLocation = append(d.productAndLocation[:index], d.productAndLocation[index+1:]...)
}

func (d *DB) AllocateSpace(width float64, length float64, heigh float64) (start *position.Position, end *position.Position) {

	return position.FindBestLocation(width, length, heigh)
}

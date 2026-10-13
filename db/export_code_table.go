package db

import "uuid"

type ExportCodeTable []*ExportCodeRow

type ExportCodeRow struct {
	exportCode uuid.UUID
	productID  uuid.UUID

	detail string
}

func (e *ExportCodeRow) SetExportCode(exportCode uuid.UUID) {
	e.exportCode = exportCode
}

func (e *ExportCodeRow) ExportCode() uuid.UUID {
	return e.exportCode
}

func (e *ExportCodeRow) ProductID() uuid.UUID {
	return e.productID
}

func NewExportCodeRow(productID uuid.UUID, detail string) *ExportCodeRow {
	return &ExportCodeRow{exportCode: uuid.New(), productID: productID, detail: detail}
}

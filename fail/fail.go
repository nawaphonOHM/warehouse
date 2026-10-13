package fail

import (
	"fmt"
	"uuid"
)

type Fail struct {
	productId uuid.UUID

	reason string
}

func (f *Fail) String() string {
	return fmt.Sprintf("productId: %v, reason: %v", f.productId, f.reason)
}

func NewFail(reason string, productId uuid.UUID) *Fail {
	return &Fail{reason: reason, productId: productId}
}

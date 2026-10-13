package fail_container

import (
	"fmt"
	"strings"

	"github.com/nawaphonOHM/warehouse/fail"
)

type FailContainer []*fail.Fail

func (c *FailContainer) String() string {

	strBuilder := strings.Builder{}

	for _, f := range *c {
		strBuilder.WriteString(fmt.Sprintf("%v", f))
	}

	return fmt.Sprintf("%v", strBuilder.String())
}

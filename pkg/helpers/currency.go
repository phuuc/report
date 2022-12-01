package helpers

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/finnpn/overview/pkg/constants"
	"github.com/shopspring/decimal"
)

func ToVND(num decimal.Decimal) string {
	return fmt.Sprintf("%s %s", humanize.Comma(num.IntPart()), constants.VND)
}

func ToIntVND(num int64) string {
	return fmt.Sprintf("%s %s", humanize.Comma(num), constants.VND)
}

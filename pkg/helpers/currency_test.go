package helpers

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestToVND(t *testing.T) {
	type args struct {
		sum decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				sum: decimal.NewFromInt(2222323243325325325),
			},
			want: decimal.NewFromInt(222232324332532535).String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToVND(tt.args.sum); got != tt.want {
				t.Errorf("ToVND() = %v, want %v", got, tt.want)
			}
		})
	}
}

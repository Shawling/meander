package meander

import (
	"errors"
	"strings"
)

// Cost 表示消费等级，数字范围小，用 int8 即可
type Cost int8

func (c Cost) String() string {
	return strings.Repeat("$", int(c))
}

func ParseCost(s string) Cost {
	return Cost(len(s))
}

// 0值不用，所以用_, 其余的自增1
const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}
func ParseCostRange(s string) (CostRange, error) {
	var r CostRange
	segs := strings.Split(s, "...")
	if len(segs) != 2 {
		return r, errors.New("invalid cost range")
	}
	r.From = ParseCost(segs[0])
	r.To = ParseCost(segs[1])
	return r, nil
}

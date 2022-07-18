package types

import (
	"fmt"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


type Periods []Period


func (p Period) Duration() time.Duration {
	return time.Duration(p.Length) * time.Second
}


func (p Period) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}


func (p Periods) TotalLength() int64 {
	var total int64
	for _, period := range p {
		total += period.Length
	}
	return total
}


func (p Periods) TotalDuration() time.Duration {
	len := p.TotalLength()
	return time.Duration(len) * time.Second
}


func (p Periods) TotalAmount() sdk.Coins {
	total := sdk.Coins{}
	for _, period := range p {
		total = total.Add(period.Amount...)
	}
	return total
}


func (p Periods) String() string {
	periodsListString := make([]string, len(p))
	for _, period := range p {
		periodsListString = append(periodsListString, period.String())
	}

	return strings.TrimSpace(fmt.Sprintf(`Vesting Periods:
		%s`, strings.Join(periodsListString, ", ")))
}

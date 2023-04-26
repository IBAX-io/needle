package vm

import (
	"github.com/IBAX-io/needle/compile"
)

const (
	// CostCall is the cost of the function calling
	CostCall = 50
	// CostContract is the cost of the contract calling
	CostContract = 100
	// CostExtend is the cost of the extend function calling
	CostExtend = 10
)

func ContractBaseCost(cb *compile.CodeBlock) int64 {
	var cost int64
	c := cb.GetContractInfo()
	if c != nil {
		cost += int64(len(cb.Objects) * CostCall)
		cost += int64(len(c.Settings) * CostCall)
		if c.Tx != nil {
			cost += int64(len(*c.Tx) * CostExtend)
		}
	}
	return cost
}

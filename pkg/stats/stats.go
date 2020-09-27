package stats

import (
	"github.com/RAZ-os/bank/pkg/types"
)

// Avg - is thah func
func Avg(payments []types.Payment) types.Money {
	sum := 0

	if len(payments) == 1 {
		return payments[0].Amount
	}
	
	for _, payment := range payments {
		sum = sum + int(payment.Amount)
	}

	avg := types.Money(sum/len(payments)) 

	return avg
}

// TotalInCategory - is thah func
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0

	if len(payments) == 1 {
		return payments[0].Amount
	}
	
	for _, payment := range payments {
		if category == payment.Category {
			sum = sum + int(payment.Amount)
		} 
	}

	return types.Money(sum)
}
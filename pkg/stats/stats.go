package stats

import (
	"github.com/RAZ-os/bank/v2/pkg/types"
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

//CategoriesAvg calculates avg payment for every category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	sumCategories := map[types.Category]types.Money{}
	countCategories := map[types.Category]types.Money{}
	avgAmount := map[types.Category]types.Money{}

	for _, payment := range payments {
		sumCategories[payment.Category] += payment.Amount
		countCategories[payment.Category]++
	}
	//fmt.Println(sumCategories, countCategories)

	for keyValue := range sumCategories {
		avgAmount[keyValue] = sumCategories[keyValue] / countCategories[keyValue]
	}

	return avgAmount
}
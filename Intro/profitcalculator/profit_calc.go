package profitcalculator

import (
	"example.com/investment_calculator/functions"
)


func CalculateProfit() (ebt, profit, ratio float64){
	var revenue, expenses, taxRate float64

	revenue = functions.GetUserInput("Enter your revenue: ");
	expenses = functions.GetUserInput("Enter your expenses: ");
	taxRate = functions.GetUserInput("Enter your tax-rate: ");
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit

	return 
}
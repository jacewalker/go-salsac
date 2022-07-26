/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculates your salary sacrifice",
	Long:  "With this tool you can calculate your salary sacrifice with your annual salary and purchase price.",
	Run: func(cmd *cobra.Command, args []string) {
		grossIncome, _ := cmd.Flags().GetInt("salary")
		price, _ := cmd.Flags().GetInt("price")
		// duration, _ := cmd.Flags().GetInt("duration")

		totalTax := calculateAnnualTax(grossIncome)
		purchasePrice := calculateSacrifice(grossIncome, totalTax, price)

		fmt.Println("##########################")
		fmt.Printf("Purchase Price: $%d\nSavings: $%d\n", purchasePrice, price-purchasePrice)
		fmt.Println("##########################")
	},
}

func calculateAnnualTax(salary int) int {
	// Calculate the annual tax based off Australian tax laws
	// https://www.ato.gov.au/individuals/income-tax/income-tax-calculator

	floatSalary := float32(salary)
	var tax float32

	switch {
	case floatSalary <= 18_200:
		tax = 0
	case floatSalary <= 45_000:
		tax = (floatSalary - 18200) * 0.19
	case floatSalary <= 120_000:
		tax = (floatSalary-45000)*0.325 + 5092
	case floatSalary <= 180_000:
		tax = (floatSalary-120000)*0.37 + 29467
	case floatSalary >= 180_001:
		tax = (floatSalary-180000)*0.45 + 51667
	}
	return int(tax)
}

func calculateSacrifice(gross int, tax int, price int) int {
	newSalary := gross - price
	newTax := calculateAnnualTax(newSalary)
	purchasePrice := (gross - tax) - (newSalary - newTax)

	return purchasePrice
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calculateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	calculateCmd.Flags().IntP("salary", "s", 0, "Your annual gross salary")
	calculateCmd.Flags().IntP("price", "p", 0, "The purchase price of the item")
	calculateCmd.Flags().IntP("duration", "d", 0, "The duration of payment in weeks (eg 1 week, 4 weeks).")
}

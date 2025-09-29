package main

import (
	"flag"
	"fmt"
	"main/utils"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' subcommand")
		return
	}

	switch os.Args[1] {
	case "add":
		expenses, err := utils.LoadExpenses()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Expense description")
		amount := addCmd.Float64("amount", 0, "Expense amount")
		category := addCmd.String("category", "General", "Expense category")

		addCmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Error: --description and positive --amount required")
			return
		}
		// expenses := []Expense{}
		all_expenses, newExpense := utils.AddExpenses(expenses, *description, *category, *amount)

		if err := utils.SaveExpenses(all_expenses); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.ID)

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

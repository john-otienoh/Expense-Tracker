package main

import (
	"flag"
	"fmt"
	"main/utils"
	"os"
	// "strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' subcommand")
		return
	}
	expenses, err := utils.LoadExpenses()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch os.Args[1] {
	case "add":

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
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listCmd.Parse(os.Args[2:])
		utils.ListExpenses(expenses)

	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.Int("month", 0, "Filter summary by month (1-12)")
		summaryCmd.Parse(os.Args[2:])
		utils.Summary(expenses, *month)

	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.Int("id", 0, "Expense ID")
		newDescription := updateCmd.String("description", "", "Expense description")
		newAmount := updateCmd.Float64("amount", 0, "Expense amount")
		newCategory := updateCmd.String("category", "General", "Expense category")
		updateCmd.Parse(os.Args[2:])
		if *newDescription == "" || *newAmount <= 0 || *id <= 0 {
			fmt.Println("Error: --description and positive --amount and --id required")
			return
		}
		expensesUpdated, ok := utils.UpdateExpenses(expenses, *id, *newDescription, *newCategory, *newAmount)
		if !ok {
			fmt.Printf("Expense with ID %d not found.\n", *id)
			return
		}
		utils.SaveExpenses(expensesUpdated)
		fmt.Println("Expenses updated successfully.")

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "Expense ID")
		deleteCmd.Parse(os.Args[2:])
		if *id <= 0 {
			fmt.Println("Error: --id must be a positive number")
			return
		}
		expensesRemained, err := utils.DeleteExpenses(expenses, *id)
		if err != nil {
			fmt.Println("error: ", err)
		}
		if err := utils.SaveExpenses(expensesRemained); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Expense deleted successfully")
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

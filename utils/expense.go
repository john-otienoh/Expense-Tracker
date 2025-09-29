package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
}

type ExpenseBook struct {
	Expenses []Expense `json:"expenses"`
}

type Budget struct {
	Month int     `json:"month"`
	Year  int     `json:"year"`
	Limit float64 `json:"limit"`
}

func GetCurrentTime() time.Time {
	return time.Now()
}
func FormatDate() string {
	return GetCurrentTime().Format("02.01.2006")
}

func AddExpenses(expenses []Expense, description, category string, amount float64) ([]Expense, Expense) {
	id := 1
	if len(expenses) > 0 {
		id = expenses[len(expenses)-1].ID + 1
	}

	expense := Expense{
		ID:          id,
		Date:        FormatDate(),
		Amount:      amount,
		Description: description,
		Category:    category,
	}
	expenses = append(expenses, expense)
	return expenses, expense
}

func UpdateExpenses(expenses []Expense, id int, newDescription, newCategory string, newAmount float64) ([]Expense, error) {
	for i := range expenses {
		if expenses[i].ID == id {
			expenses[i].Description = newDescription
			expenses[i].Category = newCategory
			expenses[i].Amount = newAmount
			expenses[i].Date = FormatDate()
			return expenses, nil
		}
	}
	return expenses, fmt.Errorf("expense with ID %d not found", id)
}

func DeleteExpenses(expenses []Expense, id int) ([]Expense, error) {
	for i := range expenses {
		if expenses[i].ID == id {
			return append(expenses[:i], expenses[i+1:]...), nil
		}
	}
	return expenses, fmt.Errorf("task with ID %d not found", id)
}

func ListExpenses(expenses []Expense) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDescription\tAmount\tCategory\tDate")
	for _, expense := range expenses {
		fmt.Fprintf(
			w, "%d\t%s\t%.2f\t%s\t%s\n",
			expense.ID, expense.Description,
			expense.Amount, expense.Category,
			expense.Date,
		)

	}
}

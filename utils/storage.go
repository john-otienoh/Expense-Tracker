package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "expenses.json"

func LoadExpenses() ([]Expense, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, fmt.Errorf("error checking file: %v", err)
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(data) == 0 {
		return []Expense{}, nil
	}

	var expenses []Expense
	unmarshalErr := json.Unmarshal(data, &expenses)
	if unmarshalErr != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", unmarshalErr)
	}
	return expenses, nil
}

func SaveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding json: %v", err)
	}
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}

package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateCSV() error {
	newFile, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)

	defer writer.Flush()

	headers := []string{
		"Name",
		"Sugar g",
		"Fiber g",
		"Sodium mg",
		"Potassium mg",
		"Total Fat g",
		"Total Sat Fat g",
		"Calories",
		"Cholestoral mg",
		"Protein g",
		"Carbs g",
	}
	writer.Write(headers)
	return nil
}

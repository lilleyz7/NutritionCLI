package handlers

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func ViewHandler() error {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleDouble)

	headerRow := table.Row{}

	for _, record := range records[0] {
		headerRow = append(headerRow, record)
	}
	t.AppendHeader(headerRow)

	innerRows := []table.Row{}

	for i := 1; i < len(records); i++ {
		localRow := table.Row{}
		for _, lrecord := range records[i] {
			localRow = append(localRow, lrecord)
		}
		innerRows = append(innerRows, localRow)
	}

	t.AppendRows(innerRows)

	t.Render()
	return nil
}

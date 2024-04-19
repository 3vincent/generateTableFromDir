package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	// Parse command-line flags
	format := flag.String("format", "csv", "Output format (csv or xlsx)")
	output := flag.String("output", "", "Output file name")
	flag.Parse()

	// Check if the correct number of command-line arguments is provided
	if len(flag.Args()) != 1 {
		fmt.Println("Usage: generateTableFromDir --format <csv|xlsx> --output <filename> <directory>")
		return
	}

	// Get the directory path from the command-line arguments
	dir := flag.Arg(0)

	switch *format {
	case "csv":
		createCSV(dir, *output)
	case "xlsx":
		createXLSX(dir, *output)
	default:
		fmt.Println("Invalid output format. Please use 'csv' or 'xlsx'.")
	}
}

func createCSV(dir, output string) {
	// Open the CSV file for writing
	file, err := createUniqueFile(output, ".csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to CSV file
	writer.Write([]string{"File Name", "Path", "Is Directory"})

	// List files and subdirectories in the specified directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files and subdirectories
	for _, file := range files {
		// Get file info
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}

		// Write file info to CSV file
		writer.Write([]string{info.Name(), filepath.Join(dir, file.Name()), fmt.Sprintf("%t", info.IsDir())})
	}

	fmt.Println("CSV file created successfully!")
}

func createXLSX(dir, output string) {
	// Open the XLSX file for writing
	file, err := createUniqueFile(output, ".xlsx")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a new XLSX file
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("Files")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return
	}

	// List files and subdirectories in the specified directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files and subdirectories
	for _, file := range files {
		// Get file info
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}

		// Add row to sheet
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = info.Name()

		cell = row.AddCell()
		cell.Value = filepath.Join(dir, file.Name())

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%t", info.IsDir())
	}

	// Save XLSX file
	err = xlsxFile.Save(output)
	if err != nil {
		fmt.Println("Error saving XLSX file:", err)
		return
	}

	fmt.Println("XLSX file created successfully!")
}

func createUniqueFile(output, ext string) (*os.File, error) {
	// If output file name is not provided, use a default name
	if output == "" {
		output = "file_list" + ext
	}

	// Check if the file already exists
	if _, err := os.Stat(output); os.IsNotExist(err) {
		// File does not exist, create it
		return os.Create(output)
	}

	// File already exists, find a unique name by appending a number
	base := strings.TrimSuffix(output, ext)
	for i := 1; ; i++ {
		newName := fmt.Sprintf("%s_%d%s", base, i, ext)
		if _, err := os.Stat(newName); os.IsNotExist(err) {
			// Unique name found, create the file
			return os.Create(newName)
		}
	}
}

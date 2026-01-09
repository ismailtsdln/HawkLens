package analytics

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

func ExportToJSON(filename string, results []plugins.Result) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

func ExportToCSV(filename string, results []plugins.Result) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{"Platform", "DataType", "Data"})

	for _, res := range results {
		dataBytes, _ := json.Marshal(res.Data)
		writer.Write([]string{res.Platform, res.DataType, string(dataBytes)})
	}

	return nil
}

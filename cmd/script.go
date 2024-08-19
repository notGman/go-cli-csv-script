package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/notgman/go-cli-script/survey"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Run a script and get a CSV file",
	Long: `Run a script and get a CSV file
	
	Example:
	$ go-cmd script 
	`,

	Run: generateCSV,
}

func init() {
	rootCmd.AddCommand(scriptCmd)
}

func generateCSV(cmd *cobra.Command, args []string) {
	fmt.Println("Running script...")

	url := survey.StringPrompt("Enter the URL: ")
	output := survey.StringPrompt("Enter the output file name: ")

	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(data.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	keys := []string{}
	for key := range result {
		keys = append(keys, key)
	}
	dataName := survey.SingleSelect("Which field do you want?", keys)

	dataItems, ok := result[dataName].([]interface{})
	if !ok {
		log.Fatalf("Could not find data with name %s", dataName)
	}

	headers := []string{}
	firstItem := dataItems[0].(map[string]interface{})
	for key := range firstItem {
		headers = append(headers, key)
	}

	headers = survey.Checkboxes("Which are the fields you want? (select multiple options)", headers)

	file, err := os.Create(output + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)

	for _, item := range dataItems {
		values := []string{}
		for _, key := range headers {
			values = append(values, fmt.Sprintf("%v", item.(map[string]interface{})[key]))
		}
		writer.Write(values)
	}

	fmt.Println("Written to file", output+".csv")

}

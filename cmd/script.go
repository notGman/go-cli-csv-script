package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Run a script and get a CSV file",
	Long: `Run a script and get a CSV file
	
	Example:
	$ go-cmd script -u <https://testurl.com> -o <script.csv>
	`,

	Run: generateCSV,
}

func init() {
	rootCmd.AddCommand(scriptCmd)
	scriptCmd.Flags().StringP("url", "u", "", "URL to run the script")
	scriptCmd.Flags().StringP("output", "o", "", "Output file")
	scriptCmd.Flags().StringP("dataName", "d", "", "Name of the data to fetch")
}

func generateCSV(cmd *cobra.Command, args []string) {
	fmt.Println("Running script...")
	url, _ := cmd.Flags().GetString("url")
	output, _ := cmd.Flags().GetString("output")
	dataName, _ := cmd.Flags().GetString("dataName")

	data,err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(data.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	dataItems,ok := result[dataName].([]interface{})
	if !ok {
		log.Fatalf("Could not find data with name %s", dataName)
	}

	file,err := os.Create(output+".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{}
	firstItem := dataItems[0].(map[string]interface{})
	for key := range firstItem {
		headers = append(headers, key)
	}
	writer.Write(headers)

	for _,item := range dataItems {
		values := []string{}
		for _,key := range headers {
			values = append(values, fmt.Sprintf("%v", item.(map[string]interface{})[key]))
		}
		writer.Write(values)
	}


	fmt.Println("Written to file", output+".csv")

}

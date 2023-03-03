package cmd

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// metricCmd represents the metric command
var metricCmd = &cobra.Command{
	Use:   "metric",
	Short: "Use Kevin to get metrics based off of the proxy",
	Long:  `A subcommand which will display all generic prometheus metrics`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/metrics")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)
		log.Printf(sb)
	},
}

func init() {
	rootCmd.AddCommand(metricCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

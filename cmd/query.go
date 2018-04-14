// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config (per viper runtime read): api_endpoint:", viper.Get("api_endpoint"))
		fmt.Println("query called")

		// create a client (safe to share across requests)
		client := graphql.NewClient(viper.GetString("api_endpoint"))

		// define a Context for the request
		// ctx := context.Background()
		ctx, cancel := context.WithTimeout(context.Background(), 3200*time.Millisecond)
		defer cancel()
		// 'query={ projects { name id } }'
		// make a request
		req := graphql.NewRequest(`
			query {
				projects {
					name
					id
				}
			}
		`)

		// run it and capture the response

		type Project struct {
			id   string
			name string
		}
		type ResponseData struct {
			projects []Project `json:"projects"`
		}

		// var prettyJSON bytes.Buffer
		var resp struct {
			data ResponseData `json:"data"`
		}

		// var respData ResponseStruct
		if err := client.Run(ctx, req, &resp); err != nil {
			log.Fatal(err)
		}

		fmt.Println("response data: ", resp.raw().toString())

	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

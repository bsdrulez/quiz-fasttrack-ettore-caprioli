/*
Copyright © 2021 Ettore Caprioli <ettore.caprioli@gmail>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
//	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/quiz-fasttrack-ettore-caprioli/appclient"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start a client to interact with the server and take the quiz",
	Long: `The main purpose of this application is to been able to take a quiz.
        
        This application is a single binary but has two distinct mode of operation:
          - when operating as a 'server' its duety is to provide questions, check the
            results and keep a list of scores
          - when operating as a 'client' it provide the command-line interface for the 
            user who wants to take the quiz.
        
        This is a development version and for now it operates only on localhost:8080.`,
	Run: func(cmd *cobra.Command, args []string) {
	    appclient.RunClient()
            //fmt.Println("client called")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

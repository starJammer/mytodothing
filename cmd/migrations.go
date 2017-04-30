// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrationsCmd represents the migrations command
var migrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		//check to make sure the migrations dir exists
		f, err := os.Open(viper.GetString("path"))
		if err != nil {
			return err
		}
		defer f.Close()

		return nil
	},
}

func init() {
	RootCmd.AddCommand(migrationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	migrationsCmd.PersistentFlags().String("path", "./migrations", "Path to where migrations are located")
	viper.BindPFlag("migrations.path", migrationsCmd.PersistentFlags().Lookup("path"))
	viper.RegisterAlias("path", "migrations.path")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

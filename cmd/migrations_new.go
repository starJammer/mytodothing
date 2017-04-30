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
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := cmd.Parent().RunE(cmd, args); err != nil {
			return err
		}

		if !viper.IsSet("name") {
			return errors.New("You must specify a file name for the migration")
		}

		name := viper.GetString("name")
		migrationPath := strings.ToLower(viper.GetString("path"))
		prefix := fmt.Sprint(time.Now().Unix())
		baseFilename := fmt.Sprint(prefix, "_", name)
		for _, suffix := range []string{".up.sql", ".down.sql"} {
			filename := path.Clean(migrationPath) + "/" + baseFilename + suffix
			f, err := os.Create(filename)
			if err != nil {
				return err
			}
			f.Close()
		}

		return nil
	},
}

func init() {
	migrationsCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	newCmd.Flags().StringP("name", "n", "", "")
	viper.BindPFlags(newCmd.Flags())
}

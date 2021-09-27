/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pedreviljoen/go-flate/core"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goflate",
	Short: "A small and simple compression CLI tool",
	Long: `A Golang compression CLI which makes use of the Golang standard library Flate for compression.

	Checks compression results, before and after
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	var (
		originalFile  string
		newFile       string
		compressLevel int
	)
	cobra.OnInitialize(initConfig)

	var compressCmd = &cobra.Command{
		Use:   "compress",
		Short: "Compress file specified in arguments",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := core.Compress(originalFile, newFile, compressLevel)
			if err != nil {
				fmt.Println("An error occured")
				os.Exit(1)
			}

			fmt.Println("Compression results: \n")
			fmt.Println("Before compression size: %d \n", res.BeforeStats.Size)
			fmt.Println("After compression size: %d \n", res.AfterStats.Size)
		},
	}
	compressCmd.PersistentFlags().StringVarP(&originalFile, "file", "f", "test.txt", "Name of the file to compress")
	compressCmd.PersistentFlags().StringVarP(&newFile, "out", "o", "test-output.compressed", "Name of the new compressed output file")
	compressCmd.PersistentFlags().IntVar(&compressLevel, "level", 5, "Level of compression. From 1 (best speed) to 9 (best compression)")
	rootCmd.AddCommand(compressCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".goflate" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".goflate")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

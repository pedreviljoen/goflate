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

	"github.com/pedreviljoen/go-flate/core"
	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
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
		fmt.Println(fmt.Sprintln("Before compression size (bytes):\n", res.BeforeStats.Size))
		fmt.Println(fmt.Sprintln("After compression size (bytes):\n", res.AfterStats.Size))
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	compressCmd.PersistentFlags().StringVarP(&originalFile, "file", "f", "test.txt", "Name of the file to compress")
	compressCmd.PersistentFlags().StringVarP(&newFile, "out", "o", "test-output.compressed", "Name of the new compressed output file")
	compressCmd.PersistentFlags().IntVar(&compressLevel, "level", 5, "Level of compression. From 1 (best speed) to 9 (best compression)")
	rootCmd.AddCommand(compressCmd)
}

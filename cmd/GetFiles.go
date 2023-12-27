/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"

	"os"
	"path/filepath"
)

// GetFilesCmd represents the GetFiles command
var GetFilesCmd = &cobra.Command{
	Use:   "GetFiles",
	Short: "Get Files from root path",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var fileType string
		var filePath string

		filePath = args[0]
		fType, _ := cmd.Flags().GetString("fileType")

		if fType != "" {
			fileType = fType
		} else {
			fileType = ""
		}

		PathExist := PathExists(filePath)

		if !PathExist {
			fmt.Printf("Path %v does not exist", filePath)
			return
		} else {
			_ = GetFileList(filePath, fileType)
		}

	},
}

func PathExists(filePath string) bool {
	_, err := os.Open(filePath)
	if err != nil {
		return false
	}
	return true
}

func GetFileList(filePath, fileType string) (out []string) {

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			match, _ := regexp.MatchString("[^\\s]+(.*?)\\."+(fileType)+"$", info.Name())
			if match || fileType == "" {
				fmt.Println(path)
				out = append(out, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return out
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	GetFilesCmd.PersistentFlags().StringP("fileType", "f", "", "File type to retrive from system")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

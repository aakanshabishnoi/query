/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var wg1 sync.WaitGroup

// SearchWordCmd represents the SearchWord command
var SearchWordCmd = &cobra.Command{
	Use:   "SearchWord",
	Short: "Search word in files from system",
	Long:  ``,

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SearchWord--- called")

		var fileType string
		var filePath string
		var searchKey string

		//give filepath and search keyword into arguments of subcommands
		filePath = args[0]
		searchKey = args[1]

		//get filetype from flag
		fType, _ := cmd.Flags().GetString("fileType")
		if fType != "" {
			fileType = fType
		} else {
			fileType = ""
		}

		presentTime := time.Now()
		var wg sync.WaitGroup
		wg.Add(1)
		//go func() {
		go getLinesFromSystem(filePath, fileType, searchKey)
		wg.Done()
		//}()

		wg.Wait()
		fmt.Println("Time since-0 :", time.Since(presentTime))

	},
}

func getLinesFromSystem(filePath, fileType, searchKey string) {

	// var wg sync.WaitGroup
	presentTime := time.Now()

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			match, _ := regexp.MatchString("[^\\s]+(.*?)\\."+(fileType)+"$", info.Name())
			if match || fileType == "" {
				// wg.Add(1)
				// go func() {
				getlinesFromFile(path, searchKey)
				// 	wg.Done()
				// }()

			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// wg.Wait()
	fmt.Println("Time since :", time.Since(presentTime))
}

func getlinesFromFile(file, searchKey string) {

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 1
	count := 0
	temp := []string{}
	for scanner.Scan() {

		temp = append(temp, scanner.Text())
		if strings.Contains(scanner.Text(), searchKey) {

			var index int
			index = strings.Index(scanner.Text(), searchKey)
			str := []string{file, "(", strconv.Itoa(line), ") ", scanner.Text()[0:index], color.Ize(color.Red, searchKey),
				scanner.Text()[len(searchKey)+index:]}

			joinSt := strings.Join(str, "")
			res := strings.Join(strings.Fields(joinSt), " ")
			// for _, value := range temp {
			// 	fmt.Println(value)
			// }
			count++
			fmt.Println(res)
			//fmt.Println("")
		}
		// if line == 1 || line == 2 || line == 3 {
		// 	//fmt.Println("inside lines")
		// 	line++
		// 	continue
		// }

		// temp = temp[1:]
		line++
	}

	//fmt.Println(count)
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	SearchWordCmd.PersistentFlags().StringP("fileType", "f", "", "File type to retrive from system")
	SearchWordCmd.PersistentFlags().Int16P("lines", "l", 0, "lines with searched lines")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SearchWordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

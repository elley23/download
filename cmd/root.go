/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"download/client"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var url string
var dir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "download",
	Short: "download a file.",
	Long:  `You can download a file using url.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(url) == 0 {
			fmt.Println("Please input the url...")
			cmd.Help()
			return
		}

		err := client.DownloadFileGo(url, dir)
		if err != nil {
			cmd.Help()
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xpower.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "url (Mandatory)")
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "", "Destination file's directory")
}

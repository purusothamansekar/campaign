package cmd

import (
	"github.com/rykov/paperboy/mail"
	"github.com/spf13/cobra"

	"fmt"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Paperboy",
	Long:  `A longer description...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Paperboy Email Engine %s\n", mail.Config.Build)
	},
}

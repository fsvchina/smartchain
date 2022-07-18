package cmd

import (
	"fmt"
	"fs.video/smartchain/core"
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "lookup version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("chain version:", core.Version)
		},
	}
	return cmd
}

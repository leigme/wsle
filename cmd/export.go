package cmd

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

import (
	"fmt"
	loki "github.com/leigme/loki/cobra"
	shell "github.com/leigme/loki/shell"
	"github.com/leigme/wsle/common"
	"github.com/spf13/cobra"
)

func init() {
	loki.Add(rootCmd, &export{}, loki.WithSkip(2))
}

type export struct {
}

func (e *export) Execute() loki.Exec {
	return func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("command is nil")
			return
		}
		exportFile := args[1]
		if len(args) < 2 {
			exportFile = common.ExportDefaultFile(args[0])
			return
		}
		c := generateExportCmd(args[0], exportFile)
		fmt.Println(c)
		s := shell.New()
		fmt.Println(s.Exe(c))
	}
}

// wsl --export [Ubuntu-22.04] [C:\Users\leig\.wsl\backup\Ubuntu-22.04_20060102150405.tar]
func generateExportCmd(linux, exportFile string) string {
	return fmt.Sprintf("wsl --export %s %s", linux, exportFile)
}

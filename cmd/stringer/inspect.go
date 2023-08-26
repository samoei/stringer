package stringer

import (
	"fmt"

	"github.com/samoei/stringer/pkg/stringer"
	"github.com/spf13/cobra"
)

var onlyDigits bool
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res, kind := stringer.Inspect(i, onlyDigits)
		plurals := "s"
		if res == 1 {
			plurals = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, plurals)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(inspectCmd)
}

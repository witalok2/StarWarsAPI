package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/witalok2/starwarsplanet/adapters/cli"
)

var (
	action        string
	planteID      uuid.UUID
	planetName    string
	planetClimate string
	planetTerrain string
	planetAppearances int
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(&planetService, action, planteID, planetName, planetClimate, planetTerrain, planetAppearances)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	var planteIDConverted string
	planteIDConverted = planteID.String()

	cliCmd.Flags().StringVarP(&action, "action", "a", "create", "command cli to create planet")
	cliCmd.Flags().StringVarP(&planteIDConverted, "id", "i", "", "command cli to get planet with ID")

}

package show

import (
	"github.com/satyazzz123/KubeHealth/cmd"
	"github.com/emicklei/go-restful/v3/log"
	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

// showCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "It show the rosurce details metion in the command (like.. pods, servcies, deployments, etc...)",
	Long:  `It show the rosurce details metion in the command (like.. pods, servcies, deployments, etc...)`,
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.Printf("Verbose mode is enabled")
		}
		log.Print("Please mention the name of resorces you want to see")
	},
}

func init() {
	cmd.RootCmd.AddCommand(ShowCmd)

	// Define flags for ShowCmd
	ShowCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose mode")
	ShowCmd.Flags().StringVarP(&Source, "source", "s", "", "Specify the source")

	// Additional flag configurations can be added as needed
}
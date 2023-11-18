package cmd
 
import(
	"github.com/satyazzz123/KubeHealth/cmd"
	"github.com/satyazzz123/KubeHealth/helper"
	"log"
	"os"
	"github.com/satyazzz123/KubeHealth/show"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

)
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Show services in a Kubernetes namespace",
	Long: `Display a list of services in a specified Kubernetes namespace, including their names,
namespaces, types, cluster IP or external IPs and ages.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		serviceInfo, err := helper.ShowServices(client, namespace)
		if err != nil {
			log.Printf("error getting services: %v", err)
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Service Name", "Namespace", "Type", "cl-IP | ex-IP", "Ports", "Age"})

			for _, service := range serviceInfo {
				row := []string{service.Name, service.Namespace, service.Type, service.IPs, service.Ports, service.Age}
				table.Append(row)
			}
			table.Render()
		}
	},
}

func init() {
	show.ShowCmd.AddCommand(servicesCmd)
	servicesCmd.Flags().StringP("ns", "n", "default", "Specify the namespace")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// servicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
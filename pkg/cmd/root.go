package cmd

import "github.com/spf13/cobra"

// Root command for the CLI
var RootCMD = &cobra.Command{
	Use:   "solar",
	Short: "🛰️ Clone all your starred GitHub repos",
	Long: `
███████╗ ██████╗ ██╗      █████╗ ██████╗
██╔════╝██╔═══██╗██║     ██╔══██╗██╔══██╗
███████╗██║   ██║██║     ███████║██████╔╝
╚════██║██║   ██║██║     ██╔══██║██╔══██╗
███████║╚██████╔╝███████╗██║  ██║██║  ██║
╚══════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝

🛰️ Clone all your starred GitHub repos
`,
}

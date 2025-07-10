// File: cmd/list-tags.go
package cmd

import (
	"fmt"
	"ghcr-cli/pkg"
	"os"

	"github.com/spf13/cobra"
)

var (
	repo     string
	token    string
	username string // FIX: The variable for username was missing.
)

var listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Returns a list of tags from ghcr.io",
	Long: `This command is used to return a list of tags for a specified 
	public container registry in the GitHub Container Registry.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// --- Get the Token (PAT) ---
		pat := token
		if pat == "" {
			pat = os.Getenv("GITHUB_TOKEN")
		}
		if pat == "" {
			return fmt.Errorf("authentication required: Please provide GitHub token via --token flag or the GITHUB_TOKEN environment variable")
		}

		// --- FIX: Get the Username ---
		// This entire block was missing.
		user := username
		if user == "" {
			user = os.Getenv("GITHUB_USER")
		}
		if user == "" {
			return fmt.Errorf("username required: Please provide your GitHub username via the --user flag or the GITHUB_USER environment variable")
		}

		// --- FIX: Call the API with the correct 3 arguments ---
		// This now matches the function definition in pkg/ghcr.go
		tags, err := pkg.ListTags(repo, user, pat)
		if err != nil {
			return err
		}

		fmt.Printf("Tags for repository %s:\n", repo)
		for _, tag := range tags {
			fmt.Println(tag)
		}
		return nil
	},
}

func init() {
	ghcrCmd.AddCommand(listTagsCmd)

	// Flag for the repository
	listTagsCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repository in format <org>/<repo-name>")
	listTagsCmd.MarkFlagRequired("repo")

	// Flag for the token (PAT)
	listTagsCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub Classic PAT (or use GITHUB_TOKEN env var)")

	// --- FIX: Define the flag for the username ---
	// This flag definition was missing.
	listTagsCmd.Flags().StringVarP(&username, "user", "u", "", "Your GitHub username (or use GITHUB_USER env var)")
	listTagsCmd.MarkFlagRequired("user")
}

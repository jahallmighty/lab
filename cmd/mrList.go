package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"github.com/zaquestion/lab/internal/git"
	lab "github.com/zaquestion/lab/internal/gitlab"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Merge Requests",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rn, err := git.RepoName()
		if err != nil {
			log.Fatal(err)
		}
		page := 0
		if len(args) == 1 {
			page, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		mrs, err := lab.ListMRs(rn, &gitlab.ListMergeRequestsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 10,
			},
			OrderBy: gitlab.String("updated_at"),
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, mr := range mrs {
			fmt.Fprintf(os.Stdout, "#%d %s\n", mr.IID, mr.Title)
		}
	},
}

type formattedMR struct {
	IID   int
	Title string
}

func init() {
	mrCmd.AddCommand(listCmd)
}

package commands

import (
	"fmt"

	"github.com/github/hub/github"
	"github.com/github/hub/utils"
)

var (
	cmdIssues = &Command{
		Run:   issues,
		Usage: "issues [-a]",
		Short: "List all issues on GitHub",
		Long:  `List summary of the open issues for all projects this user belongs to`,
	}
)

var flagAll bool

func init() {
	cmdIssues.Flag.BoolVarP(&flagAll, "all", "a", false, "ALL")
	CmdRunner.Use(cmdIssues)
}

func issues(cmd *Command, args *Args) {
	runInLocalRepo(func(localRepo *github.GitHubRepo, project *github.Project, gh *github.Client) {
		if args.Noop {
			fmt.Printf("Would request list of issues for %s\n", project)
		} else {

			user, err := gh.CurrentUser()
			utils.Check(err)
			issues, err := gh.Issues(project)
			utils.Check(err)

			for _, issue := range issues {
				if flagAll {
					fmt.Printf("%s: %s\n", issue.State, issue.Title)
				}
				if !flagAll && (issue.Assignee.ID == user.ID) {
					fmt.Printf("%s: %s\n", issue.State, issue.Title)
				}
			}
		}
	})
}

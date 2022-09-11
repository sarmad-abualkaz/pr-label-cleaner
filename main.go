package main

import (
	"flag"
	"github.com/sarmad-abualkaz/pr-label-cleaner/cmd"

	log "github.com/sirupsen/logrus"
)

func main(){
	
	action   := flag.String("action", "", "type of action to perform - only 'add' or 'remove' are acceptable")
	dryRun   := flag.Bool("dry-run", false, "dry run boolean")
	label    := flag.String("label", "", "label to add or remove")
	owner    := flag.String("owner", "", "owner of github repo")
	prNumber := flag.Int("pr-number", 0, "PR number to add or remove label")
	repo     := flag.String("repo", "", "github repoistory name")
	
	flag.Parse()

	// log program starting
	log.WithFields(log.Fields{
		"label":        *label,
		"repo":         *repo,
		"owner":        *owner,
		"pull-request": *prNumber,
		"action":       *action,   
	  }).Info("program started ...")

	cmd.LabelAction(*action, *dryRun, *label, *owner, *prNumber, *repo)

}

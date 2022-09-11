package gh

import (
	"context"
	"fmt"
	
	"github.com/google/go-github/v47/github"
	
	log "github.com/sirupsen/logrus"
)

func AddLable(ctx context.Context, client *github.Client, label string, owner string, prNumber int, repo string) error {

	log.WithFields(log.Fields{
		"label":        label,
		"repo":         repo,
		"owner":        owner,
		"pull-request": prNumber,
	}).Info("Adding a label ...")

	_, res, err := client.Issues.AddLabelsToIssue(ctx, owner, repo, prNumber, []string{label})

	if res.StatusCode == 404 {

		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
		}).Warn("PR not found ...")

		return fmt.Errorf("PR not found")

	}

	if res.StatusCode == 410 {

		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
		}).Warn("received 410 code: API Method Gone...")
		
		return fmt.Errorf("API Method Gone")
	}


	if err != nil {

		return err
	
	}

	log.WithFields(log.Fields{
		"label":        label,
		"repo":         repo,
		"owner":        owner,
		"pull-request": prNumber,
	}).Info("Label added ...")

	return nil

}

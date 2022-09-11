package cmd

import (
	"fmt"

	"github.com/sarmad-abualkaz/pr-label-cleaner/gh"
	log "github.com/sirupsen/logrus"
)

func LabelAction(action string, dryRun bool, label string, owner string, prNumber int, repo string) {

	var err error 
	
	// create client
	ctx, client := gh.CreateClient()

	switch {
    
	case action == "remove":

		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
			"action":       action,
	  	}).Info("Triggering workflow for removing label ...")

		if !dryRun{	

			err = gh.RemoveLable(ctx, client, label, owner, prNumber, repo)

		} else {

			log.WithFields(log.Fields{
				"label":        label,
				"repo":         repo,
				"owner":        owner,
				"pull-request": prNumber,
				"action":       action,
			  }).Info("[dry-run] no action: Workflow for removing label ...")

		}
    
	case action == "add":

		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
			"action":       action,
	  	}).Info("Triggering workflow for adding label ...")

		  if !dryRun{	
			
			err = gh.AddLable(ctx, client, label, owner, prNumber, repo)
		
			} else {

			log.WithFields(log.Fields{
				"label":        label,
				"repo":         repo,
				"owner":        owner,
				"pull-request": prNumber,
				"action":       action,
			  }).Info("[dry-run] no action: Workflow for removing label ...")

		}

	default: 
	
		// fatal log action not recongized
		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
			"action":       action,
	  	}).Fatal("Error: action not recgonized - expect either add or remove ...")

		panic(fmt.Errorf("Error: action not recgonized - expect either add or remove ..."))

	}

	if err != nil {

		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
			"action":       action,
		}).Error(err.Error())		

		// fatal action failed
		log.WithFields(log.Fields{
			"label":        label,
			"repo":         repo,
			"owner":        owner,
			"pull-request": prNumber,
			"action":       action,
		}).Fatal("Error: action failed ...")

		panic(err)

	}

}
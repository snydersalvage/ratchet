package util

import "github.com/snydersalvage/ratchet/logger"

// KillPipelineIfErr is an error-checking helper.
func KillPipelineIfErr(err error, killChan chan error) {
	if err != nil {
		logger.Error(err.Error())
		killChan <- err
	}
}

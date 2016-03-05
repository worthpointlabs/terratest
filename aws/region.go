package aws

import (
	"math/rand"
	"time"
)

func GetForbiddenRegions() []string {
	return []string{
		"us-west-2",
	}
}

func GetRegion() string {
	allRegions := []string{
		"us-east-1",
		"us-west-1",
		//"us-west-2",
		"eu-west-1",
		"eu-central-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"sa-east-1",
	}

	randomIndex := random(0,len(allRegions))

	return allRegions[randomIndex]
}

// Generate a random int between min and max
func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

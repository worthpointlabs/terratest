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

// Get a randomly chosen AWS region and its corresponding availability zones.
// Note that for some regions, the availability zones string is unique for each AWS account.
func GetRandomRegion() (string, string) {

	allRegions := make(map[string]string)
	allRegions["us-east-1"] = "us-east-1a,us-east-1b,us-east-1d,us-east-1e"
	allRegions["us-west-1"] = "us-west-1a,us-west-1b"
	allRegions["us-west-2"] = "us-west-2a,us-west-2b,us-west-2c"
	allRegions["eu-west-1"] = "eu-west-1a,eu-west-1b,eu-west-1c"
	allRegions["eu-central-1"] = "eu-central-1a,eu-central-1b"
	allRegions["ap-northeast-1"] = "ap-northeast-1a,ap-northeast-1c"
	allRegions["ap-northeast-2"] = "ap-northeast-2a,ap-northeast-2c"
	allRegions["ap-southeast-1"] = "ap-southeast-1a,ap-southeast-1b"
	allRegions["ap-southeast-2"] = "ap-southeast-2a,ap-southeast-2b,ap-southeast-2c"
	allRegions["sa-east-1"] = "sa-east-1a,sa-east-1b,sa-east-1c"

	// We want to select a random key in allRegions, so we create an array of the keys and
	// generate a random index value.
	var allRegionKeys []string
	for region, _ := range allRegions {
		allRegionKeys = append(allRegionKeys, region)
	}

	randomIndex := -1
	randomIndexIsValid := false

	// If our randomIndex gave us a region that's forbidden, keep iterating until we get a valid one.
	for !randomIndexIsValid {
		randomIndex = random(0,len(allRegions))
		randomIndexIsValid = true

		for _, forbiddenRegion := range GetForbiddenRegions() {
			if forbiddenRegion == allRegionKeys[randomIndex] {
				randomIndexIsValid = false
			}
		}
	}

	returnRegion := allRegionKeys[randomIndex]
	returnRegionAZs := allRegions[returnRegion]

	return returnRegion, returnRegionAZs
}

// Generate a random int between min and max
func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

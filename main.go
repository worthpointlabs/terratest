package main

import (
	"github.com/gruntwork-io/terraform-test/aws"
	"fmt"
)

func main() {
	region, azs := aws.GetRandomRegion()
	fmt.Println(region)
	fmt.Println(azs)
}

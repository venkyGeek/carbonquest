package utils

import "fmt"

func DeployFunction(function, env, version string) {
	fmt.Printf("Deploying function '%s' to '%s' with version '%s'\n", function, env, version)
	// Replace with real gcloud CLI or SDK code
}

func DescribeFunction(function string) {
	fmt.Printf("Function details for '%s' (mocked)\n", function)
	// Replace with real gcloud CLI or SDK code
}

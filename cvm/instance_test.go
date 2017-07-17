package cvm

import (
	"testing"
)

func TestDescribeProject(t *testing.T) {
	client, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	version := "2017-03-12"

	args := DescribeInstancesArgs{
		Version: &version,
	}

	response, err := client.DescribeInstances(&args)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%++v", response.Response.InstanceSet)

}

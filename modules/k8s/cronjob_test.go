//go:build kubeall || kubernetes
// +build kubeall kubernetes

package k8s

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"testing"
	"time"
)

func TestListCronJobsReturnsCronJobsInNamespace(t *testing.T) {
	t.Parallel()

	uniqueID := strings.ToLower(random.UniqueId())
	options := NewKubectlOptions("", "", uniqueID)
	configData := fmt.Sprintf(ExampleCronjobYamlTemplate, uniqueID, uniqueID)
	defer KubectlDeleteFromString(t, options, configData)
	KubectlApplyFromString(t, options, configData)

	jobs := ListCronJob(t, options, metav1.ListOptions{})
	require.Equal(t, len(jobs), 1)
	job := jobs[0]
	require.Equal(t, job.Name, "cron-job")
	require.Equal(t, job.Namespace, uniqueID)
}

func TestGetCronJobEReturnErrorForNotExistingCronJob(t *testing.T) {
	t.Parallel()

	options := NewKubectlOptions("", "", "default")
	_, err := GetJobE(t, options, random.UniqueId())
	require.Error(t, err)
}

func TestGetCronJobEReturnsCorrectJobInNamespace(t *testing.T) {
	t.Parallel()
	uniqueID := strings.ToLower(random.UniqueId())
	options := NewKubectlOptions("", "", uniqueID)
	configData := fmt.Sprintf(ExampleCronjobYamlTemplate, uniqueID, uniqueID)
	defer KubectlDeleteFromString(t, options, configData)
	KubectlApplyFromString(t, options, configData)

	job := GetCronJob(t, options, "cron-job")
	require.Equal(t, job.Name, "cron-job")
	require.Equal(t, job.Namespace, uniqueID)
}

func TestWaitUntilCronJobScheduleSuccessfullyContainer(t *testing.T) {
	t.Parallel()

	uniqueID := strings.ToLower(random.UniqueId())
	options := NewKubectlOptions("", "", uniqueID)
	configData := fmt.Sprintf(ExampleCronjobYamlTemplate, uniqueID, uniqueID)
	defer KubectlDeleteFromString(t, options, configData)
	KubectlApplyFromString(t, options, configData)

	WaitUntilCronJobSucceed(t, options, "cron-job", 60, 5*time.Second)
}

const ExampleCronjobYamlTemplate = `---
apiVersion: v1
kind: Namespace
metadata:
  name: %s
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: cron-job
  namespace: %s
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: ubuntu
            image: ubuntu:20.04
            command: ["sh", "-c", "ls"]
          restartPolicy: OnFailure
`

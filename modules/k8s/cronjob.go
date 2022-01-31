package k8s

import (
	"context"
	"fmt"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func ListCronJob(t testing.TestingT, options *KubectlOptions, filters metav1.ListOptions) []batchv1.CronJob {
	jobs, err := ListCronJobsE(t, options, filters)
	require.NoError(t, err)
	return jobs
}

func ListCronJobsE(t testing.TestingT, options *KubectlOptions, filters metav1.ListOptions) ([]batchv1.CronJob, error) {
	clientset, err := GetKubernetesClientFromOptionsE(t, options)
	if err != nil {
		return nil, err
	}
	resp, err := clientset.BatchV1().CronJobs(options.Namespace).List(context.Background(), filters)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func GetCronJob(t testing.TestingT, options *KubectlOptions, cronJobName string) *batchv1.CronJob {
	job, err := GetCronJobE(t, options, cronJobName)
	require.NoError(t, err)
	return job
}

func GetCronJobE(t testing.TestingT, options *KubectlOptions, cronJobName string) (*batchv1.CronJob, error) {
	clientset, err := GetKubernetesClientFromOptionsE(t, options)
	if err != nil {
		return nil, err
	}
	return clientset.BatchV1().CronJobs(options.Namespace).Get(context.Background(), cronJobName, metav1.GetOptions{})
}

func WaitUntilCronJobSucceed(t testing.TestingT, options *KubectlOptions, cronJobName string, retries int, sleepBetweenRetries time.Duration) {
	require.NoError(t, WaitUntilCronJobSucceedE(t, options, cronJobName, retries, sleepBetweenRetries))
}

func WaitUntilCronJobSucceedE(t testing.TestingT, options *KubectlOptions, cronJobName string, retries int, sleepBetweenRetries time.Duration) error {
	statusMsg := fmt.Sprintf("Wait for CronJob %s to successfully schedule container", cronJobName)
	message, err := retry.DoWithRetryE(
		t,
		statusMsg,
		retries,
		sleepBetweenRetries,
		func() (string, error) {
			job, err := GetCronJobE(t, options, cronJobName)
			if err != nil {
				return "", err
			}
			if !IsCronJobSucceeded(job) {
				return "", NewCronJobNotSucceeded(job)
			}
			return "CronJob scheduled container", nil
		},
	)
	if err != nil {
		logger.Logf(t, "Timed out waiting for CronJob to schedule job: %s", err)
		return err
	}
	logger.Logf(t, message)
	return nil
}

func IsCronJobSucceeded(cronJob *batchv1.CronJob) bool {
	return cronJob.Status.LastSuccessfulTime == nil
}

package k8s

import (
	"context"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func GetCronJob(t testing.TestingT, options *KubectlOptions, jobName string) *batchv1.CronJob {
	job, err := GetCronJobE(t, options, jobName)
	require.NoError(t, err)
	return job
}

func GetCronJobE(t testing.TestingT, options *KubectlOptions, jobName string) (*batchv1.CronJob, error) {
	clientset, err := GetKubernetesClientFromOptionsE(t, options)
	if err != nil {
		return nil, err
	}
	return clientset.BatchV1().CronJobs(options.Namespace).Get(context.Background(), jobName, metav1.GetOptions{})
}

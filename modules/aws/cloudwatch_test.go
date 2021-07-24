package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCloudWatchClient(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	client := NewCloudWatchClient(t, region)
	assert.NotEmpty(t, client)
}
func TestNewCloudWatchClientE(t *testing.T) {
	t.Parallel()

	region := GetRandomStableRegion(t, nil, nil)
	client, err := NewCloudWatchClientE(t, region)
	require.NoError(t, err)
	assert.NotEmpty(t, client)
}

// +build azure

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/require"
)

func TestDiagnosticsSettingsResourceExists(t *testing.T) {
	t.Parallel()

	diagnosticsSettingResourceName := "fakename"
	_, err := azure.DiagnosticSettingsResourceExistsE(diagnosticsSettingResourceName)
	require.Error(t, err)
}

package terraform

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/stretchr/testify/assert"
)

func TestInitBackendConfigFile(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("../../test/fixtures/terraform-backend-file", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	customBackendFile := filepath.Join(testFolder, "terraform.backend")

	options := &Options{
		TerraformDir: testFolder,
		BackendConfig: map[string]interface{}{
			customBackendFile: nil,
		},
	}

	InitAndApply(t, options)

	assert.FileExists(t, customBackendFile)
}

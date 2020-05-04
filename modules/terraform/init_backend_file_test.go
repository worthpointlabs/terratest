package terraform

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/stretchr/testify/assert"
)

func TestInitBackendConfig(t *testing.T) {
	t.Parallel()

	stateDirectory, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	customBackendFile := filepath.Join(stateDirectory, "terraform.backend")

	testFolder, err := files.CopyTerraformFolderToTemp("../../test/fixtures/terraform-backend-file", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	options := &Options{
		TerraformDir: testFolder,
		BackendConfig: map[string]interface{}{
			customBackendFile: nil,
		},
	}

	InitAndApply(t, options)

	assert.FileExists(t, remoteStateFile)
}

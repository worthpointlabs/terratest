package terratest

import (
	"testing"
	"path"
	"errors"
)

func TestGetOutput(t *testing.T) {
	t.Parallel()

	randomResourceCollectionOptions := NewRandomResourceCollectionOptions()
	randomResourceCollection, err := CreateRandomResourceCollection(randomResourceCollectionOptions)
	defer randomResourceCollection.DestroyResources()
	if err != nil {
		t.Errorf("Failed to create random resource collection: %s\n", err.Error())
	}

	options := NewTerratestOptions()
	options.UniqueId = randomResourceCollection.UniqueId
	options.TestName = "Test - TestGetOutput"
	options.TemplatePath = path.Join(fixtureDir, "local-resources-only-example")

	if _, err := Apply(options); err != nil {
		t.Fatal(err)
	}

	testOutput(options, "template1", "template1", t)
	testOutput(options, "template2", "template2", t)
}

func testOutput(terratestOptions *TerratestOptions, key string, expectedOutput string, t *testing.T) {
	actualOutput, err := Output(terratestOptions, key)

	if err != nil {
		t.Fatal(err)
	}

	if actualOutput != expectedOutput {
		t.Fatal(errors.New("Got unexpected output for key " + key + ". Expected: " + expectedOutput + ". Actual: " + actualOutput + "."))
	}
}


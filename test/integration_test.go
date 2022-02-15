package go_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"gen/github.com/woven-planet/go-zserio/test/go/reference_modules/testobject1/testobject"
	"github.com/cucumber/godog"
	zserio "github.com/woven-planet/go-zserio"
)

func testWorkspace(filePath string) string {
	return path.Join(os.Getenv("TEST_SRCDIR"), os.Getenv("TEST_WORKSPACE"), filePath)
}

var (
	// ReferenceFilePath is the path to the input data.
	ReferenceFilePath string = testWorkspace(os.Getenv("TESTDATA_BIN"))

	// ReencodedFilePath is the path where to write the data again.
	ReencodedFilePath string = path.Join(
		os.Getenv("TEST_UNDECLARED_OUTPUTS_DIR"),
		`testdata_reencoded.bin`,
	)

	// ReferenceBinaryContent is used to store binary content from the reference file.
	ReferenceBinaryContent []byte
)

func TestIntegration(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "Integration",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
		},
	}
	if suite.Run() != 0 {
		t.Error("integration tests failed")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have a zserio binary created using the reference implementation`, func(ctx context.Context) (context.Context, error) {
		return ctx, readReferenceZserioFile()
	})

	ctx.Step(`^I decode this binary with go-zerio and encode it again`, func(ctx context.Context) (context.Context, error) {
		return ctx, reencodeZserioTestBinary()
	})

	ctx.Step(`^The binary content should be the same`, func(ctx context.Context) (context.Context, error) {
		return ctx, verifyFilesAreEqual()
	})
}

func readReferenceZserioFile() error {
	var err error
	ReferenceBinaryContent, err = os.ReadFile(ReferenceFilePath)
	if err != nil {
		return fmt.Errorf("read reference binary: %w", err)
	}

	return nil
}

func reencodeZserioTestBinary() error {
	// read the binary file...
	var testObject testobject.TestObject
	if err := zserio.Unmarshal(ReferenceBinaryContent, &testObject); err != nil {
		return fmt.Errorf("unmarshal reference: %w", err)
	}

	got, err := zserio.Marshal(&testObject)
	if err != nil {
		return fmt.Errorf("re-encode: %w", err)
	}

	if err := os.WriteFile(ReencodedFilePath, got, 0664); err != nil {
		return fmt.Errorf("write file: %w")
	}
	return nil
}

func verifyFilesAreEqual() error {
	fileA, err := os.ReadFile(ReferenceFilePath)
	if err != nil {
		return fmt.Errorf("read reference: %w", err)
	}

	fileB, err := os.ReadFile(ReencodedFilePath)
	if err != nil {
		return fmt.Errorf("read re-encoded: %w", err)
	}
	if bytes.Equal(fileA, fileB) {
		return nil
	}

	diff := len(fileA) - len(fileB)

	if diff > 0 {
		if bytes.Equal(fileA[:len(fileB)], fileB) {
			return fmt.Errorf("reference file has %d extra bytes at the end: %q", diff, fileA[len(fileB):])
		}

		if bytes.Equal(fileA[diff:], fileB) {
			return fmt.Errorf("reference file has %d extra bytes at the beginning: %q", diff, fileA[:diff])
		}
	}

	if diff < 0 {
		diff = -diff
		if bytes.Equal(fileA, fileB[:len(fileA)]) {
			return fmt.Errorf("re-encoded file has %d extra bytes at the end: %q", diff, fileB[len(fileA):])
		}

		if bytes.Equal(fileA, fileB[diff:]) {
			return fmt.Errorf("re-encoded file has %d extra bytes at the beginning: %q", diff, fileB[:diff])
		}
	}

	return errors.New("contents are of equal length but do not match")
}

package go_test

import (
	"bytes"
	"context"
	"path"
	"os"
	"errors"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/test/go/reference_modules/testobject1/testobject"
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
	r := bitio.NewCountReader(bytes.NewBuffer(ReferenceBinaryContent))

	if err := testObject.UnmarshalZserio(r); err != nil {
		return fmt.Errorf("unmarshal reference: %w", err)
	}

	f, err := os.OpenFile(ReencodedFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bitio.NewCountWriter(f)
	defer w.Close()

	if err := testObject.MarshalZserio(w); err != nil {
		return fmt.Errorf("marshal: %w", err)
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

package go_test

import (
	"bytes"
	"context"
	"path"
	"errors"
	"os"
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
	// ReferenceFilePath is the path to the input data
	// TODO @aignas 2022-02-02: use env variables to inject this into the test
	ReferenceFilePath string = testWorkspace(os.Getenv("TESTDATA_BIN"))
)

const (
	// ReencodedFilePath is the path where to write the data again
	// TODO @aignas 2022-02-02: put this into the bazel test directory for undefined outputs
	ReencodedFilePath string = `testdata_reencoded.bin`

	// TODO @aignas 2022-02-02: can the following constants be injected by bazel?
	GoZserioOutputDirectory string = `test/go-schema`

	GoZserioRoot string = `github.com/woven-planet/go-zserio/` + GoZserioOutputDirectory
)

var ReferenceBinaryContent []byte

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

	var buf bytes.Buffer
	w := bitio.NewCountWriter(&buf)
	defer w.Close()

	if err := testObject.MarshalZserio(w); err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return os.WriteFile(ReencodedFilePath, buf.Bytes(), 0644)
}

func verifyFilesAreEqual() error {
	fileA, err := os.ReadFile(ReferenceFilePath)
	if err != nil {
		return err
	}

	fileB, err := os.ReadFile(ReencodedFilePath)
	if err != nil {
		return err
	}
	if !bytes.Equal(fileA, fileB) {
		return errors.New("The reencoded file does not match the reference file")
	}
	return nil
}

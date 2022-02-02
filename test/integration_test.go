package go_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/cucumber/godog"
	"github.com/icza/bitio"
	"github.com/woven-planet/go-zserio/test/go-schema/reference_modules/testobject1/testobject"
)

const (
	// ReferenceFilePath is the path to the input data
	ReferenceFilePath string = `bin/testdata.bin`

	// ReencodedFilePath is the path where to write the data again
	ReencodedFilePath string = `bin/testdata_reencoded.bin`

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
	ReferenceBinaryContent, err = io.ReadFile(ReferenceFilePath)
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

	return io.WriteFile(ReencodedFilePath, buf.Bytes(), 0644)
}

func verifyFilesAreEqual() error {
	fileA, err := io.ReadFile(ReferenceFilePath)
	if err != nil {
		return err
	}

	fileB, err := io.ReadFile(ReencodedFilePath)
	if err != nil {
		return err
	}
	if !bytes.Equal(fileA, fileB) {
		return errors.New("The reencoded file does not match the reference file")
	}
	return nil
}

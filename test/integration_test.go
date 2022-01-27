package go_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"os/exec"
	"syscall"
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
		return ctx, generateReferenceZserioFile()
	})

	ctx.Step(`^I decode this binary with go-zerio and encode it again`, func(ctx context.Context) (context.Context, error) {
		return ctx, reencodeZserioTestBinary()
	})

	ctx.Step(`^The binary content should be the same`, func(ctx context.Context) (context.Context, error) {
		return ctx, verifyFilesAreEqual()
	})
}

func generateReferenceZserioFile() error {
	cmd := exec.Command("python", "python/write_test_data.py")
	_, err := cmd.CombinedOutput()
	return err
}

func reencodeZserioTestBinary() error {
	// read the binary file...
	binaryContent, err := ioutil.ReadFile(ReferenceFilePath)
	if err != nil {
		return err
	}

	var testObject testobject.TestObject
	r := bitio.NewCountReader(bytes.NewBuffer(binaryContent))
	err = testObject.UnmarshalZserio(r)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	w := bitio.NewCountWriter(&buf)
	err = testObject.MarshalZserio(w)
	if err != nil {
		return err
	}
	w.Close()
	return ioutil.WriteFile(ReencodedFilePath, buf.Bytes(), syscall.O_RDWR|syscall.O_CREAT|syscall.O_TRUNC)
}

func verifyFilesAreEqual() error {
	fileA, err := ioutil.ReadFile(ReferenceFilePath)
	if err != nil {
		return err
	}

	fileB, err := ioutil.ReadFile(ReencodedFilePath)
	if err != nil {
		return err
	}
	if !bytes.Equal(fileA, fileB) {
		return errors.New("The reencoded file does not match the reference file")
	}
	return nil
}

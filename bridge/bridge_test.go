package bridge

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	assert.NoError(t, err)
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = fmt.Errorf("content received for writer is empty")
	return
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}
	err := api2.PrintMessage("hello")
	if err != nil {
		expectedErrorMessage := "you need to pass an io.writer to printerimpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("error message was not correct.\n actual:%s\n expected: %s\n", err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{
		Writer: &testWriter,
	}
	expectedMessage := "hello"
	err = api2.PrintMessage(expectedMessage)

	if err != nil {
		t.Errorf("err trying to use the api2 implementation: %s\n", err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Fatalf("api2 did not write correctly on io.writer. \n actual: %s\nexpected: %s\n ", testWriter.Msg, expectedMessage)
	}

}

func TestNormalPrinter(t *testing.T) {
	expectedMessage := "Hello io.Writer"
	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("the expected message on the io.writer doesn't match actual.\n actual: %s\n expected:%s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter(t *testing.T) {
	expectedMessage := "Message from Packt:Hello io.Writer"
	normal := PacktPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	testWriter := TestWriter{}
	normal = PacktPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("the expected message on the io.writer doesn't match actual.\n actual: %s\n expected:%s\n", testWriter.Msg, expectedMessage)
	}
}

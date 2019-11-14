package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/lungria/mono-cli/statementsPdf"
)

func TestGenerate(t *testing.T) {
	var file *os.File
	var fileReader *bufio.Reader

	testData := [][]string{
		{"qweqweqwe", "1569903114", "Нарахування відсотків за вересень", "4829", "true", "22", "22", "980", "0", "0", "1930"},
		{"qwesdfqwe", "1569903114", "CARD2CARD UAMAB", "4829", "true", "5066", "5066", "980", "0", "0", "193550"},
	}

	err := statementsPdf.Generate("statementsPdf/testPdf/statement_test_actual.pdf", testData)
	if err != nil {
		t.Errorf("Pdf isn't generated, Error appear %v ", err)
	}

	// Open statement_test_actual.pdf
	file, err = os.Open("statementsPdf/testPdf/statement_test_actual.pdf")
	if err != nil {
		t.Errorf("Open statement_test_actual.pdf failed. Error: %v \n", err)
	}
	fileReader = bufio.NewReader(file)
	actualPdf, err := ioutil.ReadAll(fileReader)
	if err != nil {
		t.Errorf("Read from statement_test_actual.pdf failed. Error: %v \n", err)
	}

	// Open statement_test_expected.pdf
	file, err = os.Open("statementsPdf/testPdf/statement_test_expected.pdf")
	if err != nil {
		t.Errorf("Open statement_test_expected,pdf failed. Error: %v \n", err)
	}
	fileReader = bufio.NewReader(file)
	expectedPdf, err := ioutil.ReadAll(fileReader)
	if err != nil {
		t.Errorf("Read from statement_test_actual.pdf failed. Error: %v \n", err)
	}

	// Compare content of pdf files
	result := bytes.Compare(actualPdf, expectedPdf)
	if result != 1 {
		t.Errorf("Pdf files aren't sames")
	}
}

func TestGeneratePagination(t *testing.T) {
	const maxRows = 20
	var file *os.File
	var fileReader *bufio.Reader
	var multipliedTestData = make([][]string, maxRows)

	testData := [][]string{
		{"qweqweqwe", "1569903114", "Нарахування відсотків за вересень", "4829", "true", "22", "22", "980", "0", "0", "1930"},
	}

	for i := 0; i < maxRows; i++ {
		multipliedTestData[i] = append(multipliedTestData[i], testData[0]...)
	}

	err := statementsPdf.Generate("statementsPdf/testPdf/statement_test_actual_pagination.pdf", multipliedTestData)
	if err != nil {
		t.Errorf("Pdf isn't generated, Error appear %v ", err)
	}

	// Open statement_test_actual.pdf
	file, err = os.Open("statementsPdf/testPdf/statement_test_actual_pagination.pdf")
	if err != nil {
		t.Errorf("Open statement_test_actual.pdf failed. Error: %v \n", err)
	}
	fileReader = bufio.NewReader(file)
	actualPdf, err := ioutil.ReadAll(fileReader)
	if err != nil {
		t.Errorf("Read from statement_test_actual.pdf failed. Error: %v \n", err)
	}

	// Open statement_test_expected.pdf
	file, err = os.Open("statementsPdf/testPdf/statement_test_expected_pagination.pdf")
	if err != nil {
		t.Errorf("Open statement_test_expected,pdf failed. Error: %v \n", err)
	}
	fileReader = bufio.NewReader(file)
	expectedPdf, err := ioutil.ReadAll(fileReader)
	if err != nil {
		t.Errorf("Read from statement_test_actual.pdf failed. Error: %v \n", err)
	}

	// Compare content of pdf files
	result := bytes.Compare(actualPdf, expectedPdf)
	if result != 1 {
		t.Errorf("Pdf files aren't sames")
	}
}

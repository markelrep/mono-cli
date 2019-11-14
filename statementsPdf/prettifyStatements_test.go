package statementsPdf

import (
	"errors"
	"testing"
)

func TestConvertCoins(t *testing.T) {
	var actual string

	tests := map[string]string{
		"1":      "0.01",
		"99":     "0.99",
		"110":    "1.10",
		"200":    "2.00",
		"2500":   "25.00",
		"894561": "8945.61",
		"-500":   "-5.00",
	}

	for got, exp := range tests {
		actual = convertCoins(got)
		if actual != exp {
			t.Errorf("TestConvertCoins failed. Expected %v, Got %v \n", exp, actual)
		}
	}
}

func TestConvertTimestamp(t *testing.T) {
	var actual string

	tests := map[string]string{
		"1573754138": "2019-11-14 19:55:38",
		"946770296":  "2000-01-02 01:44:56",
		"0":          "1970-01-01 03:00:00",
		"3502914296": "2081-01-01 01:44:56",
	}

	for got, exp := range tests {
		actual = convertTimestamp(got)
		if actual != exp {
			t.Errorf("TestConvertTimestamp failed. Expected %v, Got %v \n", exp, actual)
		}
	}
}

func TestTrimCsv(t *testing.T) {
	testDataPositive := [][]string{
		{"qweqweqwe", "1569903114", "Нарахування відсотків за вересень", "4829", "true", "22", "22", "980", "0", "0", "1930"},
		{"qwesdfqwe", "1569903114", "CARD2CARD UAMAB", "4829", "true", "5066", "5066", "980", "0", "0", "193550"},
	}
	testDataNegative := [][]string{
		{"qweqweqwe", "1569903114", "Нарахування відсотків за вересень", "4829", "true", "22", "22", "980", "0", "0", "1930"},
		{"qwesdfqwe", "1569903114", "CARD2CARD UAMAB", "4829", "true", "5066", "5066", "980", "0", "0", "193550", "193550"},
	}

	expected := [][]string{
		{"1569903114", "Нарахування відсотків за вересень", "22", "1930"},
		{"1569903114", "CARD2CARD UAMAB", "5066", "193550"},
	}
	actual, _ := trimCsv(testDataPositive)
	for row := 0; row < len(expected); row++ {
		for col := 0; col < len(actual[row]); col++ {
			if actual[row][col] != expected[row][col] {
				t.Errorf("TestTrimCsv failed. Expected %v, Got %v \n", expected[row], actual[row])
			}
		}
	}
	_, err := trimCsv(testDataNegative)
	expErr := errors.New("Count of column is incorrect ")
	if err.Error() != expErr.Error() {
		t.Errorf("TestTrimCsv failed. Expected error %v, Got %v \n", err, expErr)
	}
}

func TestPrettyCsvArr(t *testing.T) {
	testData := [][]string{
		{"qweqweqwe", "1569903114", "Нарахування відсотків за вересень", "4829", "true", "22", "22", "980", "0", "0", "1930"},
		{"qwesdfqwe", "1569903114", "CARD2CARD UAMAB", "4829", "true", "5066", "5066", "980", "0", "0", "193550"},
	}

	expected := [][]string{
		{"2019-10-01 07:11:54", "Нарахування відсотків за вересень", "0.22", "19.30"},
		{"2019-10-01 07:11:54", "CARD2CARD UAMAB", "50.66", "1935.50"},
	}

	actual, _ := prettyCsvArr(testData)
	for row := 0; row < len(expected); row++ {
		for col := 0; col < len(actual[row]); col++ {
			if actual[row][col] != expected[row][col] {
				t.Errorf("TestPrettyCsvArr failed. Expected %v, Got %v \n", expected[row], actual[row])
			}
		}
	}
}

package statementsPdf

import (
	"strconv"
	"time"
)

const (
	dateColumn = iota
	descriptionColumn
	amountColumn
	balanceColumn
)
// TODO: flexible headers value
var headers = []string{"Time", "Description", "Amount", "Balance"}

func convertCoins(coins string) string {
	amountCoins, _ := strconv.ParseFloat(coins, 64)
	amount := amountCoins / 100
	return strconv.FormatFloat(amount, 'f', 1, 64)
}

func convertTimestamp(timestamp string) string {
	timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
	timeUnix := time.Unix(timestampInt, 0)
	return timeUnix.Format("2006-01-02 15:04:05")
}

func prettyCsvArr(statements [][]string) [][]string {
	statements = trimCsv(statements)
	for row := 0; row < len(statements); row++ {
		statements[row][dateColumn] = convertTimestamp(statements[row][dateColumn])
		statements[row][amountColumn] = convertCoins(statements[row][amountColumn])
		statements[row][balanceColumn] = convertCoins(statements[row][balanceColumn])
	}
	return statements
}

func trimCsv(statements [][]string) [][]string {
	var trimmedStatements [][]string
	var newRow []string
	for row := 0; row < len(statements); row++ {
		newRow := append(newRow, statements[row][1:3]...)
		newRow = append(newRow, statements[row][5:6]...)
		newRow = append(newRow, statements[row][10:11]...)
		trimmedStatements = append(trimmedStatements, newRow)
	}
	return trimmedStatements
}

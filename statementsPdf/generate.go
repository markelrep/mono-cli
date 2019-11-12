package statementsPdf

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

const (
	colCount = 4
	colWd    = 45.0
	marginH  = 15.0
	lineHt   = 5.5
	cellGap  = 2.0
)

const (
	imgFolder   = "statementsPdf/img/"
	fontsFolder = "statementsPdf/fonts/"
)

type cpdf struct {
	gofpdf.Pdf
	*gofpdf.Fpdf
	gofpdf.FontLoader
}

// var colStrList [colCount]string
type cellType struct {
	str  string
	list [][]byte
	ht   float64
}

var (
	cellList [colCount]cellType
	cell     cellType
)

// Generate returns pdf file
func Generate(filename string, data [][]string) error {
	statements := prettyCsvArr(data)
	//Create new PDF
	pdf := cpdf{Fpdf: gofpdf.New("P", "mm", "A4", "")}

	// Add font with support cyrillic characters
	pdf.AddUTF8Font("Alice-Regular", "", fontsFolder+"Alice-Regular.ttf")

	// Set global margins
	pdf.Fpdf.SetMargins(marginH, 35, marginH)
	pdf.Fpdf.AddPage()
	pdf.Fpdf.PageNo()

	pdf.renderHeader()
	pdf.renderTableHeaders(headers)
	// Render table with statements with pagination
	pdf.renderStatementsTable(statements, headers)

	return pdf.Fpdf.OutputFileAndClose(filename)
}

func (p cpdf) renderStatementsTable(data [][]string, headers []string) {
	y := p.Fpdf.GetY()
	count := 0
	for row := 1; row < len(data); row++ {
		p.Fpdf.SetTextColor(24, 24, 24)
		p.Fpdf.SetFillColor(255, 255, 255)
		maxHt := lineHt
		// Cell height calculation loop
		for col := 0; col < colCount; col++ {
			count++
			if count > len(data) {
				count = 1
			}
			cell.str = strings.Join(data[(len(data)-1)-row][col:col+1], " ")
			cell.list = p.Fpdf.SplitLines([]byte(cell.str), colWd-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[col] = cell
		}

		// Cell render loop
		x := marginH
		for colJ := 0; colJ < colCount; colJ++ {
			//pdf.Rect(x, y, colWd, maxHt+cellGap+cellGap, "D")
			cell = cellList[colJ]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				p.Fpdf.SetXY(x+cellGap, cellY)
				p.Fpdf.CellFormat(colWd-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
					"C", false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap + cellGap

		p.addNextPages(y, headers)
	}
}

func (p cpdf) renderHeader() {
	p.Fpdf.SetFont("Alice-Regular", "", 24)
	p.Fpdf.Text(marginH*3, 20, "Header of statements")
	p.Fpdf.ImageOptions(
		imgFolder+"mono-logo.jpg",
		170, 8,
		20, 20,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)

}

func (p cpdf) renderTableHeaders(headers []string) {
	p.Fpdf.SetFont("Alice-Regular", "", 9)
	p.Fpdf.SetTextColor(224, 224, 224)
	p.Fpdf.SetFillColor(45, 45, 45)
	// Fill headers to cells
	for colJ := 0; colJ < colCount; colJ++ {
		p.Fpdf.CellFormat(colWd, 10, headers[colJ], "1", 0, "CM", true, 0, "")
	}
	p.Fpdf.Ln(-1)
}

func (p cpdf) addNextPages(y float64, headers []string) {
	_, ht, _ := p.Fpdf.PageSize(p.Fpdf.PageNo())
	if y+cell.ht*lineHt > ht {
		p.Fpdf.SetMargins(marginH, 15, marginH)
		p.Fpdf.AddPage()
		y = p.Fpdf.GetY()
		p.Fpdf.SetFont("Alice-Regular", "", 9)

		p.renderTableHeaders(headers)

		y = p.Fpdf.GetY()

	}
}

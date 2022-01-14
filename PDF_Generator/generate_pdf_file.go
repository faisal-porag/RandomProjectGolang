package PDF_Generator

import (
	"errors"
	"github.com/jung-kurt/gofpdf"
)

/*CHECK OUT THIS LINK */
/* TODO ANOTHER RESOURCE: https://pdfcrowd.com/doc/api/html-to-pdf/go/ */

func GeneratePdf(filename string) error {

	if filename != "" {
		filename = "./PDF_Generator/" + filename
	} else {
		return errors.New("file name with .pdf expected")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "Welcome to golangcode.com", "0", 0, "CM", false, 0, "")

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"./dummy_img/avatar.jpg",
		80, 20,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)

	return pdf.OutputFileAndClose(filename)
}

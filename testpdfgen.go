package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"strconv"

	"codeberg.org/go-pdf/fpdf"
)

type Colour struct {
	R, G, B int
}

func main() {
	// Define and parse command-line flags
	pages := flag.Int("pages", 10, "Number of pages to generate (10 default)")
	width := flag.Float64("width", 210, "Page width in mm (A4 default)")
	height := flag.Float64("height", 297, "Page height in mm (A4 default)")
	outputFile := flag.String("output", "output.pdf", "Output PDF file name")
	flag.Parse()

	creatorreference := "github.com/nickheaphy/testpdfgen"

	// Create a new PDF document
	pdf := fpdf.New("P", "mm", "", "")
	pdf.SetAutoPageBreak(false, 0)

	// Get the current user for metadata
	currentUser, _ := user.Current()
	if currentUser != nil {
		pdf.SetAuthor(currentUser.Username, false)
	}
	// Set document metadata
	pdf.SetTitle("Test PDF Generation", false)
	pdf.SetCreator(creatorreference, false)

	// Set the default fill colour for the document
	defaultFillColour := Colour{R: 255, G: 220, B: 130}

	// Set a scaling factor based on A4 page size
	xscaling := *width / 210.0
	yscaling := *height / 297.0
	scaling := min(xscaling, yscaling, 1.0)

	// Add pages with content
	for i := 1; i <= *pages; i++ {
		// Add a new page with the specified dimensions
		pdf.AddPageFormat("P", fpdf.SizeType{Wd: *width, Ht: *height})

		// Fill color for the background
		pdf.SetFillColor(defaultFillColour.R, defaultFillColour.G, defaultFillColour.B)
		pdf.Rect(0, 0, *width, *height, "F")
		// Set line width
		pdf.SetLineWidth(0.5 * scaling)

		// Draw a line from top-left to bottom-right
		pdf.Line(0, 0, *width, *height)
		// Draw a line from top-right to bottom-left
		pdf.Line(*width, 0, 0, *height)

		// Draw a circle at the center of the page
		pdf.SetFillColor(255, 255, 255) // white fill
		pdf.Circle(*width/2, *height/2, 30*scaling, "DF")

		// Add the page number to the top-center
		// The CellFormat function allows for precise positioning
		pdf.SetXY(0, *height*0.1) // Position 10% from the top
		pdf.SetFont("Arial", "B", 25*scaling)
		pageStr := strconv.Itoa(i)
		pdf.CellFormat(*width, 10, "Page "+pageStr+" of "+strconv.Itoa(*pages), "", 0, "C", false, 0, "")

		// Add the filename under the page number
		pdf.SetXY(0, *height*0.15) // Position 15% from the top
		pdf.SetFont("Arial", "", 16*scaling)
		pdf.CellFormat(*width, 10, *outputFile, "", 0, "C", false, 0, "")

		// Add registration circle marks to the middle of each edge
		pdf.SetFillColor(defaultFillColour.R, defaultFillColour.G, defaultFillColour.B)
		o := 10.0          //circle offset from edge
		r := 2.0 * scaling //circle radius
		// Top edge
		pdf.Circle(*width/2, o, r, "DF")
		pdf.Line(*width/2-r, o, *width/2+r, o)
		pdf.Line(*width/2, o-r, *width/2, o+r)
		// Bottom edge
		pdf.Circle(*width/2, *height-o, r, "DF")
		pdf.Line(*width/2-r, *height-o, *width/2+r, *height-o)
		pdf.Line(*width/2, *height-o-r, *width/2, *height-o+r)
		// Left edge
		pdf.Circle(o, *height/2, r, "DF")
		pdf.Line(o-r, *height/2, o+r, *height/2)
		pdf.Line(o, *height/2-r, o, *height/2+r)
		// Right edge
		pdf.Circle(*width-o, *height/2, r, "DF")
		pdf.Line(*width-o-r, *height/2, *width-o+r, *height/2)
		pdf.Line(*width-o, *height/2-r, *width-o, *height/2+r)

		// Add the page number to the center of the circle
		pdf.SetXY(0, *height/2-5) // Center vertically in the circle
		pdf.SetFont("Arial", "B", 100*scaling)
		pdf.CellFormat(*width, 10, pageStr, "", 0, "C", false, 0, "")

		// Add the github reference to the bottom-center
		pdf.SetXY(0, *height-*height*0.15)
		pdf.SetFont("Arial", "", 6*scaling)
		pdf.CellFormat(*width, 10, creatorreference, "", 0, "C", false, 0, "")
	}

	// Save the PDF to a file
	err := pdf.OutputFileAndClose(*outputFile)
	if err != nil {
		log.Fatalf("Failed to generate PDF: %v", err)
	}

	fmt.Printf("✅ Successfully generated '%s' with %d pages (%.2f x %.2f mm).\n", *outputFile, *pages, *width, *height)
}

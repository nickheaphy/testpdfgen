package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	"strconv"

	"codeberg.org/go-pdf/fpdf"
)

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

	// Add pages with content
	for i := 1; i <= *pages; i++ {
		// Add a new page with the specified dimensions
		pdf.AddPageFormat("P", fpdf.SizeType{Wd: *width, Ht: *height})

		// Fill background with light pink
		pdf.SetFillColor(255, 220, 130) // Light pink
		pdf.Rect(0, 0, *width, *height, "F")
		// Set line width
		pdf.SetLineWidth(0.5)

		// Draw a line from top-left to bottom-right
		pdf.Line(0, 0, *width, *height)
		// Draw a line from top-right to bottom-left
		pdf.Line(*width, 0, 0, *height)

		// Draw a circle at the center of the page
		pdf.SetFillColor(255, 255, 255) // Light gray fill
		pdf.Circle(*width/2, *height/2, 30, "DF")

		// Add the page number to the top-center
		// The CellFormat function allows for precise positioning
		pdf.SetXY(0, *height*0.1) // Position 10% from the top
		pdf.SetFont("Arial", "B", 25)
		pageStr := strconv.Itoa(i)
		pdf.CellFormat(*width, 10, "Page "+pageStr, "", 0, "C", false, 0, "")

		// Add the page number to the center of the circle
		pdf.SetXY(0, *height/2-5) // Center vertically in the circle
		pdf.SetFont("Arial", "B", 100)
		pdf.CellFormat(*width, 10, pageStr, "", 0, "C", false, 0, "")

		// Add the github reference to the bottom-center
		pdf.SetXY(0, *height-*height*0.06) // Position 15mm from the bottom
		pdf.SetFont("Arial", "", 6)
		pdf.CellFormat(*width, 10, creatorreference, "", 0, "C", false, 0, "")
	}

	// Save the PDF to a file
	err := pdf.OutputFileAndClose(*outputFile)
	if err != nil {
		log.Fatalf("Failed to generate PDF: %v", err)
	}

	fmt.Printf("✅ Successfully generated '%s' with %d pages (%.2f x %.2f mm).\n", *outputFile, *pages, *width, *height)
}

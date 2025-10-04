# TestPDFGen CLI

TestPDFGen is a simple command-line tool built with Go to generate PDF files with specified dimensions and page counts. Each page is generated with diagonal lines and a page number, making it useful for print testing.

It uses the excellent `fpdf` library from `codeberg.org/go-pdf/fpdf`.

-----

## Features

  - Generate a PDF with a custom number of pages.
  - Specify custom page dimensions (width and height) in millimeters.
  - Set a custom output filename.
  - Each page includes diagonal lines and a page number for easy identification.

-----

## Prerequisites

  - **Go**: You need to have Go installed and configured on your system (version 1.18 or higher recommended). You can download it from [golang.org](https://golang.org/dl/).

-----

## Installation & Setup

1.  **Clone or download the code:**
    Save the `testpdfgen.go` file to a directory on your computer.

2.  **Install the `fpdf` dependency:**
    Open your terminal and run the following command to download the required library:

    ```bash
    go get codeberg.org/go-pdf/fpdf
    ```

3.  **Build the executable:**
    Navigate to the directory containing `testpdfgen.go` and run the build command:

    ```bash
    go build testpdfgen.go
    ```

    This will create an executable file named `testpdfgen` (or `testpdfgen.exe` on Windows).

-----

## Usage

You can run the tool directly from your terminal.

### **Basic Command**

To generate a default PDF (10 pages, A4 dimensions, named `output.pdf`), simply run:

```bash
./testpdfgen
```

✅ You'll see the confirmation message: `Successfully generated 'output.pdf' with 10 pages (210.00 x 297.00 mm).`

### **Command-Line Flags**

You can customize the output using the following flags:

| Flag      | Description                               | Default      |
| :-------- | :---------------------------------------- | :----------- |
| `-pages`  | The total number of pages to generate.    | `10`         |
| `-width`  | The width of each page in millimeters.    | `210` (A4)   |
| `-height` | The height of each page in millimeters.   | `297` (A4)   |
| `-output` | The name of the output PDF file.          | `output.pdf` |

### **Examples**

  * **Generate a 50-page document:**

    ```bash
    ./testpdfgen -pages=50
    ```

  * **Generate a US Letter sized PDF (215.9 x 279.4 mm) named `letter.pdf`:**

    ```bash
    ./testpdfgen -pages=20 -width=215.9 -height=279.4 -output="letter.pdf"
    ```

  * **Generate a square 15-page PDF:**

    ```bash
    ./testpdfgen -pages=15 -width=200 -height=200 -output="square_doc.pdf"
    ```

### **Getting Help**

To view all available options and their default values, use the `-h` or `-help` flag:

```bash
./testpdfgen -h
```

-----

## License

This project is open-source and available under the [MIT License](https://www.google.com/search?q=LICENSE).
package main
import (
"fmt"
"log"
"os"
"github.com/unidoc/unipdf/v3/common/license"
"github.com/unidoc/unipdf/v3/extractor"
"github.com/unidoc/unipdf/v3/model"
)
func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(`UNIDOC_LICENSE_KEY`)
	if err != nil {
		panic(err)
	}
}
func main() {
f, err := os.Open("input.pdf")
if err != nil {
log.Fatalf("Failed to open PDF: %v\n", err)
}
defer f.Close()
pdfReader, err := model.NewPdfReader(f)
if err != nil {
log.Fatalf("Failed to read PDF: %v\n", err)
}
numOfPages, err := pdfReader.GetNumPages()
if err != nil {
log.Fatal("Failed to retrieve the number of pages: %v\n", err)
}
for i := 0; i < numOfPages; i++{
	pageNum := i + 1

	page, err :=pdfReader.GetPage(pageNum)
	if err != nil{
		log.Fatal("Error %v",err)
	}
	textExtractor, err := extractor.New(page)
	if err != nil{
		log.Fatal("Error %v",err)
	}
	text, err := textExtractor.ExtractText()
	if err !=nil{
		log.Fatal("Error %v",err)
	}
	fmt.Printf("This is the text of PDF",text)

}



}

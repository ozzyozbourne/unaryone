package excelize

import (
	"github.com/ozzyozbourne/unaryone/pbout"
	"github.com/xuri/excelize/v2"
	"log"
	"path/filepath"
	"strconv"
)

func WriteToFile(xlsxVal *pbout.XlsxValues) {
	log.Printf("File to be created	-> %s\n"+
		"Sheet to be created	-> %s\n", xlsxVal.FileName, xlsxVal.SheetName)
	f := excelize.NewFile()

	defer func() {
		log.Printf("Closing the file")
		if err := f.Close(); err != nil {
			log.Fatalf("Error occurred in closing the xlsx file -> %s\n", err.Error())
		}
		log.Printf("Closed successfully\n")
	}()

	log.Printf("Creating the\n sheet nameed -> %s\n", xlsxVal.SheetName)
	index, err := f.NewSheet(xlsxVal.SheetName)
	if err != nil {
		log.Fatalf("Error occurred in creating sheet named -> %s\nError -> %s\n", xlsxVal.SheetName, err.Error())
	}

	log.Printf("Created Successfully\n")

	log.Printf("Deleting the default generated sheet\n")
	if err := f.DeleteSheet("Sheet1"); err != nil {
		log.Fatalf("Error occurred in delete default sheet names -> sheet1\n")
	}
	log.Printf("Deleted sucessfuly\n")

	log.Printf("Setting with index %d as active\n", index)
	f.SetActiveSheet(index)
	log.Printf("Success\n")

	log.Printf("Looping through the repeated files and save to xlsx\n")
	for _, v := range xlsxVal.Cols {

		log.Printf("Saving in row -> %d\nFirst name -> %s Last name -> %s Age %d\n",
			v.RowNo, v.FirstName, v.LastName, v.Age)

		row := int(v.RowNo)
		if err := f.SetCellStr(xlsxVal.SheetName, "A"+strconv.Itoa(row), v.FirstName); err != nil {
			log.Fatalf("Error occurred in writing First name -> %s\n", v)
		}

		row++
		if err := f.SetCellStr(xlsxVal.SheetName, "B"+strconv.Itoa(row), v.LastName); err != nil {
			log.Fatalf("Error occurred in writing Last name -> %s\n", v)
		}

		row++
		if err := f.SetCellInt(xlsxVal.SheetName, "C"+strconv.Itoa(row), int(v.Age)); err != nil {
			log.Fatalf("Error occurred in writing Age-> %s\n", v)
		}
		log.Printf("ROW Saved Successfully\n")
	}

	path := filepath.Join("savedxlsxfiles", xlsxVal.FileName+".xlsx")
	log.Printf("Savint file to disk with path -> %s\n", path)
	if err := f.SaveAs(path); err != nil {
		log.Fatalf("Error in saving to disk with file name %s\n", xlsxVal.FileName)
	}
	log.Printf("Saved Successfully\n")

}

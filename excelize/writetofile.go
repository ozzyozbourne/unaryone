package excelize

import (
	"github.com/ozzyozbourne/unaryone/pbout"
	"github.com/ozzyozbourne/unaryone/utils"
	"github.com/xuri/excelize/v2"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func WriteToFile(xlsxVal *pbout.XlsxValues) error {
	log.Printf("checking to if savexlsxfiles dir is present or not\n")
	utils.MakeDir()
	log.Printf("File to be created	-> %s Sheet to be created	-> %s\n", xlsxVal.FileName, xlsxVal.SheetName)
	f := excelize.NewFile()

	defer func() error {
		log.Printf("Closing the file")
		if err := f.Close(); err != nil {
			log.Printf("Error occurred in closing the xlsx file -> %s\n", err.Error())
			return err
		}
		log.Printf("Closed successfully\n")
		return nil
	}()

	log.Printf("Creating the\n sheet nameed -> %s\n", xlsxVal.SheetName)
	index, err := f.NewSheet(xlsxVal.SheetName)
	if err != nil {
		log.Printf("Error occurred in creating sheet named -> %s\nError -> %s\n", xlsxVal.SheetName, err.Error())
		return err
	}

	log.Printf("Created Successfully\n")

	if strings.Compare(xlsxVal.SheetName, "sheet1") != 0 {
		log.Printf("Deleting the default generated sheet\n")
		if err := f.DeleteSheet("Sheet1"); err != nil {
			log.Printf("Error occurred in delete default sheet names -> sheet1\n")
			return err
		}
		log.Printf("Deleted sucessfuly\n")
	} else {
		log.Printf("Sheet name is default sheet1 and skiping deletion")
	}

	log.Printf("Setting with index %d as active\n", index)
	f.SetActiveSheet(index)
	log.Printf("Success\n")

	log.Printf("Looping through the repeated files and save to xlsx\n")
	for _, v := range xlsxVal.Cols {

		log.Printf("\nSaving in row -> %d\nFirst name -> %s\nLast name -> %s\nAge -> %d\n",
			v.RowNo, v.FirstName, v.LastName, v.Age)

		row := int(v.RowNo)
		if err := f.SetCellStr(xlsxVal.SheetName, "A"+strconv.Itoa(row), v.FirstName); err != nil {
			log.Printf("Error occurred in writing First name -> %s\n", v)
			return err
		}

		if err := f.SetCellStr(xlsxVal.SheetName, "B"+strconv.Itoa(row), v.LastName); err != nil {
			log.Printf("Error occurred in writing Last name -> %s\n", v)
			return err
		}

		if err := f.SetCellInt(xlsxVal.SheetName, "C"+strconv.Itoa(row), int(v.Age)); err != nil {
			log.Printf("Error occurred in writing Age-> %s\n", v)
			return err
		}
		log.Printf("ROW Saved Successfully\n")
	}

	path := filepath.Join("../", "savedxlsxfiles", xlsxVal.FileName+".xlsx")
	log.Printf("Saving file to disk [LOC] -> %s\n", path)
	if err := f.SaveAs(path); err != nil {
		log.Printf("Error in saving to disk with file name %s\n", xlsxVal.FileName)
		return err
	}
	log.Printf("Saved Successfully\n")
	return nil

}

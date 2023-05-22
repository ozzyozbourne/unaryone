package excelize

import (
	"github.com/ozzyozbourne/unaryone/pbout"
	schemas "github.com/ozzyozbourne/unaryone/pbout/schemas"
	"testing"
)

func getXlsxValue() *pbout.XlsxValues {

	s1 := schemas.Columns{
		RowNo:     1,
		FirstName: "osaid",
		LastName:  "khan",
		Age:       29,
	}
	s2 := schemas.Columns{
		RowNo:     2,
		FirstName: "suvaid",
		LastName:  "khan",
		Age:       27,
	}

	return &pbout.XlsxValues{
		FileName:  "books",
		SheetName: "sheet1",
		Cols:      []*schemas.Columns{&s1, &s2},
	}
}

func TestWriteToFile(t *testing.T) {
	t.Logf("%+v\n", getXlsxValue())
	if err := WriteToFile(getXlsxValue()); err != nil {
		t.Logf("Error occured\n")
		t.Error(err)
	}
}

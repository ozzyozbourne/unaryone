package utils

import (
	"github.com/ozzyozbourne/unaryone/pbout"
	"github.com/ozzyozbourne/unaryone/pbout/schemas"
)

func GetXlsxValue() *pbout.XlsxValues {
	s1 := schemas.Columns{
		RowNo:     1,
		FirstName: "Schools",
		LastName:  "Out in Summer",
		Age:       60,
	}

	s2 := schemas.Columns{
		RowNo:     2,
		FirstName: "Poison",
		LastName:  "Poison",
		Age:       75,
	}

	s3 := schemas.Columns{
		RowNo:     3,
		FirstName: "thunder",
		LastName:  "happens",
		Age:       85,
	}

	s4 := schemas.Columns{
		RowNo:     3,
		FirstName: "All the world",
		LastName:  "A stage",
		Age:       185,
	}

	return &pbout.XlsxValues{
		FileName:  "Alice",
		SheetName: "Copper",
		Cols:      []*schemas.Columns{&s1, &s2, &s3, &s4},
	}
}

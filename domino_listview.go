package keep_sdk

import (
	"errors"

	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseListView = new(BaseDominoListView)

type BaseDominoListView struct {
	Name             string `json:"name"`
	SelectionFormula string `json:"selectionFormula"`
	Columns          []t.DesignColumnSimple
}

type DominoRestListView struct {
	t.DominoBaseListView
}

func DominoListView(doc t.ListViewBody) error {

	// basis field list view
	if h.HasFieldStructProps(doc.Name, "name") &&
		doc.Name != "" &&
		len(doc.Name) > 0 {

		baseListView.Name = doc.Name

	}

	if h.HasFieldStructProps(doc.SelectionFormula, "selectionFormula") &&
		doc.SelectionFormula != "" &&
		len(doc.SelectionFormula) > 0 {

		baseListView.SelectionFormula = doc.SelectionFormula

	} else {
		return errors.New("Domino list needs selectionFormula value")
	}

	if h.HasFieldStructProps(doc.Columns, "columns") &&
		len(doc.Columns) > 0 {

		var columnArr []t.DesignColumnSimple

		for i := 0; i < len(doc.Columns); i++ {
			err := ValidateDesignColumnSimple(doc.Columns[i])
			if err != nil {
				return errors.New(err.Error())
			}
			break
		}

		baseListView.Columns = columnArr

	} else {
		return errors.New("Domino list needs correct column value.")
	}

	return nil
}

func ValidateDesignColumnSimple(obj t.DesignColumnSimple) error {
	if obj.Name == "" {
		return errors.New("Required property name is missing")
	}
	if obj.Formula == "" {
		return errors.New("Required property formula is missing")
	}

	return nil
}

func ToListViewJson() (t.ListViewBody, error) {
	json := t.ListViewBody{}

	if baseListView.Name != "" &&
		baseListView.SelectionFormula != "" &&
		len(baseListView.Columns) > 0 {

		json.Name = baseListView.Name
		json.SelectionFormula = baseListView.SelectionFormula
		json.Columns = baseListView.Columns

		return json, nil
	}

	return t.ListViewBody{}, errors.New("Failed to convert DominoListView Object to ListViewBody because of having a invalid required fields in Domino List View (name, selectionFormula and columns)")

}

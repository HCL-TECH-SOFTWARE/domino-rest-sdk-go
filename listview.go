/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"errors"
	"strings"
)

// BaseListView - REST API list view properties
// ListType		Domino list type options. [folder/views]
// Alias 		Alias of the view
// IsFolder		Identifies if the view is Folder or not
// Title		Title of the view
// Unid			Unique id of the view
// Noteid		The note ID, which is uniquely identifies this view within a particular database.
type BaseListView struct {
	Type     ListType `json:"type"`
	Alias    []string `json:"alias"`
	IsFolder bool     `json:"isFolder"`
	Title    string   `json:"title"`
	Unid     string   `json:"unid"`
	Noteid   string   `json:"noteid"`
}

// ListType - Domino list type options. [folder/views]
type ListType struct {
	FOLDER string `default:"folder"`
	VIEW   string `default:"view"`
}

// ListViewBody Domino REST API list view properties for getting all views
// Name					Name of the view
// SelectionFormula 	The formula of getting the view
// Columns				Columns that comprises the view to 3which the document fields will be based
type ListViewBody struct {
	BaseListView
	Name             string               `json:"name"`
	SelectionFormula string               `json:"selectionFormula"`
	Columns          []DesignColumnSimple `json:"columns"`
}

type DesignColumnSimple struct {
	Name                   string `json:"name"`
	Title                  string `json:"title"`
	Formula                string `json:"formula"`
	SeparateMultipleValues bool   `json:"separateMultipleValues"`
	Sort                   string `json:"sort"`
}

type KeyTypeRepresentation struct {
	NUMBER string `json:"number" default:"number"`
	TEXT   string `json:"text" default:"text"`
	TIME   string `json:"time" default:"time"`
}

// SortType constants
const (
	ASCENDING  = "ascending"
	DESCENDING = "descending"
	NONE       = "none"
)

// ListViewEntryOptions are object property response from domino server.
// Mode					Document mode to retrieve the documents with. (Every Form can have multiple modes, each can be different from other modes). Also, Current logged-in user must have access for the specified mode.
// Meta					When set to false, all metadata Json items on the top level of an object get suppressed. Default is true
// StartsWith			A character combination to perform a partial match to identify a starting point. The character combination will be applied to the "column" and "direction" passed in the query string. This cannot be combined with the "keys" parameter, i.e. you cannot filter on keys and startwith within that key.
// PivotColumn			Name of the column to provide the data for the pivot aggregator
// MetaAdditional		Additional metadata that is not included in the View. This may have a slight performance cost so use only if necessary.
// Category				To restrict view queries
// Column 				For alternative sorting. This requires the list to be designed for indexing on this column
// DistinctDocuments	Determines whether, when documents=true, only distinct documents should be retrieved if they exist multiple times in the list.
// FTSearchQuery		Full-text search query to filter the contents of the list
// Count				How many entries shall be returned, default = Integer.MaxInteger
// UnreadOnly			Retrieve only unread entries. Cannot be combined with documents=true, documentsOnly=true, or methods to select or query documents
// KeyAllowPartial		Select by partial Key. Default is false (key match is exact)
// Documents			Shall the query return documents instead of view entries
// Key					Useful for categorized or sorted lists. Limits return values to entries matching the key or keys. Use multiple key parameter items to specify multiple keys in request URL. The keys specified must be in the same order as the sorted columns from left to right. Unsorted columns will be ignored.
// Directorn			The direction for alternative sorting. This is ignored unless "column" query parameter is passed as well. This requires the list to be designed for indexing on this column in the desired direction. Defaults to ascending if column is set.
// Scope				 Determines what shall the view return:
//   - document entries
//   - category names
//   - all
//
// RichTextAs			The format RichText fields will be returned when retrieving documents instead of view entries. The default if unspecified is mime.
// MarkRead				When retrieving documents instead of view entries, mark them as read by the current user after retrieval
// MarkUnread			When retrieving documents instead of view entries, mark them as unread by the current user after retrieval
// Start				At which entry should return values start (zero based), default = 0
type GetListViewEntryOptions struct {
	Mode              string                 `json:"mode"`
	Meta              bool                   `json:"meta"`
	StartsWith        string                 `json:"startsWith"`
	MetaAdditional    bool                   `json:"metaAdditional"`
	Category          []string               `json:"category"`
	Column            string                 `json:"column"`
	DistinctDocuments bool                   `json:"distinctDocuments"`
	FTSearchQuery     string                 `json:"ftSearchQUery"`
	Count             int                    `json:"count"`
	StartKey          string                 `json:"startKey"`
	UntilKey          string                 `json:"untilKey"`
	KeyType           KeyTypeRepresentation  `json:"keyType"`
	UnreadOnly        bool                   `json:"unreadOnly"`
	KeyAllowPartial   bool                   `json:"keyAllowPartial"`
	Documents         bool                   `json:"documents"`
	Key               []string               `json:"key"`
	Direction         SortShort              `json:"direction"`
	Scope             ViewEntryScopes        `json:"scope"`
	RichTextAs        RichTextRepresentation `json:"richTextAs"`
	MarkRead          bool                   `json:"markRead"`
	MarkUnread        bool                   `json:"markUnread"`
	Start             int                    `json:"start"`
}

// GetListViewEntryOptions duplicate copy of ListViewEntryOptions selecting the following
// fields, [mode, startsWith, column, count, direction, key, scope, start]
type GetListPivotViewEntryOptions struct {
	Mode       string          `json:"mode"`
	StartsWith string          `json:"startsWith"`
	Column     string          `json:"column"`
	Count      int             `json:"count"`
	Key        []string        `json:"key"`
	Direction  SortShort       `json:"direction"`
	Scope      ViewEntryScopes `json:"scope"`
	Start      int             `json:"start"`
}

// GetListViewOptions options for get /list document operation.
// Type allows to specify views, folders, all
// Column returns information about the column. Use with caution, slows down the API.
type GetListViewOptions struct {
	Type    string `json:"type"`
	Columns bool   `json:"columns"`
	Filter  string `json:"filter"`
}

// SortShort pertains with sort type in getting list of objects
type SortShort struct {
	ASC  string `default:"asc" json:"asc"`
	DESC string `default:"desc" json:"desc"`
}

// Fetch view entries scope of what the view returns.
type ViewEntryScopes struct {
	ALL        string `default:"all"`
	CATEGORIES string `default:"categories"`
	DOCUMENTS  string `default:"documents"`
}

// CreateUpdateDesignOptions for PUT /design/designName/designName document operation.
type CreateUpdateDesignOptions struct {
	Raw     bool   `json:"raw"`
	NSFPath string `json:"nsfPath"`
}

// Options for GET /design/designName/designName document operation
type GetDesignOptions struct {
	Raw     bool   `json:"raw"`
	NSFPath string `json:"nsfPath"`
}

func DominoListView(doc ListViewBody) (*ListViewBody, error) {
	listView := new(ListViewBody)

	if len(strings.Trim(doc.Name, "")) == 0 {
		return nil, errors.New("Domino lists needs name value")
	}

	if len(strings.Trim(doc.SelectionFormula, "")) == 0 {
		return nil, errors.New("Domino list needs selectionFormula")
	}

	if len(doc.Columns) > 0 {
		arr := []DesignColumnSimple{}
		for _, value := range doc.Columns {
			err := validateDesignColumnSimple(value)
			if err != nil {
				return nil, err
			}
			arr = append(arr, value)
		}
		listView.Columns = arr
	} else {
		return nil, errors.New("Domino lists needs correct columns value")
	}

	listView.Name = doc.Name
	listView.SelectionFormula = doc.SelectionFormula

	return listView, nil
}

func validateDesignColumnSimple(obj DesignColumnSimple) error {

	if len(strings.Trim(obj.Name, "")) == 0 {
		return errors.New("Required property 'name' is missing.")
	}

	if len(strings.Trim(obj.Formula, "")) == 0 {
		return errors.New("Required property 'formula' is missing.")
	}

	return nil
}

func (doc *ListViewBody) ToListViewJson() (map[string]interface{}, error) {

	json := make(map[string]interface{})

	if len(strings.Trim(doc.Name, "")) > 0 &&
		len(strings.Trim(doc.SelectionFormula, "")) > 0 &&
		len(doc.Columns) > 0 {

		json["name"] = doc.Name
		json["selectionFormula"] = doc.SelectionFormula
		json["columns"] = doc.Columns

	} else {
		return nil, errors.New("Failed to convert DominoListView Object to ListViewBody because of having a invalid required fields in Domino List View (name, selectionFormula and columns)")
	}

	return json, nil
}

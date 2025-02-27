/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

// QueryActions constants are available methods from domino server to handle request with queries.
const (
	EXECUTE = "execute"
	EXPLAIN = "explain"
	PARSE   = "parse"
)

// GetDocumentOptions are optional parameter for retrieving documents
// to domino rest.
type GetDocumentOptions struct {
	ComputeWithForm bool   `json:"computeWithForm"`
	Meta            bool   `json:"meta"`
	RichTextAs      string `json:"richTextAs"`
	MarkRead        bool   `json:"markRead"`
	MarkUnread      bool   `json:"markUnread"`
	Mode            string `json:"mode"`
}

// CreateDocumentOptions are optional parameter for inserting documents
// to domino rest.
type CreateDocumentOptions struct {
	RichTextAs string `json:"richTextAs"`
	ParentUnid string `json:"parentUnid"`
}

// UpdateDocumentOptions are optional parameter for updating specific document
// to domino rest.
type UpdateDocumentOptions struct {
	RichTextAs string `json:"richTextAs"`
	Mode       string `json:"mode"`
	ParentUnid string `json:"parentUnid"`
	Revision   string `json:"revision"`
	MarkUnread bool   `json:"markUnread"`
}

// BulkGetDocumentOptions are optional parameter for getting list of documents
// from domino rest.
type BulkGetDocumentOptions struct {
	Meta       bool   `json:"meta"`
	RichTextAs string `json:"richTextAs"`
}

// GetDocumentByQueryOptions are optional parameter for retrieving documents
// from domino rest with specified queries.
type GetDocumentByQueryOptions struct {
	Count      int    `json:"count"`
	RichTextAs string `json:"richTextAs"`
	Start      int    `json:"start"`
}

// GetDocumentByQueryRequest are result request from domino rest mapped as an object.
type GetDocumentByQueryRequest struct {
	MaxScanDocs      int                    `json:"maxScanDocs"`
	MaxScanEntries   int                    `json:"maxScanEntries"`
	Mode             string                 `json:"mode"`
	NoViews          bool                   `json:"noViews"`
	Query            string                 `json:"query"`
	ReplaceItems     map[string]interface{} `json:"replaceItems"`
	Variables        map[string]interface{} `json:"variables"`
	TimeoutSecs      int                    `json:"timeoutSecs"`
	ViewRefresh      bool                   `json:"viewRefresh"`
	Start            int                    `json:"start"`
	Count            int                    `json:"count"`
	Forms            []string               `json:"forms"`
	IncludeFormAlias bool                   `json:"includeFormAlias"`
	MarkRead         bool                   `json:"markRead"`
}

// RichTextRepresentation is a list of type of richtext available in domino rest.
type RichTextRepresentation struct {
	HTML     string `json:"html" default:"html"`
	MIME     string `json:"mime" default:"mime"`
	MARKDOWN string `json:"markdown" default:"markdown"`
	PLAIN    string `json:"plain" default:"plain"`
}

// BulkUpdateDocumentsByQueryRequest are result request from domino rest mapped as object.
type BulkUpdateDocumentsByQueryRequest struct {
	MaxScanDocs           int                    `json:"maxScanDocs"`
	MaxScanEntries        int                    `json:"maxScanEntries"`
	Mode                  string                 `json:"mode"`
	NoViews               bool                   `json:"noViews"`
	Query                 string                 `json:"query"`
	ReplaceItems          map[string]interface{} `json:"replaceItems"`
	Variables             map[string]interface{} `json:"variables"`
	ReturnUpdatedDocument bool                   `json:"returnUpdatedDocument"`
	TimeoutSecs           int                    `json:"timeoutSecs"`
	ViewRefresh           bool                   `json:"viewRefresh"`
	Start                 int                    `json:"start"`
	Count                 int                    `json:"count"`
}

// DocumentMeta is a sub response object used to mapped json data result
// from domino rest.
type DocumentMeta struct {
	NoteId             int      `json:"noteid"`
	UNID               string   `json:"unid"`
	Created            string   `json:"created"`
	LastModified       string   `json:"lastmodified"`
	LastAccessed       string   `json:"lastaccessed"`
	LastModifiedInFile string   `json:"lastmodifiedinfile"`
	AddedToFile        string   `json:"addedtofile"`
	NoteClass          []string `json:"noteclass"`
	Unread             bool     `json:"unread"`
	Editable           bool     `json:"editable"`
	Revision           string   `json:"revision"`
	Size               int      `json:"size"`
	Etag               string   `json:"etag"`
	TopLevelChildUNIDs []string `json:"toplevelchildunids"`
}

type DocumentJSON struct {
	Form   string                 `json:"Form"`
	Fields map[string]interface{} `json:"fields"`
}

// DocumentInfo a parameter object needed to request in domino rest.
type DocumentInfo struct {
	Meta     DocumentMeta           `json:"@meta"`
	Form     string                 `json:"Form"`
	Warnings []string               `json:"@warnings"`
	Fields   map[string]interface{} `json:"fields"`
}

// DocumentMethods provides sets of available functions that can be used externally
type DocumentMethods struct {
	ToDocumentJSON func() map[string]interface{}
	GetUNID        func() string
	SetUNID        func(unid string)
	GetRevision    func() string
}

func DominoDocument(doc map[string]interface{}) (*DocumentInfo, error) {

	if doc["Form"] != nil && len(doc["Form"].(string)) == 0 {
		return nil, errors.New("document needs form value")
	}

	methodInfo := new(DocumentInfo)
	methodInfo.Fields = make(map[string]interface{})
	nonFieldKeys := []string{"@meta", "Form", "@warnings"}
	for key, value := range doc {

		if !slices.Contains(nonFieldKeys, key) {
			methodInfo.Fields[key] = value
		}

	}

	if doc["@meta"] != nil {
		data, err := json.Marshal(doc["@meta"])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal doc[\"@meta\"]: %w", err)
		}

		if err := json.Unmarshal(data, &methodInfo.Meta); err != nil {
			return nil, fmt.Errorf("failed to unmarshal doc[\"@meta\"] data: %w", err)
		}
	}
	if doc["@warnings"] != nil {
		for _, v := range doc["@warnings"].([]interface{}) {
			str, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("element %v is not a string", v)
			}
			methodInfo.Warnings = append(methodInfo.Warnings, str)
		}
	}
	methodInfo.Form = doc["Form"].(string)

	return methodInfo, nil
}

func (d *DocumentInfo) toDocumentJSON() map[string]interface{} {

	json := make(map[string]interface{})
	json["Form"] = d.Form

	for key, value := range d.Fields {
		json[key] = value
	}

	return json
}

func (d *DocumentInfo) getUNID() string {
	return d.Meta.UNID
}

func (d *DocumentInfo) setUNID(unid string) {
	d.Meta.UNID = unid
}

func (d *DocumentInfo) getRevision() string {
	return d.Meta.Revision
}

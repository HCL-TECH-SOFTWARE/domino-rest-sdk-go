/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package types

type DominoBaseDocument struct {
	Meta     DominoDocumentMeta `json:"@meta"`
	Form     string             `json:"Form"`
	Warnings []string           `json:"@warnings"`
}
type DocumentBody struct {
	Fields map[string]interface{} `json:"fields"`
	DominoBaseDocument
}

type DocumentJSON struct {
	Form   string                 `json:"Form"`
	Fields map[string]interface{} `json:"fields"`
}

type DominoDocumentMeta struct {
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
	Size               string   `json:"size"`
}

type DominoRestDocument struct {
	DominoBaseDocument
	Fields    map[string]interface{}        `json:"fields"`
	GetUNID   func() string                 `json:"getUNID"`
	SetUNID   func(unid string)             `json:"setUNID"`
	ToDocJson func() map[string]interface{} `json:"toDocJson"`
	ToJson    func() DocumentBody           `json:"toJson"`
}

/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	"errors"
	"slices"

	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseDominoDoc = new(BaseDominoDocument)

type BaseDominoDocument struct {
	Meta     t.DominoDocumentMeta   `json:"@meta"`
	Form     string                 `json:"Form"`
	Warnings []string               `json:"@warnings"`
	Fields   map[string]interface{} `json:"fields"`
	t.DominoRestDocument
}

func DominoDocument(doc t.DocumentBody) (*BaseDominoDocument, error) {

	baseDominoDoc.Fields = make(map[string]interface{})

	if doc.Form == "" || len(doc.Form) == 0 {
		return nil, errors.New("domino document needs form value")
	}

	nonFieldKeys := []string{"@meta", "Form", "@warnings"}

	docInfo, err := h.StructTranspose(doc)
	if err != nil {
		return nil, err
	}

	for key, value := range docInfo["fields"].(map[string]interface{}) {
		if !slices.Contains(nonFieldKeys, key) {
			baseDominoDoc.Fields[key] = value
		}

	}

	baseDominoDoc.Meta = doc.Meta
	baseDominoDoc.Form = doc.Form
	baseDominoDoc.Warnings = doc.Warnings
	baseDominoDoc.GetUNID = GetUNID
	baseDominoDoc.SetUNID = SetUNID
	baseDominoDoc.ToDocJson = ToDocJson
	baseDominoDoc.ToJson = ToJson

	return baseDominoDoc, nil
}

func GetUNID() string {
	return baseDominoDoc.Meta.UNID
}

func SetUNID(unid string) {
	baseDominoDoc.Meta.UNID = unid
}

func ToDocJson() map[string]interface{} {

	json := make(map[string]interface{})
	json["Form"] = baseDominoDoc.Form

	for key, value := range baseDominoDoc.Fields {
		json[key] = value
	}

	return json
}

func ToJson() t.DocumentBody {
	var json t.DocumentBody

	json.Fields["@meta"] = baseDominoDoc.Meta
	json.Fields["@warnings"] = baseDominoDoc.Warnings
	json.Form = baseDominoDoc.Form

	for key, value := range baseDominoDoc.Fields {
		json.Fields[key] = value
	}

	return json
}

/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baselvOperations = new(BaseDominoListViewOperations)

type BaseDominoListViewOperations struct {
}

func getListViewEntry(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, listViewName string, options t.GetListViewEntryOptions) {

	if !options.Documents && options.Subscriber == nil {

	}

	if options.Documents {

	}

}

// func getListViewEntry(dataSource string, da t.DominoRestAccess, dc t.DominoRestConnector, listViewName string, options t.GetListViewEntryOptions) {

// }

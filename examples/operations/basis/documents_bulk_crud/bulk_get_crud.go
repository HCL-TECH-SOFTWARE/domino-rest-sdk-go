/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package documentsbulkcrud

import (
	"fmt"

	gosdk "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go"
)

func BulkGetDocumentSample(session *gosdk.SessionMethods) {

	unids := []string{"10BA715F2EB3B02885258A7B005D6DA2"}

	// you have the option to use the options variable which contains all of
	// the parameters or the POST /bulk/unid API (example: meta), refer to the
	// swagger of Domino REST API for more info.
	options := new(gosdk.BulkGetDocumentOptions)

	result, err := session.BulkGetDocument("customersdb", unids, *options)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package main

import (
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	"github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/logger"
)

func TestStructTranspose() {

	type Person struct {
		Id   string `json:"id"`
		Name string
		Age  int
	}

	result, _ := helpers.StructTranspose(Person{})
	logger.Pretty(result)
}

func TestOmit() {

	type Person struct {
		Id   string
		Name string
		Age  int
	}

	//result := helpers.Omit(Person{}, "Name")

	var result = helpers.Omit(Person{}, "Name")

	logger.Pretty(result.(Person))
}

func TestPick() {

}

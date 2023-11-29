/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package helpers

import (
	"time"
)

func IsJWTExpired(expiry int) bool {
	exp := int64(expiry)
	now := (time.Now().UnixNano() / int64(time.Millisecond))
	return exp <= now
}

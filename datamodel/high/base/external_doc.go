// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"github.com/pb33f/libopenapi/datamodel/high"
	low "github.com/pb33f/libopenapi/datamodel/low/base"
)

// ExternalDoc represents a high-level External Documentation object as defined by OpenAPI 2 and 3
//
// Allows referencing an external resource for extended documentation.
//  v2 - https://swagger.io/specification/v2/#externalDocumentationObject
//  v3 - https://spec.openapis.org/oas/v3.1.0#external-documentation-object
type ExternalDoc struct {
	Description string
	URL         string
	Extensions  map[string]any
	low         *low.ExternalDoc
}

// NewExternalDoc will create a new high-level External Documentation object from a low-level one.
func NewExternalDoc(extDoc *low.ExternalDoc) *ExternalDoc {
	d := new(ExternalDoc)
	d.low = extDoc
	if !extDoc.Description.IsEmpty() {
		d.Description = extDoc.Description.Value
	}
	if !extDoc.URL.IsEmpty() {
		d.URL = extDoc.URL.Value
	}
	d.Extensions = high.ExtractExtensions(extDoc.Extensions)
	return d
}

// GoLow returns the low-level ExternalDoc instance used to create the high-level one.
func (e *ExternalDoc) GoLow() *low.ExternalDoc {
	return e.low
}

func (e *ExternalDoc) GetExtensions() map[string]any {
	return e.Extensions
}
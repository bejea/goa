// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// storage HTTP server types
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design

package server

import (
	"unicode/utf8"

	goa "goa.design/goa"
	storage "goa.design/goa/examples/cellar/gen/storage"
	storageviews "goa.design/goa/examples/cellar/gen/storage/views"
)

// AddRequestBody is the type of the "storage" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// MultiUpdateRequestBody is the type of the "storage" service "multi_update"
// endpoint HTTP request body.
type MultiUpdateRequestBody struct {
	// Array of bottle info that matches the ids attribute
	Bottles []*BottleRequestBody `form:"bottles,omitempty" json:"bottles,omitempty" xml:"bottles,omitempty"`
}

// ListResponseBody is the type of the "storage" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredBottleTinyResponseBody

// ShowResponseBody is the type of the "storage" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryResponseBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponseBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "storage" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing bottle
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredBottleTinyResponseBody is used to define fields on response body types.
type StoredBottleTinyResponseBody struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryTinyResponseBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
}

// WineryTinyResponseBody is used to define fields on response body types.
type WineryTinyResponseBody struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// WineryResponseBody is used to define fields on response body types.
type WineryResponseBody struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentResponseBody is used to define fields on response body types.
type ComponentResponseBody struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// WineryRequestBody is used to define fields on request body types.
type WineryRequestBody struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentRequestBody is used to define fields on request body types.
type ComponentRequestBody struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// BottleRequestBody is used to define fields on request body types.
type BottleRequestBody struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryRequestBody `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentRequestBody `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "storage" service.
func NewListResponseBody(res storage.StoredBottleTinyCollection) ListResponseBody {
	body := make([]*StoredBottleTinyResponseBody, len(res))
	for i, val := range res {
		body[i] = &StoredBottleTinyResponseBody{
			ID:   &val.ID,
			Name: &val.Name,
		}
		if val.Winery != nil {
			body[i].Winery = marshalWineryTinyToWineryTinyResponseBody(val.Winery)
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "storage" service.
func NewShowResponseBody(res *storageviews.StoredBottleView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          res.Projected.ID,
		Name:        res.Projected.Name,
		Vintage:     res.Projected.Vintage,
		Description: res.Projected.Description,
		Rating:      res.Projected.Rating,
	}
	if res.Projected.Composition != nil {
		body.Composition = make([]*ComponentResponseBody, len(res.Projected.Composition))
		for j, val := range res.Projected.Composition {
			body.Composition[j] = &ComponentResponseBody{
				Varietal:   val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	if res.Projected.Winery != nil {
		body.Winery = marshalViewedWineryViewToWineryResponseBody(res.Projected.Winery)
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "storage" service.
func NewShowNotFoundResponseBody(res *storage.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: &res.Message,
		ID:      &res.ID,
	}
	return body
}

// NewShowShowPayload builds a storage service show endpoint payload.
func NewShowShowPayload(id string, view *string) *storage.ShowPayload {
	return &storage.ShowPayload{
		ID:   id,
		View: view,
	}
}

// NewAddBottle builds a storage service add endpoint payload.
func NewAddBottle(body *AddRequestBody) *storage.Bottle {
	v := &storage.Bottle{
		Name:        *body.Name,
		Vintage:     *body.Vintage,
		Description: body.Description,
		Rating:      body.Rating,
	}
	v.Winery = unmarshalWineryRequestBodyToWinery(body.Winery)
	if body.Composition != nil {
		v.Composition = make([]*storage.Component, len(body.Composition))
		for j, val := range body.Composition {
			v.Composition[j] = &storage.Component{
				Varietal:   *val.Varietal,
				Percentage: val.Percentage,
			}
		}
	}
	return v
}

// NewRemoveRemovePayload builds a storage service remove endpoint payload.
func NewRemoveRemovePayload(id string) *storage.RemovePayload {
	return &storage.RemovePayload{
		ID: id,
	}
}

// NewMultiAddBottle builds a storage service multi_add endpoint payload.
func NewMultiAddBottle(body []*BottleRequestBody) []*storage.Bottle {
	v := make([]*storage.Bottle, len(body))
	for i, val := range body {
		v[i] = &storage.Bottle{
			Name:        *val.Name,
			Vintage:     *val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		v[i].Winery = unmarshalWineryRequestBodyToWinery(val.Winery)
		if val.Composition != nil {
			v[i].Composition = make([]*storage.Component, len(val.Composition))
			for j, val := range val.Composition {
				v[i].Composition[j] = &storage.Component{
					Varietal:   *val.Varietal,
					Percentage: val.Percentage,
				}
			}
		}
	}
	return v
}

// NewMultiUpdateMultiUpdatePayload builds a storage service multi_update
// endpoint payload.
func NewMultiUpdateMultiUpdatePayload(body *MultiUpdateRequestBody, ids []string) *storage.MultiUpdatePayload {
	v := &storage.MultiUpdatePayload{}
	if body.Bottles != nil {
		v.Bottles = make([]*storage.Bottle, len(body.Bottles))
		for j, val := range body.Bottles {
			v.Bottles[j] = &storage.Bottle{
				Name:        *val.Name,
				Vintage:     *val.Vintage,
				Description: val.Description,
				Rating:      val.Rating,
			}
			v.Bottles[j].Winery = unmarshalWineryRequestBodyToWinery(val.Winery)
			if val.Composition != nil {
				v.Bottles[j].Composition = make([]*storage.Component, len(val.Composition))
				for k, val := range val.Composition {
					v.Bottles[j].Composition[k] = &storage.Component{
						Varietal:   *val.Varietal,
						Percentage: val.Percentage,
					}
				}
			}
		}
	}
	v.Ids = ids
	return v
}

// Validate runs the validations defined on AddRequestBody
func (body *AddRequestBody) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := body.Winery.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}

// Validate runs the validations defined on MultiUpdateRequestBody
func (body *MultiUpdateRequestBody) Validate() (err error) {
	for _, e := range body.Bottles {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Validate runs the validations defined on StoredBottleTinyResponseBody
func (body *StoredBottleTinyResponseBody) Validate() (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := body.Winery.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Validate runs the validations defined on WineryTinyResponseBody
func (body *WineryTinyResponseBody) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// Validate runs the validations defined on WineryRequestBody
func (body *WineryRequestBody) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "(?i)[a-z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "(?i)[a-z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// Validate runs the validations defined on ComponentRequestBody
func (body *ComponentRequestBody) Validate() (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}

// Validate runs the validations defined on BottleRequestBody
func (body *BottleRequestBody) Validate() (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := body.Winery.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}

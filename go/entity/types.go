// Typed models for the EvilInsultGenerator SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/evil-insult-generator-sdk/go/core"
)

// GenerateInsult is the typed data model for the generate_insult entity.
type GenerateInsult struct {
	Active *bool `json:"active,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Created *string `json:"created,omitempty"`
	Createdby *string `json:"createdby,omitempty"`
	Insult *string `json:"insult,omitempty"`
	Language *string `json:"language,omitempty"`
	Number *string `json:"number,omitempty"`
	Shown *string `json:"shown,omitempty"`
}

// GenerateInsultLoadMatch is the typed request payload for GenerateInsult.LoadTyped.
type GenerateInsultLoadMatch struct {
	Active *bool `json:"active,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Created *string `json:"created,omitempty"`
	Createdby *string `json:"createdby,omitempty"`
	Insult *string `json:"insult,omitempty"`
	Language *string `json:"language,omitempty"`
	Number *string `json:"number,omitempty"`
	Shown *string `json:"shown,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

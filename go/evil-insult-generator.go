package voxgigevilinsultgeneratorsdk

import (
	"github.com/voxgig-sdk/evil-insult-generator-sdk/core"
	"github.com/voxgig-sdk/evil-insult-generator-sdk/entity"
	"github.com/voxgig-sdk/evil-insult-generator-sdk/feature"
	_ "github.com/voxgig-sdk/evil-insult-generator-sdk/utility"
)

// Type aliases preserve external API.
type EvilInsultGeneratorSDK = core.EvilInsultGeneratorSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type EvilInsultGeneratorEntity = core.EvilInsultGeneratorEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type EvilInsultGeneratorError = core.EvilInsultGeneratorError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewGenerateInsultEntityFunc = func(client *core.EvilInsultGeneratorSDK, entopts map[string]any) core.EvilInsultGeneratorEntity {
		return entity.NewGenerateInsultEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewEvilInsultGeneratorSDK = core.NewEvilInsultGeneratorSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature

package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewGenerateInsultEntityFunc func(client *EvilInsultGeneratorSDK, entopts map[string]any) EvilInsultGeneratorEntity


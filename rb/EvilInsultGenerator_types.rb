# frozen_string_literal: true

# Typed models for the EvilInsultGenerator SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# GenerateInsult entity data model.
#
# @!attribute [rw] active
#   @return [Boolean, nil]
#
# @!attribute [rw] comment
#   @return [String, nil]
#
# @!attribute [rw] created
#   @return [String, nil]
#
# @!attribute [rw] createdby
#   @return [String, nil]
#
# @!attribute [rw] insult
#   @return [String, nil]
#
# @!attribute [rw] language
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [String, nil]
#
# @!attribute [rw] shown
#   @return [String, nil]
GenerateInsult = Struct.new(
  :active,
  :comment,
  :created,
  :createdby,
  :insult,
  :language,
  :number,
  :shown,
  keyword_init: true
)

# Match filter for GenerateInsult#load (any subset of GenerateInsult fields).
#
# @!attribute [rw] active
#   @return [Boolean, nil]
#
# @!attribute [rw] comment
#   @return [String, nil]
#
# @!attribute [rw] created
#   @return [String, nil]
#
# @!attribute [rw] createdby
#   @return [String, nil]
#
# @!attribute [rw] insult
#   @return [String, nil]
#
# @!attribute [rw] language
#   @return [String, nil]
#
# @!attribute [rw] number
#   @return [String, nil]
#
# @!attribute [rw] shown
#   @return [String, nil]
GenerateInsultLoadMatch = Struct.new(
  :active,
  :comment,
  :created,
  :createdby,
  :insult,
  :language,
  :number,
  :shown,
  keyword_init: true
)


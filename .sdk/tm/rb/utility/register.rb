# EvilInsultGenerator SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

EvilInsultGeneratorUtility.registrar = ->(u) {
  u.clean = EvilInsultGeneratorUtilities::Clean
  u.done = EvilInsultGeneratorUtilities::Done
  u.make_error = EvilInsultGeneratorUtilities::MakeError
  u.feature_add = EvilInsultGeneratorUtilities::FeatureAdd
  u.feature_hook = EvilInsultGeneratorUtilities::FeatureHook
  u.feature_init = EvilInsultGeneratorUtilities::FeatureInit
  u.fetcher = EvilInsultGeneratorUtilities::Fetcher
  u.make_fetch_def = EvilInsultGeneratorUtilities::MakeFetchDef
  u.make_context = EvilInsultGeneratorUtilities::MakeContext
  u.make_options = EvilInsultGeneratorUtilities::MakeOptions
  u.make_request = EvilInsultGeneratorUtilities::MakeRequest
  u.make_response = EvilInsultGeneratorUtilities::MakeResponse
  u.make_result = EvilInsultGeneratorUtilities::MakeResult
  u.make_point = EvilInsultGeneratorUtilities::MakePoint
  u.make_spec = EvilInsultGeneratorUtilities::MakeSpec
  u.make_url = EvilInsultGeneratorUtilities::MakeUrl
  u.param = EvilInsultGeneratorUtilities::Param
  u.prepare_auth = EvilInsultGeneratorUtilities::PrepareAuth
  u.prepare_body = EvilInsultGeneratorUtilities::PrepareBody
  u.prepare_headers = EvilInsultGeneratorUtilities::PrepareHeaders
  u.prepare_method = EvilInsultGeneratorUtilities::PrepareMethod
  u.prepare_params = EvilInsultGeneratorUtilities::PrepareParams
  u.prepare_path = EvilInsultGeneratorUtilities::PreparePath
  u.prepare_query = EvilInsultGeneratorUtilities::PrepareQuery
  u.result_basic = EvilInsultGeneratorUtilities::ResultBasic
  u.result_body = EvilInsultGeneratorUtilities::ResultBody
  u.result_headers = EvilInsultGeneratorUtilities::ResultHeaders
  u.transform_request = EvilInsultGeneratorUtilities::TransformRequest
  u.transform_response = EvilInsultGeneratorUtilities::TransformResponse
}

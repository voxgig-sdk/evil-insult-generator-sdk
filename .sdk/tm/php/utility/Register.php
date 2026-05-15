<?php
declare(strict_types=1);

// EvilInsultGenerator SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

EvilInsultGeneratorUtility::setRegistrar(function (EvilInsultGeneratorUtility $u): void {
    $u->clean = [EvilInsultGeneratorClean::class, 'call'];
    $u->done = [EvilInsultGeneratorDone::class, 'call'];
    $u->make_error = [EvilInsultGeneratorMakeError::class, 'call'];
    $u->feature_add = [EvilInsultGeneratorFeatureAdd::class, 'call'];
    $u->feature_hook = [EvilInsultGeneratorFeatureHook::class, 'call'];
    $u->feature_init = [EvilInsultGeneratorFeatureInit::class, 'call'];
    $u->fetcher = [EvilInsultGeneratorFetcher::class, 'call'];
    $u->make_fetch_def = [EvilInsultGeneratorMakeFetchDef::class, 'call'];
    $u->make_context = [EvilInsultGeneratorMakeContext::class, 'call'];
    $u->make_options = [EvilInsultGeneratorMakeOptions::class, 'call'];
    $u->make_request = [EvilInsultGeneratorMakeRequest::class, 'call'];
    $u->make_response = [EvilInsultGeneratorMakeResponse::class, 'call'];
    $u->make_result = [EvilInsultGeneratorMakeResult::class, 'call'];
    $u->make_point = [EvilInsultGeneratorMakePoint::class, 'call'];
    $u->make_spec = [EvilInsultGeneratorMakeSpec::class, 'call'];
    $u->make_url = [EvilInsultGeneratorMakeUrl::class, 'call'];
    $u->param = [EvilInsultGeneratorParam::class, 'call'];
    $u->prepare_auth = [EvilInsultGeneratorPrepareAuth::class, 'call'];
    $u->prepare_body = [EvilInsultGeneratorPrepareBody::class, 'call'];
    $u->prepare_headers = [EvilInsultGeneratorPrepareHeaders::class, 'call'];
    $u->prepare_method = [EvilInsultGeneratorPrepareMethod::class, 'call'];
    $u->prepare_params = [EvilInsultGeneratorPrepareParams::class, 'call'];
    $u->prepare_path = [EvilInsultGeneratorPreparePath::class, 'call'];
    $u->prepare_query = [EvilInsultGeneratorPrepareQuery::class, 'call'];
    $u->result_basic = [EvilInsultGeneratorResultBasic::class, 'call'];
    $u->result_body = [EvilInsultGeneratorResultBody::class, 'call'];
    $u->result_headers = [EvilInsultGeneratorResultHeaders::class, 'call'];
    $u->transform_request = [EvilInsultGeneratorTransformRequest::class, 'call'];
    $u->transform_response = [EvilInsultGeneratorTransformResponse::class, 'call'];
});

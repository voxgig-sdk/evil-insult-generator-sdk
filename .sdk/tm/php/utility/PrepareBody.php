<?php
declare(strict_types=1);

// EvilInsultGenerator SDK utility: prepare_body

class EvilInsultGeneratorPrepareBody
{
    public static function call(EvilInsultGeneratorContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}

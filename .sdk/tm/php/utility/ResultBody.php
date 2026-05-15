<?php
declare(strict_types=1);

// EvilInsultGenerator SDK utility: result_body

class EvilInsultGeneratorResultBody
{
    public static function call(EvilInsultGeneratorContext $ctx): ?EvilInsultGeneratorResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}

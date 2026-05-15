<?php
declare(strict_types=1);

// EvilInsultGenerator SDK utility: result_headers

class EvilInsultGeneratorResultHeaders
{
    public static function call(EvilInsultGeneratorContext $ctx): ?EvilInsultGeneratorResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}

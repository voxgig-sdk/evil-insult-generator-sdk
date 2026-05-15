<?php
declare(strict_types=1);

// EvilInsultGenerator SDK utility: feature_add

class EvilInsultGeneratorFeatureAdd
{
    public static function call(EvilInsultGeneratorContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}

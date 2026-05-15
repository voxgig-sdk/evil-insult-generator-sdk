<?php
declare(strict_types=1);

// EvilInsultGenerator SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class EvilInsultGeneratorMakeContext
{
    public static function call(array $ctxmap, ?EvilInsultGeneratorContext $basectx): EvilInsultGeneratorContext
    {
        return new EvilInsultGeneratorContext($ctxmap, $basectx);
    }
}

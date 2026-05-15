<?php
declare(strict_types=1);

// EvilInsultGenerator SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class EvilInsultGeneratorFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new EvilInsultGeneratorBaseFeature();
            case "test":
                return new EvilInsultGeneratorTestFeature();
            default:
                return new EvilInsultGeneratorBaseFeature();
        }
    }
}

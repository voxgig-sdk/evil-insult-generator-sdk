<?php
declare(strict_types=1);

// EvilInsultGenerator SDK exists test

require_once __DIR__ . '/../evilinsultgenerator_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = EvilInsultGeneratorSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}

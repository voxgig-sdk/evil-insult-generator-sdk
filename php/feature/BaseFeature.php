<?php
declare(strict_types=1);

// EvilInsultGenerator SDK base feature

class EvilInsultGeneratorBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(EvilInsultGeneratorContext $ctx, array $options): void {}
    public function PostConstruct(EvilInsultGeneratorContext $ctx): void {}
    public function PostConstructEntity(EvilInsultGeneratorContext $ctx): void {}
    public function SetData(EvilInsultGeneratorContext $ctx): void {}
    public function GetData(EvilInsultGeneratorContext $ctx): void {}
    public function GetMatch(EvilInsultGeneratorContext $ctx): void {}
    public function SetMatch(EvilInsultGeneratorContext $ctx): void {}
    public function PrePoint(EvilInsultGeneratorContext $ctx): void {}
    public function PreSpec(EvilInsultGeneratorContext $ctx): void {}
    public function PreRequest(EvilInsultGeneratorContext $ctx): void {}
    public function PreResponse(EvilInsultGeneratorContext $ctx): void {}
    public function PreResult(EvilInsultGeneratorContext $ctx): void {}
    public function PreDone(EvilInsultGeneratorContext $ctx): void {}
    public function PreUnexpected(EvilInsultGeneratorContext $ctx): void {}
}

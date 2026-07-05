<?php
declare(strict_types=1);

// Typed models for the EvilInsultGenerator SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** GenerateInsult entity data model. */
class GenerateInsult
{
    public ?bool $active = null;
    public ?string $comment = null;
    public ?string $created = null;
    public ?string $createdby = null;
    public ?string $insult = null;
    public ?string $language = null;
    public ?string $number = null;
    public ?string $shown = null;
}

/** Request payload for GenerateInsult#load. */
class GenerateInsultLoadMatch
{
    public ?bool $active = null;
    public ?string $comment = null;
    public ?string $created = null;
    public ?string $createdby = null;
    public ?string $insult = null;
    public ?string $language = null;
    public ?string $number = null;
    public ?string $shown = null;
}


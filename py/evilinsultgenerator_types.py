# Typed models for the EvilInsultGenerator SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class GenerateInsult:
    active: Optional[bool] = None
    comment: Optional[str] = None
    created: Optional[str] = None
    createdby: Optional[str] = None
    insult: Optional[str] = None
    language: Optional[str] = None
    number: Optional[str] = None
    shown: Optional[str] = None


@dataclass
class GenerateInsultLoadMatch:
    active: Optional[bool] = None
    comment: Optional[str] = None
    created: Optional[str] = None
    createdby: Optional[str] = None
    insult: Optional[str] = None
    language: Optional[str] = None
    number: Optional[str] = None
    shown: Optional[str] = None


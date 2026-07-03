# EvilInsultGenerator SDK

Evil Insult Generator client, generated from the OpenAPI spec.

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## Try it

**TypeScript**
```bash
npm install evil-insult-generator
```

**Python**
```bash
pip install evil-insult-generator-sdk
```

**PHP**
```bash
composer require voxgig/evil-insult-generator-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/evil-insult-generator-sdk/go
```

**Ruby**
```bash
gem install evil-insult-generator-sdk
```

**Lua**
```bash
luarocks install evil-insult-generator-sdk
```

## Quickstart

### TypeScript

```ts
import { EvilInsultGeneratorSDK } from 'evil-insult-generator'

const client = new EvilInsultGeneratorSDK({
  apikey: process.env.EVIL-INSULT-GENERATOR_APIKEY,
})

// Load generateinsult data
const generateinsult = await client.GenerateInsult().load({})
console.log(generateinsult.data)
```

See the [TypeScript README](ts/README.md) for the full guide.

## Surfaces

| Surface | Path |
| --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | `go-cli/` |
| **MCP server** | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o evil-insult-generator-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "evil-insult-generator": {
      "command": "/abs/path/to/evil-insult-generator-mcp"
    }
  }
}
```

## Entities

The API exposes one entity:

| Entity | Description | API path |
| --- | --- | --- |
| **GenerateInsult** |  | `/generate_insult.php` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
import os
from evilinsultgenerator_sdk import EvilInsultGeneratorSDK

client = EvilInsultGeneratorSDK({
    "apikey": os.environ.get("EVIL-INSULT-GENERATOR_APIKEY"),
})


# Load a specific generateinsult
generateinsult, err = client.GenerateInsult().load({"id": "example_id"})
print(generateinsult)
```

### PHP

```php
<?php
require_once 'evilinsultgenerator_sdk.php';

$client = new EvilInsultGeneratorSDK([
    "apikey" => getenv("EVIL-INSULT-GENERATOR_APIKEY"),
]);


// Load a specific generateinsult
[$generateinsult, $err] = $client->GenerateInsult()->load(["id" => "example_id"]);
print_r($generateinsult);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/evil-insult-generator-sdk/go"

client := sdk.NewEvilInsultGeneratorSDK(map[string]any{
    "apikey": os.Getenv("EVIL-INSULT-GENERATOR_APIKEY"),
})

// Load generateinsult data
generateinsult, err := client.GenerateInsult(nil).Load(map[string]any{}, nil)
fmt.Println(generateinsult)
```

### Ruby

```ruby
require_relative "EvilInsultGenerator_sdk"

client = EvilInsultGeneratorSDK.new({
  "apikey" => ENV["EVIL-INSULT-GENERATOR_APIKEY"],
})


# Load a specific generateinsult
generateinsult, err = client.GenerateInsult().load({ "id" => "example_id" })
puts generateinsult
```

### Lua

```lua
local sdk = require("evil-insult-generator_sdk")

local client = sdk.new({
  apikey = os.getenv("EVIL-INSULT-GENERATOR_APIKEY"),
})


-- Load a specific generateinsult
local generateinsult, err = client:GenerateInsult():load({ id = "example_id" })
print(generateinsult)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = EvilInsultGeneratorSDK.test()
const result = await client.GenerateInsult().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = EvilInsultGeneratorSDK.test()
result, err = client.GenerateInsult().load({"id": "test01"})
```

### PHP

```php
$client = EvilInsultGeneratorSDK::test();
[$result, $err] = $client->GenerateInsult()->load(["id" => "test01"]);
```

### Golang

```go
client := sdk.Test()
result, err := client.GenerateInsult(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = EvilInsultGeneratorSDK.test
result, err = client.GenerateInsult().load({ "id" => "test01" })
```

### Lua

```lua
local client = sdk.test()
local result, err = client:GenerateInsult():load({ id = "test01" })
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

---

Generated from the Evil Insult Generator OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

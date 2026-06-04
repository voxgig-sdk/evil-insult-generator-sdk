# EvilInsultGenerator SDK

Fetch humorous, over-the-top insults in plain text, XML, or JSON across several languages

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Evil Insult Generator

[Evil Insult Generator](https://evilinsult.com) is a long-running novelty service run by the Evil Insult Generator Team that returns randomly generated, intentionally rude insults. The same backend powers the website, an Android app, a Windows 10 app, and Telegram/Facebook bots, and the project also has an open-source code repository on GitHub.

What you get from the API:

- A single `GET` endpoint at `/generate_insult.php` that returns one insult per call.
- A `lang` query parameter to pick the language. The site lists support for `cn`, `de`, `el`, `en`, `es`, `fr`, `ru`, and `sw` (Chinese, German, Greek, English, Spanish, French, Russian, Swahili).
- A response-format selector so the same endpoint can return plain text, XML, or JSON.

The API is open and requires no authentication or API key. Community monitoring at [freepublicapis.com](https://freepublicapis.com/evil-insult-generator) reports high uptime and sub-200 ms response times, and notes that CORS is not enabled, so browser callers may need to proxy requests through their own backend.

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

## 30-second quickstart

### TypeScript

```ts
import { EvilInsultGeneratorSDK } from 'evil-insult-generator'

const client = new EvilInsultGeneratorSDK({})

```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

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
| **GenerateInsult** | Wraps the single insult endpoint `GET /generate_insult.php`, with options for language (`lang`) and response format (plain text, XML, or JSON). | `/generate_insult.php` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from evilinsultgenerator_sdk import EvilInsultGeneratorSDK

client = EvilInsultGeneratorSDK({})


# Load a specific generateinsult
generateinsult, err = client.GenerateInsult(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'evilinsultgenerator_sdk.php';

$client = new EvilInsultGeneratorSDK([]);


// Load a specific generateinsult
[$generateinsult, $err] = $client->GenerateInsult(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/evil-insult-generator-sdk/go"

client := sdk.NewEvilInsultGeneratorSDK(map[string]any{})

```

### Ruby

```ruby
require_relative "EvilInsultGenerator_sdk"

client = EvilInsultGeneratorSDK.new({})


# Load a specific generateinsult
generateinsult, err = client.GenerateInsult(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("evil-insult-generator_sdk")

local client = sdk.new({})


-- Load a specific generateinsult
local generateinsult, err = client:GenerateInsult(nil):load(
  { id = "example_id" }, nil
)
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
client = EvilInsultGeneratorSDK.test(None, None)
result, err = client.GenerateInsult(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = EvilInsultGeneratorSDK::test(null, null);
[$result, $err] = $client->GenerateInsult(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.GenerateInsult(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = EvilInsultGeneratorSDK.test(nil, nil)
result, err = client.GenerateInsult(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:GenerateInsult(nil):load(
  { id = "test01" }, nil
)
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

## Using the Evil Insult Generator

- Upstream: [https://evilinsult.com](https://evilinsult.com)

- The Evil Insult Generator site links to a `legal.pdf` document but does not publish explicit licence terms on its homepage.
- No attribution, auth key, or sign-up is required to call the public endpoint.
- Treat the generated text as adult humour: insults are intentionally crude and may be offensive.
- Confirm acceptable use with the maintainers before redistributing the content commercially.

---

Generated from the Evil Insult Generator OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

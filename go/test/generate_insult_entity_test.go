package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/evil-insult-generator-sdk"
	"github.com/voxgig-sdk/evil-insult-generator-sdk/core"

	vs "github.com/voxgig/struct"
)

func TestGenerateInsultEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.GenerateInsult(nil)
		if ent == nil {
			t.Fatal("expected non-nil GenerateInsultEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := generate_insultBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "generate_insult." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set EVILINSULTGENERATOR_TEST_GENERATE_INSULT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		generateInsultRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.generate_insult", setup.data)))
		var generateInsultRef01Data map[string]any
		if len(generateInsultRef01DataRaw) > 0 {
			generateInsultRef01Data = core.ToMapAny(generateInsultRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = generateInsultRef01Data

		// LOAD
		generateInsultRef01Ent := client.GenerateInsult(nil)
		generateInsultRef01MatchDt0 := map[string]any{}
		generateInsultRef01DataDt0Loaded, err := generateInsultRef01Ent.Load(generateInsultRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if generateInsultRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func generate_insultBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "generate_insult", "GenerateInsultTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read generate_insult test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse generate_insult test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"generate_insult01", "generate_insult02", "generate_insult03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("EVILINSULTGENERATOR_TEST_GENERATE_INSULT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"EVILINSULTGENERATOR_TEST_GENERATE_INSULT_ENTID": idmap,
		"EVILINSULTGENERATOR_TEST_LIVE":      "FALSE",
		"EVILINSULTGENERATOR_TEST_EXPLAIN":   "FALSE",
		"EVILINSULTGENERATOR_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["EVILINSULTGENERATOR_TEST_GENERATE_INSULT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["EVILINSULTGENERATOR_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["EVILINSULTGENERATOR_APIKEY"],
			},
			extra,
		})
		client = sdk.NewEvilInsultGeneratorSDK(core.ToMapAny(mergedOpts))
	}

	live := env["EVILINSULTGENERATOR_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["EVILINSULTGENERATOR_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}

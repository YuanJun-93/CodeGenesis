//go:build integration

package generator

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/constant"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "../../../configs/code-genesis-api.yaml", "the config file")

func TestUseGenerator_Integration(t *testing.T) {
	// Parse flags to allow overriding config path if needed
	flag.Parse()

	// 1. Load Config
	var c config.Config
	// Resolve absolute path for config to avoid relative path issues during test execution
	absPath, err := filepath.Abs(*configFile)
	if err != nil {
		t.Fatalf("Failed to resolve config path: %v", err)
	}

	// Check if config file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		t.Skipf("Config file not found at %s, skipping integration test", absPath)
		return
	}

	conf.MustLoad(absPath, &c)

	// 2. Mock ServiceContext (We only need Config for this logic)
	// Note: We do NOT call svc.NewServiceContext(c) to avoid connecting to DB/Redis
	svcCtx := &svc.ServiceContext{
		Config: c,
	}

	// 3. Initialize Logic
	logic := NewUseGeneratorLogic(context.Background(), svcCtx)
	// Disable logx output for cleaner test results
	logx.Disable()

	// 4. Prepare Request
	req := &types.GeneratorRequest{
		Message: "Generate a simple Hello World HTML page with red text.",
		Type:    constant.GeneratorTypeHtml,
	}

	// 5. Execute
	t.Logf("Sending request to AI Provider: %s (Model: %s)...", c.Ai.Provider, c.Ai.Model)
	resp, err := logic.UseGenerator(req)

	// 6. Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result, "Result path should not be empty")

	t.Logf("Generation Successful! Saved to: %s", resp.Result)

	// Verify file actually exists
	_, err = os.Stat(resp.Result)
	assert.NoError(t, err, "Saved file should exist on disk")
}

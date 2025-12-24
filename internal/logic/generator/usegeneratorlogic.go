package generator

import (
	"context"
	"fmt"

	"github.com/YuanJun-93/CodeGenesis/internal/pkg/ai"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/constant"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/parser"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/saver"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseGeneratorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseGeneratorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseGeneratorLogic {
	return &UseGeneratorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseGeneratorLogic) UseGenerator(req *types.GeneratorRequest) (resp *types.GeneratorResponse, err error) {
	// 1. Init AI Client
	client := ai.NewAiClient(l.svcCtx.Config)

	// 2. Call AI (Synchronous)
	// Construct a simple system prompt based on type
	systemPrompt := "You are a helpful code generator."
	if req.Type == constant.GeneratorTypeHtml {
		systemPrompt += " Please generate a single HTML file."
	}

	rawContent, err := client.DoRequest(systemPrompt, req.Message)
	if err != nil {
		return nil, fmt.Errorf("AI request failed: %w", err)
	}

	// 3. Parse Code (Extract all parts)
	parseResult := parser.ParseMultiFile(rawContent)

	// 4. Save (Logic handled inside Saver based on type)
	// req.Type matches "html" or "multi_file" (from frontend/java logic)
	savedPath, err := saver.SaveCode(parseResult, req.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to save code: %w", err)
	}

	// 5. Return Result
	return &types.GeneratorResponse{
		Result: savedPath,
	}, nil
}

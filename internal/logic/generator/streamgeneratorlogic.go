package generator

import (
	"context"
	"strings"

	"github.com/YuanJun-93/CodeGenesis/internal/pkg/ai"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/constant"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/parser"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/saver"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StreamGeneratorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStreamGeneratorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StreamGeneratorLogic {
	return &StreamGeneratorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// StreamGenerator returns a channel that streams content, and handles saving internally
func (l *StreamGeneratorLogic) StreamGenerator(req *types.GeneratorRequest) (chan string, error) {
	client := ai.NewAiClient(l.svcCtx.Config)

	systemPrompt := "You are a helpful code generator."
	if req.Type == constant.GeneratorTypeHtml {
		systemPrompt += " Please generate a single HTML file."
	}

	// 1. Get raw stream from AI
	aiStream, err := client.DoStreamRequest(systemPrompt, req.Message)
	if err != nil {
		return nil, err
	}

	// 2. Create a channel for the Handler
	handlerStream := make(chan string)

	// 3. Process stream in background (Accumulate + Forward)
	go func() {
		defer close(handlerStream)

		var fullContentBuilder strings.Builder

		for chunk := range aiStream {
			// Forward to handler
			handlerStream <- chunk
			// Accumulate for saving
			fullContentBuilder.WriteString(chunk)
		}

		// Stream finished, now save
		fullContent := fullContentBuilder.String()
		parseResult := parser.ParseMultiFile(fullContent)

		savePath, err := saver.SaveCode(parseResult, req.Type)
		if err != nil {
			logx.Errorf("Failed to save stream code: %v", err)
		} else {
			logx.Infof("Stream code saved to: %s", savePath)
		}
	}()

	return handlerStream, nil
}

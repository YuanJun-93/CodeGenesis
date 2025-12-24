package generator

import (
	"fmt"
	"net/http"

	"github.com/YuanJun-93/CodeGenesis/internal/logic/generator"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StreamGeneratorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeneratorRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Transfer-Encoding", "chunked")

		l := generator.NewStreamGeneratorLogic(r.Context(), svcCtx)
		stream, err := l.StreamGenerator(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		flusher, ok := w.(http.Flusher)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("streaming not supported"))
			return
		}

		for chunk := range stream {
			// SSE format: data: <content>\n\n
			// We might need to escape newlines if strict SSE, but efficient AI stream usually sends safe chunks
			// For simplicity, just sending raw string as data
			fmt.Fprintf(w, "data: %s\n\n", chunk)
			flusher.Flush()
		}

		// End of stream
		fmt.Fprintf(w, "data: [DONE]\n\n")
		flusher.Flush()
	}
}

package groq

import (
	"io"
	"net/http"
)

type ReadWriter struct {
	w http.ResponseWriter
	r *http.Response
}

func NewReadWriter(w http.ResponseWriter, r *http.Response) *ReadWriter {
	return &ReadWriter{
		w: w,
		r: r,
	}
}

func (r *ReadWriter) setResponseHeader() {
	r.w.Header().Set("Content-Type", r.r.Header.Get("Content-Type"))
	// 这些头部设置通常用于提高性能，控制缓存行为，以及管理持久连接和流式传输。
	r.w.Header().Set("Cache-Control", "no-cache")
	r.w.Header().Set("Connection", "keep-alive")
	r.w.Header().Set("Transfer-Encoding", "chunked")
	r.w.Header().Set("X-Accel-Buffering", "no")

	// Set headers for CORS
	r.w.Header().Set("Access-Control-Allow-Origin", "*")
	r.w.Header().Set("Access-Control-Allow-Methods", "*")
	r.w.Header().Set("Access-Control-Allow-Headers", "*")
}

func (r *ReadWriter) StreamHandler() {
	defer r.r.Body.Close()
	r.setResponseHeader()
	buf := make([]byte, 4*1024)
	for {
		n, err := r.r.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		if n == 0 {
			continue
		}
		r.w.Write(buf[:n])
	}
}

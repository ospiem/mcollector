// Package compress provides middleware for compressing and decompressing HTTP requests and responses.
package compress

import (
	"fmt"
	"io"
	"net/http"

	gzip "github.com/klauspost/compress/gzip"
	"github.com/rs/zerolog"
)

// compressFunc represents the compression function used.
const compressFunc = "gzip"

// contentEncoding represents the content encoding type.
const contentEncoding = "Content-Encoding"

var allowedContentTypes = []string{
	"application/javascript",
	"application/json",
	"text/css",
	"text/html",
	"text/plain",
	"text/xml",
}

// gzipWriter wraps http.ResponseWriter to write compressed data.
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Write writes compressed data to underlying writer.
func (w gzipWriter) Write(b []byte) (int, error) {
	ww, err := w.Writer.Write(b)
	if err != nil {
		return 0, fmt.Errorf("cannot write with gzip: %w", err)
	}
	return ww, nil
}

// DecompressRequest returns a middleware that decompresses incoming requests if necessary.
func DecompressRequest(log zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const wrapErr = "middleware compressor"

			// If content type does not match with allowedContentTypes stop processing and return to next handler
			if !matchContentTypes(r.Header.Values("Content-Type"), allowedContentTypes) {
				log.Debug().Msgf("did not match Content-Type in DecompressRequest")
				next.ServeHTTP(w, r)
				return
			}
			log.Debug().Msgf("matched Content-Type in DecompressRequest")

			// If compress function does not match with compressFunc stop processing and return to next handler
			if !matchCompressFunc(r.Header.Values(contentEncoding), compressFunc) {
				log.Debug().Msgf("did not match %s", contentEncoding)
				next.ServeHTTP(w, r)
				return
			}

			decompressed, err := decompressGzip(r.Body, log)
			if err != nil {
				log.Error().Err(err).Msg(wrapErr)
				http.Error(w, "failed to decompress data", http.StatusInternalServerError)
				return
			}

			r.Body = io.NopCloser(decompressed)

			next.ServeHTTP(w, r)
		})
	}
}

// CompressResponse returns a middleware that compress outgoing responses if necessary.
func CompressResponse(log zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const wrapError = "middleware compressor"

			// If client does not support compressed body stop processing and return to next handler
			if !matchCompressFunc(r.Header.Values("Accept-Encoding"), compressFunc) {
				log.Debug().Msg("did not match  Accept-Encoding")
				next.ServeHTTP(w, r)
				return
			}
			log.Debug().Msg("match Accept-Encoding")
			gz, err := gzip.NewWriterLevel(w, gzip.BestCompression)
			if err != nil {
				log.Error().Err(err).Msg(wrapError)
				return
			}

			defer func() {
				if err := gz.Close(); err != nil {
					log.Error().Err(err).Msg("cannot close gzip in compress response")
				}
			}()
			w.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
		})
	}
}

func matchContentTypes(headers []string, allowedTypes []string) bool {
	if len(headers) == 0 {
		return false
	}
	for _, curType := range headers {
		for _, allowType := range allowedTypes {
			if curType == allowType {
				return true
			}
		}
	}
	return false
}

func matchCompressFunc(headers []string, compressFunc string) bool {
	if len(headers) == 0 {
		return false
	}
	for _, h := range headers {
		if h == compressFunc {
			return true
		}
	}
	return false
}

func decompressGzip(data io.Reader, log zerolog.Logger) (io.Reader, error) {
	r, err := gzip.NewReader(data)
	if err != nil {
		return nil, fmt.Errorf("decompressGzip error: %w", err)
	}
	func() {
		err := r.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close gzip.NewReader")
		}
	}()

	return r, nil
}

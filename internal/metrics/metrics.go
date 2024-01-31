package metrics

import "net/http"

type MetricsResponseWriter struct {
	Wrapped       http.ResponseWriter
	StatusCode    int
	HeaderWritten bool
}

func (mw *MetricsResponseWriter) Header() http.Header {
	return mw.Wrapped.Header()
}

func (mw *MetricsResponseWriter) WriteHeader(statusCode int) {
	mw.Wrapped.WriteHeader(statusCode)

	if !mw.HeaderWritten {
		mw.StatusCode = statusCode
		mw.HeaderWritten = true
	}
}

func (mw *MetricsResponseWriter) Write(b []byte) (int, error) {
	if !mw.HeaderWritten {
		mw.StatusCode = http.StatusOK
		mw.HeaderWritten = true
	}

	return mw.Wrapped.Write(b)
}

func (mw *MetricsResponseWriter) Unwrap() http.ResponseWriter {
	return mw.Wrapped
}

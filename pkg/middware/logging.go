package middleware

//type bodyLogWriter struct {
//	gin.ResponseWriter
//	body *bytes.Buffer
//}
//
//func (w bodyLogWriter) Write(b []byte) (int, error) {
//	w.body.Write(b)
//	return w.ResponseWriter.Write(b)
//}
//
//// Logging is a middleware function that logs the each request.
//func Logging() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now().UTC()
//		path := c.Request.URL.Path
//
//		// Skip for the health check requests.
//		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
//			return
//		}
//
//		// Read the Body content
//		var bodyBytes []byte
//		if c.Request.Body != nil {
//			bodyBytes, _ = io.ReadAll(c.Request.Body)
//		}
//
//		// Restore the io.ReadCloser to its original state
//		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
//
//		// The basic informations.
//		method := c.Request.Method
//		ip := c.ClientIP()
//
//		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
//		blw := &bodyLogWriter{
//			body:           bytes.NewBufferString(""),
//			ResponseWriter: c.Writer,
//		}
//		c.Writer = blw
//
//		// Continue.
//		c.Next()
//
//		// Calculates the latency.
//		end := time.Now().UTC()
//		cost := end.Sub(start)
//
//		code, message := -1, ""
//
//		// get code and message
//		var response handler.Response
//		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
//			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
//			code = errno.InternalServerError.Code
//			message = err.Error()
//		} else {
//			code = response.Code
//			message = response.Message
//		}
//
//		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
//	}
//}

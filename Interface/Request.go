package Interface

type Request interface {
	GET(url string, args ...interface{}) (response string, err error)
	GETWithHeaders(url string, headers map[string]string, args ...interface{}) (response string, err error)
	GETWithCookies(url string, cookies map[string]string, args ...interface{}) (response string, err error)
	POST(url string, args ...interface{}) (response string, err error)
}

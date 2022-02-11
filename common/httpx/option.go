package httpx

import (
	"time"
)

// Options contains configuration options for the client
type Options struct {
	RandomAgent      bool
	DefaultUserAgent string
	HTTPProxy        string
	SocksProxy       string
	Threads          int
	CdnCheck         bool
	ExcludeCdn       bool
	// Timeout is the maximum time to wait for the request
	Timeout time.Duration
	// RetryMax is the maximum number of retries
	RetryMax      int
	CustomHeaders map[string]string
	// VHostSimilarityRatio 1 - 100
	VHostSimilarityRatio int
	FollowRedirects      bool
	FollowHostRedirects  bool
	MaxRedirects         int
	Unsafe               bool
	TLSGrab              bool
	// VHOSTs options
	VHostIgnoreStatusCode     bool
	VHostIgnoreContentLength  bool
	VHostIgnoreNumberOfWords  bool
	VHostIgnoreNumberOfLines  bool
	VHostStripHTML            bool
	Allow                     []string
	Deny                      []string
	MaxResponseBodySizeToSave int64
	MaxResponseBodySizeToRead int64
	UnsafeURI                 string
}

// DefaultOptions contains the default options
var DefaultOptions = Options{
	RandomAgent:  true,
	Threads:      25,
	Timeout:      30 * time.Second,
	RetryMax:     5,
	MaxRedirects: 10,
	Unsafe:       false,
	CdnCheck:     true,
	ExcludeCdn:   false,
	// VHOSTs options
	VHostIgnoreStatusCode:    false,
	VHostIgnoreContentLength: true,
	VHostIgnoreNumberOfWords: false,
	VHostIgnoreNumberOfLines: false,
	VHostStripHTML:           false,
	VHostSimilarityRatio:     85,
	DefaultUserAgent:         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36 ppbsecurity/httpx",
}

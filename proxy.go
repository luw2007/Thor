package thor

import (
	"net/http"
	"net/url"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/levigross/grequests"
)

const (
	defaultCheckURI = "http://baidu.com"
)

type proxy struct {
	addr       *url.URL
	usable     bool
	checkTiker *time.Ticker
	client     *grequests.Session
	entry      *logrus.Entry
}

func NewProxy(addr *url.URL, entry *logrus.Entry) *proxy {
	ro := &grequests.RequestOptions{
		Proxies:   map[string]*url.URL{"http": addr, "https": addr},
		UserAgent: "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36",
	}

	return &proxy{
		addr:       addr,
		usable:     false,
		checkTiker: time.NewTicker(time.Millisecond),
		client:     grequests.NewSession(ro),
		entry: entry.WithFields(logrus.Fields{
			"addr": addr.String(),
			"type": "proxy",
		}),
	}
}

func (p *proxy) check() bool {
	resp, err := p.client.Get(defaultCheckURI, nil)
	if err != nil {
		return false
	}
	p.usable = false
	// body := resp.String()
	// p.entry.Println("get", defaultCheckURI, resp.Header, len(body))
	// fmt.Printf("is HTTP2: %v (%s)\n\n", resp.RawResponse.ProtoAtLeast(2, 0), resp.RawResponse.Proto)
	if resp.StatusCode == http.StatusOK {
		p.usable = true
		return true
	}
	return false
}

func (p *proxy) Delay() (bool, float64) {
	if !p.check() {
		return false, 0
	}
	// 开始计算平均值
	var (
		s   time.Time
		sum = 0.0
		n   = 5.0
	)
	for i := 0.0; i < n; i++ {
		s = time.Now()
		if !p.check() {
			return false, 0
		}
		sum += time.Since(s).Seconds()
		time.Sleep(time.Millisecond * 50)
	}
	return true, sum / n
}

func (p *proxy) Post(url, ua string, params map[string]string) (int, []byte) {
	ro := &grequests.RequestOptions{
		Params:    params,
		UserAgent: ua,
	}
	resp, err := p.client.Post(url, ro)
	if err != nil {
		p.entry.WithFields(logrus.Fields{
			"url":    url,
			"params": params,
		}).WithError(err)
	}
	return resp.StatusCode, resp.Bytes()
}

func (p *proxy) Close() {
	p.client.CloseIdleConnections()
}

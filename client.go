package thor

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"crypto/tls"

	"github.com/Sirupsen/logrus"
	"github.com/luw2007/thor/utils"
)

type direct struct {
	client *http.Client
	entry  *logrus.Entry
}

const (
	defaultCheckURI = "https://baidu.com"
)

func NewDirect(addr string, proxy *url.URL, entry *logrus.Entry) *direct {
	tr := &http.Transport{
		DialContext:           utils.DialContext(1*time.Second, 1*time.Second, addr),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:          100,
		IdleConnTimeout:       50 * time.Second,
		TLSHandshakeTimeout:   1 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
		DisableCompression:    true,
		MaxIdleConnsPerHost:   100,
	}
	if proxy.Host != "" {
		tr.Proxy = http.ProxyURL(proxy)
	}
	return &direct{
		client: &http.Client{
			Transport: tr,
			Timeout:   time.Second * 2,
		},
		entry: entry.WithFields(logrus.Fields{
			"addr":  addr,
			"Proxy": proxy.String(),
			"type":  "direct",
		}),
	}
}

func (p *direct) Post(uri string, ua string, params map[string]string) (int, []byte) {
	v := &url.Values{}
	for key, item := range params {
		v.Add(key, item)
	}
	p.entry.WithFields(logrus.Fields{
		"uri":    uri,
		"params": params,
	}).Debug("post")
	req, err := http.NewRequest("POST", uri, strings.NewReader(v.Encode()))
	if err != nil {
		p.entry.WithError(err).WithField("uri", uri).Error("request error")
		return 404, nil
	}
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := p.client.Do(req)
	if err != nil {
		p.entry.WithError(err).WithFields(logrus.Fields{
			"uri": uri,
		}).Error("http error")
		return 503, nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.entry.WithError(err).WithFields(logrus.Fields{
			"uri":  uri,
			"data": data,
		}).Error("http error")
		return resp.StatusCode, nil
	}
	return resp.StatusCode, data
}
func (p *direct) check() bool {
	resp, err := p.client.Get(defaultCheckURI)
	if err != nil {
		return false
	}
	// body := resp.String()
	// fmt.Println("get", defaultCheckURI, resp.Header, len(body))
	// fmt.Printf("is HTTP2: %v (%s)\n\n", resp.RawResponse.ProtoAtLeast(2, 0), resp.RawResponse.Proto)
	if resp.StatusCode != http.StatusOK {
		return true
	}
	return true
}
func (p *direct) Delay() (bool, float64) {
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

func (p *direct) Close() {
	return
}

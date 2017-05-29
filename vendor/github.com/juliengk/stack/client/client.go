package client

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Config struct {
	//        UserAgent       string
	Scheme  string
	Host    string
	Port    int
	Path    string
	Timeout int
}

type BasicAuth struct {
	Username string
	Password string
}

type Request struct {
	Url       string
	Headers   map[string]string
	BasicAuth BasicAuth
	Values    map[string][]string
	Timeout   int
}

type Result struct {
	Response *http.Response
	Body     []byte
	Error    error
}

var EmptyHeader = http.Header{}

var methods = []string{
	"OPTIONS",
	"GET",
	"HEAD",
	"POST",
	"PUT",
	"DELETE",
}

func New(c *Config) (Request, error) {
	txtUrl := buildUrl(c)

	return Request{
		Url:     txtUrl,
		Timeout: c.Timeout,
	}, nil
}

func (r *Request) HeaderAdd(name, value string) {
	if len(r.Headers) == 0 {
		r.Headers = make(map[string]string)
	}

	if _, ok := r.Headers[name]; !ok {
		r.Headers[name] = value
	}
}

func (r *Request) SetBasicAuth(username, password string) {
	ba := BasicAuth{
		Username: username,
		Password: password,
	}

	r.BasicAuth = ba
}

func (r *Request) ValueAdd(name, value string) {
	if len(r.Values) == 0 {
		r.Values = make(map[string][]string)
	}

	if _, ok := r.Values[name]; ok {
		r.Values[name] = append(r.Values[name], value)
	} else {
		r.Values[name] = []string{value}
	}
}

func (r *Request) Do(method string, data io.Reader) Result {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	timeout := time.Second * time.Duration(r.Timeout)

	clnt := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	req, err := http.NewRequest(method, r.Url, data)
	if err != nil {
		return Result{Error: err}
	}

	if len(r.Headers) > 0 {
		for key, value := range r.Headers {
			req.Header.Add(key, value)
		}
	}

	if r.BasicAuth != (BasicAuth{}) {
		req.SetBasicAuth(r.BasicAuth.Username, r.BasicAuth.Password)
	}

	if len(r.Values) > 0 {
		q := req.URL.Query()

		for k, values := range r.Values {
			for _, v := range values {
				q.Add(k, v)
			}
		}
	}

	resp, err := clnt.Do(req)
	if err != nil {
		r := Result{Error: err}
		if resp != nil {
			r.Response = resp
		}

		return r
	}

	if method == "OPTIONS" || method == "HEAD" {
		return Result{Response: resp}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{Response: resp, Error: err}
	}

	return Result{Response: resp, Body: body}
}

func (r *Request) Options() Result {
	return r.Do("OPTIONS", nil)
}

func (r *Request) Get() Result {
	return r.Do("GET", nil)
}

func (r *Request) Head() Result {
	return r.Do("HEAD", nil)
}

func (r *Request) Post(data io.Reader) Result {
	return r.Do("POST", data)
}

/*
func (r *Request) Put(data io.Reader) Result {
	return r.Do("PUT", data)
}

func (r *Request) Delete() Result {
	return r.Do("DELETE", nil)
}
*/

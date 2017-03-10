package gopherillamail

import (
    "fmt"
    "net/http"
    "time"
)

const (
    VERSION = "0.1"
    STANDARD_IP = "127.0.0.1"
)

var (
    STANDARD_USERAGENT = fmt.Sprintf("GopherillaMail v%s", VERSION)
)

type Mail struct {
    guid string
    subject string
    sender string
    time time.Time
    read bool
    excerpt string
    body string
}

type GuerrillamailClient struct {
    email string
    ip string
    httpclient *http.Client
    SUBSCR bool
    PHPSESSID string
}

func NewGuerrillamail(email string) (*GuerrillamailClient, error) {
    clnt := &GuerrillamailClient{}

    clnt.ip = STANDARD_IP //TODO: maybe make this variable?
    clnt.email = email

    hc := &http.Client{}
    hc.Transport = &transport{http.DefaultTransport, STANDARD_USERAGENT}
    clnt.httpclient = hc
    return clnt, nil
}

func AnonymousGuerrillamail(user_agent, ip string) (*GuerrillamailClient, error) {
    clnt := &GuerrillamailClient{}

    clnt.ip = STANDARD_IP //TODO: maybe make this variable?

    hc := &http.Client{}
    hc.Transport = &transport{http.DefaultTransport, STANDARD_USERAGENT}
    clnt.httpclient = hc
}

func (c *GuerrillamailClient) SetUserAgent(user_agent string){
    hc := &http.Client{}
    hc.Transport = &transport{http.DefaultTransport, user_agent}
    c.httpclient = hc
}

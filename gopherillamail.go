package gopherillamail

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	VERSION     = "0.1"
	STANDARD_IP = "127.0.0.1"
)

var (
	STANDARD_USERAGENT = fmt.Sprintf("GopherillaMail v%s", VERSION)
)

type Mail struct {
	guid    string
	subject string
	sender  string
	time    time.Time
	read    bool
	excerpt string
	body    string
}

type Inbox struct {
	User_agent      string
	Ip              string
	email_timestamp uint
	httpclient      *http.Client
}

// Returns an Inbox without any email
func blankInbox() (*Inbox, error) {
	cjar, _ := cookiejar.New(nil)
	hcl := &http.Client{
		Jar: cjar,
	}
	inb := &Inbox{
		httpclient: hcl
	}

	inb.ip = STANDARD_IP
	inb.user_agent = STANDARD_USERAGENT

	return inb, nil
}

// Returns an Inbox with a custom email
func NewInbox(email string) (*Inbox, error) {
	inb, err := blankInbox()
	if err != nil {
		return err
	}
	err = inb.setEmail(email)
	if err != nil {
		return err
	}
	return &inb
}

// Returns an Inbox with a random email
func AnonymousInbox() (*Inbox, error) {
	inb, err := blankInbox()
	if err != nil {
		return err
	}
	err = inb.randomEmail()
	if err != nil {
		return err
	}
	return &inb
}

// Does a function call to Guerrillamail's api
func (c *Inbox) doRequest(function_name string, args map[string]string) error {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"http://api.guerrillamail.com/ajax.php?f=%s&ip=%s&agent=%s",
			function_name,
			c.ip,
			c.user_agent,
		),
		nil,
	)
	if err != nil {
		return err
	}

	// Build the querystring from the arguments
	q := req.URL.Query()
	for key, val := range args {
		q.Add(key, val)
	}
	// Set the querystring
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
}

// Sets the user agent
func (c *Inbox) SetUserAgent(user_agent string) {
	c.user_agent = user_agent
}

// Sets the email address
func (c *Inbox) setEmail(user_agent string) error {
	err := c.doRequest()
	if err != nil {
		return err
	}
}

func (c *Inbox) GetEmail(user_agent string) error {
	err := c.doRequest(
		"get_email_address",
		map[string]string{
			"lang": "en",
			""
		},
	)
	if err != nil {
		return err
	}
}

// Asks Guerrillamail for a random email address
func (c *Inbox) randomEmail(user_agent string) error {
	err := c.doRequest()
	if err != nil {
		return err
	}
}

// Sets the client IP
func (c *Inbox) SetIP(ip string) {
	c.ip = ip
}

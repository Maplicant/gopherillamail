package gopherillamail

import (
	"fmt"
	"net/http"
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
	email           string
	user_agent      string
	ip              string
	email_timestamp uint
	httpclient      *http.Client
	SUBSCR          bool
	PHPSESSID       string
}

// Returns an Inbox without any email
func blankInbox() (*Inbox, error) {
	inb := &Inbox{}

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
	inb.setEmail(email)
	return &inb
}

// Returns an Inbox with a random email
func AnonymousInbox() (*Inbox, error) {
	inb, err := blankInbox()
	if err != nil {
		return err
	}
	inb.randomEmail()
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

	// Build the querystring
	q := req.URL.Query()
	for key, val := range args {
		q.Add(key, val)
	}
	// q.Add("api_key", "key_from_environment_or_flag")
	// q.Add("another_thing", "foo & bar")
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
	c.doRequest()
}

// Asks Guerrillamail for a random email address
func (c *Inbox) randomEmail(user_agent string) error {
	c.doRequest()
}

// Sets the IP
func (c *Inbox) SetIP(ip string) error {
	c.ip = ip
}

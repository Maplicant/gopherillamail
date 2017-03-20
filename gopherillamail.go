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

// Mail is an e-mail from GuerrillaMail
type Mail struct {
	guid    string
	subject string
	sender  string
	time    time.Time
	read    bool
	excerpt string
	body    string
}

// Inbox is a struct that allows you to retrieve e-mails from GuerrillaMail
type Inbox struct {
	UserAgent      string
	IP             string
	sid_token      string
	Email          string
	EmailList      []Mail
	emailTimestamp uint
	httpclient     *http.Client
}

// Returns an Inbox without any email
func blankInbox() (*Inbox, error) {
	cjar, _ := cookiejar.New(nil)
	hcl := &http.Client{
		Jar: cjar,
	}
	inb := &Inbox{
		httpclient: hcl,
		IP:         STANDARD_IP,
		UserAgent:  STANDARD_USERAGENT,
	}

	return inb, nil
}

// NewInbox returns an Inbox with a custom email
func NewInbox(email string) (*Inbox, error) {
	inb, err := blankInbox()
	if err != nil {
		return nil, err
	}
	err = inb.setEmail(email)
	if err != nil {
		return nil, err
	}
	return inb, nil
}

// AnonymousInbox returns an Inbox with a random email
func AnonymousInbox() (*Inbox, error) {
	inb, err := blankInbox()
	if err != nil {
		return inb, fmt.Errorf("could not create blank inbox: %v", err)
	}

	err = inb.randomEmail()
	if err != nil {
		return inb, fmt.Errorf("could not create random email: %v", err)
	}

	err = inb.getEmail() // You have to call this at least once to set the sid_token and the Email in the struct
	if err != nil {
		return inb, fmt.Errorf("could not get initial email list: %v", err)
	}

	return inb, nil
}

// Does a function call to Guerrillamail's api
func (c *Inbox) doRequest(functionName string, args map[string]string) error {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"http://api.guerrillamail.com/ajax.php?f=%s&ip=%s&agent=%s",
			functionName,
			c.IP,
			c.UserAgent,
		),
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not build request to GuerrillaMail: %v", err)
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
		return fmt.Errorf("could not do request to GuerrillaMail: %v", err)
	}

	defer resp.Body.Close()
	return nil
}

// SetUserAgent sets the user agent
func (c *Inbox) SetUserAgent(userAgent string) {
	c.UserAgent = userAgent
}

// Sets the email address
func (c *Inbox) setEmail(userAgent string) error {
	err := c.doRequest()
	if err != nil {
		return fmt.Errorf("could not set email address: %v", err)
	}
	return nil
}

// getEmail Gets the e-mail address of your inbox and sets the sid_token
func (c *Inbox) getEmail() error {
	err := c.doRequest(
		"get_email_address",
		map[string]string{
			"lang": "en",
			// "",
		},
	)
	if err != nil {
		return fmt.Errorf("could not get email address: %v", err)
	}
	return nil
}

// getEmailList does the initial call to initialize the EmailList. Shouldn't be used to check for new e-mails.
func (c *Inbox) getEmailList() error {
	err := c.doRequest()
	if err != nil {
		return fmt.Errorf("could not get initial emails: %v", err)
	}
	return nil
}

// Asks Guerrillamail for a random email address
func (c *Inbox) randomEmail() error {
	err := c.doRequest()
	if err != nil {
		return fmt.Errorf("could not generate random email: %v", err)
	}
	return nil
}

// SetIP sets the client IP
func (c *Inbox) SetIP(IP string) {
	c.IP = IP
}

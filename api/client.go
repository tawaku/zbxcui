package api

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync/atomic"
)

type Client struct {
	// Members
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
	token      *string
	id         uint32
}

func MakeClient(rawurl, usr, pwd string, logger *log.Logger) (*Client, error) {
	// Initialize
	c := new(Client)

	// Set Logger
	if logger != nil {
		c.Logger = logger
	} else {
		c.Logger = log.New(os.Stdout, "[ZBXCUI]", log.LstdFlags|log.LUTC)
	}

	// Set url
	if u, err := url.Parse(rawurl); err != nil {
		return nil, errors.Wrap(err, "Failed to parse URL.")
	} else {
		c.URL = u
	}

	// Set HTTP client
	c.HTTPClient = new(http.Client)

	// Login
	if r := c.userLogin(usr, pwd, false); r.Error != nil {
		return nil, errors.Wrap(r.Error.Error(), "Login failed.")
	} else {
		if r.Result == nil {
			return nil, errors.New("Login failed: authentication token is not returned.")
		} else {
			// Assert interface to string(token in Client struct)
			token := r.Result.(string)
			c.token = &token
		}
	}

	return c, nil
}

func (self *Client) Close() {
	self.userLogout()
}

func (self *Client) request(send, recv interface{}) {
	// Request
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(send)
	// Send data log
	self.Logger.Printf("Send: %s", b.String())

	req, err := http.NewRequest("POST", self.URL.String(), b)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json-rpc")
	req.Close = true
	res, err := self.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Receive data log
	body, _ := ioutil.ReadAll(res.Body)
	self.Logger.Printf("Recv: %s", string(body))
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	json.NewDecoder(res.Body).Decode(recv)

	// Error data log
	if recv.(*Response).Error != nil {
		self.Logger.Printf("Error: %+v", recv.(*Response).Error)
	}
}

func (self *Client) makeRequest(m string, p interface{}) *Request {
	atomic.AddUint32(&self.id, 1)
	r := &Request{
		Jsonrpc: "2.0",
		Method:  m,
		Params:  p,
		Auth:    self.token,
		Id:      self.id,
	}
	return r
}

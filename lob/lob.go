package lob

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// LibraryVersion represents this library version
	LibraryVersion = "0.1"

	// BaseURL represents Lob API base URL
	BaseURL = "https://api.lob.com/v1/"

	// APIVersion ...
	APIVersion = "2017-09-08"
)

// Client manages communication with the Lob API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Application API key
	APIKey string

	// Services
	Address *AddressService
}

// NewClient returns a new Lob API client. if a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
		APIKey:  apiKey,
		// UserAgent: UserAgent,
	}
	c.Address = &AddressService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	// req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/json")
	}

	req.SetBasicAuth(c.APIKey, "")
	req.Header.Add("Lob-Version", APIVersion)
	req.Header.Add("Accept", "application/json")

	// req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// spew.Dump(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// ErrorResponse ...
type ErrorResponse struct {
	ErrorType struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %v",
		r.ErrorType.StatusCode, r.ErrorType.Message)
}

// CheckResponse checks the API response for error, and returns it
// if present. A response is considered an error if it has non StatusOK
// code.
func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusOK {
		return nil
	}

	resp := new(ErrorResponse)

	if r.StatusCode == http.StatusInternalServerError || r.StatusCode == http.StatusNotFound {
		return resp
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}

	return resp
}

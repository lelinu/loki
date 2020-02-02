package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Client struct{
	httpClient *http.Client
	baseUrl url.URL
	enableMock bool
	enableDebug bool
	mocks map[string]*mock
}

func NewClient(httpClient *http.Client, baseUrl url.URL) *Client{

	if httpClient == nil {
		httpClient = getHTTPClient()
	}

	client := &Client{
		httpClient: httpClient,
	}

	client.SetBaseUrl(baseUrl)
	client.ResetMocks()

	return client
}

func (c *Client) SetMock(enableMock bool){
	c.enableMock = enableMock
}

func (c *Client) SetDebug(enabledDebug bool){
	c.enableDebug = enabledDebug
}

func (c *Client) SetBaseUrl(url url.URL){
	c.baseUrl = url
}

func (c *Client) AddMock(mock *mock){
	c.mocks[getMockId(mock.httpMethod, mock.url, mock.action)] = mock
}

func (c *Client) ResetMocks(){
	c.mocks = make(map[string]*mock)
}

func (c *Client) Post(action string, body interface{}) (*http.Response, error) {

	if c.enableMock{
		
		mock := c.mocks[getMockId(http.MethodPost, c.baseUrl.String(), action)]
		if mock == nil{
			return nil, errors.New("soap client : no mock up found for given action")
		}
		return mock.response, mock.err
	}

	soapReq := newSoapRequest(body)
	payload, err := xml.MarshalIndent(soapReq, "", "  ")

	if err != nil{
		log.Printf("soap client : an error had occurred during marhsal indent %v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.baseUrl.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/xml, multipart/related")
	req.Header.Set("SOAPAction", "http://tempuri.org/" + action)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	if c.enableDebug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("soap client : an error had occurred during post %v", err)
		return nil, err
	}

	if c.enableDebug == true {
		dump, _ := httputil.DumpResponse(response, true)
		log.Println(string(dump))
	}

	return response, nil
}

func getMockId(httpMethod string, url string, action string) string {
	return fmt.Sprintf("%s_%s_%s", httpMethod, url, action)
}

func getHTTPClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
		DisableCompression:  true,
	}

	client := &http.Client{Transport: tr}
	return client
}


package sendcloud_email_go

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const APIBase string = "https://api.sendcloud.net"

type  SendCloud struct {
	apiUser    string
	apiKey     string
	apiBase	   string
	client  *http.Client
}

func (client *SendCloud)validateConfig() error {
	if len(client.apiBase) == 0 {
		client.apiBase = APIBase
	}
	switch {
	case len(client.apiKey) == 0:
		return errors.New("apiKey cannot be empty")
	case len(client.apiKey) == 0:
		return errors.New("apiKey cannot be empty")
	}
	return nil
}

type Response struct {
	*http.Response
}

func NewSendCloud(apiUser string, apiKey string) (*SendCloud, error) {
	switch {
	case len(apiUser) == 0:
		return nil,errors.New("smsUser cannot be empty")
	case len(apiKey) == 0:
		return nil,errors.New("smsKey cannot be empty")
	}

	sc := &SendCloud{
		apiUser: apiUser,
		apiKey:  apiKey,
		apiBase: APIBase,
		client:  http.DefaultClient,
	}
	return sc, nil
}

func (sc *SendCloud) request(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := sc.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
		_, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return response, readErr
		}
	}

	return response, err
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusOK {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}


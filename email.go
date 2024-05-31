package sendcloud_email_go

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	sendCommonPath  =  "/apiv2/mail/send"
	sendTemplatePath  =  "/apiv2/mail/sendtemplate"
)

func (client *SendCloud) SendCommonEmail(ctx context.Context, args *SendEmailArgs) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil,fmt.Errorf("failed to send message: %w", err)
	}
	if err := validateSendCommonEmail(args); err != nil {
		return nil,fmt.Errorf("failed to send message: %w", err)
	}
	if args.Attachments == nil {

	}else {
		payload, err := args.Marshal()
		if err!= nil {
            return nil,fmt.Errorf("failed to send message: %w", err)
        }
		req, err := http.NewRequest("POST", client.apiBase + sendCommonPath, payload)
		if err!= nil {
            return nil,fmt.Errorf("failed to send message: %w", err)
        }
        req.Header.Set("Content-Type", "multipart/form-data")
		resp, err := client.request(ctx, req, &SendEmailResult{})
		if err!= nil {
			return nil,fmt.Errorf("failed to send message: %w", err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil,fmt.Errorf("failed to read response body: %w", err)
		}
		var responseData SendEmailResult
		if err := json.Unmarshal(body, &responseData); err != nil {
			return nil,fmt.Errorf("failed to unmarshal response body: %w", err)
		}
		result := &responseData
		if !result.Result {
			return result,fmt.Errorf("Email sending failed: %s", result.Message)
		}
		return result,nil
	}

}

func validateSendCommonEmail(args *SendEmailArgs) error {
	switch {
	case len(args.From) == 0:
		return errors.New("from cannot be empty")
	case len(args.Subject) == 0:
		return errors.New("subject cannot be empty")
	}
	return nil
}
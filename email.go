package sendcloud_email_go

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
)

const (
	sendCommonPath   = "/apiv2/mail/send"
	sendTemplatePath = "/apiv2/mail/sendtemplate"
)

func (client *SendCloud) SendCommonEmail(ctx context.Context, args *SendEmailArgs) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	if err := validateSendCommonEmail(args); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	var req *http.Request
	var err error

	if args.Attachments == nil {
		params, err := client.PrepareSendCommonEmailParams(args)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		formDataEncoded := params.Encode()
		req, err = http.NewRequest("POST", client.apiBase+sendCommonPath, bytes.NewBufferString(formDataEncoded))
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		payload, err := client.MarshalSendEmailArgs(args)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req, err = http.NewRequest("POST", client.apiBase+sendCommonPath, payload)
		if err != nil {
			return nil, fmt.Errorf("failed to send email: %w", err)
		}
		req.Header.Set("Content-Type", "multipart/form-data")
	}
	responseData := new(SendEmailResult)
	err = client.Request(ctx, req, &responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
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

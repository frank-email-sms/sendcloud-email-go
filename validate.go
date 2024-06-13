package sendcloud_email_go

import (
	"errors"
	"strings"
)

const MAX_RECEIVERS = 100
const MAX_MAILLIST = 5
const MAX_CONTENT_SIZE = 10 * 1024 * 1024

func (args *SendEmailTemplateArgs)validateSendEmailTemplate() error {
	if err := args.validateEmailCommonFields(); err != nil {
		return err
	}
	if args.TemplateInvokeName == "" {
		return errors.New("templateInvokeName cannot be empty")
	}
	return nil
}

func (args *SendEmailArgs)validateSendCommonEmail() error {
	if err := args.validateEmailCommonFields(); err != nil {
		return err
	}
	switch {
	case len(args.Html) == 0 && len(args.Plain) == 0:
		return errors.New("html or plain cannot be empty")
	case len(args.Html) > 0 && len(args.Html) > MAX_CONTENT_SIZE:
		return errors.New("html content is too long")
	case len(args.Plain) > 0 && len(args.Plain) > MAX_CONTENT_SIZE:
		return errors.New("plain Content is too long")
	}
	return nil
}

func (e *EmailCommonFields)validateEmailCommonFields() error {
	switch {
	case len(e.From) == 0:
		return errors.New("from cannot be empty")
	case len(e.Subject) == 0:
		return errors.New("subject cannot be empty")
	case e.To == "":
		return errors.New("to cannot be empty")
	}
	if e.UseAddressList {
		to := strings.Split(e.To, ";")
		if len(to) > MAX_MAILLIST {
			return errors.New("address list exceeds limit")
		}
	}else {
		to := strings.Split(e.To, ";")
		cc := strings.Split(e.CC, ";")
		bcc := strings.Split(e.BCC, ";")
		receivers := len(to)
		receivers += len(cc)
		receivers += len(bcc)
		// Check if the total number of receivers exceeds the maximum allowed
		if receivers > MAX_RECEIVERS {
			return errors.New("the total number of receivers exceeds the maximum allowed")
		}
	}
	return nil
}
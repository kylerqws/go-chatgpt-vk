package registry

import ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"

type LoggerRegistry interface {
	Register(name string, log ctrlog.Logger)
	Logger(name string) (ctrlog.Logger, error)
}

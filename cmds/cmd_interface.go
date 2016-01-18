package cmds

import (
	"github.com/krishicks/i18n4go/common"
)

type CommandInterface interface {
	common.PrinterInterface
	Options() common.Options
	Run() error
}

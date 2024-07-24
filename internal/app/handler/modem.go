package handler

import (
	"fmt"
	"strings"

	"github.com/damonto/telegram-sms/internal/pkg/modem"
	"github.com/damonto/telegram-sms/internal/pkg/util"
	"gopkg.in/telebot.v3"
)

func HandleModemsCommand(c telebot.Context) error {
	modems := modem.GetManager().GetModems()
	if len(modems) == 0 {
		return c.Send("No modems found")
	}

	template := `
Manufaturer: %s
Model: %s
Revision: %s
IMEI: %s
Signal: %d
Network: %s
ICCID: %s
EID: %s
`
	var message string
	for _, m := range modems {
		manufacturer, _ := m.GetManufacturer()
		model, _ := m.GetModel()
		revision, _ := m.GetRevision()
		imei, _ := m.GetImei()
		signal, _ := m.GetSignalQuality()
		network, _ := m.GetOperatorName()
		ICCID, _ := m.GetICCID()

		message += fmt.Sprintf(
			template,
			manufacturer,
			model,
			revision,
			fmt.Sprintf("`%s`", imei),
			signal,
			network,
			fmt.Sprintf("`%s`", ICCID),
			fmt.Sprintf("`%s`", m.Eid))
	}
	return c.Send(util.EscapeText(strings.TrimRight(message, "\n")), &telebot.SendOptions{ParseMode: telebot.ModeMarkdownV2})
}

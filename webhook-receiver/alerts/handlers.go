package alerts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"webhook-receiver/api"
	"webhook-receiver/telegram"
)

const (
	Firing   = "firing"
	Resolved = "resolved"

	Fire         = "%F0%9F%94%A5"
	ThumbsUp     = "%F0%9F%91%8D"
	NewLine      = "%0A"
	MoonFront    = "%F0%9F%8C%95"
	MoonBack     = "%F0%9F%8C%91"
	MoonHalf     = "%F0%9F%8C%97"
	Server       = "%E2%9A%99%EF%B8%8F"
	Service      = "%F0%9F%9A%80"
	Down         = "%E2%98%A0%EF%B8%8F"
	LowRAM       = "%F0%9F%92%AD"
	LowDisk      = "%F0%9F%92%BF"
	NotAvailable = "%F0%9F%98%B5"
	WarningLevel = "%E2%9A%A0%EF%B8%8F"
	ErrorLevel   = "%E2%9D%8C"
	ServiceError = "%E2%9B%94%EF%B8%8F"
	LowCPU       = "%F0%9F%92%94"
	LicenseLimit = "💸"

	Critical = "critical"
	Warning  = "warning"
)

func codeText(text string) string {
	return "<code>" + text + "</code>"
}

func boldText(text string) string {
	return "<b>" + text + "</b>"
}

func HandleAlert(response http.ResponseWriter, request *http.Request) {
	var requestBody AlertsRequestBody
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		log.Println(err)
		api.HandleError(response, http.StatusBadRequest, err.Error())
		return
	}

	for _, alert := range requestBody.Alerts {
		level := alert.GetLevel()
		department := alert.GetTeam()
		target := alert.GetTarget()
		instance := alert.GetInstance()
		tp := alert.GetType()
		summary := alert.GetSummary()

		message := codeText(fmt.Sprintf(
			"Level:      | %s"+
				"Department: | %s"+
				"Target:     | %s"+
				"Location:   | %s"+
				"Type:       | %s",
			level+NewLine,
			department+NewLine,
			target+NewLine,
			instance+NewLine,
			tp+NewLine,
		))

		message = message + NewLine + boldText("SUMMARY: ") + summary + NewLine

		// Send
		bot := telegram.GetBot()
		channelId := chi.URLParam(request, "channelId")

		if err := bot.SendMessage(channelId, message); err != nil {
			log.Println(err)
			api.HandleError(response, http.StatusBadRequest, err.Error())
			return
		}
	}

	api.HandleSuccess(response)
}

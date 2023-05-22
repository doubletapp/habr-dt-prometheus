package alerts

import "strings"

type AlertsRequestBody struct {
	// Version string `json:"version"`

	// ключ, определяющий группу алертов (например для дедубликации)
	// GroupKey string `json:"groupKey"`

	// сколько алертов были отброшены из-за параметра "max_alerts"
	// TruncatedAlerts int `json:"truncatedAlerts"`

	// "resolved" или "firing"
	// Status            string            `json:"status"`
	// Receiver          string            `json:"receiver"`
	// GroupLabels       map[string]string `json:"groupLabels"`
	// CommonLabels      map[string]string `json:"commonLabels"`
	// CommonAnnotations map[string]string `json:"commonAnnotations"`

	// обратная ссылка на Alertmanager
	// ExternalURL string  `json:"externalURL"`
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	// "resolved" или "firing"
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`

	// StartsAt string `json:"startsAt"`
	// EndsAt string `json:"endsAt"`

	// сущность, из-за которой алерт был создан
	// GeneratorURL string `json:"generatorURL"`

	// идентификатор алерта
	// Fingerprint string `json:"fingerprint"`
}

func (alert *Alert) GetLevel() string {
	level := "Unknown"

	if t, exists := alert.Labels["alertname"]; exists {
		if t == "ServiceError" {
			return ErrorLevel + "[ERROR]"
		}
	}

	if severity, exists := alert.Labels["severity"]; exists {
		switch severity {
		case Warning:
			level = WarningLevel + " [WARNING]"
		case Critical:
			level = Fire + " [CRITICAL]"
		}
	}
	if alert.Status == Resolved {
		level = ThumbsUp + " [RESOLVED]"
	}

	return level
}

func (alert *Alert) GetTeam() string {
	department := MoonHalf + " [BACK+FRONT]"

	if team, exists := alert.Labels["team"]; exists {
		switch strings.ToUpper(team) {
		case "FRONTEND":
			department = MoonFront + " [FRONT]"
		case "BACKEND":
			department = MoonBack + " [BACK]"
		}
	}

	return department
}

func (alert *Alert) GetSummary() string {
	summary := "Summary wasn't provided"
	if text, exists := alert.Annotations["summary"]; exists {
		summary = text
	}

	return summary
}

func (alert *Alert) GetTarget() string {
	target := "Target wasn't provided"
	if t, exists := alert.Labels["instance_type"]; exists {
		switch strings.ToUpper(t) {
		case "SERVICE":
			target = Service + " [SERVICE]"
		case "NODE":
			target = Server + " [SERVER]"
		default:

		}
	}

	return target
}

func (alert *Alert) GetType() string {
	tp := "Type unknown"
	if t, exists := alert.Labels["alertname"]; exists {
		switch t {
		case "ServiceDown":
			tp = Down + " [DOWN]"
		case "NodeDown":
			tp = Down + " [DOWN]"
		case "NodeLowDiskSpace":
			tp = LowDisk + " [LOW DISK]"
		case "NodeLowMemory":
			tp = LowRAM + " [LOW RAM]"
		case "ServiceNotAvailable":
			tp = NotAvailable + " [SERVICE NOT AVAILABLE]"
		case "ServiceError":
			tp = ServiceError + " [5xx THROWN]"
		case "NodeHighCPUUsage":
			tp = LowCPU + " [LOW CPU]"
		}
	}

	return tp
}

func (alert *Alert) GetInstance() string {
	instance := "Instance wasn't provided"
	if inst, exists := alert.Labels["instance"]; exists {
		instance = inst
	}

	return instance
}

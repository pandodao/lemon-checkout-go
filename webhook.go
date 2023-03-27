package lemon

type (
	WebhookAttributes struct {
		DiscountTotalUSD int    `json:"discount_total_usd"`
		TotalUSD         int    `json:"total_usd"`
		Status           string `json:"status"`
	}

	WebhookPayload struct {
		Meta struct {
			TestMode   bool   `json:"test_mode"`
			EventName  string `json:"event_name"`
			CustomData any    `json:"custom_data"`
		} `json:"meta"`

		Data struct {
			Type       string            `json:"type"`
			ID         string            `json:"id"`
			Attributes WebhookAttributes `json:"attributes"`
		} `json:"data"`
	}
)

package types


type WebhookOrder struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Token string `json:"token"`
	SubtotalPrice string `json:"subtotal_price"`
	TotalPrice string `json:"total_price"`
	Currency string `json:"currency"`
	ContactEmail string `json:"contact_email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	OrderNumber int `json:"order_number"`
	CustomerLocale string `json:"customer_locale"`
	OrderStatusUrl string `json:"order_status_url"`
}
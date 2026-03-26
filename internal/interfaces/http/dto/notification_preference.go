package dto

type NotificationPreferenceResponse struct {
	InApp bool `json:"in_app"`
	Push  bool `json:"push"`
	Email bool `json:"email"`
	// Future extensibility: add more fields as needed (e.g., SMS, WhatsApp)
}

type SetNotificationPreferenceRequest struct {
	InApp bool `json:"in_app"`
	Push  bool `json:"push"`
	Email bool `json:"email"`
	// Future extensibility: add more fields as needed (e.g., SMS, WhatsApp)
}

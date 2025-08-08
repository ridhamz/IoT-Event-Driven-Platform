package domain

type Device struct {
	ID        int64  `json:"id",omitempty`
	Name      string `json:"first_name"`
	UserID    int64  `json:"user_id", omitempty`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type DeviceEvent struct {
	DeviceID  int64   `json:"device_id,omitempty"`
	UserID    int64   `json:"user_id,omitempty"`    // associated user ID (optional)
	Name      string  `json:"name,omitempty"`       // device name or event name
	Email     string  `json:"email,omitempty"`      // user email (optional)
	EventType string  `json:"event_type"`           // e.g., "temperature", "motion_detected"
	Timestamp string  `json:"timestamp"`            // ISO8601 timestamp of event
	Data      float64 `json:"data,omitempty"`       // event-specific data, e.g., temperature value
	CreatedAt string  `json:"created_at,omitempty"` // when record was created in your system
}

type APIKey struct {
	ID        int64  `json:"id"`
	APIKey    string `json:"api_key"`
	CreatedAt string `json:"created_at"`
}

type DeviceWithAPIKeys struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	UserID    int64    `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	APIKeys   []APIKey `json:"api_keys"`
}

package models

// Action describes an action that the ScanBadge client performs, based on the result of the specified condition.
type Action struct {
	ID          int64      `db:"action_id" json:"action_id"`
	UserID      int64      `db:"user_id" json:"user_id"`
	DeviceID    int64      `db:"device_id" json:"device_id"`
	Name        string     `db:"action_name" json:"action_name" form:"name"`
	Description string     `db:"action_description" json:"action_description" form:"description"`
	Value       string     `db:"action_value" json:"action_value" form:"value"`
	Device      Device     `db:"-" json:"device"`
	Type        ActionType `db:"-" json:"action_type"`
}

// ActionType describes an action type. The action type can be used to determine what driver to use when the action is performed.
type ActionType struct {
	ID          int64  `db:"action_type_id" json:"id"`
	Name        string `db:"action_type_name" json:"name" form:"name"`
	Description string `db:"action_type_sgdescription" json:"description" form:"description"`
}

// Condition describes a condition. When a condition evaluates to TRUE, the related action should be performed.
type Condition struct {
	ID          int64         `db:"condition_id" json:"id"`
	UserID      int64         `db:"user_id" json:"user_id"`
	DeviceID    int64         `db:"device_id" json:"device_id"`
	Name        string        `db:"condition_name" json:"name" form:"name"`
	Description string        `db:"condition_description" json:"description" form:"description"`
	Value       string        `db:"condition_value" json:"value" form:"value"`
	Action      Action        `db:"-" json:"action"`
	Device      Device        `db:"-" json:"device"`
	Type        ConditionType `db:"-" json:"condition_type"`
}

// ConditionType describes a condition type.
type ConditionType struct {
	ID          int64  `db:"condition_type_id" json:"id"`
	Name        string `db:"condition_type_name" json:"name" form:"name"`
	Description string `db:"condition_type_description" json:"description" form:"description"`
	ExecuteArgs string `db:"condition_type_execute_args" json:"execute_args" form:"execute_args"`
}

type Count struct {
	Endpoint string `json:"endpoint"`
	Count    int64  `json:"count"`
}

// Device describes a device.
// When adding a device, only name, description and key fields can be entered.
type Device struct {
	ID          int64  `db:"device_id" json:"device_id"`
	UserID      int64  `db:"user_id"`
	Name        string `db:"device_name" json:"name" form:"name"`
	Description string `db:"device_description" json:"description" form:"description"`
	Key         string `db:"device_key" json:"key" form:"key"`
	User        User   `db:"-" json:"user"`
}

// Log describes a log entry.
type Log struct {
	ID          int64  `db:"log_id" json:"id"`
	Date        int64  `db:"log_date" json:"date"`
	UserID      int64  `db:"user_id"`
	Type        string `db:"log_type" json:"type" form:"type"`
	Description string `db:"log_message" json:"message" form:"message"`
	Origin      string `db:"log_origin" json:"origin" form:"origin"`
	Object      string `db:"log_object" json:"object" form:"object"`
	User        User   `db:"-" json:"user"`
}

// User describes a user.
type User struct {
	ID        int64  `db:"user_id" json:"id"`
	Username  string `db:"user_username" json:"username" form:"username"`
	Email     string `db:"user_email" json:"email" form:"email"`
	Password  string `db:"user_password" json:"password,omitempty" form:"password"`
	FirstName string `db:"user_first_name" json:"first_name" form:"first_name"`
	LastName  string `db:"user_last_name" json:"last_name" form:"last_name"`
	Roles     Role   `db:"-" json:"roles,omitempty"`
}

// Role describes a user role.
type Role struct {
	ID          int64  `db:"role_id" json:"id"`
	Level       int64  `db:"role_Level" json:"level"`
	Name        string `db:"role_name" json:"name"`
	Description string `db:"role_description" json:"description"`
}

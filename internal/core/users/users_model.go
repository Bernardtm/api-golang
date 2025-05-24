package users

import (
	"time"
)

type UserFund struct {
	UsersFundsUUID   string     `json:"users_funds_uuid" db:"users_funds_uuid"`             // UUID do registro de associação entre usuário e fundo (chave primária)
	UserUUID         string     `json:"user_uuid" db:"user_uuid"`                           // UUID do usuário (chave estrangeira)
	FundUUID         string     `json:"fund_uuid" db:"fund_uuid"`                           // UUID do fundo (chave estrangeira)
	CreationDate     *time.Time `json:"creation_date,omitempty" db:"creation_date"`         // Data de criação (padrão: data atual)
	ModificationDate *time.Time `json:"modification_date,omitempty" db:"modification_date"` // Data de modificação (opcional)
	StatusUUID       string     `json:"status_uuid" db:"status_uuid"`                       // UUID do status (chave estrangeira)
}

type UserNotification struct {
	UserNotificationUUID string     `json:"user_notification_uuid" db:"user_notification_uuid"` // UUID da notificação do usuário (chave primária)
	UserUUID             string     `json:"user_uuid" db:"user_uuid"`                           // UUID do usuário (chave estrangeira)
	NotificationUUID     string     `json:"notification_uuid" db:"notification_uuid"`           // UUID da notificação (chave estrangeira)
	IsRead               *bool      `json:"is_read,omitempty" db:"is_read"`                     // Indica se a notificação foi lida (opcional, padrão: false)
	ReadDate             *time.Time `json:"read_date,omitempty" db:"read_date"`                 // Data de leitura da notificação (opcional)
	IsExecuted           *bool      `json:"is_executed,omitempty" db:"is_executed"`             // Indica se a notificação foi executada (opcional, padrão: false)
	ExecutedDate         *time.Time `json:"executed_date,omitempty" db:"executed_date"`         // Data de execução (opcional)
	ExecutionLog         *string    `json:"execution_log,omitempty" db:"execution_log"`         // Log de execução (opcional)
	CreationDate         *time.Time `json:"creation_date,omitempty" db:"creation_date"`         // Data de criação (padrão: data atual)
	ModificationDate     *time.Time `json:"modification_date,omitempty" db:"modification_date"` // Data de modificação (opcional)
	StatusUUID           string     `json:"status_uuid" db:"status_uuid"`                       // UUID do status (chave estrangeira)
}

type User struct {
	UserUUID         string     `json:"user_uuid" db:"user_uuid"`                           // UUID do usuário (chave primária)
	Username         *string    `json:"username,omitempty" db:"username"`                   // Nome de usuário (único, opcional)
	Email            string     `json:"email" db:"email"`                                   // Email do usuário (único)
	Password         *string    `json:"password,omitempty" db:"password"`                   // Senha do usuário (opcional)
	IsComplete       *bool      `json:"is_complete,omitempty" db:"is_complete"`             // Indica se o perfil está completo (padrão: false, opcional)
	CreationDate     *time.Time `json:"creation_date,omitempty" db:"creation_date"`         // Data de criação (padrão: data atual)
	ModificationDate *time.Time `json:"modification_date,omitempty" db:"modification_date"` // Data de modificação (opcional)
	StatusUUID       string     `json:"status_uuid" db:"status_uuid"`                       // UUID do status (chave estrangeira)
}

type UserProfileUser struct {
	UserProfileUsersUUID string     `json:"user_profile_users_uuid" db:"user_profile_users_uuid"` // UUID do registro de associação entre usuário e perfil de usuário (chave primária)
	UserUUID             *string    `json:"user_uuid,omitempty" db:"user_uuid"`                   // UUID do usuário (chave estrangeira, opcional)
	UserProfileUUID      *string    `json:"user_profile_uuid,omitempty" db:"user_profile_uuid"`   // UUID do perfil de usuário (chave estrangeira, opcional)
	CreationDate         *time.Time `json:"creation_date,omitempty" db:"creation_date"`           // Data de criação (padrão: data atual)
	ModificationDate     *time.Time `json:"modification_date,omitempty" db:"modification_date"`   // Data de modificação (opcional)
	StatusUUID           string     `json:"status_uuid" db:"status_uuid"`                         // UUID do status (chave estrangeira)
}

type UsersCreatedResponse struct {
	UserUUID string `json:"user_uuid"`
}

package auth

import (
	"bernardtm/backend/configs"
	"bernardtm/backend/internal/core/auth/token"
	"bernardtm/backend/internal/core/email"
	"bernardtm/backend/internal/core/users"
	"bernardtm/backend/internal/utils"
	"bernardtm/backend/pkg/providers/emails"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email string, password string) (string, error)
	Send2FACode(to string, otp string) error
	Login2Step(twoFactorCodeID string, otp string) (LoginResponse, error)
	RequestPasswordReset(requestData RecoverPasswordRequest) error
	ResetPassword(userUUID string, requestData PasswordResetRequest) []error
}

type authService struct {
	userRepo              users.UserRepository
	emailService          email.EmailService
	twoFactorCodesService TwoFactorCodesService
	tokenService          token.TokenService
	frontendURL           string
}

// NewUserService creates a new UserService instance
func NewAuthService(
	userRepo users.UserRepository,
	config *configs.AppConfig,
	emailService email.EmailService,
	twoFactorCodesService TwoFactorCodesService,
	tokenService token.TokenService,
) *authService {
	return &authService{
		userRepo:              userRepo,
		emailService:          emailService,
		twoFactorCodesService: twoFactorCodesService,
		tokenService:          tokenService,
		frontendURL:           config.FrontendURL,
	}
}

// Login authenticate a user and returns a JWT token
func (s *authService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// generate 8 digits numeric otp
	twoFactorRequest := TwoFactorCodesRequest{
		Id:              user.Id,
		Size:            6,
		IsAlphanumeric:  false,
		MinutesToExpiry: 15,
	}
	twoFactor, err := s.twoFactorCodesService.GenerateTwoFactorCode(twoFactorRequest)
	if err != nil {
		return "", err
	}
	// send 2fa code by email
	err = s.Send2FACode(email, twoFactor.Code)
	if err != nil {
		return "", err
	}

	// generate token using twoFactorCodeID
	claims := &token.Claims{
		ID: twoFactor.TwoFactorCodeUUID,
	}
	token, err := s.tokenService.GenerateToken(claims, "2step_verification", 15*time.Minute)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Send2FACode(to string, otp string) error {
	var mail emails.EmailDto
	mail.To = []string{to}
	mail.Sender = "no-reply@company.com"
	mail.Subject = "Código de Verificação"
	mail.Body = fmt.Sprintf(`
	Prezado(a),

	Este é o seu código de verificação:

	%s
	O código de verificação só é válido por 15 minutos. Por favor, não compartilhe este código com ninguém.

	Nota: Caso não tenha iniciado esta solicitação, por favor, altere a sua senha imediatamente. Vá para Alterar Senha.`, otp)
	err := s.emailService.SendEmail(mail)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) Login2Step(twoFactorCodeID string, otp string) (LoginResponse, error) {
	twoFactorCodeResponse, err := s.twoFactorCodesService.ValidateTwoFactorCode(twoFactorCodeID, otp)
	if err != nil {
		return LoginResponse{}, err
	}
	user, err := s.userRepo.GetByID(twoFactorCodeResponse.Id)
	if err != nil {
		return LoginResponse{}, errors.New("invalid code")
	}

	claims := &token.APIClaims{
		ID:          user.Id,
		Email:       user.Email,
		Name:        user.Username,
		Role:        "api",
		Permissions: []string{"read", "write"},
	}
	token, err := s.tokenService.GenerateToken(claims, "api", 24*time.Hour)
	if err != nil {
		return LoginResponse{}, err
	}
	// invalidate two factor only at the end
	if err = s.twoFactorCodesService.InvalidateTwoFactorCode(twoFactorCodeID); err != nil {
		return LoginResponse{}, err
	}
	loginResponse := LoginResponse{
		Token:  token,
		Name:   user.Username,
		Email:  user.Email,
		Avatar: user.ProfileImageLink,
	}
	return loginResponse, err
}

// RequestPasswordReset identify user by email, generate and send a verification code
func (s *authService) RequestPasswordReset(requestData RecoverPasswordRequest) error {
	user, err := s.userRepo.GetByEmail(requestData.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	claims := &token.Claims{
		ID: user.Id,
	}
	token, err := s.tokenService.GenerateToken(claims, "password_reset_verification", 15*time.Minute)
	if err != nil {
		return err
	}
	if err = s.SendRecoveryPasswordLinkEmail(requestData.Email, token); err != nil {
		return err
	}
	return nil
}

func (s *authService) SendRecoveryPasswordLinkEmail(to string, token string) error {

	var mail emails.EmailDto
	mail.To = []string{to}
	mail.Sender = "no-reply@company.com"
	mail.Subject = "Solicitação de Redefinição de Senha"
	mail.IsHTML = true
	mail.Body = fmt.Sprintf(`
	<p>Prezado(a),</p>

	<p>Você solicitou a redefinição de sua senha de acesso.</p>

	<p>Caso você não tenha feito essa solicitação, por favor ignore este e-mail.</p>

	<p>Acesse o link abaixo para redefinir sua senha: <br />
	<a href="%s/recovery-password?token=%s">Clique aqui para redefinir a senha</a>
	</p>

	<p>Este link só é válido por 15 minutos. Por favor, não compartilhe este link com ninguém.</p>`,
		s.frontendURL,
		token,
	)
	err := s.emailService.SendEmail(mail)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) ResetPassword(userid string, requestData PasswordResetRequest) []error {
	var errorsList []error
	user, err := s.userRepo.GetByID(userid)
	if err != nil {
		return append(errorsList, errors.New("invalid code"))
	}

	// validate password
	if errors := utils.ValidatePassword(requestData.Password); errors != nil {
		return errors
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		return append(errorsList, err)
	}

	// update password
	userRequest := users.UserRequest{
		Password:   string(hashedPassword),
		Username:   user.Username,
		Email:      user.Email,
		TaxNumber:  user.TaxNumber,
		StatusUUID: user.StatusUUID,
	}
	s.userRepo.Update(user.Id, userRequest)

	// enviar email informando que a senha foi alterada
	if err = s.SendPasswordResetEmail(user.Email, user.Username); err != nil {
		return append(errorsList, err)
	}
	return nil
}

func (s *authService) SendPasswordResetEmail(to string, username string) error {
	var mail emails.EmailDto
	mail.To = []string{to}
	mail.Sender = "no-reply@company.com"
	mail.Subject = "Senha Redefinida"
	mail.Body = fmt.Sprintf(`
	Prezado(a) %s,

	Sua Senha foi redefinida com sucesso.`,
		username,
	)
	err := s.emailService.SendEmail(mail)
	if err != nil {
		return err
	}
	return nil
}

package auth

import (
	"bernardtm/backend/internal/core/status"
	"bernardtm/backend/internal/utils"
	"errors"
)

type TwoFactorCodesService interface {
	GenerateTwoFactorCode(entity TwoFactorCodesRequest) (TwoFactorCodesResponse, error)
	ValidateTwoFactorCode(twoFactorCodeID string, otp string) (TwoFactorCodesResponse, error)
	InvalidateTwoFactorCode(twoFactorCodeID string) error
}

type twoFactorCodesService struct {
	repo       TwoFactorCodesRepository
	repoStatus status.StatusRepository
}

func NewTwoFactorCodesService(repo TwoFactorCodesRepository, repoStatus status.StatusRepository) *twoFactorCodesService {
	return &twoFactorCodesService{
		repo:       repo,
		repoStatus: repoStatus,
	}
}

// GenerateTwoFactorCode generates a 2FA code and persists it in the database
func (s *twoFactorCodesService) GenerateTwoFactorCode(entity TwoFactorCodesRequest) (TwoFactorCodesResponse, error) {
	twoFactorCodeResponse := TwoFactorCodesResponse{}
	// generate otp
	otp, err := utils.GenerateOTP(entity.Size, entity.IsAlphanumeric)
	if err != nil {
		return twoFactorCodeResponse, err
	}
	// persist otp, 15 minutes expiration
	twoFactorCodeID, err := s.repo.Create(entity.Id, otp, entity.MinutesToExpiry)
	if err != nil {
		return twoFactorCodeResponse, err
	}
	twoFactorCodeResponse.TwoFactorCodeUUID = twoFactorCodeID
	twoFactorCodeResponse.Code = otp
	return twoFactorCodeResponse, nil
}

// ValidateTwoFactorCode validates a 2FA code
func (s *twoFactorCodesService) ValidateTwoFactorCode(twoFactorCodeID string, otp string) (TwoFactorCodesResponse, error) {
	twoFactor, err := s.repo.GetByID(twoFactorCodeID)
	if err != nil {
		return TwoFactorCodesResponse{}, errors.New("invalid 2fa code")
	}

	status, err := s.repoStatus.GetByName("Actived")
	if err != nil {
		return TwoFactorCodesResponse{}, errors.New("invalid 2fa code, inactived")
	}
	if twoFactor.StatusUUID != status.StatusUUID {
		return TwoFactorCodesResponse{}, errors.New("invalid 2fa code, inactived")
	}
	if twoFactor.Code != otp {
		return TwoFactorCodesResponse{}, errors.New("invalid 2fa code")
	}
	if twoFactor.ExpirationDate.Before(twoFactor.CurrentTimestamp) {
		return TwoFactorCodesResponse{}, errors.New("expired 2fa code")
	}
	return twoFactor, nil
}

func (s *twoFactorCodesService) InvalidateTwoFactorCode(twoFactorCodeID string) error {
	return s.repo.Update(twoFactorCodeID)
}

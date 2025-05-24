package shareds

type OnboardingQueueRequest struct {
	FundUUID     string `json:"fund_uuid"`
	FundFileUUID string `json:"fund_file_uuid"`
	FileLink     string `json:"file_link"`
}

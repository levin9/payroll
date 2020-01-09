package models

type AttImportInput struct {
	BizType  string `json:"biz_type"`
	FilePath string `json:"file_path"`
	TenantId string `json:"tenant_id"`
}

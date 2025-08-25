package dto



type Attachment struct {
    AttachmentId   string ` json:"AttachmentId" `
    FileName   string ` json:"FileName" `
    FileType   string ` json:"FileType" `
    UploadedBy   string ` json:"UploadedBy" `
    UploadDate   string ` json:"UploadDate" `
    ChangeLogId   string ` json:"ChangeLogId" `
    Deleted bool `json:"deleted"`}


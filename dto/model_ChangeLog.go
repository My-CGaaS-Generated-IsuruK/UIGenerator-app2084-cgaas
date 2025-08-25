package dto



type ChangeLog struct {
    ChangeLogId   string ` json:"ChangeLogId" `
    Description   string ` json:"Description" `
    Date   string ` json:"Date" `
    Type   string ` json:"Type" `
    ProjectId   string ` json:"ProjectId" `
    Deleted bool `json:"deleted"`}


package dto



type Project struct {
    ProjectId   string ` json:"ProjectId" `
    Name   string ` json:"Name" `
    Description   string ` json:"Description" `
    StartDate   string ` json:"StartDate" `
    EndDate   string ` json:"EndDate" `
    Status   string ` json:"Status" `
    WorkspaceId   string ` json:"WorkspaceId" `
    Deleted bool `json:"deleted"`}


package rpc

type InstructionCommon struct {
	Accounts    []string `json:"accounts"`
	Data        string   `json:"data"`
	ProgramId   string   `json:"programId"`
	StackHeight uint64   `json:"stackHeight"`
}

type InstructionParsed struct {
	Parsed      any    `json:"parsed"`
	Program     string `json:"program"`
	ProgramId   string `json:"programId"`
	StackHeight uint64 `json:"stackHeight"`
}

type InstructionFull struct {
	Accounts    []string `json:"accounts"`
	Data        string   `json:"data"`
	ProgramId   string   `json:"programId"`
	StackHeight uint64   `json:"stackHeight"`
	Parsed      any      `json:"parsed"`
	Program     string   `json:"program"`
}

type ParsedObjInfo struct {
	Amount      string `json:"amount"`
	Authority   string `json:"authority"`
	Destination string `json:"destination"`
	Source      string `json:"source"`
}

type ParsedObj struct {
	Info ParsedObjInfo `json:"info"`
	Type string        `json:"type"`
}

type InstructionInnerParsed struct {
	Parsed      ParsedObj `json:"parsed"`
	Program     string    `json:"program"`
	ProgramId   string    `json:"programId"`
	StackHeight uint64    `json:"stackHeight"`
}

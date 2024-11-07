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

type InfoTransfer struct {
	Amount      string `json:"amount"`
	Authority   string `json:"authority"`
	Destination string `json:"destination"`
	Source      string `json:"source"`
}

type ParsedObj struct {
	Info any    `json:"info"`
	Type string `json:"type"`
}

type InstructionInnerParsed struct {
	Parsed      any    `json:"parsed"`
	Program     string `json:"program"`
	ProgramId   string `json:"programId"`
	StackHeight uint64 `json:"stackHeight"`
}

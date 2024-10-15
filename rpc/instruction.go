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

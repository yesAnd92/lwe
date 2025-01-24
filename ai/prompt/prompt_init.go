package prompt

import _ "embed"

//go:embed git_log_summary.prompt
var LogSummaryPrompt string

//go:embed git_diff_summary.prompt
var GitDiffPrompt string

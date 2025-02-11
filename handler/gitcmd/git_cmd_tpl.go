package gitcmd

//git command

var (

	//git log
	LOG_TPL            = "git -C %s --no-pager log %s --no-merges "
	LOG_FORMAT_TPL     = `--format=format:'%h*-*%an*-*%ct*-*%s' ` //使用*-*作为分隔符
	LOG_AUTHOR_TPL     = `--author=%s `
	LOG_START_DATE_TPL = `--since=%s `
	LOG_END_DATE_TPL   = `--until=%s `
	LOG_RECENTN_TPL    = `-n %d `

	// git show
	SHOW_TPL = "git show %s"

	// git status
	STATUS_TPL = "git -C %s status"

	// git status short
	STATUS_TPL_SHORT = "git -C %s status --short | grep '^[ MADRCU]' | awk '{print $2}'"

	// git status in short result
	STATUS_CHECK_TPL = "git -C %s status -s"

	//git pull
	GIT_PULL = `git -C %s pull --rebase`

	// git clone
	CLONE_TPL = "git clone %s %s"

	// git lab group api
	// gitlab doc  https://docs.gitlab.com/ee/api/groups.html#details-of-a-group
	GITLAB_GROUP_DETAIL = "%s://%s/api/v4/groups%s?private_token=%s"

	// Determine if the current directory is a git repository
	EXIST_GIT_REPO = "git rev-parse --is-inside-work-tree"

	//git Branch
	GIT_BRANCH = `git branch`

	//git diff
	GIT_DIFF = `git -C %s diff -w -b --ignore-cr-at-eol --diff-filter=d | grep -vE '^(index|diff|\+\+\+)'`

	//git add
	GIT_ADD = `git -C %s add .`

	//git commit
	GIT_COMMIT = `git -C %s commit -m "%s"`

	//git push
	GIT_PUSH = `git -C %s push`
)

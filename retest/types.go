package retest

type GHRetest struct {
	Name   string
	Url    string
	Config interface{}
}

type Runtime struct {
	CommentId int64
	Debug     bool
	Pr        string
	Nwo       string
	Owner     string
	Repo      string
}

type GHRetestResult struct {
	Retested int
	Error    int
}

type PullRequest struct {
	Number int
	Branch string
	Commit string
}

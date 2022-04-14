package git

var (
	_            = Branch
	_ pushOption = (*optionBranch)(nil)
)

type optionBranch struct {
	branch string
}

// Branch 分支
func Branch(branch string) *optionBranch {
	return &optionBranch{
		branch: branch,
	}
}

func (b *optionBranch) applyPush(options *pushOptions) {
	options.branch = b.branch
}

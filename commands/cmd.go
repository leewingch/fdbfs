package commands

type Command interface {
	Execute(context *Context) error
}

type CmdParser struct {
}

func (cmd *CmdParser) Parse(cmdString string) (cmd *CmdParser, err error) {
	return nil, nil
}

type LsCmd struct {
	path string
}

func (cmd *LsCmd) Execute(context *Context) error {
	entries, err := path.Entries(cmd.path)
	if err != nil {
		return err
	}
}

type CatCmd struct {
	path string
}

func (cmd *CatCmd) Execute(context *Context) error {
	ok, err := path.IsDir(cmd.path)
	if err != nil {
		return err
	}
	if ok {
		fmt.Fprintf(os.Stderr, "%s is not a dir", cmd.path)
		return nil
	}

	data, err := path.Read(cmd.path)
	if err != nil {
		return err
	}

	fmt.Println(data)
}

type CdCmd struct {
	path string
}

func (cmd *CdCmd) Execute(context *Context) error {
	ok, err := path.IsDir(cmd.path)
	if err != nil {
		return err
	}

	if ok {
		context.CurrentPath = cmd.path;
	} else {
		fmt.Fprintf(os.Stderr, "%s is not a dir", cmd.path)
	}
}

type QuitCmd struct {
}

func (cmd *QuitCmd) Execute(context *Context) error {
	os.Exit(0);
}

type PwdCmd struct {
}

func (cmd *PwdCmd) Execute(context *Context) error {
	fmt.Println(context.CurrentPath);
}

type StatCmd struct {
	path string
}

func (cmd *StatCmd) Execute(context *Context) error {
	stat, err := path.Stat(cmd.path)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", stat);
}

type RmCmd struct {
	path string
}

func (cmd *RmCmd) Execute(context *Context) error {
	return path.Remove(cmd.path)
}

type CpCmd struct {
	from string
	to string
}

func (cmd *CpCmd) Execute(context *Context) error {
	return path.Copy(cmd.from, cmd.to)
}

type MvCmd struct {
	from string
	to string
}

func (cmd *MvCmd) Execute(context *Context) error {
	return path.Move(cmd.from, cmd.to)
}







package path

func Entries(path string) (entries []string, err error) {
}

func IsDir(path string) (ok bool, err error) {
}

func Read(path string) ([]byte, err error) {
}

type StatInfo struct {
}

func Stat(path string) (info *StatInfo, err error) {
}

func Remove(path string) error {
}

func Copy(from, to string) error {
}

func Move(from, to string) error {
}


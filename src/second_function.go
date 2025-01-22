package src

var (
	tes string
)

func Open() error {
	tes = "missutsan"
	return nil
}

func Tes() string {
	return tes
}

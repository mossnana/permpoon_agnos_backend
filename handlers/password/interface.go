package password

type MPA interface {
	GetMinimumActionToValid(password string) uint
}

package types

// 类型接口
type Object interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

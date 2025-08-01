package reductor

// пустая модель если надо вернуть по умолчанию не известно что
type ModelGeneric struct{}

// формат сообщения
type Message struct {
	Sender string
	Page   ModelType
	Model  interface{}
}

package driver

type Message struct {
	Payload interface{}
}

type SetCommand struct {
	Key string
	Value interface{}
}

type GetCommand struct {
	Key string
}

type DelCommand struct {
	Key string
}

type ExitCommand struct {

}

type SetResult struct {
	InsertedCount int
}

type DelResult struct {
	DeletedCount int
}

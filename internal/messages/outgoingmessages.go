package messages

// HitMessage send to player when he hits stuff
type HitMessage struct {
	BaseMessage
	Coordinate Coordinate `json:"coordinate"`
}

// NewHitMessage constructor function
func NewHitMessage(username string, coordinate Coordinate) HitMessage {
	return HitMessage{
		BaseMessage: BaseMessage{Username: username, Type: "HIT"},
		Coordinate:  coordinate,
	}
}

// MissMessage send to player when he misses stuff
type MissMessage struct {
	BaseMessage
	Coordinate Coordinate `json:"coordinate"`
}

// NewMissMessage constructor function
func NewMissMessage(username string, coordinate Coordinate) MissMessage {
	return MissMessage{
		BaseMessage: BaseMessage{Username: username, Type: "MISS"},
		Coordinate:  coordinate,
	}
}

type ShipDestroyedMessage struct {
	BaseMessage
	Coordinate Coordinate `json:"coordinate"`
	ShipSize   int        `json:"shipSize"`
	Vertical   bool       `json:"vertical"`
}

func NewShipDestroyedMessage(username string, coordinate Coordinate, shipSize int, vertical bool) ShipDestroyedMessage {
	return ShipDestroyedMessage{
		BaseMessage: BaseMessage{Username: username, Type: "SHIP_DESTROYED"},
		Coordinate:  coordinate,
		ShipSize:    shipSize,
		Vertical:    vertical,
	}
}

type VictoryMessage struct {
	BaseMessage
}

func NewVictoryMessage(username string) VictoryMessage {
	return VictoryMessage{
		BaseMessage{Username: username, Type: "VICTORY"},
	}
}

type LossMessage struct {
	BaseMessage
}

func NewLossMessage(username string) LossMessage {
	return LossMessage{
		BaseMessage{Username: username, Type: "LOSS"},
	}
}

type BoardMessage struct {
	BaseMessage
	ShipSizes []int `json:"shipSizes"`
}

func NewBoardMessage(username string, shipSizes []int) BoardMessage {
	return BoardMessage{
		BaseMessage: BaseMessage{Username: username, Type: "BOARD"},
		ShipSizes:   shipSizes,
	}
}

type awaitingPlayersMessage struct {
	BaseMessage
}

func NewAwaitingPlayersMessage(username string) awaitingPlayersMessage {
	return awaitingPlayersMessage{
		BaseMessage{Username: username, Type: "AWAITING_PLAYERS"},
	}
}

type gameStartedMessage struct {
	BaseMessage
	TurnMessage
}

func NewGameStartedMessage(username string, turn bool, duration int) gameStartedMessage {
	return gameStartedMessage{
		BaseMessage: BaseMessage{Username: username, Type: "GAME_STARTED"},
		TurnMessage: TurnMessage{Turn: turn, Duration: duration},
	}
}

type TurnMessage struct {
	BaseMessage
	Turn     bool `json:"turn"`
	Duration int  `json:"duration"`
}

func NewTurnMessage(username string, turn bool, duration int) TurnMessage {
	return TurnMessage{
		BaseMessage: BaseMessage{Username: username, Type: "TURN"},
		Turn:        turn,
		Duration:    duration,
	}
}

type TurnExtendedMessage struct {
	BaseMessage
	Turn     bool `json:"turn"`
	Duration int  `json:"duration"`
}

func NewTurnExtendedMessage(username string, duration int) TurnExtendedMessage {
	return TurnExtendedMessage{
		BaseMessage: BaseMessage{Username: username, Type: "TURN_EXTENDED"},
		Turn:        true,
		Duration:    duration,
	}
}

type boardStateMessage struct {
	BaseMessage
	Hits    []HitMessage           `json:"hits"`
	Misses  []MissMessage          `json:"misses"`
	Destoys []ShipDestroyedMessage `json:"destroys"`
	Board   BoardMessage           `json:"board"`
}

func NewBoardStateMessage(username string, hits []HitMessage, misses []MissMessage, destroys []ShipDestroyedMessage, board BoardMessage) boardStateMessage {
	return boardStateMessage{
		BaseMessage: BaseMessage{Username: username, Type: "BOARD_STATE"},
		Board:       board,
		Hits:        hits,
		Misses:      misses,
		Destoys:     destroys,
	}
}

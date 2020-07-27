package messages

// HitMessage send to player when he hits stuff
type hitMessage struct {
	BaseMessage
	Coordinate Coordinate `json:"coordinate"`
}

// NewHitMessage constructor function
func NewHitMessage(coordinate Coordinate) hitMessage {
	return hitMessage{
		BaseMessage: BaseMessage{Type: "HIT"},
		Coordinate:  coordinate,
	}
}

// MissMessage send to player when he misses stuff
type missMessage struct {
	BaseMessage
	Coordinate Coordinate `json:"coordinate"`
}

// NewMissMessage constructor function
func NewMissMessage(coordinate Coordinate) missMessage {
	return missMessage{
		BaseMessage: BaseMessage{Type: "MISS"},
		Coordinate:  coordinate,
	}
}

type shipDestroyedMessage struct {
	BaseMessage
	Coordinate Coordinate `json:"coordinate"`
	ShipSize   int        `json:"shipSize"`
	Vertical   bool       `json:"vertical"`
}

func NewShipDestroyedMessage(coordinate Coordinate, shipSize int, vertical bool) shipDestroyedMessage {
	return shipDestroyedMessage{
		BaseMessage: BaseMessage{Type: "SHIP_DESTROYED"},
		Coordinate:  coordinate,
		ShipSize:    shipSize,
		Vertical:    vertical,
	}
}

type VictoryMessage struct {
	BaseMessage
}

func NewVictoryMessage() VictoryMessage {
	return VictoryMessage{
		BaseMessage{Type: "VICTORY"},
	}
}

type LossMessage struct {
	BaseMessage
}

func NewLossMessage() LossMessage {
	return LossMessage{
		BaseMessage{Type: "LOSS"},
	}
}

type boardMessage struct {
	BaseMessage
	ShipSizes []int `json:"shipSizes"`
}

func NewBoardMessage(shipSizes []int) boardMessage {
	return boardMessage{
		BaseMessage: BaseMessage{Type: "BOARD"},
		ShipSizes:   shipSizes,
	}
}

type awaitingPlayersMessage struct {
	BaseMessage
}

func NewAwaitingPlayersMessage() awaitingPlayersMessage {
	return awaitingPlayersMessage{
		BaseMessage{Type: "AWAITING_PLAYERS"},
	}
}

type gameStartedMessage struct {
	BaseMessage
	TurnMessage
}

func NewGameStartedMessage(turn bool, duration int) gameStartedMessage {
	return gameStartedMessage{
		BaseMessage: BaseMessage{Type: "GAME_STARTED"},
		TurnMessage: TurnMessage{Turn: turn, Duration: duration},
	}
}

type TurnMessage struct {
	BaseMessage
	Turn     bool `json:"turn"`
	Duration int  `json:"duration"`
}

func NewTurnMessage(turn bool, duration int) TurnMessage {
	return TurnMessage{
		BaseMessage: BaseMessage{Type: "TURN"},
		Turn:        turn,
		Duration:    duration,
	}
}

type TurnExtendedMessage struct {
	BaseMessage
	Turn     bool `json:"turn"`
	Duration int  `json:"duration"`
}

func NewTurnExtendedMessage(duration int) TurnExtendedMessage {
	return TurnExtendedMessage{
		BaseMessage: BaseMessage{Type: "TURN_EXTENDED"},
		Turn:        true,
		Duration:    duration,
	}
}

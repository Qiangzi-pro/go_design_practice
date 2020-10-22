package article64_state_machine

type State int32

const (
	SMALL State = iota
	SUPER
	FIRE
	CAPE
)

// 查表法
type Event int32

const (
	GOT_MUSHROOT = iota
	GOT_CAPE
	GOT_FIRE
	MET_MONSTER
)

type MarioStateMachine struct {
	score int
	currentState State
	transitionTable [][]State
	actionTable [][]int
}

func (fsm *MarioStateMachine) ObtainMushRoom() {
	fsm.executeEvent(GOT_MUSHROOT)
}

func (fsm *MarioStateMachine) ObtainCape() {
	fsm.executeEvent(GOT_CAPE)
}

func (fsm *MarioStateMachine) ObtainFireFlower() {
	fsm.executeEvent(GOT_FIRE)
}

func (fsm *MarioStateMachine) MeetMonster() {
	fsm.executeEvent(MET_MONSTER)
}

var tt = [][]State{
	{SUPER, CAPE, FIRE, SMALL},
	{SUPER, CAPE, FIRE, SMALL},
	{CAPE, CAPE, CAPE, SMALL},
	{FIRE, FIRE, FIRE, SMALL},
}

var at = [][]int{
	{+100, +200, +300, +0},
	{+0, +200, +300, -100},
	{+0, +0, +0, -200},
	{+0, +0, +0, -300},
}

type IMario interface {
	ObtainMushRoom()
	ObtainCape()
	ObtainFireFlower()
	MeetMonster()
}

func NewMarioStateMachine() *MarioStateMachine {
	return &MarioStateMachine{
		score: 0,
		currentState: SMALL,
		transitionTable: tt,
		actionTable: at,
	}
}

func (fsm *MarioStateMachine) executeEvent(event Event)  {
	stateValue := fsm.currentState
	fsm.currentState = tt[stateValue][event]
	fsm.score = at[stateValue][event]
}

package lobby

import (
	missionloader "problem_solver/mission_loader"
	"problem_solver/packet/types"
	"sync"
)

type Lobby struct {
	Code            string
	Players         []*Player
	Questions       []missionloader.Question
	CurrentQuestion int
	State           types.State
	mu              sync.RWMutex
}

func NewLobby(code string) *Lobby {
	return &Lobby{
		Code:            code,
		Players:         make([]*Player, 0),
		Questions:       make([]missionloader.Question, 0),
		CurrentQuestion: -1,
		State:           types.StateWaiting,
	}
}

func (l *Lobby) AddQuestion(q missionloader.Question) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Questions = append(l.Questions, q)
}

func (l *Lobby) NextQuestion() missionloader.Question {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.Questions) == 0 {
		return nil
	}

	l.CurrentQuestion++
	if l.CurrentQuestion >= len(l.Questions) {
		return nil
	}

	return l.Questions[l.CurrentQuestion]
}

func (l *Lobby) BroadcastToAll(packet types.CBBasePacket) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, player := range l.Players {
		player.SendPacket(packet)
	}
}

func (l *Lobby) AddPlayer(player *Player) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Players = append(l.Players, player)
}

func (l *Lobby) RemovePlayer(name string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, p := range l.Players {
		if p.Name == name {
			l.Players = append(l.Players[:i], l.Players[i+1:]...)
			break
		}
	}
}

func (l *Lobby) GetPlayer(name string) *Player {
	l.mu.RLock()
	defer l.mu.RUnlock()
	for _, p := range l.Players {
		if p.Name == name {
			return p
		}
	}
	return nil
}

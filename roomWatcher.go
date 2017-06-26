package main

import (
	"encoding/json"
	"strconv"
)

// Observation is the struct that represent the state of a room at a fixed time
// it's also the stuct we export in JSON and send to the webclient through
// the websocket
type Observation struct {
	Name            string
	RoomNumber      string
	LastEvent       string
	Tmc             int
	Activity        []int
	InactivityCount int
	BathRoomCount   int
}

func isEventActif(event string) bool {
	switch event {
	case
		"BEDROOM",
		"BATHROOM",
		"FALL":
		return true
	}
	return false
}

func (o *Observation) updateWith(r RoomJSON) {

	o.RoomNumber = strconv.Itoa(r.Room)
	o.BathRoomCount = r.BathroomCount
	o.LastEvent = r.LastEvent

	if isEventActif(o.LastEvent) {
		o.Activity = append(o.Activity, 1)
		o.InactivityCount = 0
	} else {
		o.Activity = append(o.Activity, 0)
		o.InactivityCount++
	}

	if o.InactivityCount >= 5 {
		o.InactivityCount = 0
		o.Activity = nil
		o.Tmc = 0
	} else {
		// int(len(room['acti']) / 5) * 5
		// Hum / 5 * 5 ?
		o.Tmc = len(o.Activity)
	}
}

// RoomWatcher is the struct who interface with all the remote update and
// local storage logic
type RoomWatcher struct {
	number          int
	LastObservation Observation
}

// NewRoomWatcher is the factory for RoomWatcher
func NewRoomWatcher(number int) RoomWatcher {

	r := RoomWatcher{}
	r.number = number
	return r
}

func (r RoomWatcher) persistanceKey() string {

	return strconv.Itoa(r.number)
}

// UpdateDataInStore ask for an update of the previous observation. It's keeping
// track of changes in the key value store
func (r *RoomWatcher) UpdateDataInStore() {

	t := GetTarketService()
	result, _ := t.GetRoomInfos(r.number)

	v, ok := GetKVStore().Read(r.persistanceKey())

	var o Observation

	if ok {
		json.Unmarshal([]byte(v), &o)
	} else {
		o = Observation{}
	}

	o.updateWith(result.Room)
	r.LastObservation = o

	serialized, err := json.Marshal(o)
	if err == nil {
		GetKVStore().Write(r.persistanceKey(), serialized)
	}
}

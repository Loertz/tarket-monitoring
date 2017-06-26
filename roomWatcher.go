package main

import (
	"encoding/json"
	"strconv"
)

type Observation struct {
	Name            string
	RoomNumber      string
	LastEvent       string
	Tmc             int
	Activity        []int
	InactivityCount int
	BathRoomCount   int
}

func IsEventActif(event string) bool {
	switch event {
	case
		"BEDROOM",
		"BATHROOM",
		"FALL":
		return true
	}
	return false
}

func (o *Observation) UpdateWith(r RoomJSON) {

	o.RoomNumber = strconv.Itoa(r.Room)
	o.BathRoomCount = r.BathroomCount
	o.LastEvent = r.LastEvent

	if IsEventActif(o.LastEvent) {
		o.Activity = append(o.Activity, 1)
		o.InactivityCount = 0
	} else {
		o.Activity = append(o.Activity, 0)
		o.InactivityCount += 1
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

type RoomWatcher struct {
	number          int
	LastObservation Observation
}

func NewRoomWatcher(number int) RoomWatcher {

	r := RoomWatcher{}
	r.number = number
	return r
}

func (r RoomWatcher) persistanceKey() string {

	return strconv.Itoa(r.number)
}

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

	o.UpdateWith(result.Room)
	r.LastObservation = o

	serialized, err := json.Marshal(o)
	if err == nil {
		GetKVStore().Write(r.persistanceKey(), serialized)
	}
}

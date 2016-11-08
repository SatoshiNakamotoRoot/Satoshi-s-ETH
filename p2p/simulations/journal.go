package simulations

import (
	"bytes"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger/glog"
)

// Journal is an instance of a guaranteed no-loss subscription using event.TypeMux
// Network components POST events to the TypeMux, which then is read by the journal
// Each journal belongs to a subscription
type Journal struct {
	lock    sync.Mutex
	counter int
	cursor  int
	sub     event.Subscription
	events  []*event.Event
}

// func (self *Journal) SnapshotAt(pos int) {}

// NewJournal constructor takes eventer and types to subscribe to
func NewJournal(eventer *event.TypeMux, types ...interface{}) *Journal {
	self := &Journal{}
	self.sub = eventer.Subscribe(types...)
	go func() {
		self.Write()
	}()
	return self
}

func (self *Journal) Close() {
	self.sub.Unsubscribe()
}

// Write is a forever loop
func (self *Journal) Write() {
	for {
		select {
		case ev, ok := <-self.sub.Chan():
			if !ok {
				return
			}
			self.append(ev)
		}
	}
}

func (self *Journal) append(evs ...*event.Event) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.events = append(self.events, evs...)
}

func (self *Journal) WaitEntries(n int) {
	for self.NewEntries() < n {
	}
}

func (self *Journal) NewEntries() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	return len(self.events) - self.cursor
}

func (self *Journal) Read(f func(*event.Event) bool) (read int, err error) {
	self.lock.Lock()
	defer self.lock.Unlock()
	for self.cursor < len(self.events) && f(self.events[self.cursor]) {
		read++
		self.cursor++
	}
	self.reset(self.cursor)
	return read, nil
}

func (self *Journal) Reset(n int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.reset(n)
}

func (self *Journal) reset(n int) {
	if n > self.counter {
		n = self.counter
	}
	self.events = self.events[self.cursor:]
	self.cursor = 0
}

func (self *Journal) Counter() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.counter
}

// type History()

func (self *Journal) Cursor() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.cursor
}

// deltas: changes in the number of cumulative actions: non-negative integers.
// base unit is the fixed minimal interval  between two measurements (time quantum)
// acceleration : to slow down you just set the base unit higher.
// to speed up: skip x number of base units
// frequency: given as the (constant or average) number of base units between measurements
// if resolution is expressed as the inverse of frequency  = preserved information
// setting the acceleration
// beginning of the record (lifespan) of the network is index 0
// acceleration means that snapshots are rarer so the same span can be generated by the journal
// then update logs can be compressed (toonly one state transition per affected node)
// epoch, epochcount

type Delta struct {
	On  int
	Off int
}

func oneOutOf(n int) int {
	t := rand.Intn(n)
	if t == 0 {
		return 1
	}
	return 0
}

func deltas(i int) (d []*Delta) {
	if i == 0 {
		return []*Delta{
			&Delta{10, 0},
			&Delta{20, 0},
		}
	}
	return []*Delta{
		&Delta{oneOutOf(10), oneOutOf(10)},
		&Delta{oneOutOf(2), oneOutOf(2)},
	}
}

func mockJournalTest(nw *Network, ticker *<-chan time.Time) {

	ids := RandomNodeIDs(100)
	action := "Off"
	for n := 0; ; n++ {
		select {
		case <-*ticker:
			var entries []*Entry

			if n == 0 {

				entries = []*Entry{
					&Entry{
						Type:   "Node",
						Action: "On",
						Object: &SimNode{ID: ids[0]},
					},
					&Entry{
						Type:   "Node",
						Action: "On",
						Object: &SimNode{ID: ids[1]},
					},
				}
			} else {
				sc := &SimConn{
					Caller: ids[0],
					Callee: ids[1],
				}
				if n%3 == 0 {
					if action == "On" {
						action = "Off"
					} else {
						action = "On"
					}
					entries = append(entries, &Entry{
						Type:   "Conn",
						Action: action,
						Object: sc,
					})
				}
			}

			glog.V(6).Info("entries: %v", entries)
			nw.AppendEntries(entries...)
		}
	}
}

// MockNetwork generates random connectivity events and posts them
// to the eventer
// The journal using the eventer can then be read to visualise or
// drive connections
func MockNetwork(eventer *event.TypeMux, ticker <-chan time.Time) {
	ids := RandomNodeIDs(100)
	var onNodes []*SimNode
	offNodes := ids
	var onConns []*SimConn

	n := 0
	for _ = range ticker {
		ds := deltas(n)
		for i := 0; len(offNodes) > 0 && i < ds[0].On; i++ {
			c := rand.Intn(len(offNodes))
			sn := &SimNode{ID: offNodes[c]}
			err := eventer.Post(&Entry{
				Type:   "Node",
				Action: "On",
				Object: sn,
			})
			if err != nil {
				panic(err.Error())
			}
			onNodes = append(onNodes, sn)
			offNodes = append(offNodes[0:c], offNodes[c+1:]...)
		}
		for i := 0; len(onNodes) > 0 && i < ds[0].Off; i++ {
			c := rand.Intn(len(onNodes))
			sn := onNodes[c]
			err := eventer.Post(&Entry{
				Type:   "Node",
				Action: "Off",
				Object: sn,
			})
			if err != nil {
				panic(err.Error())
			}
			onNodes = append(onNodes[0:c], onNodes[c+1:]...)
			offNodes = append(offNodes, sn.ID)
		}
		for i := 0; len(onNodes) > 1 && i < ds[1].On; i++ {
			caller := onNodes[rand.Intn(len(onNodes))].ID
			callee := onNodes[rand.Intn(len(onNodes))].ID
			if bytes.Compare(caller[:], callee[:]) >= 0 {
				i--
				continue
			}
			sc := &SimConn{
				Caller: caller,
				Callee: callee,
			}
			err := eventer.Post(&Entry{
				Type:   "Conn",
				Action: "On",
				Object: sc,
			})
			if err != nil {
				panic(err.Error())
			}
			onConns = append(onConns, sc)
		}
		for i := 0; len(onConns) > 0 && i < ds[1].Off; i++ {
			c := rand.Intn(len(onConns))
			err := eventer.Post(&Entry{
				Type:   "Conn",
				Action: "Off",
				Object: onConns[c],
			})
			if err != nil {
				panic(err.Error())
			}
			onConns = append(onConns[0:c], onConns[c+1:]...)
		}
		n++
	}

}

package core

import (
	"github.com/Adeithe/go-twitch/irc"
	"github.com/nicklaw5/helix/v2"

	"whisprite/query"
)

type Handlers []Handler

type Dispatch struct {
	handlers map[string]Handlers

	Twitch *helix.Client
	Writer *irc.Conn
}

func (d *Dispatch) GuardHandlers() {
	if d.handlers == nil {
		d.handlers = make(map[string]Handlers)
	}
}

func (d *Dispatch) Register(pool string, handlers ...Handler) {
	d.GuardHandlers()
	d.handlers[pool] = append(d.handlers[pool], handlers...)
}

func (d *Dispatch) RegisterCounters() {
	var (
		c = query.Counter
	)
	d.GuardHandlers()

	pool := "counts"
	d.FlushPool(pool)

	counters, _ := c.Find()
	handlers := make(Handlers, len(counters))

	for idx, counter := range counters {
		handlers[idx] = Handler{
			Name: counter.Name,
			Run: func(d Dispatch, e Event, self Handler) {
				c.Where(c.Name.Eq(e.Name)).Update(c.Value, c.Value.Add(1))
				e.Sayf(e.Channel, "The %s counter is now at %d.", e.Name, counter.Value+1)
			},
		}
	}

	d.Register(pool, handlers...)
}

func (d *Dispatch) Handle(event Event) {
	for pool, handlers := range d.handlers {
		for idx, handler := range handlers {
			if handler.Authenticate(event) && handler.RespondsTo(event.Name) {
				log.Debugf("%s (%s[%v]) matched handler %s, firing", event.Name, pool, idx, handler.Name)
				handler.Run(*d, event, handler)
			}
		}
	}
}

func (d *Dispatch) FlushPool(pool string) {
	log.Debugf("Flushing pool %s", pool)
	d.handlers[pool] = make(Handlers, 0)
}

package handlers

import (
	"strings"
	"whisprite/core"
	"whisprite/model"
	"whisprite/query"
)

var Counters = []core.Handler{
	{
		Name:        "newcounter",
		VipRequired: true,
		Run: func(d core.Dispatch, e core.Event, self core.Handler) {
			var (
				name  = e.Args[0]
				iname = strings.ToLower(name)
				c     = query.Counter
			)

			existing, _ := c.Where(c.Name.Eq(iname)).First()
			if existing != nil {
				e.Sayf(e.Channel, "We already have a counter named %s, with %d on it!", iname, existing.Value)
				return
			}

			c.Create(&model.Counter{Name: iname, Value: 1})
			e.Sayf(e.Channel, "Created counter %s", iname)
		},
	},

	{
		Name:        "rmcounter",
		ModRequired: true,
		Run: func(d core.Dispatch, e core.Event, self core.Handler) {
			var (
				c     = query.Counter
				iname = strings.ToLower(e.Args[0])
			)

			info, err := c.Where(c.Name.Eq(iname)).Delete(&model.Counter{})
			if err != nil {
				e.Sayf(e.Channel, "%v", err)
				return
			}

			if info.RowsAffected == 0 {
				e.Say(e.Channel, "No matching counters found :shrug:")
				return
			}

			e.Sayf(e.Channel, "Deleted counter %s", iname)
		},
	},

	{
		Name:        "uncount",
		ModRequired: true,
		Run: func(d core.Dispatch, e core.Event, self core.Handler) {
			var (
				c     = query.Counter
				iname = strings.ToLower(e.Args[0])
			)

			counter, _ := c.Where(c.Name.Eq(iname)).First()
			if counter == nil {
				e.Say(e.Channel, "No such counter :shrug:")
				return
			}

			c.Where(c.Name.Eq(iname)).Update(c.Value, c.Value.Sub(1))
			e.Sayf(e.Channel, "The %s counter is now at %d", iname, counter.Value-1)
		},
	},
}

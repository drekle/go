package main

import (
	"encoding/json"

	"github.com/drekle/go/rss/gopherjs/actions"
	"github.com/drekle/go/rss/gopherjs/components"
	"github.com/drekle/go/rss/gopherjs/dispatcher"
	"github.com/drekle/go/rss/gopherjs/store"
	"github.com/drekle/go/rss/gopherjs/store/model"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

func main() {
	attachLocalStorage()

	vecty.SetTitle("GopherJS • TodoMVC")
	vecty.AddStylesheet("node_modules/todomvc-common/base.css")
	vecty.AddStylesheet("node_modules/todomvc-app-css/index.css")
	p := &components.PageView{}
	store.Listeners.Add(p, func() {
		p.Items = store.Items
		vecty.Rerender(p)
	})
	vecty.RenderBody(p)
}

func attachLocalStorage() {
	store.Listeners.Add(nil, func() {
		data, err := json.Marshal(store.Items)
		if err != nil {
			println("failed to store items: " + err.Error())
		}
		js.Global.Get("localStorage").Set("items", string(data))
	})

	if data := js.Global.Get("localStorage").Get("items"); data != js.Undefined {
		var items []*model.Item
		if err := json.Unmarshal([]byte(data.String()), &items); err != nil {
			println("failed to load items: " + err.Error())
		}
		dispatcher.Dispatch(&actions.ReplaceItems{
			Items: items,
		})
	}
}

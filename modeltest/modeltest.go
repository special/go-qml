package main

import (
	"fmt"
	"github.com/go-qml/qml"
	"log"
	"time"
)

func main() {
	if err := qml.Run(run); err != nil {
		log.Panic(err)
	}
}

type MyModel struct {
	qml.QmlModelBase
	Count int
}

func (this *MyModel) RowCount() int {
	return this.Count
}

func (this *MyModel) RoleNames() map[int]string {
	re := make(map[int]string)
	re[10] = "magic"
	return re
}

func (this *MyModel) Data(row, role int) interface{} {
	return fmt.Sprintf("row %d role %d", row, role)
}

func run() error {
	engine := qml.NewEngine()
	model := &MyModel{}
	go func() {
		time.Sleep(10 * time.Second) // XXX hack to avoid crashing before _addr is set
		for {
			model.BeginInsertRows(model.Count, model.Count)
			model.Count++
			model.EndInsertRows()
			time.Sleep(2 * time.Second)
		}
	}()

	engine.Context().SetVar("goModel", model)
	component, err := engine.LoadFile("model.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()
	return nil
}

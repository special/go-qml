package qml

// #include "capi.h"
//
import "C"

import (
	"fmt"
	"unsafe"
)

type QmlModelBase struct {
	// XXX Embed Common instead?
	_addr   unsafe.Pointer
	_engine *Engine
}

type QmlModel interface {
	// Embed QmlModelBase for implementation
	addr(impl QmlModel, engine *Engine) unsafe.Pointer
	engine() *Engine

	RowCount() int
	RoleNames() map[int]string
	Data(row, role int) interface{}
}

var models = make(map[unsafe.Pointer]QmlModel)

func (this *QmlModelBase) addr(impl QmlModel, engine *Engine) unsafe.Pointer {
	if this._addr == nil {
		RunMain(func() {
			this._addr = C.createModel()
			this._engine = engine
			models[this._addr] = impl
		})
	}
	return this._addr
}

func (this *QmlModelBase) engine() *Engine {
	return this._engine
}

//export hookModelRowCount
func hookModelRowCount(model unsafe.Pointer) C.int {
	return C.int(models[model].RowCount())
}

//export hookModelRoleNames
func hookModelRoleNames(model unsafe.Pointer, re *C.DataValue) {
	roles := models[model].RoleNames()
	var value string
	for role, name := range roles {
		value += fmt.Sprintf("%d=%s;", role, name)
	}
	packDataValue(value, re, models[model].engine(), cppOwner)
}

//export hookModelData
func hookModelData(model unsafe.Pointer, row, role C.int, re *C.DataValue) {
	value := models[model].Data(int(row), int(role))
	packDataValue(value, re, models[model].engine(), cppOwner)
}

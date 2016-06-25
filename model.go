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

	// Modification signals
	BeginResetModel()
	EndResetModel()
	BeginInsertRows(start, end int)
	EndInsertRows()
	BeginRemoveRows(start, end int)
	EndRemoveRows()
	DataChanged(start, end int)

	// Abstract API
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

func (this *QmlModelBase) BeginResetModel() {
	RunMain(func() {
		C.beginResetModel(this._addr)
	})
}

func (this *QmlModelBase) EndResetModel() {
	RunMain(func() {
		C.endResetModel(this._addr)
	})
}

func (this *QmlModelBase) BeginInsertRows(start, end int) {
	RunMain(func() {
		C.beginInsertRows(this._addr, C.int(start), C.int(end))
	})
}

func (this *QmlModelBase) EndInsertRows() {
	RunMain(func() {
		C.endInsertRows(this._addr)
	})
}

func (this *QmlModelBase) BeginRemoveRows(start, end int) {
	RunMain(func() {
		C.beginRemoveRows(this._addr, C.int(start), C.int(end))
	})
}

func (this *QmlModelBase) EndRemoveRows() {
	RunMain(func() {
		C.endRemoveRows(this._addr)
	})
}

func (this *QmlModelBase) DataChanged(start, end int) {
	RunMain(func() {
		C.modelDataChanged(this._addr, C.int(start), C.int(end))
	})
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

// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package foundation

import (
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
)

/*
#include <stdint.h>

// Note: these functions have a different signature but because they are only
// used as function pointers (and never called) and because they use C name
// mangling, the signature doesn't really matter.
void winrt_TypedEventHandler_Invoke(void);
void winrt_TypedEventHandler_QueryInterface(void);
uint64_t winrt_TypedEventHandler_AddRef(void);
uint64_t winrt_TypedEventHandler_Release(void);

// The Vtable structure for WinRT TypedEventHandler interfaces.
typedef struct {
	void *QueryInterface;
	void *AddRef;
	void *Release;
	void *Invoke;
} TypedEventHandlerVtbl_t;

// The Vtable itself. It can be kept constant.
static const TypedEventHandlerVtbl_t winrt_TypedEventHandlerVtbl = {
	(void*)winrt_TypedEventHandler_QueryInterface,
	(void*)winrt_TypedEventHandler_AddRef,
	(void*)winrt_TypedEventHandler_Release,
	(void*)winrt_TypedEventHandler_Invoke,
};

// A small helper function to get the Vtable.
const TypedEventHandlerVtbl_t * winrt_getTypedEventHandlerVtbl(void) {
	return &winrt_TypedEventHandlerVtbl;
}
*/
import "C"

const GUIDTypedEventHandler string = "9de1c534-6ae1-11e0-84e1-18a905bcc53f"
const SignatureTypedEventHandler string = "delegate({9de1c534-6ae1-11e0-84e1-18a905bcc53f})"

type TypedEventHandler struct {
	ole.IUnknown
	sync.Mutex
	refs uint64
	IID  ole.GUID
}

type TypedEventHandlerCallback func(instance *TypedEventHandler, sender unsafe.Pointer, args unsafe.Pointer)

var callbacksTypedEventHandler map[unsafe.Pointer]TypedEventHandlerCallback = make(map[unsafe.Pointer]TypedEventHandlerCallback)

func NewTypedEventHandler(iid *ole.GUID, callback TypedEventHandlerCallback) *TypedEventHandler {
	inst := (*TypedEventHandler)(C.malloc(C.size_t(unsafe.Sizeof(TypedEventHandler{}))))
	inst.RawVTable = (*interface{})((unsafe.Pointer)(C.winrt_getTypedEventHandlerVtbl()))
	inst.IID = *iid // copy contents

	callbacksTypedEventHandler[unsafe.Pointer(inst)] = callback

	inst.addRef()
	return inst
}

// addRef increments the reference counter by one
func (r *TypedEventHandler) addRef() uint64 {
	r.Lock()
	defer r.Unlock()
	r.refs++
	return r.refs
}

// removeRef decrements the reference counter by one. If it was already zero, it will just return zero.
func (r *TypedEventHandler) removeRef() uint64 {
	r.Lock()
	defer r.Unlock()

	if r.refs > 0 {
		r.refs--
	}

	return r.refs
}

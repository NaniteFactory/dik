package dik

import (
	"errors"
	"sync"
	"unsafe"
)

// ----------------------------------------------------------------------------
// Class KeyObserver

// KeyObserver is a joypad for a robot.
// I believe it is thread-safe.
type KeyObserver interface {
	// All advanced methods are here hidden. Optional stuff.
	Self() *KeyObserved

	// Hook
	Out(lpvData uintptr) // GetDeviceState() must call this function.

	// Read
	IsReleased(scanCode uint8) (ret bool, err error)
	IsProbablyPressed(scanCode uint8) (ret bool, err error) // IsPressed

	// Write
	PressKey(scanCode uint8)
	ReleaseKey(scanCode uint8)
}

// KeyObserved is DIKeys being watched and watches itself as it implements KeyObserver interface.
// To understand what it really does, refer to its essential methods In() and Out().
// I believe it is thread-safe.
type KeyObserved struct {
	actualKeys     *[256]uint8     // Where the 3rd argument of GetDeviceState() points to. // lazy
	availableKeys  map[uint8]uint8 // map[ScanCode]Value // dikeys observed
	nAvailableKeys int             // This member is there because accessing it can be cheaper than len(availableKeys).
	mutex          sync.Mutex      // A locker for thread safety. Beware of ptrd objects. You should never read/write a map at the same time.
}

// NewKeyObserver is a contructor and where you start out everything.
// Register dikeys to an observer, so those can be observed. xD
//
// e.g.
// var joypad = dik.NewKeyObserver(dik.Z, dik.X, dik.C, dik.V, dik.Left, dik.Right, dik.Up, dik.Down)
//
// Then call #Out from GetDeviceState() hook in order to manipulate those key states programmatically.
// Also note that you won't be able to use them in any normal way once you get that #Out() hook working.
func NewKeyObserver(scanCodes ...uint8) (itf KeyObserver) {
	getScanCodesNoDup := func(elements []uint8) []uint8 {
		// Use map to record duplicates as we find them.
		encountered := map[uint8]bool{}
		result := []uint8{}
		for v := range elements {
			if encountered[elements[v]] == true {
				// Do not add duplicate.
			} else {
				// Record this element as an encountered element.
				encountered[elements[v]] = true
				// Append to result slice.
				result = append(result, elements[v])
			}
		}
		// Return the new slice.
		return result
	} // a nice snippet of code found on the internet. xD

	// evaluate availableKeys from scan codes
	scanCodesNoDup := getScanCodesNoDup(scanCodes)
	keys := map[uint8]uint8{}
	for _, v := range scanCodesNoDup {
		keys[v] = KeyReleased
	}

	// Create our object to be returned.
	ko := KeyObserved{}
	// These two members will not be overwritten. And when if that ever happens, you need a mutex for that.
	ko.availableKeys = keys
	ko.nAvailableKeys = len(ko.availableKeys)
	return &ko
}

// ----------------------------------------------------------------------------
// Read

// Self () in order to convert an interface to a concrete struct.
func (ko *KeyObserved) Self() *KeyObserved {
	return ko
}

// NumberOfKeys returns the number (count) of keys this KeyObserved watches.
func (ko *KeyObserved) NumberOfKeys() int {
	return ko.nAvailableKeys
}

// IsAvailable determines whether a dikey of a scan code is being observed or not.
func (ko *KeyObserved) IsAvailable(scanCode uint8) bool {
	ko.mutex.Lock()
	defer ko.mutex.Unlock()

	if _, ok := ko.availableKeys[scanCode]; ok { // The map contains that.
		return true
	}
	return false
}

// StateRaw returns a pointer to the lowest-level direct-input-key states.
// That is where the 3rd argument of GetDeviceState() points to.
// Normally you would not want to use this method.
// This returns nil if #Out() hasn't been called yet.
func (ko *KeyObserved) StateRaw() (actualKeys *[256]uint8) {
	return actualKeys
}

// StateAll returns a copy of a map of keys being observed by this object.
// Notice that there is a cost for copying the whole map thing.
func (ko *KeyObserved) StateAll() (copyAvailableKeys map[uint8]uint8) {
	ko.mutex.Lock()
	defer ko.mutex.Unlock()

	copyAvailableKeys = map[uint8]uint8{}
	for k, v := range ko.availableKeys {
		copyAvailableKeys[k] = v
	}
	return copyAvailableKeys
}

// State is a getter that returns a state of a key.
// A state whether it's pressed or released.
// An error will be returned when the scan code is invalid.
func (ko *KeyObserved) State(scanCode uint8) (diKeyValue uint8, err error) {
	ko.mutex.Lock()
	defer ko.mutex.Unlock()
	if v, ok := ko.availableKeys[scanCode]; ok {
		return v, nil
	}
	return 0, errors.New("This observer doesn't look for that scan code")
}

// IsReleased determines whether a key of a scan code is being released or not.
// An error will be returned when the scan code is invalid.
func (ko *KeyObserved) IsReleased(scanCode uint8) (ret bool, err error) {
	val, err := ko.State(scanCode)
	if err != nil {
		return ret, err
	}
	if val == KeyReleased {
		return true, nil
	}
	return false, nil
}

// IsProbablyPressed performs a reverse operation of IsReleased(), which determines whether a key of a scan code is being released or not.
// An error will be returned when the scan code is invalid.
func (ko *KeyObserved) IsProbablyPressed(scanCode uint8) (ret bool, err error) {
	ret, err = ko.IsReleased(scanCode)
	return !ret, err
}

// ----------------------------------------------------------------------------
// Write

// In () is like a listener or a callback. This function is a controller (an interface) for a robot.
func (ko *KeyObserved) In(scanCode, value uint8) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		outputdbg.LogPrintln(r)
	// 	}
	// }()

	if ko.IsAvailable(scanCode) {
		ko.mutex.Lock()
		// You gotta get a lock because this could be an unhandleable concurrent map write.
		// When that happens it generates a fatal error that can't be recovered or catched any way.
		ko.availableKeys[scanCode] = value
		ko.mutex.Unlock()
	}
}

// Out () is a kind of OnGetDeviceState(). Call this function upon GetDeviceState() hook getting called.
// This method is like a plug out of a game controller, when the receiver is seen as a joypad.
// Parameter lpvData of this function is the 3rd argument of GetDeviceState().
func (ko *KeyObserved) Out(lpvData uintptr) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		outputdbg.LogPrintln(r)
	// 	}
	// }()

	// our manipulation
	ko.mutex.Lock() // lock up your daughter
	ko.actualKeys = (*[256]uint8)(unsafe.Pointer(lpvData))
	for scanCode, diKeyValue := range ko.availableKeys {
		ko.actualKeys[scanCode] = diKeyValue
	}
	ko.mutex.Unlock() // the man is back in town

	// ko.mutex.Lock()
	// outputdbg.LogPrintln(ko.availableKeys) //
	// ko.mutex.Unlock()
}

// PressKey of a scan code.
func (ko *KeyObserved) PressKey(scanCode uint8) {
	ko.In(scanCode, KeyPressed)
}

// ReleaseKey of a scan code.
func (ko *KeyObserved) ReleaseKey(scanCode uint8) {
	ko.In(scanCode, KeyReleased)
}

// ----------------------------------------------------------------------------

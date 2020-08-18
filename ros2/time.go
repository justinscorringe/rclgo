package ros2

import (
	"time"
	gotime "time"
)

const maxUint32 = int64(^uint32(0))

// ROS TIME IMPLEMENTATION

//Time struct contains a temporal value {sec,nsec}
type Time struct {
	temporal
}

//NewTime creates a Time object of given integers {sec,nsec}
func NewTime(sec uint32, nsec uint32) Time {
	sec, nsec = normalizeTemporal(int64(sec), int64(nsec))
	return Time{temporal{sec, nsec}}
}

//Now creates a Time object of value Now
func Now() Time {
	var t Time
	t.FromNSec(uint64(gotime.Now().UnixNano()))
	return t
}

//Diff returns difference of two Time objects as a Duration
func (t *Time) Diff(from Time) Duration {
	sec, nsec := normalizeTemporal(int64(t.Sec)-int64(from.Sec),
		int64(t.NSec)-int64(from.NSec))
	return Duration{temporal{sec, nsec}}
}

//Add returns sum of Time and Duration given
func (t *Time) Add(d Duration) Time {
	sec, nsec := normalizeTemporal(int64(t.Sec)+int64(d.Sec),
		int64(t.NSec)+int64(d.NSec))
	return Time{temporal{sec, nsec}}
}

//Sub returns subtraction of Time and Duration given
func (t *Time) Sub(d Duration) Time {
	sec, nsec := normalizeTemporal(int64(t.Sec)-int64(d.Sec),
		int64(t.NSec)-int64(d.NSec))
	return Time{temporal{sec, nsec}}
}

//Cmp returns int comparison of two Time objects
func (t *Time) Cmp(other Time) int {
	return cmpUint64(t.ToNSec(), other.ToNSec())
}

// ROS DURATION IMPLEMENTATION

//Duration type which is a wrapper for a temporal value of {sec,nsec}
type Duration struct {
	temporal
}

//NewDuration instantiates a new Duration item with given sec and nsec integers
func NewDuration(sec uint32, nsec uint32) Duration {
	sec, nsec = normalizeTemporal(int64(sec), int64(nsec))
	return Duration{temporal{sec, nsec}}
}

//Add function for adding two durations together
func (d *Duration) Add(other Duration) Duration {
	sec, nsec := normalizeTemporal(int64(d.Sec)+int64(other.Sec),
		int64(d.NSec)+int64(other.NSec))
	return Duration{temporal{sec, nsec}}
}

//Sub function for subtracting a duration from another
func (d *Duration) Sub(other Duration) Duration {
	sec, nsec := normalizeTemporal(int64(d.Sec)-int64(other.Sec),
		int64(d.NSec)-int64(other.NSec))
	return Duration{temporal{sec, nsec}}
}

//Cmp function to compare two durations
func (d *Duration) Cmp(other Duration) int {
	return cmpUint64(d.ToNSec(), other.ToNSec())
}

//Sleep function pauses go routine for duration d
func (d *Duration) Sleep() error {
	if !d.IsZero() {
		time.Sleep(time.Duration(d.ToNSec()) * time.Nanosecond)
	}
	return nil
}

// ROS TEMPORAL IMPLEMENTATION

func normalizeTemporal(sec int64, nsec int64) (uint32, uint32) {
	const SecondInNanosecond = 1000000000
	if nsec > SecondInNanosecond {
		sec += nsec / SecondInNanosecond
		nsec = nsec % SecondInNanosecond
	} else if nsec < 0 {
		sec += nsec/SecondInNanosecond - 1
		nsec = nsec%SecondInNanosecond + SecondInNanosecond
	}

	if sec < 0 || sec > maxUint32 {
		panic("Time is out of range")
	}

	return uint32(sec), uint32(nsec)
}

func cmpUint64(lhs, rhs uint64) int {
	var result int
	if lhs > rhs {
		result = 1
	} else if lhs < rhs {
		result = -1
	} else {
		result = 0
	}
	return result
}

type temporal struct {
	Sec  uint32
	NSec uint32
}

func (t *temporal) IsZero() bool {
	return t.Sec == 0 && t.NSec == 0
}

func (t *temporal) ToSec() float64 {
	return float64(t.Sec) + float64(t.NSec)*1e-9
}

func (t *temporal) ToNSec() uint64 {
	return uint64(t.Sec)*1000000000 + uint64(t.NSec)
}

func (t *temporal) FromSec(sec float64) {
	nsec := uint64(sec * 1e9)
	t.FromNSec(nsec)
}

func (t *temporal) FromNSec(nsec uint64) {
	t.Sec, t.NSec = normalizeTemporal(0, int64(nsec))
}

func (t *temporal) Normalize() {
	t.Sec, t.NSec = normalizeTemporal(int64(t.Sec), int64(t.NSec))
}

package fronius

import "fmt"

type Current int64

func (c Current) String() string {
	return fmt.Sprintf("%vA", float64(c)/float64(Ampere))
}

const (
	Nanoampere  Current = 1
	Microampere         = 1000 * Nanoampere
	Milliampere         = 1000 * Microampere
	Ampere              = 1000 * Milliampere
	Kiloampere          = 1000 * Ampere
)

type Voltage int64

func (c Voltage) String() string {
	return fmt.Sprintf("%vV", float64(c)/float64(Volt))
}

const (
	Nanovolt  Voltage = 1
	Microvolt         = 1000 * Nanovolt
	Millivolt         = 1000 * Microvolt
	Volt              = 1000 * Millivolt
	Kilovolt          = 1000 * Volt
)

type Effect int64

const (
	Nanowatt  Effect = 1
	Microwatt        = 1000 * Nanowatt
	Milliwatt        = 1000 * Microwatt
	Watt             = 1000 * Milliwatt
	Kilowatt         = 1000 * Watt
)

type Energi int64

const (
	Nanowatthour  Energi = 1
	Microwatthour        = 1000 * Nanowatthour
	Milliwatthour        = 1000 * Microwatthour
	Watthour             = 1000 * Milliwatthour
	Kilowatthour         = 1000 * Watthour
	Megewatthour         = 1000 * Kilowatthour
	Gigawatthour         = 1000 * Megewatthour
)

type Frequency int64

const (
	Nanohertz  Frequency = 1
	Microhertz           = 1000 * Nanohertz
	Millihertz           = 1000 * Microhertz
	Hertz                = 1000 * Millihertz
	Kilohertz            = 1000 * Hertz
)

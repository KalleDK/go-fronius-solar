package fronius

import "fmt"

type Field struct {
	Unit  string
	Value float64
}

func (f Field) ToCurrent() (Current, error) {
	switch f.Unit {
	case "A":
		return Current(f.Value * float64(Ampere)), nil
	}
	return 0, fmt.Errorf("invalid current unit %v %v", f.Unit, f.Value)
}

func (f Field) ToVoltage() (Voltage, error) {
	switch f.Unit {
	case "V":
		return Voltage(f.Value * float64(Volt)), nil
	}
	return 0, fmt.Errorf("invalid voltage unit %v %v", f.Unit, f.Value)
}

func (f Field) ToEnergi() (Energi, error) {
	switch f.Unit {
	case "Wh":
		return Energi(f.Value * float64(Watthour)), nil
	}
	return 0, fmt.Errorf("invalid energi unit %v %v", f.Unit, f.Value)
}

func (f Field) ToEffect() (Effect, error) {
	switch f.Unit {
	case "W":
		return Effect(f.Value * float64(Watt)), nil
	}
	return 0, fmt.Errorf("invalid effect unit %v %v", f.Unit, f.Value)
}

func (f Field) ToFrequency() (Frequency, error) {
	switch f.Unit {
	case "Hz":
		return Frequency(f.Value * float64(Hertz)), nil
	}
	return 0, fmt.Errorf("invalid frequency unit %v %v", f.Unit, f.Value)
}

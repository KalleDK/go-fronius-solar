package fronius

import (
	"encoding/json"
	"time"
)

type Phase struct {
	Current Current
	Voltage Voltage
}

type Head struct {
	RequestArguments map[string]string
	Status           struct {
		Code        int
		Reason      string
		UserMessage string
	}
	Timestamp time.Time
}

type InverterDataPacket struct {
	Body struct {
		Data InverterData
	}
	Head Head
}

type InverterData struct {
	Phase1 Phase
	Phase2 Phase
	Phase3 Phase
}

func (d *InverterData) UnmarshalJSON(p []byte) (err error) {
	var i Current
	var u Voltage

	fields := map[string]Field{}

	if err = json.Unmarshal(p, &fields); err != nil {
		return err
	}

	if i, err = fields["IAC_L1"].ToCurrent(); err != nil {
		return err
	}

	if u, err = fields["UAC_L1"].ToVoltage(); err != nil {
		return err
	}

	d.Phase1 = Phase{
		Current: i,
		Voltage: u,
	}

	if i, err = fields["IAC_L2"].ToCurrent(); err != nil {
		return err
	}

	if u, err = fields["UAC_L2"].ToVoltage(); err != nil {
		return err
	}

	d.Phase2 = Phase{
		Current: i,
		Voltage: u,
	}

	if i, err = fields["IAC_L3"].ToCurrent(); err != nil {
		return err
	}

	if u, err = fields["UAC_L3"].ToVoltage(); err != nil {
		return err
	}

	d.Phase3 = Phase{
		Current: i,
		Voltage: u,
	}
	return nil
}

type CommonInverterDataPacket struct {
	Body struct {
		Data CommonInverterData
	}
	Head Head
}

type CommonInverterData struct {
	DayEnergy   Energi
	YearEnergy  Energi
	TotalEnergy Energi
	Frequency   Frequency
	ACCurrent   Current
	ACVoltage   Voltage
	ACEffect    Effect
	DCCurrent   Current
	DCVoltage   Voltage
}

func (d *CommonInverterData) UnmarshalJSON(p []byte) (err error) {
	fields := map[string]Field{}

	if err = json.Unmarshal(p, &fields); err != nil {
		return err
	}

	if d.Frequency, err = fields["FAC"].ToFrequency(); err != nil {
		return err
	}

	if d.DayEnergy, err = fields["DAY_ENERGY"].ToEnergi(); err != nil {
		return err
	}

	if d.YearEnergy, err = fields["YEAR_ENERGY"].ToEnergi(); err != nil {
		return err
	}

	if d.TotalEnergy, err = fields["TOTAL_ENERGY"].ToEnergi(); err != nil {
		return err
	}

	if d.ACCurrent, err = fields["IAC"].ToCurrent(); err != nil {
		return err
	}

	if d.ACVoltage, err = fields["UAC"].ToVoltage(); err != nil {
		return err
	}

	if d.ACEffect, err = fields["PAC"].ToEffect(); err != nil {
		return err
	}

	if d.DCCurrent, err = fields["IDC"].ToCurrent(); err != nil {
		return err
	}

	if d.DCVoltage, err = fields["UDC"].ToVoltage(); err != nil {
		return err
	}

	return nil
}

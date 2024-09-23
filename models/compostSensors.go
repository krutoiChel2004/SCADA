package models

// ContainerData содержит данные для всех измерений одного контейнера с типами float32
type CompostData struct {
	Container1TempTop     float32 `db:"container_1_temp_top"`
	Container1TempBottom  float32 `db:"container_1_temp_bottom"`
	Container2TempTop     float32 `db:"container_2_temp_top"`
	Container2TempBottom  float32 `db:"container_2_temp_bottom"`
	Container3TempTop     float32 `db:"container_3_temp_top"`
	Container3TempBottom  float32 `db:"container_3_temp_bottom"`
	Container4TempTop     float32 `db:"container_4_temp_top"`
	Container4TempBottom  float32 `db:"container_4_temp_bottom"`
	Container5TempTop     float32 `db:"container_5_temp_top"`
	Container5TempBottom  float32 `db:"container_5_temp_bottom"`
	Container6TempTop     float32 `db:"container_6_temp_top"`
	Container6TempBottom  float32 `db:"container_6_temp_bottom"`
	Container7TempTop     float32 `db:"container_7_temp_top"`
	Container7TempBottom  float32 `db:"container_7_temp_bottom"`
	Container8TempTop     float32 `db:"container_8_temp_top"`
	Container8TempBottom  float32 `db:"container_8_temp_bottom"`
	Container9TempTop     float32 `db:"container_9_temp_top"`
	Container9TempBottom  float32 `db:"container_9_temp_bottom"`
	Container1PressureIn  float32 `db:"container_1_pressure_in"`
	Container1PressureOut float32 `db:"container_1_pressure_out"`
	Container2PressureIn  float32 `db:"container_2_pressure_in"`
	Container2PressureOut float32 `db:"container_2_pressure_out"`
	Container3PressureIn  float32 `db:"container_3_pressure_in"`
	Container3PressureOut float32 `db:"container_3_pressure_out"`
	Container4PressureIn  float32 `db:"container_4_pressure_in"`
	Container4PressureOut float32 `db:"container_4_pressure_out"`
	Container5PressureIn  float32 `db:"container_5_pressure_in"`
	Container5PressureOut float32 `db:"container_5_pressure_out"`
	Container6PressureIn  float32 `db:"container_6_pressure_in"`
	Container6PressureOut float32 `db:"container_6_pressure_out"`
	Container7PressureIn  float32 `db:"container_7_pressure_in"`
	Container7PressureOut float32 `db:"container_7_pressure_out"`
	Container8PressureIn  float32 `db:"container_8_pressure_in"`
	Container8PressureOut float32 `db:"container_8_pressure_out"`
	Container9PressureIn  float32 `db:"container_9_pressure_in"`
	Container9PressureOut float32 `db:"container_9_pressure_out"`
	Container1HeaterTemp  float32 `db:"container_1_heater_temp"`
	Container2HeaterTemp  float32 `db:"container_2_heater_temp"`
	Container3HeaterTemp  float32 `db:"container_3_heater_temp"`
	Container4HeaterTemp  float32 `db:"container_4_heater_temp"`
	Container5HeaterTemp  float32 `db:"container_5_heater_temp"`
	Container6HeaterTemp  float32 `db:"container_6_heater_temp"`
	Container7HeaterTemp  float32 `db:"container_7_heater_temp"`
	Container8HeaterTemp  float32 `db:"container_8_heater_temp"`
	Container9HeaterTemp  float32 `db:"container_9_heater_temp"`
	OutsideTemp           float32 `db:"outside_temp"`
	ControlUnitTemp       float32 `db:"control_unit_temp"`
}

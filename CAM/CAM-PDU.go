package CAM

import "github.com/pprizbul/go-ITS/protocols/CDD"

type CAM struct {
	Header CDD.ItsPduHeader
	Cam    CamPayload
}

type CamPayload struct {
	GenerationDeltaTime CDD.GenerationDeltaTime
	CamParameters       CamParameters
}

type CamParameters struct {
	BasicContainer         CDD.BasicContainer
	HighFrequencyContainer HighFrequencyContainer
	//LowFrequencyContainer    LowFrequencyContainer OPTIONAL
	//SpecialVehicleContainer  SpecialVehicleContainer OPTIONAL
}

type HighFrequencyContainer struct { //CHOICE
	BasicVehicleContainerHighFrequency BasicVehicleContainerHighFrequency
	//RsuContainerHighFrequency          RSUContainerHighFrequency
}

//LowFrequencyContainer

//SpecialVehicleContainer

type BasicVehicleContainerHighFrequency struct {
	Heading                  CDD.Heading
	Speed                    CDD.Speed
	DriveDirection           CDD.DriveDirection
	VehicleLength            CDD.VehicleLength
	VehicleWidth             CDD.VehicleWidth
	LongitudinalAcceleration CDD.AccelerationComponent
	Curvature                CDD.Curvature
	CurvatureCalculationMode CDD.CurvatureCalculationMode
	YawRate                  CDD.YawRate
	//accelerationControl AccelerationControl OPTIONAL
	//lanePosition LanePosition OPTIONAL
	//steeringWheelAngle SteeringWheelAngle OPTIONAL
	//lateralAcceleration AccelerationComponent OPTIONAL
	//verticalAcceleration AccelerationComponent OPTIONAL
	//performanceClass PerformanceClass OPTIONAL
	//cenDsrcTollingZone CenDsrcTollingZone OPTIONAL
}

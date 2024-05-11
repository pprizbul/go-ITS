/*
Common data ditctionary ASN implementation (CDD.asn), ETSI standard TS 102 894-2 V2.2.1
Not all field are implemented, just fields that CAM protocol use. Order of implementations are
same as in ASN file.

ENUM, const, etc.
*/

package CDD

type AccelerationConfidence int //51
const (
	AccelerationConfidence_outOfRange  AccelerationConfidence = 101
	AccelerationConfidence_unavailable AccelerationConfidence = 102
)

type AccelerationValue int //122
const (
	AccelerationValue_negativeOutOfRange AccelerationValue = -160
	AccelerationValue_positiveOutOfRange AccelerationValue = 160
	AccelerationValue_unavailable        AccelerationValue = 161
)

type AltitudeConfidence int //341 Enum
const (
	AltitudeConfidence_alt_000_01 AltitudeConfidence = iota
	AltitudeConfidence_alt_000_02
	AltitudeConfidence_alt_000_05
	AltitudeConfidence_alt_000_10
	AltitudeConfidence_alt_000_20
	AltitudeConfidence_alt_000_50
	AltitudeConfidence_alt_001_00
	AltitudeConfidence_alt_002_00
	AltitudeConfidence_alt_005_00
	AltitudeConfidence_alt_010_00
	AltitudeConfidence_alt_020_00
	AltitudeConfidence_alt_050_00
	AltitudeConfidence_alt_100_00
	AltitudeConfidence_alt_200_00
	AltitudeConfidence_outOfRange
	AltitudeConfidence_unavailable
)

type AltitudeValue int //375
const (
	AltitudeValue_negativeOutOfRange AltitudeValue = -100000
	AltitudeValue_postiveOutOfRange  AltitudeValue = 800000 //misspell is according to standard
	AltitudeValue_unavailable        AltitudeValue = 800001
)

type CurvatureCalculationMode int //887 Enum
const (
	CurvatureCalculationMode_yawRateUsed CurvatureCalculationMode = iota
	CurvatureCalculationMode_yawRateNotUsed
	CurvatureCalculationMode_unavailable
)

type CurvatureConfidence int //921 Enum
const (
	CurvatureConfidence_onePerMeter_0_00002 CurvatureConfidence = iota
	CurvatureConfidence_onePerMeter_0_0001
	CurvatureConfidence_onePerMeter_0_0005
	CurvatureConfidence_onePerMeter_0_002
	CurvatureConfidence_onePerMeter_0_01
	CurvatureConfidence_onePerMeter_0_1
	CurvatureConfidence_outOfRange
	CurvatureConfidence_unavailable
)

type CurvatureValue int //956
const (
	CurvatureValue_outOfRangeNegative CurvatureValue = -1023
	CurvatureValue_straight           CurvatureValue = 0
	CurvatureValue_outOfRangePositive CurvatureValue = 1022
	CurvatureValue_unavailable        CurvatureValue = 1023
)

type DriveDirection int //1216 Enum
const (
	DriveDirection_forward DriveDirection = iota
	DriveDirection_backward
	DriveDirection_unavailable
)

type GenerationDeltaTime int //1394

type HeadingConfidence int

const (
	HeadingConfidence_outOfRange  HeadingConfidence = 126
	HeadingConfidence_unavailable HeadingConfidence = 127
)

type HeadingValue int //1573
const (
	HeadingValue_wgs84North  HeadingValue = 0
	HeadingValue_wgs84East   HeadingValue = 900
	HeadingValue_wgs84South  HeadingValue = 1800
	HeadingValue_wgs84West   HeadingValue = 2700
	HeadingValue_doNotUse    HeadingValue = 3600
	HeadingValue_unavailable HeadingValue = 3601
)

type Latitude int //2016
const (
	Latitude_unavailable Latitude = 900000001
)

type Longitude int //2081
const (
	Longitude_valueNotUsed Longitude = -1800000000
	Longitude_unavailable  Longitude = 1800000001
)

type MessageId int //2215
const (
	Denm MessageId = iota + 1
	Cam
	poim
	spatem
	mapem
	ivim
	rfu1
	rfu2
	srem
	ssem
	evcsn
	saem
	rtcmem
	cpm
	imzm
	vam
	dsm
	pcim
	pcvm
	mcm
	pam
)

type OrdinalNumber1B int //2345

type SemiAxisLength int //2936
const (
	SemiAxisLength_doNotUse    SemiAxisLength = 0
	SemiAxisLength_outOfRange  SemiAxisLength = 4094
	SemiAxisLength_unavailable SemiAxisLength = 4095
)

type SpeedConfidence int //3125
const (
	SpeedConfidence_outOfRange  SpeedConfidence = 126
	SpeedConfidence_unavailable SpeedConfidence = 127
)

type SpeedValue int //3154
const (
	SpeedValue_standstill  SpeedValue = 0
	SpeedValue_outOfRange  SpeedValue = 16382
	SpeedValue_unavailable SpeedValue = 16383
)

type StationId int //3340

type SteeringWheelAngleValue int //3445
const (
	SteeringWheelAngleValue_negativeOutOfRange = -511
	SteeringWheelAngleValue_positiveOutOfRange = 511
	SteeringWheelAngleValue_unavailable        = 512
)

type TrafficParticipantType int //3576
const (
	Unknown TrafficParticipantType = iota
	Pedestrian
	Cyclist
	Moped
	Motorcycle
	PassengerCar
	Bus
	LightTruck
	HeavyTruck
	Trailer
	SpecialVehicle
	Tram
	LightVruVehicle
	Animal
	Agricultural
	RoadSideUnit
)

type VehicleLengthConfidenceIndication int //3825 Enum
const (
	VehicleLengthConfidenceIndication_noTrailerPresent VehicleLengthConfidenceIndication = iota
	VehicleLengthConfidenceIndication_trailerPresentWithKnownLength
	VehicleLengthConfidenceIndication_trailerPresentWithUnknownLength
	VehicleLengthConfidenceIndication_trailerPresenceIsUnknown
	VehicleLengthConfidenceIndication_unavailable
)

type VehicleLengthValue int //3846
const (
	VehicleLengthValue_outOfRange  VehicleLengthValue = 1022
	VehicleLengthValue_unavailable VehicleLengthValue = 1023
)

type VehicleWidth int //3925
const (
	VehicleWidth_outOfRange  VehicleWidth = 61
	VehicleWidth_unavailable VehicleWidth = 62
)

type Wgs84AngleValue int //4246
const (
	Wgs84AngleValue_wgs84North  Wgs84AngleValue = 0
	Wgs84AngleValue_wgs84East   Wgs84AngleValue = 900
	Wgs84AngleValue_wgs84South  Wgs84AngleValue = 1800
	Wgs84AngleValue_wgs84West   Wgs84AngleValue = 2700
	Wgs84AngleValue_doNotUse    Wgs84AngleValue = 3600
	Wgs84AngleValue_unavailable Wgs84AngleValue = 3601
)

type YawRateConfidence int //4310 Enum
const (
	YawRateConfidence_degSec_000_01 YawRateConfidence = iota
	YawRateConfidence_degSec_000_05
	YawRateConfidence_degSec_000_10
	YawRateConfidence_degSec_001_00
	YawRateConfidence_degSec_005_00
	YawRateConfidence_degSec_010_00
	YawRateConfidence_degSec_100_00
	YawRateConfidence_outOfRange
	YawRateConfidence_unavailable
)

type YawRateValue int //4346
const (
	YawRateValue_negativeOutOfRange YawRateValue = -32766
	YawRateValue_positiveOutOfRange YawRateValue = 32766
	YawRateValue_unavailable        YawRateValue = 32767
)

type AccelerationComponent struct { //4425
	Value      AccelerationValue
	Confidence AccelerationConfidence
}

type Altitude struct { //4522
	AltitudeValue      AltitudeValue
	AltitudeConfidence AltitudeConfidence
}

type BasicContainer struct { //4539
	StationType       TrafficParticipantType
	ReferencePosition ReferencePositionWithConfidence
}

type Curvature struct {
	CurvatureValue      CurvatureValue
	CurvatureConfidence CurvatureConfidence
}

type Heading struct { //5344
	HeadingValue      HeadingValue
	HeadingConfidence HeadingConfidence
}

type ItsPduHeader struct { //5518
	ProtocolVersion OrdinalNumber1B
	MessageId       MessageId
	StationId       StationId
}

type PositionConfidenceEllipse struct { //6318
	SemiMajorAxisLength      SemiAxisLength
	SemiMinorAxisLength      SemiAxisLength
	SemiMajorAxisOrientation Wgs84AngleValue
}

type ReferencePositionWithConfidence struct { //6619
	Latitude                  Latitude
	Longitude                 Longitude
	PositionConfidenceEllipse PositionConfidenceEllipse
	Altitude                  Altitude
}

type Speed struct { //6822
	SpeedValue      SpeedValue
	SpeedConfidence SpeedConfidence
}

type VehicleLength struct { //7029
	VehicleLengthValue                VehicleLengthValue
	VehicleLengthConfidenceIndication VehicleLengthConfidenceIndication
}

type YawRate struct { //7224
	YawRateValue      YawRateValue
	YawRateConfidence YawRateConfidence
}

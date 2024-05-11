package CVM

import "github.com/pprizbul/go-ITS/CDD"

const Cvm CDD.MessageId = 45 //MessageId for CVM

type CVM struct {
	Header CDD.ItsPduHeader
	Cvm    CvmPayload
}

type CvmPayload struct {
	Direction     CDD.DriveDirection
	Speed         CDD.SpeedValue
	Acceleration  CDD.AccelerationValue
	SteeringAngle CDD.SteeringWheelAngleValue
}

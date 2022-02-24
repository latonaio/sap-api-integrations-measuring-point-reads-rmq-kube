package sap_api_output_formatter

type MeasuringPoint struct {
	 ConnectionKey    string `json:"connection_key"`
	 Result           bool   `json:"result"`
	 RedisKey         string `json:"redis_key"`
	 Filepath         string `json:"filepath"`
	 APISchema        string `json:"api_schema"`
	 MeasuringPoint     string `json:"planned_order"`
	 Deleted          bool   `json:"deleted"`
}

type Header struct {
	MeasuringPoint                 string        `json:"MeasuringPoint"`
	MeasuringPointDescription      string        `json:"MeasuringPointDescription"`
	MeasuringPointObjectIdentifier string        `json:"MeasuringPointObjectIdentifier"`
	TechnicalObjectType            string        `json:"TechnicalObjectType"`
	MeasuringPointPositionNumber   string        `json:"MeasuringPointPositionNumber"`
	MeasuringPointCategory         string        `json:"MeasuringPointCategory"`
	CreationDate                   string        `json:"CreationDate"`
	LastChangeDate                 string        `json:"LastChangeDate"`
	MeasuringPointIsCounter        bool          `json:"MeasuringPointIsCounter"`
	MsrgPtInternalCharacteristic   string        `json:"MsrgPtInternalCharacteristic"`
	CharcValueUnit                 string        `json:"CharcValueUnit"`
	MeasuringPointDecimalPlaces    int64         `json:"MeasuringPointDecimalPlaces"`
	MeasuringPointExponent         int64         `json:"MeasuringPointExponent"`
	MeasuringPointCodeGroup        string        `json:"MeasuringPointCodeGroup"`
	ValuationCodeIsSufficient      bool          `json:"ValuationCodeIsSufficient"`
	Assembly                       string        `json:"Assembly"`
	MeasuringPointIsInactive       bool          `json:"MeasuringPointIsInactive"`
	MeasuringPointShortText        string        `json:"MeasuringPointShortText"`
	MeasurementRangeUnit           string        `json:"MeasurementRangeUnit"`
	MsmtRdngSourceMeasuringPoint   string        `json:"MsmtRdngSourceMeasuringPoint"`
	MeasuringPointTargetValue      int64         `json:"MeasuringPointTargetValue"`
	MeasuringPointMaximumThreshold float64       `json:"MeasuringPointMaximumThreshold"`
	MeasuringPointMinimumThreshold int64         `json:"MeasuringPointMinimumThreshold"`
	MeasuringPointAnnualEstimate   float64       `json:"MeasuringPointAnnualEstimate"`
	CounterOverflowRdngThreshold   int64         `json:"CounterOverflowRdngThreshold"`
	MsrgPtIsCountingBackwards      bool          `json:"MsrgPtIsCountingBackwards"`
	MeasurementTransferIsSupported bool          `json:"MeasurementTransferIsSupported"`
	FunctionalLocation             string        `json:"FunctionalLocation"`
	Equipment                      string        `json:"Equipment"`
	MsmtRdngTransferMode           string        `json:"MsmtRdngTransferMode"`
}

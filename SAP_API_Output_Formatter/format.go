package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-measuring-point-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		header = append(header, Header{
			MeasuringPoint:                 data.MeasuringPoint,
			MeasuringPointDescription:      data.MeasuringPointDescription,
			MeasuringPointObjectIdentifier: data.MeasuringPointObjectIdentifier,
			TechnicalObjectType:            data.TechnicalObjectType,
			MeasuringPointPositionNumber:   data.MeasuringPointPositionNumber,
			MeasuringPointCategory:         data.MeasuringPointCategory,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			MeasuringPointIsCounter:        data.MeasuringPointIsCounter,
			MsrgPtInternalCharacteristic:   data.MsrgPtInternalCharacteristic,
			CharcValueUnit:                 data.CharcValueUnit,
			MeasuringPointDecimalPlaces:    data.MeasuringPointDecimalPlaces,
			MeasuringPointExponent:         data.MeasuringPointExponent,
			MeasuringPointCodeGroup:        data.MeasuringPointCodeGroup,
			ValuationCodeIsSufficient:      data.ValuationCodeIsSufficient,
			Assembly:                       data.Assembly,
			MeasuringPointIsInactive:       data.MeasuringPointIsInactive,
			MeasuringPointShortText:        data.MeasuringPointShortText,
			MeasurementRangeUnit:           data.MeasurementRangeUnit,
			MsmtRdngSourceMeasuringPoint:   data.MsmtRdngSourceMeasuringPoint,
			MeasuringPointTargetValue:      data.MeasuringPointTargetValue,
			MeasuringPointMaximumThreshold: data.MeasuringPointMaximumThreshold,
			MeasuringPointMinimumThreshold: data.MeasuringPointMinimumThreshold,
			MeasuringPointAnnualEstimate:   data.MeasuringPointAnnualEstimate,
			CounterOverflowRdngThreshold:   data.CounterOverflowRdngThreshold,
			MsrgPtIsCountingBackwards:      data.MsrgPtIsCountingBackwards,
			MeasurementTransferIsSupported: data.MeasurementTransferIsSupported,
			FunctionalLocation:             data.FunctionalLocation,
			Equipment:                      data.Equipment,
			MsmtRdngTransferMode:           data.MsmtRdngTransferMode,
		})
	}

	return header, nil
}

package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-production-order-confirmation-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		ProductionOrder:                          data.ProductionOrder,
		ProductionOrderItem:                      data.ProductionOrderItem,
		Operations:                               data.Operations,
		OperationsItem:                           data.OperationsItem,
		OperationID:                              data.OperationID,
		ConfirmationCountingID:                   data.ConfirmationCountingID,
		OperationPlannedQuantityInBaseUnit:       *data.OperationPlannedQuantityInBaseUnit,
		OperationPlannedQuantityInProductionUnit: *data.OperationPlannedQuantityInProductionUnit,
		OperationPlannedQuantityInOperationUnit:  *data.OperationPlannedQuantityInOperationUnit,
		ProductBaseUnit:                          *data.ProductBaseUnit,
		ProductProductionUnit:                    *data.ProductProductionUnit,
		ProductOperationUnit:                     *data.ProductOperationUnit,
		OperationPlannedScrapInPercent:           data.OperationPlannedScrapInPercent,
		ConfirmationEntryDate:                    data.ConfirmationEntryDate,
		ConfirmationEntryTime:                    data.ConfirmationEntryTime,
		ConfirmationText:                         data.ConfirmationText,
		IsFinalConfirmation:                      data.IsFinalConfirmation,
		WorkCenter:                               *data.WorkCenter,
		EmployeeIDWhoConfirmed:                   *data.EmployeeIDWhoConfirmed,
		ConfirmedExecutionStartDate:              data.ConfirmedExecutionStartDate,
		ConfirmedExecutionStartTime:              data.ConfirmedExecutionStartTime,
		ConfirmedSetupStartDate:                  data.ConfirmedSetupStartDate,
		ConfirmedSetupStartTime:                  data.ConfirmedSetupStartTime,
		ConfirmedProcessingStartDate:             data.ConfirmedProcessingStartDate,
		ConfirmedProcessingStartTime:             data.ConfirmedProcessingStartTime,
		ConfirmedExecutionEndDate:                data.ConfirmedExecutionEndDate,
		ConfirmedExecutionEndTime:                data.ConfirmedExecutionEndTime,
		ConfirmedSetupEndDate:                    data.ConfirmedSetupEndDate,
		ConfirmedSetupEndTime:                    data.ConfirmedSetupEndTime,
		ConfirmedProcessingEndDate:               data.ConfirmedProcessingEndDate,
		ConfirmedProcessingEndTime:               data.ConfirmedProcessingEndTime,
		ConfirmedWaitDuration:                    data.ConfirmedWaitDuration,
		WaitDurationUnit:                         data.WaitDurationUnit,
		ConfirmedQueueDuration:                   data.ConfirmedQueueDuration,
		QueueDurationUnit:                        data.QueueDurationUnit,
		ConfirmedMoveDuration:                    data.ConfirmedMoveDuration,
		MoveDurationUnit:                         data.MoveDurationUnit,
		ConfirmedYieldQuantity:                   data.ConfirmedYieldQuantity,
		ConfirmedScrapQuantity:                   data.ConfirmedScrapQuantity,
		OperationVarianceReason:                  data.OperationVarianceReason,
		CreationDate:                             *data.CreationDate,
		CreationTime:                             *data.CreationTime,
		LastChangeDate:                           *data.LastChangeDate,
		LastChangeTime:                           *data.LastChangeTime,
		IsCancelled:                              data.IsCancelled,
	}
}

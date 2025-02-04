// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package scsi

// Generated from Table F.1
const (
	NoAdditionalSenseInformation                            AdditionalSenseCode = 0x0000
	FilemarkDetected                                        AdditionalSenseCode = 0x0001
	EndOfPartitionmediumDetected                            AdditionalSenseCode = 0x0002
	SetmarkDetected                                         AdditionalSenseCode = 0x0003
	BeginningOfPartitionmediumDetected                      AdditionalSenseCode = 0x0004
	EndOfDataDetected                                       AdditionalSenseCode = 0x0005
	IoProcessTerminated                                     AdditionalSenseCode = 0x0006
	ProgrammableEarlyWarningDetected                        AdditionalSenseCode = 0x0007
	AudioPlayOperationInProgress                            AdditionalSenseCode = 0x0011
	AudioPlayOperationPaused                                AdditionalSenseCode = 0x0012
	AudioPlayOperationSuccessfullyCompleted                 AdditionalSenseCode = 0x0013
	AudioPlayOperationStoppedDueToError                     AdditionalSenseCode = 0x0014
	NoCurrentAudioStatusToReturn                            AdditionalSenseCode = 0x0015
	OperationInProgress                                     AdditionalSenseCode = 0x0016
	CleaningRequested                                       AdditionalSenseCode = 0x0017
	EraseOperationInProgress                                AdditionalSenseCode = 0x0018
	LocateOperationInProgress                               AdditionalSenseCode = 0x0019
	RewindOperationInProgress                               AdditionalSenseCode = 0x001a
	SetCapacityOperationInProgress                          AdditionalSenseCode = 0x001b
	VerifyOperationInProgress                               AdditionalSenseCode = 0x001c
	AtaPassThroughInformationAvailable                      AdditionalSenseCode = 0x001d
	ConflictingSaCreationRequest                            AdditionalSenseCode = 0x001e
	LogicalUnitTransitioningToAnotherPowerCondition         AdditionalSenseCode = 0x001f
	ExtendedCopyInformationAvailable                        AdditionalSenseCode = 0x0020
	AtomicCommandAbortedDueToAca                            AdditionalSenseCode = 0x0021
	DeferredMicrocodeIsPending                              AdditionalSenseCode = 0x0022
	NoIndexsectorSignal                                     AdditionalSenseCode = 0x0100
	NoSeekComplete                                          AdditionalSenseCode = 0x0200
	PeripheralDeviceWriteFault                              AdditionalSenseCode = 0x0300
	NoWriteCurrent                                          AdditionalSenseCode = 0x0301
	ExcessiveWriteErrors                                    AdditionalSenseCode = 0x0302
	LogicalUnitNotReadyCauseNotReportable                   AdditionalSenseCode = 0x0400
	LogicalUnitIsInProcessOfBecomingReady                   AdditionalSenseCode = 0x0401
	LogicalUnitNotReadyInitializingCommandRequired          AdditionalSenseCode = 0x0402
	LogicalUnitNotReadyManualInterventionRequired           AdditionalSenseCode = 0x0403
	LogicalUnitNotReadyFormatInProgress                     AdditionalSenseCode = 0x0404
	LogicalUnitNotReadyRebuildInProgress                    AdditionalSenseCode = 0x0405
	LogicalUnitNotReadyRecalculationInProgress              AdditionalSenseCode = 0x0406
	LogicalUnitNotReadyOperationInProgress                  AdditionalSenseCode = 0x0407
	LogicalUnitNotReadyLongWriteInProgress                  AdditionalSenseCode = 0x0408
	LogicalUnitNotReadySelfTestInProgress                   AdditionalSenseCode = 0x0409
	LogicalUnitNotAccessibleAsymmetricAccessStateTransition AdditionalSenseCode = 0x040a
	LogicalUnitNotAccessibleTargetPortInStandbyState        AdditionalSenseCode = 0x040b
	LogicalUnitNotAccessibleTargetPortInUnavailableState    AdditionalSenseCode = 0x040c
	LogicalUnitNotReadyStructureCheckRequired               AdditionalSenseCode = 0x040d
	LogicalUnitNotReadySecuritySessionInProgress            AdditionalSenseCode = 0x040e
	LogicalUnitNotReadyAuxiliaryMemoryNotAccessible         AdditionalSenseCode = 0x0410
	LogicalUnitNotReadyNotifyenableSpinupRequired           AdditionalSenseCode = 0x0411
	LogicalUnitNotReadyOffline                              AdditionalSenseCode = 0x0412
	LogicalUnitNotReadySaCreationInProgress                 AdditionalSenseCode = 0x0413
	LogicalUnitNotReadySpaceAllocationInProgress            AdditionalSenseCode = 0x0414
	LogicalUnitNotReadyRoboticsDisabled                     AdditionalSenseCode = 0x0415
	LogicalUnitNotReadyConfigurationRequired                AdditionalSenseCode = 0x0416
	LogicalUnitNotReadyCalibrationRequired                  AdditionalSenseCode = 0x0417
	LogicalUnitNotReadyADoorIsOpen                          AdditionalSenseCode = 0x0418
	LogicalUnitNotReadyOperatingInSequentialMode            AdditionalSenseCode = 0x0419
	LogicalUnitNotReadyStartStopUnitCommandInProgress       AdditionalSenseCode = 0x041a
	LogicalUnitNotReadySanitizeInProgress                   AdditionalSenseCode = 0x041b
	LogicalUnitNotReadyAdditionalPowerUseNotYetGranted      AdditionalSenseCode = 0x041c
	LogicalUnitNotReadyConfigurationInProgress              AdditionalSenseCode = 0x041d
	LogicalUnitNotReadyMicrocodeActivationRequired          AdditionalSenseCode = 0x041e
	LogicalUnitNotReadyMicrocodeDownloadRequired            AdditionalSenseCode = 0x041f
	LogicalUnitNotReadyLogicalUnitResetRequired             AdditionalSenseCode = 0x0420
	LogicalUnitNotReadyHardResetRequired                    AdditionalSenseCode = 0x0421
	LogicalUnitNotReadyPowerCycleRequired                   AdditionalSenseCode = 0x0422
	LogicalUnitNotReadyAffiliationRequired                  AdditionalSenseCode = 0x0423
	DepopulationInProgress                                  AdditionalSenseCode = 0x0424
	LogicalUnitDoesNotRespondToSelection                    AdditionalSenseCode = 0x0500
	NoReferencePositionFound                                AdditionalSenseCode = 0x0600
	MultiplePeripheralDevicesSelected                       AdditionalSenseCode = 0x0700
	LogicalUnitCommunicationFailure                         AdditionalSenseCode = 0x0800
	LogicalUnitCommunicationTimeOut                         AdditionalSenseCode = 0x0801
	LogicalUnitCommunicationParityError                     AdditionalSenseCode = 0x0802
	LogicalUnitCommunicationCrcErrorultraDma32              AdditionalSenseCode = 0x0803
	UnreachableCopyTarget                                   AdditionalSenseCode = 0x0804
	TrackFollowingError                                     AdditionalSenseCode = 0x0900
	TrackingServoFailure                                    AdditionalSenseCode = 0x0901
	FocusServoFailure                                       AdditionalSenseCode = 0x0902
	SpindleServoFailure                                     AdditionalSenseCode = 0x0903
	HeadSelectFault                                         AdditionalSenseCode = 0x0904
	VibrationInducedTrackingError                           AdditionalSenseCode = 0x0905
	ErrorLogOverflow                                        AdditionalSenseCode = 0x0a00
	Warning                                                 AdditionalSenseCode = 0x0b00
	WarningSpecifiedTemperatureExceeded                     AdditionalSenseCode = 0x0b01
	WarningEnclosureDegraded                                AdditionalSenseCode = 0x0b02
	WarningBackgroundSelfTestFailed                         AdditionalSenseCode = 0x0b03
	WarningBackgroundPreScanDetectedMediumError             AdditionalSenseCode = 0x0b04
	WarningBackgroundMediumScanDetectedMediumError          AdditionalSenseCode = 0x0b05
	WarningNonVolatileCacheNowVolatile                      AdditionalSenseCode = 0x0b06
	WarningDegradedPowerToNonVolatileCache                  AdditionalSenseCode = 0x0b07
	WarningPowerLossExpected                                AdditionalSenseCode = 0x0b08
	WarningDeviceStatisticsNotificationActive               AdditionalSenseCode = 0x0b09
	WarningHighCriticalTemperatureLimitExceeded             AdditionalSenseCode = 0x0b0a
	WarningLowCriticalTemperatureLimitExceeded              AdditionalSenseCode = 0x0b0b
	WarningHighOperatingTemperatureLimitExceeded            AdditionalSenseCode = 0x0b0c
	WarningLowOperatingTemperatureLimitExceeded             AdditionalSenseCode = 0x0b0d
	WarningHighCriticalHumidityLimitExceeded                AdditionalSenseCode = 0x0b0e
	WarningLowCriticalHumidityLimitExceeded                 AdditionalSenseCode = 0x0b0f
	WarningHighOperatingHumidityLimitExceeded               AdditionalSenseCode = 0x0b10
	WarningLowOperatingHumidityLimitExceeded                AdditionalSenseCode = 0x0b11
	WarningMicrocodeSecurityAtRisk                          AdditionalSenseCode = 0x0b12
	WarningMicrocodeDigitalSignatureValidationFailure       AdditionalSenseCode = 0x0b13
	WarningPhysicalElementStatusChange                      AdditionalSenseCode = 0x0b14
	WriteError                                              AdditionalSenseCode = 0x0c00
	WriteErrorRecoveredWithAutoReallocation                 AdditionalSenseCode = 0x0c01
	WriteErrorAutoReallocationFailed                        AdditionalSenseCode = 0x0c02
	WriteErrorRecommendReassignment                         AdditionalSenseCode = 0x0c03
	CompressionCheckMiscompareError                         AdditionalSenseCode = 0x0c04
	DataExpansionOccurredDuringCompression                  AdditionalSenseCode = 0x0c05
	BlockNotCompressible                                    AdditionalSenseCode = 0x0c06
	WriteErrorRecoveryNeeded                                AdditionalSenseCode = 0x0c07
	WriteErrorRecoveryFailed                                AdditionalSenseCode = 0x0c08
	WriteErrorLossOfStreaming                               AdditionalSenseCode = 0x0c09
	WriteErrorPaddingBlocksAdded                            AdditionalSenseCode = 0x0c0a
	AuxiliaryMemoryWriteError                               AdditionalSenseCode = 0x0c0b
	WriteErrorUnexpectedUnsolicitedData                     AdditionalSenseCode = 0x0c0c
	WriteErrorNotEnoughUnsolicitedData                      AdditionalSenseCode = 0x0c0d
	MultipleWriteErrors                                     AdditionalSenseCode = 0x0c0e
	DefectsInErrorWindow                                    AdditionalSenseCode = 0x0c0f
	IncompleteMultipleAtomicWriteOperations                 AdditionalSenseCode = 0x0c10
	WriteErrorRecoveryScanNeeded                            AdditionalSenseCode = 0x0c11
	WriteErrorInsufficientZoneResources                     AdditionalSenseCode = 0x0c12
	ErrorDetectedByThirdPartyTemporaryInitiator             AdditionalSenseCode = 0x0d00
	ThirdPartyDeviceFailure                                 AdditionalSenseCode = 0x0d01
	CopyTargetDeviceNotReachable                            AdditionalSenseCode = 0x0d02
	IncorrectCopyTargetDeviceType                           AdditionalSenseCode = 0x0d03
	CopyTargetDeviceDataUnderrun                            AdditionalSenseCode = 0x0d04
	CopyTargetDeviceDataOverrun                             AdditionalSenseCode = 0x0d05
	InvalidInformationUnit                                  AdditionalSenseCode = 0x0e00
	InformationUnitTooShort                                 AdditionalSenseCode = 0x0e01
	InformationUnitTooLong                                  AdditionalSenseCode = 0x0e02
	InvalidFieldInCommandInformationUnit                    AdditionalSenseCode = 0x0e03
	IdCrcOrEccError                                         AdditionalSenseCode = 0x1000
	LogicalBlockGuardCheckFailed                            AdditionalSenseCode = 0x1001
	LogicalBlockApplicationTagCheckFailed                   AdditionalSenseCode = 0x1002
	LogicalBlockReferenceTagCheckFailed                     AdditionalSenseCode = 0x1003
	LogicalBlockProtectionErrorOnRecoverBufferedData        AdditionalSenseCode = 0x1004
	LogicalBlockProtectionMethodError                       AdditionalSenseCode = 0x1005
	UnrecoveredReadError                                    AdditionalSenseCode = 0x1100
	ReadRetriesExhausted                                    AdditionalSenseCode = 0x1101
	ErrorTooLongToCorrect                                   AdditionalSenseCode = 0x1102
	MultipleReadErrors                                      AdditionalSenseCode = 0x1103
	UnrecoveredReadErrorAutoReallocateFailed                AdditionalSenseCode = 0x1104
	LEcUncorrectableError                                   AdditionalSenseCode = 0x1105
	CircUnrecoveredError                                    AdditionalSenseCode = 0x1106
	DataReSynchronizationError                              AdditionalSenseCode = 0x1107
	IncompleteBlockRead                                     AdditionalSenseCode = 0x1108
	NoGapFound                                              AdditionalSenseCode = 0x1109
	MiscorrectedError                                       AdditionalSenseCode = 0x110a
	UnrecoveredReadErrorRecommendReassignment               AdditionalSenseCode = 0x110b
	UnrecoveredReadErrorRecommendRewriteTheData             AdditionalSenseCode = 0x110c
	DeCompressionCrcError                                   AdditionalSenseCode = 0x110d
	CannotDecompressUsingDeclaredAlgorithm                  AdditionalSenseCode = 0x110e
	ErrorReadingUpceanNumber                                AdditionalSenseCode = 0x110f
	ErrorReadingIsrcNumber                                  AdditionalSenseCode = 0x1110
	ReadErrorLossOfStreaming                                AdditionalSenseCode = 0x1111
	AuxiliaryMemoryReadError                                AdditionalSenseCode = 0x1112
	ReadErrorFailedRetransmissionRequest                    AdditionalSenseCode = 0x1113
	ReadErrorLbaMarkedBadByApplicationClient                AdditionalSenseCode = 0x1114
	WriteAfterSanitizeRequired                              AdditionalSenseCode = 0x1115
	AddressMarkNotFoundForIdField                           AdditionalSenseCode = 0x1200
	AddressMarkNotFoundForDataField                         AdditionalSenseCode = 0x1300
	RecordedEntityNotFound                                  AdditionalSenseCode = 0x1400
	RecordNotFound                                          AdditionalSenseCode = 0x1401
	FilemarkOrSetmarkNotFound                               AdditionalSenseCode = 0x1402
	EndOfDataNotFound                                       AdditionalSenseCode = 0x1403
	BlockSequenceError                                      AdditionalSenseCode = 0x1404
	RecordNotFoundRecommendReassignment                     AdditionalSenseCode = 0x1405
	RecordNotFoundDataAutoReallocated                       AdditionalSenseCode = 0x1406
	LocateOperationFailure                                  AdditionalSenseCode = 0x1407
	RandomPositioningError                                  AdditionalSenseCode = 0x1500
	MechanicalPositioningError                              AdditionalSenseCode = 0x1501
	PositioningErrorDetectedByReadOfMedium                  AdditionalSenseCode = 0x1502
	DataSynchronizationMarkError                            AdditionalSenseCode = 0x1600
	DataSyncErrorDataRewritten                              AdditionalSenseCode = 0x1601
	DataSyncErrorRecommendRewrite                           AdditionalSenseCode = 0x1602
	DataSyncErrorDataAutoReallocated                        AdditionalSenseCode = 0x1603
	DataSyncErrorRecommendReassignment                      AdditionalSenseCode = 0x1604
	RecoveredDataWithNoErrorCorrectionApplied               AdditionalSenseCode = 0x1700
	RecoveredDataWithRetries                                AdditionalSenseCode = 0x1701
	RecoveredDataWithPositiveHeadOffset                     AdditionalSenseCode = 0x1702
	RecoveredDataWithNegativeHeadOffset                     AdditionalSenseCode = 0x1703
	RecoveredDataWithRetriesAndorCircApplied                AdditionalSenseCode = 0x1704
	RecoveredDataUsingPreviousSectorId                      AdditionalSenseCode = 0x1705
	RecoveredDataWithoutEccDataAutoReallocated              AdditionalSenseCode = 0x1706
	RecoveredDataWithoutEccRecommendReassignment            AdditionalSenseCode = 0x1707
	RecoveredDataWithoutEccRecommendRewrite                 AdditionalSenseCode = 0x1708
	RecoveredDataWithoutEccDataRewritten                    AdditionalSenseCode = 0x1709
	RecoveredDataWithErrorCorrectionApplied                 AdditionalSenseCode = 0x1800
	RecoveredDataWithErrorCorrRetriesApplied                AdditionalSenseCode = 0x1801
	RecoveredDataDataAutoReallocated                        AdditionalSenseCode = 0x1802
	RecoveredDataWithCirc                                   AdditionalSenseCode = 0x1803
	RecoveredDataWithLEc                                    AdditionalSenseCode = 0x1804
	RecoveredDataRecommendReassignment                      AdditionalSenseCode = 0x1805
	RecoveredDataRecommendRewrite                           AdditionalSenseCode = 0x1806
	RecoveredDataWithEccDataRewritten                       AdditionalSenseCode = 0x1807
	RecoveredDataWithLinking                                AdditionalSenseCode = 0x1808
	DefectListError                                         AdditionalSenseCode = 0x1900
	DefectListNotAvailable                                  AdditionalSenseCode = 0x1901
	DefectListErrorInPrimaryList                            AdditionalSenseCode = 0x1902
	DefectListErrorInGrownList                              AdditionalSenseCode = 0x1903
	ParameterListLengthError                                AdditionalSenseCode = 0x1a00
	SynchronousDataTransferError                            AdditionalSenseCode = 0x1b00
	DefectListNotFound                                      AdditionalSenseCode = 0x1c00
	PrimaryDefectListNotFound                               AdditionalSenseCode = 0x1c01
	GrownDefectListNotFound                                 AdditionalSenseCode = 0x1c02
	MiscompareDuringVerifyOperation                         AdditionalSenseCode = 0x1d00
	MiscompareVerifyOfUnmappedLba                           AdditionalSenseCode = 0x1d01
	RecoveredIdWithEccCorrection                            AdditionalSenseCode = 0x1e00
	PartialDefectListTransfer                               AdditionalSenseCode = 0x1f00
	InvalidCommandOperationCode                             AdditionalSenseCode = 0x2000
	AccessDeniedInitiatorPendingEnrolled                    AdditionalSenseCode = 0x2001
	AccessDeniedNoAccessRights                              AdditionalSenseCode = 0x2002
	AccessDeniedInvalidMgmtIdKey                            AdditionalSenseCode = 0x2003
	IllegalCommandWhileInWriteCapableState                  AdditionalSenseCode = 0x2004
	IllegalCommandWhileInExplicitAddressMode                AdditionalSenseCode = 0x2006
	IllegalCommandWhileInImplicitAddressMode                AdditionalSenseCode = 0x2007
	AccessDeniedEnrollmentConflict                          AdditionalSenseCode = 0x2008
	AccessDeniedInvalidLuIdentifier                         AdditionalSenseCode = 0x2009
	AccessDeniedInvalidProxyToken                           AdditionalSenseCode = 0x200a
	AccessDeniedAclLunConflict                              AdditionalSenseCode = 0x200b
	IllegalCommandWhenNotInAppendOnlyMode                   AdditionalSenseCode = 0x200c
	NotAnAdministrativeLogicalUnit                          AdditionalSenseCode = 0x200d
	NotASubsidiaryLogicalUnit                               AdditionalSenseCode = 0x200e
	NotAConglomerateLogicalUnit                             AdditionalSenseCode = 0x200f
	LogicalBlockAddressOutOfRange                           AdditionalSenseCode = 0x2100
	InvalidElementAddress                                   AdditionalSenseCode = 0x2101
	InvalidAddressForWrite                                  AdditionalSenseCode = 0x2102
	InvalidWriteCrossingLayerJump                           AdditionalSenseCode = 0x2103
	UnalignedWriteCommand                                   AdditionalSenseCode = 0x2104
	WriteBoundaryViolation                                  AdditionalSenseCode = 0x2105
	AttemptToReadInvalidData                                AdditionalSenseCode = 0x2106
	ReadBoundaryViolation                                   AdditionalSenseCode = 0x2107
	MisalignedWriteCommand                                  AdditionalSenseCode = 0x2108
	IllegalFunctionuse20002400Or2600                        AdditionalSenseCode = 0x2200
	InvalidTokenOperationCauseNotReportable                 AdditionalSenseCode = 0x2300
	InvalidTokenOperationUnsupportedTokenType               AdditionalSenseCode = 0x2301
	InvalidTokenOperationRemoteTokenUsageNotSupported       AdditionalSenseCode = 0x2302
	InvalidTokenOperationRemoteRodTokenCreationNotSupported AdditionalSenseCode = 0x2303
	InvalidTokenOperationTokenUnknown                       AdditionalSenseCode = 0x2304
	InvalidTokenOperationTokenCorrupt                       AdditionalSenseCode = 0x2305
	InvalidTokenOperationTokenRevoked                       AdditionalSenseCode = 0x2306
	InvalidTokenOperationTokenExpired                       AdditionalSenseCode = 0x2307
	InvalidTokenOperationTokenCancelled                     AdditionalSenseCode = 0x2308
	InvalidTokenOperationTokenDeleted                       AdditionalSenseCode = 0x2309
	InvalidTokenOperationInvalidTokenLength                 AdditionalSenseCode = 0x230a
	InvalidFieldInCdb                                       AdditionalSenseCode = 0x2400
	CdbDecryptionError                                      AdditionalSenseCode = 0x2401
	SecurityAuditValueFrozen                                AdditionalSenseCode = 0x2404
	SecurityWorkingKeyFrozen                                AdditionalSenseCode = 0x2405
	NonceNotUnique                                          AdditionalSenseCode = 0x2406
	NonceTimestampOutOfRange                                AdditionalSenseCode = 0x2407
	InvalidXcdb                                             AdditionalSenseCode = 0x2408
	InvalidFastFormat                                       AdditionalSenseCode = 0x2409
	LogicalUnitNotSupported                                 AdditionalSenseCode = 0x2500
	InvalidFieldInParameterList                             AdditionalSenseCode = 0x2600
	ParameterNotSupported                                   AdditionalSenseCode = 0x2601
	ParameterValueInvalid                                   AdditionalSenseCode = 0x2602
	ThresholdParametersNotSupported                         AdditionalSenseCode = 0x2603
	InvalidReleaseOfPersistentReservation                   AdditionalSenseCode = 0x2604
	DataDecryptionError                                     AdditionalSenseCode = 0x2605
	TooManyTargetDescriptors                                AdditionalSenseCode = 0x2606
	UnsupportedTargetDescriptorTypeCode                     AdditionalSenseCode = 0x2607
	TooManySegmentDescriptors                               AdditionalSenseCode = 0x2608
	UnsupportedSegmentDescriptorTypeCode                    AdditionalSenseCode = 0x2609
	UnexpectedInexactSegment                                AdditionalSenseCode = 0x260a
	InlineDataLengthExceeded                                AdditionalSenseCode = 0x260b
	InvalidOperationForCopySourceOrDestination              AdditionalSenseCode = 0x260c
	CopySegmentGranularityViolation                         AdditionalSenseCode = 0x260d
	InvalidParameterWhilePortIsEnabled                      AdditionalSenseCode = 0x260e
	InvalidDataOutBufferIntegrityCheckValue                 AdditionalSenseCode = 0x260f
	DataDecryptionKeyFailLimitReached                       AdditionalSenseCode = 0x2610
	IncompleteKeyAssociatedDataSet                          AdditionalSenseCode = 0x2611
	VendorSpecificKeyReferenceNotFound                      AdditionalSenseCode = 0x2612
	ApplicationTagModePageIsInvalid                         AdditionalSenseCode = 0x2613
	TapeStreamMirroringPrevented                            AdditionalSenseCode = 0x2614
	CopySourceOrCopyDestinationNotAuthorized                AdditionalSenseCode = 0x2615
	WriteProtected                                          AdditionalSenseCode = 0x2700
	HardwareWriteProtected                                  AdditionalSenseCode = 0x2701
	LogicalUnitSoftwareWriteProtected                       AdditionalSenseCode = 0x2702
	AssociatedWriteProtect                                  AdditionalSenseCode = 0x2703
	PersistentWriteProtect                                  AdditionalSenseCode = 0x2704
	PermanentWriteProtect                                   AdditionalSenseCode = 0x2705
	ConditionalWriteProtect                                 AdditionalSenseCode = 0x2706
	SpaceAllocationFailedWriteProtect                       AdditionalSenseCode = 0x2707
	ZoneIsReadOnly                                          AdditionalSenseCode = 0x2708
	NotReadyToReadyChangeMediumMayHaveChanged               AdditionalSenseCode = 0x2800
	ImportOrExportElementAccessed                           AdditionalSenseCode = 0x2801
	FormatLayerMayHaveChanged                               AdditionalSenseCode = 0x2802
	ImportexportElementAccessedMediumChanged                AdditionalSenseCode = 0x2803
	PowerOnResetOrBusDeviceResetOccurred                    AdditionalSenseCode = 0x2900
	PowerOnOccurred                                         AdditionalSenseCode = 0x2901
	ScsiBusResetOccurred                                    AdditionalSenseCode = 0x2902
	BusDeviceResetFunctionOccurred                          AdditionalSenseCode = 0x2903
	DeviceInternalReset                                     AdditionalSenseCode = 0x2904
	TransceiverModeChangedToSingleEnded                     AdditionalSenseCode = 0x2905
	TransceiverModeChangedToLvd                             AdditionalSenseCode = 0x2906
	ITNexusLossOccurred                                     AdditionalSenseCode = 0x2907
	ParametersChanged                                       AdditionalSenseCode = 0x2a00
	ModeParametersChanged                                   AdditionalSenseCode = 0x2a01
	LogParametersChanged                                    AdditionalSenseCode = 0x2a02
	ReservationsPreempted                                   AdditionalSenseCode = 0x2a03
	ReservationsReleased                                    AdditionalSenseCode = 0x2a04
	RegistrationsPreempted                                  AdditionalSenseCode = 0x2a05
	AsymmetricAccessStateChanged                            AdditionalSenseCode = 0x2a06
	ImplicitAsymmetricAccessStateTransitionFailed           AdditionalSenseCode = 0x2a07
	PriorityChanged                                         AdditionalSenseCode = 0x2a08
	CapacityDataHasChanged                                  AdditionalSenseCode = 0x2a09
	ErrorHistoryITNexusCleared                              AdditionalSenseCode = 0x2a0a
	ErrorHistorySnapshotReleased                            AdditionalSenseCode = 0x2a0b
	ErrorRecoveryAttributesHaveChanged                      AdditionalSenseCode = 0x2a0c
	DataEncryptionCapabilitiesChanged                       AdditionalSenseCode = 0x2a0d
	TimestampChanged                                        AdditionalSenseCode = 0x2a10
	DataEncryptionParametersChangedByAnotherITNexus         AdditionalSenseCode = 0x2a11
	DataEncryptionParametersChangedByVendorSpecificEvent    AdditionalSenseCode = 0x2a12
	DataEncryptionKeyInstanceCounterHasChanged              AdditionalSenseCode = 0x2a13
	SaCreationCapabilitiesDataHasChanged                    AdditionalSenseCode = 0x2a14
	MediumRemovalPreventionPreempted                        AdditionalSenseCode = 0x2a15
	ZoneResetWritePointerRecommended                        AdditionalSenseCode = 0x2a16
	CopyCannotExecuteSinceHostCannotDisconnect              AdditionalSenseCode = 0x2b00
	CommandSequenceError                                    AdditionalSenseCode = 0x2c00
	TooManyWindowsSpecified                                 AdditionalSenseCode = 0x2c01
	InvalidCombinationOfWindowsSpecified                    AdditionalSenseCode = 0x2c02
	CurrentProgramAreaIsNotEmpty                            AdditionalSenseCode = 0x2c03
	CurrentProgramAreaIsEmpty                               AdditionalSenseCode = 0x2c04
	IllegalPowerConditionRequest                            AdditionalSenseCode = 0x2c05
	PersistentPreventConflict                               AdditionalSenseCode = 0x2c06
	PreviousBusyStatus                                      AdditionalSenseCode = 0x2c07
	PreviousTaskSetFullStatus                               AdditionalSenseCode = 0x2c08
	PreviousReservationConflictStatus                       AdditionalSenseCode = 0x2c09
	PartitionOrCollectionContainsUserObjects                AdditionalSenseCode = 0x2c0a
	NotReserved                                             AdditionalSenseCode = 0x2c0b
	OrwriteGenerationDoesNotMatch                           AdditionalSenseCode = 0x2c0c
	ResetWritePointerNotAllowed                             AdditionalSenseCode = 0x2c0d
	ZoneIsOffline                                           AdditionalSenseCode = 0x2c0e
	StreamNotOpen                                           AdditionalSenseCode = 0x2c0f
	UnwrittenDataInZone                                     AdditionalSenseCode = 0x2c10
	DescriptorFormatSenseDataRequired                       AdditionalSenseCode = 0x2c11
	OverwriteErrorOnUpdateInPlace                           AdditionalSenseCode = 0x2d00
	InsufficientTimeForOperation                            AdditionalSenseCode = 0x2e00
	CommandTimeoutBeforeProcessing                          AdditionalSenseCode = 0x2e01
	CommandTimeoutDuringProcessing                          AdditionalSenseCode = 0x2e02
	CommandTimeoutDuringProcessingDueToErrorRecovery        AdditionalSenseCode = 0x2e03
	CommandsClearedByAnotherInitiator                       AdditionalSenseCode = 0x2f00
	CommandsClearedByPowerLossNotification                  AdditionalSenseCode = 0x2f01
	CommandsClearedByDeviceServer                           AdditionalSenseCode = 0x2f02
	SomeCommandsClearedByQueuingLayerEvent                  AdditionalSenseCode = 0x2f03
	IncompatibleMediumInstalled                             AdditionalSenseCode = 0x3000
	CannotReadMediumUnknownFormat                           AdditionalSenseCode = 0x3001
	CannotReadMediumIncompatibleFormat                      AdditionalSenseCode = 0x3002
	CleaningCartridgeInstalled                              AdditionalSenseCode = 0x3003
	CannotWriteMediumUnknownFormat                          AdditionalSenseCode = 0x3004
	CannotWriteMediumIncompatibleFormat                     AdditionalSenseCode = 0x3005
	CannotFormatMediumIncompatibleMedium                    AdditionalSenseCode = 0x3006
	CleaningFailure                                         AdditionalSenseCode = 0x3007
	CannotWriteApplicationCodeMismatch                      AdditionalSenseCode = 0x3008
	CurrentSessionNotFixatedForAppend                       AdditionalSenseCode = 0x3009
	CleaningRequestRejected                                 AdditionalSenseCode = 0x300a
	WormMediumOverwriteAttempted                            AdditionalSenseCode = 0x300c
	WormMediumIntegrityCheck                                AdditionalSenseCode = 0x300d
	MediumNotFormatted                                      AdditionalSenseCode = 0x3010
	IncompatibleVolumeType                                  AdditionalSenseCode = 0x3011
	IncompatibleVolumeQualifier                             AdditionalSenseCode = 0x3012
	CleaningVolumeExpired                                   AdditionalSenseCode = 0x3013
	MediumFormatCorrupted                                   AdditionalSenseCode = 0x3100
	FormatCommandFailed                                     AdditionalSenseCode = 0x3101
	ZonedFormattingFailedDueToSpareLinking                  AdditionalSenseCode = 0x3102
	SanitizeCommandFailed                                   AdditionalSenseCode = 0x3103
	DepopulationFailed                                      AdditionalSenseCode = 0x3104
	NoDefectSpareLocationAvailable                          AdditionalSenseCode = 0x3200
	DefectListUpdateFailure                                 AdditionalSenseCode = 0x3201
	TapeLengthError                                         AdditionalSenseCode = 0x3300
	EnclosureFailure                                        AdditionalSenseCode = 0x3400
	EnclosureServicesFailure                                AdditionalSenseCode = 0x3500
	UnsupportedEnclosureFunction                            AdditionalSenseCode = 0x3501
	EnclosureServicesUnavailable                            AdditionalSenseCode = 0x3502
	EnclosureServicesTransferFailure                        AdditionalSenseCode = 0x3503
	EnclosureServicesTransferRefused                        AdditionalSenseCode = 0x3504
	EnclosureServicesChecksumError                          AdditionalSenseCode = 0x3505
	RibbonInkOrTonerFailure                                 AdditionalSenseCode = 0x3600
	RoundedParameter                                        AdditionalSenseCode = 0x3700
	EventStatusNotification                                 AdditionalSenseCode = 0x3800
	EsnPowerManagementClassEvent                            AdditionalSenseCode = 0x3802
	EsnMediaClassEvent                                      AdditionalSenseCode = 0x3804
	EsnDeviceBusyClassEvent                                 AdditionalSenseCode = 0x3806
	ThinProvisioningSoftThresholdReached                    AdditionalSenseCode = 0x3807
	SavingParametersNotSupported                            AdditionalSenseCode = 0x3900
	MediumNotPresent                                        AdditionalSenseCode = 0x3a00
	MediumNotPresentTrayClosed                              AdditionalSenseCode = 0x3a01
	MediumNotPresentTrayOpen                                AdditionalSenseCode = 0x3a02
	MediumNotPresentLoadable                                AdditionalSenseCode = 0x3a03
	MediumNotPresentMediumAuxiliaryMemoryAccessible         AdditionalSenseCode = 0x3a04
	SequentialPositioningError                              AdditionalSenseCode = 0x3b00
	TapePositionErrorAtBeginningOfMedium                    AdditionalSenseCode = 0x3b01
	TapePositionErrorAtEndOfMedium                          AdditionalSenseCode = 0x3b02
	TapeOrElectronicVerticalFormsUnitNotReady               AdditionalSenseCode = 0x3b03
	SlewFailure                                             AdditionalSenseCode = 0x3b04
	PaperJam                                                AdditionalSenseCode = 0x3b05
	FailedToSenseTopOfForm                                  AdditionalSenseCode = 0x3b06
	FailedToSenseBottomOfForm                               AdditionalSenseCode = 0x3b07
	RepositionError                                         AdditionalSenseCode = 0x3b08
	ReadPastEndOfMedium                                     AdditionalSenseCode = 0x3b09
	ReadPastBeginningOfMedium                               AdditionalSenseCode = 0x3b0a
	PositionPastEndOfMedium                                 AdditionalSenseCode = 0x3b0b
	PositionPastBeginningOfMedium                           AdditionalSenseCode = 0x3b0c
	MediumDestinationElementFull                            AdditionalSenseCode = 0x3b0d
	MediumSourceElementEmpty                                AdditionalSenseCode = 0x3b0e
	EndOfMediumReached                                      AdditionalSenseCode = 0x3b0f
	MediumMagazineNotAccessible                             AdditionalSenseCode = 0x3b11
	MediumMagazineRemoved                                   AdditionalSenseCode = 0x3b12
	MediumMagazineInserted                                  AdditionalSenseCode = 0x3b13
	MediumMagazineLocked                                    AdditionalSenseCode = 0x3b14
	MediumMagazineUnlocked                                  AdditionalSenseCode = 0x3b15
	MechanicalPositioningOrChangerError                     AdditionalSenseCode = 0x3b16
	ReadPastEndOfUserObject                                 AdditionalSenseCode = 0x3b17
	ElementDisabled                                         AdditionalSenseCode = 0x3b18
	ElementEnabled                                          AdditionalSenseCode = 0x3b19
	DataTransferDeviceRemoved                               AdditionalSenseCode = 0x3b1a
	DataTransferDeviceInserted                              AdditionalSenseCode = 0x3b1b
	TooManyLogicalObjectsOnPartitionToSupportOperation      AdditionalSenseCode = 0x3b1c
	InvalidBitsInIdentifyMessage                            AdditionalSenseCode = 0x3d00
	LogicalUnitHasNotSelfConfiguredYet                      AdditionalSenseCode = 0x3e00
	LogicalUnitFailure                                      AdditionalSenseCode = 0x3e01
	TimeoutOnLogicalUnit                                    AdditionalSenseCode = 0x3e02
	LogicalUnitFailedSelfTest                               AdditionalSenseCode = 0x3e03
	LogicalUnitUnableToUpdateSelfTestLog                    AdditionalSenseCode = 0x3e04
	TargetOperatingConditionsHaveChanged                    AdditionalSenseCode = 0x3f00
	MicrocodeHasBeenChanged                                 AdditionalSenseCode = 0x3f01
	ChangedOperatingDefinition                              AdditionalSenseCode = 0x3f02
	InquiryDataHasChanged                                   AdditionalSenseCode = 0x3f03
	ComponentDeviceAttached                                 AdditionalSenseCode = 0x3f04
	DeviceIdentifierChanged                                 AdditionalSenseCode = 0x3f05
	RedundancyGroupCreatedOrModified                        AdditionalSenseCode = 0x3f06
	RedundancyGroupDeleted                                  AdditionalSenseCode = 0x3f07
	SpareCreatedOrModified                                  AdditionalSenseCode = 0x3f08
	SpareDeleted                                            AdditionalSenseCode = 0x3f09
	VolumeSetCreatedOrModified                              AdditionalSenseCode = 0x3f0a
	VolumeSetDeleted                                        AdditionalSenseCode = 0x3f0b
	VolumeSetDeassigned                                     AdditionalSenseCode = 0x3f0c
	VolumeSetReassigned                                     AdditionalSenseCode = 0x3f0d
	ReportedLunsDataHasChanged                              AdditionalSenseCode = 0x3f0e
	EchoBufferOverwritten                                   AdditionalSenseCode = 0x3f0f
	MediumLoadable                                          AdditionalSenseCode = 0x3f10
	MediumAuxiliaryMemoryAccessible                         AdditionalSenseCode = 0x3f11
	IscsiIpAddressAdded                                     AdditionalSenseCode = 0x3f12
	IscsiIpAddressRemoved                                   AdditionalSenseCode = 0x3f13
	IscsiIpAddressChanged                                   AdditionalSenseCode = 0x3f14
	InspectReferralsSenseDescriptors                        AdditionalSenseCode = 0x3f15
	MicrocodeHasBeenChangedWithoutReset                     AdditionalSenseCode = 0x3f16
	ZoneTransitionToFull                                    AdditionalSenseCode = 0x3f17
	BindCompleted                                           AdditionalSenseCode = 0x3f18
	BindRedirected                                          AdditionalSenseCode = 0x3f19
	SubsidiaryBindingChanged                                AdditionalSenseCode = 0x3f1a
	RamFailureshouldUse40Nn                                 AdditionalSenseCode = 0x4000
	DataPathFailureshouldUse40Nn                            AdditionalSenseCode = 0x4100
	PowerOnOrSelfTestFailureshouldUse40Nn                   AdditionalSenseCode = 0x4200
	MessageError                                            AdditionalSenseCode = 0x4300
	InternalTargetFailure                                   AdditionalSenseCode = 0x4400
	PersistentReservationInformationLost                    AdditionalSenseCode = 0x4401
	AtaDeviceFailedSetFeatures                              AdditionalSenseCode = 0x4471
	SelectOrReselectFailure                                 AdditionalSenseCode = 0x4500
	UnsuccessfulSoftReset                                   AdditionalSenseCode = 0x4600
	ScsiParityError                                         AdditionalSenseCode = 0x4700
	DataPhaseCrcErrorDetected                               AdditionalSenseCode = 0x4701
	ScsiParityErrorDetectedDuringStDataPhase                AdditionalSenseCode = 0x4702
	InformationUnitIucrcErrorDetected                       AdditionalSenseCode = 0x4703
	AsynchronousInformationProtectionErrorDetected          AdditionalSenseCode = 0x4704
	ProtocolServiceCrcError                                 AdditionalSenseCode = 0x4705
	PhyTestFunctionInProgress                               AdditionalSenseCode = 0x4706
	SomeCommandsClearedByIscsiProtocolEvent                 AdditionalSenseCode = 0x477f
	InitiatorDetectedErrorMessageReceived                   AdditionalSenseCode = 0x4800
	InvalidMessageError                                     AdditionalSenseCode = 0x4900
	CommandPhaseError                                       AdditionalSenseCode = 0x4a00
	DataPhaseError                                          AdditionalSenseCode = 0x4b00
	InvalidTargetPortTransferTagReceived                    AdditionalSenseCode = 0x4b01
	TooMuchWriteData                                        AdditionalSenseCode = 0x4b02
	AcknakTimeout                                           AdditionalSenseCode = 0x4b03
	NakReceived                                             AdditionalSenseCode = 0x4b04
	DataOffsetError                                         AdditionalSenseCode = 0x4b05
	InitiatorResponseTimeout                                AdditionalSenseCode = 0x4b06
	ConnectionLost                                          AdditionalSenseCode = 0x4b07
	DataInBufferOverflowDataBufferSize                      AdditionalSenseCode = 0x4b08
	DataInBufferOverflowDataBufferDescriptorArea            AdditionalSenseCode = 0x4b09
	DataInBufferError                                       AdditionalSenseCode = 0x4b0a
	DataOutBufferOverflowDataBufferSize                     AdditionalSenseCode = 0x4b0b
	DataOutBufferOverflowDataBufferDescriptorArea           AdditionalSenseCode = 0x4b0c
	DataOutBufferError                                      AdditionalSenseCode = 0x4b0d
	PcieFabricError                                         AdditionalSenseCode = 0x4b0e
	PcieCompletionTimeout                                   AdditionalSenseCode = 0x4b0f
	PcieCompleterAbort                                      AdditionalSenseCode = 0x4b10
	PciePoisonedTlpReceived                                 AdditionalSenseCode = 0x4b11
	PcieEcrcCheckFailed                                     AdditionalSenseCode = 0x4b12
	PcieUnsupportedRequest                                  AdditionalSenseCode = 0x4b13
	PcieAcsViolation                                        AdditionalSenseCode = 0x4b14
	PcieTlpPrefixBlocked                                    AdditionalSenseCode = 0x4b15
	LogicalUnitFailedSelfConfiguration                      AdditionalSenseCode = 0x4c00
	OverlappedCommandsAttempted                             AdditionalSenseCode = 0x4e00
	WriteAppendError                                        AdditionalSenseCode = 0x5000
	WriteAppendPositionError                                AdditionalSenseCode = 0x5001
	PositionErrorRelatedToTiming                            AdditionalSenseCode = 0x5002
	EraseFailure                                            AdditionalSenseCode = 0x5100
	EraseFailureIncompleteEraseOperationDetected            AdditionalSenseCode = 0x5101
	CartridgeFault                                          AdditionalSenseCode = 0x5200
	MediaLoadOrEjectFailed                                  AdditionalSenseCode = 0x5300
	UnloadTapeFailure                                       AdditionalSenseCode = 0x5301
	MediumRemovalPrevented                                  AdditionalSenseCode = 0x5302
	MediumRemovalPreventedByDataTransferElement             AdditionalSenseCode = 0x5303
	MediumThreadOrUnthreadFailure                           AdditionalSenseCode = 0x5304
	VolumeIdentifierInvalid                                 AdditionalSenseCode = 0x5305
	VolumeIdentifierMissing                                 AdditionalSenseCode = 0x5306
	DuplicateVolumeIdentifier                               AdditionalSenseCode = 0x5307
	ElementStatusUnknown                                    AdditionalSenseCode = 0x5308
	DataTransferDeviceErrorLoadFailed                       AdditionalSenseCode = 0x5309
	DataTransferDeviceErrorUnloadFailed                     AdditionalSenseCode = 0x530a
	DataTransferDeviceErrorUnloadMissing                    AdditionalSenseCode = 0x530b
	DataTransferDeviceErrorEjectFailed                      AdditionalSenseCode = 0x530c
	DataTransferDeviceErrorLibraryCommunicationFailed       AdditionalSenseCode = 0x530d
	ScsiToHostSystemInterfaceFailure                        AdditionalSenseCode = 0x5400
	SystemResourceFailure                                   AdditionalSenseCode = 0x5500
	SystemBufferFull                                        AdditionalSenseCode = 0x5501
	InsufficientReservationResources                        AdditionalSenseCode = 0x5502
	InsufficientResources                                   AdditionalSenseCode = 0x5503
	InsufficientRegistrationResources                       AdditionalSenseCode = 0x5504
	InsufficientAccessControlResources                      AdditionalSenseCode = 0x5505
	AuxiliaryMemoryOutOfSpace                               AdditionalSenseCode = 0x5506
	QuotaError                                              AdditionalSenseCode = 0x5507
	MaximumNumberOfSupplementalDecryptionKeysExceeded       AdditionalSenseCode = 0x5508
	MediumAuxiliaryMemoryNotAccessible                      AdditionalSenseCode = 0x5509
	DataCurrentlyUnavailable                                AdditionalSenseCode = 0x550a
	InsufficientPowerForOperation                           AdditionalSenseCode = 0x550b
	InsufficientResourcesToCreateRod                        AdditionalSenseCode = 0x550c
	InsufficientResourcesToCreateRodToken                   AdditionalSenseCode = 0x550d
	InsufficientZoneResources                               AdditionalSenseCode = 0x550e
	InsufficientZoneResourcesToCompleteWrite                AdditionalSenseCode = 0x550f
	MaximumNumberOfStreamsOpen                              AdditionalSenseCode = 0x5510
	InsufficientResourcesToBind                             AdditionalSenseCode = 0x5511
	UnableToRecoverTableOfContents                          AdditionalSenseCode = 0x5700
	GenerationDoesNotExist                                  AdditionalSenseCode = 0x5800
	UpdatedBlockRead                                        AdditionalSenseCode = 0x5900
	OperatorRequestOrStateChangeInput                       AdditionalSenseCode = 0x5a00
	OperatorMediumRemovalRequest                            AdditionalSenseCode = 0x5a01
	OperatorSelectedWriteProtect                            AdditionalSenseCode = 0x5a02
	OperatorSelectedWritePermit                             AdditionalSenseCode = 0x5a03
	LogException                                            AdditionalSenseCode = 0x5b00
	ThresholdConditionMet                                   AdditionalSenseCode = 0x5b01
	LogCounterAtMaximum                                     AdditionalSenseCode = 0x5b02
	LogListCodesExhausted                                   AdditionalSenseCode = 0x5b03
	RplStatusChange                                         AdditionalSenseCode = 0x5c00
	SpindlesSynchronized                                    AdditionalSenseCode = 0x5c01
	SpindlesNotSynchronized                                 AdditionalSenseCode = 0x5c02
	FailurePredictionThresholdExceeded                      AdditionalSenseCode = 0x5d00
	MediaFailurePredictionThresholdExceeded                 AdditionalSenseCode = 0x5d01
	LogicalUnitFailurePredictionThresholdExceeded           AdditionalSenseCode = 0x5d02
	SpareAreaExhaustionPredictionThresholdExceeded          AdditionalSenseCode = 0x5d03
	HardwareImpendingFailureGeneralHardDriveFailure         AdditionalSenseCode = 0x5d10
	HardwareImpendingFailureDriveErrorRateTooHigh           AdditionalSenseCode = 0x5d11
	HardwareImpendingFailureDataErrorRateTooHigh            AdditionalSenseCode = 0x5d12
	HardwareImpendingFailureSeekErrorRateTooHigh            AdditionalSenseCode = 0x5d13
	HardwareImpendingFailureTooManyBlockReassigns           AdditionalSenseCode = 0x5d14
	HardwareImpendingFailureAccessTimesTooHigh              AdditionalSenseCode = 0x5d15
	HardwareImpendingFailureStartUnitTimesTooHigh           AdditionalSenseCode = 0x5d16
	HardwareImpendingFailureChannelParametrics              AdditionalSenseCode = 0x5d17
	HardwareImpendingFailureControllerDetected              AdditionalSenseCode = 0x5d18
	HardwareImpendingFailureThroughputPerformance           AdditionalSenseCode = 0x5d19
	HardwareImpendingFailureSeekTimePerformance             AdditionalSenseCode = 0x5d1a
	HardwareImpendingFailureSpinUpRetryCount                AdditionalSenseCode = 0x5d1b
	HardwareImpendingFailureDriveCalibrationRetryCount      AdditionalSenseCode = 0x5d1c
	HardwareImpendingFailurePowerLossProtectionCircuit      AdditionalSenseCode = 0x5d1d
	ControllerImpendingFailureGeneralHardDriveFailure       AdditionalSenseCode = 0x5d20
	ControllerImpendingFailureDriveErrorRateTooHigh         AdditionalSenseCode = 0x5d21
	ControllerImpendingFailureDataErrorRateTooHigh          AdditionalSenseCode = 0x5d22
	ControllerImpendingFailureSeekErrorRateTooHigh          AdditionalSenseCode = 0x5d23
	ControllerImpendingFailureTooManyBlockReassigns         AdditionalSenseCode = 0x5d24
	ControllerImpendingFailureAccessTimesTooHigh            AdditionalSenseCode = 0x5d25
	ControllerImpendingFailureStartUnitTimesTooHigh         AdditionalSenseCode = 0x5d26
	ControllerImpendingFailureChannelParametrics            AdditionalSenseCode = 0x5d27
	ControllerImpendingFailureControllerDetected            AdditionalSenseCode = 0x5d28
	ControllerImpendingFailureThroughputPerformance         AdditionalSenseCode = 0x5d29
	ControllerImpendingFailureSeekTimePerformance           AdditionalSenseCode = 0x5d2a
	ControllerImpendingFailureSpinUpRetryCount              AdditionalSenseCode = 0x5d2b
	ControllerImpendingFailureDriveCalibrationRetryCount    AdditionalSenseCode = 0x5d2c
	DataChannelImpendingFailureGeneralHardDriveFailure      AdditionalSenseCode = 0x5d30
	DataChannelImpendingFailureDriveErrorRateTooHigh        AdditionalSenseCode = 0x5d31
	DataChannelImpendingFailureDataErrorRateTooHigh         AdditionalSenseCode = 0x5d32
	DataChannelImpendingFailureSeekErrorRateTooHigh         AdditionalSenseCode = 0x5d33
	DataChannelImpendingFailureTooManyBlockReassigns        AdditionalSenseCode = 0x5d34
	DataChannelImpendingFailureAccessTimesTooHigh           AdditionalSenseCode = 0x5d35
	DataChannelImpendingFailureStartUnitTimesTooHigh        AdditionalSenseCode = 0x5d36
	DataChannelImpendingFailureChannelParametrics           AdditionalSenseCode = 0x5d37
	DataChannelImpendingFailureControllerDetected           AdditionalSenseCode = 0x5d38
	DataChannelImpendingFailureThroughputPerformance        AdditionalSenseCode = 0x5d39
	DataChannelImpendingFailureSeekTimePerformance          AdditionalSenseCode = 0x5d3a
	DataChannelImpendingFailureSpinUpRetryCount             AdditionalSenseCode = 0x5d3b
	DataChannelImpendingFailureDriveCalibrationRetryCount   AdditionalSenseCode = 0x5d3c
	ServoImpendingFailureGeneralHardDriveFailure            AdditionalSenseCode = 0x5d40
	ServoImpendingFailureDriveErrorRateTooHigh              AdditionalSenseCode = 0x5d41
	ServoImpendingFailureDataErrorRateTooHigh               AdditionalSenseCode = 0x5d42
	ServoImpendingFailureSeekErrorRateTooHigh               AdditionalSenseCode = 0x5d43
	ServoImpendingFailureTooManyBlockReassigns              AdditionalSenseCode = 0x5d44
	ServoImpendingFailureAccessTimesTooHigh                 AdditionalSenseCode = 0x5d45
	ServoImpendingFailureStartUnitTimesTooHigh              AdditionalSenseCode = 0x5d46
	ServoImpendingFailureChannelParametrics                 AdditionalSenseCode = 0x5d47
	ServoImpendingFailureControllerDetected                 AdditionalSenseCode = 0x5d48
	ServoImpendingFailureThroughputPerformance              AdditionalSenseCode = 0x5d49
	ServoImpendingFailureSeekTimePerformance                AdditionalSenseCode = 0x5d4a
	ServoImpendingFailureSpinUpRetryCount                   AdditionalSenseCode = 0x5d4b
	ServoImpendingFailureDriveCalibrationRetryCount         AdditionalSenseCode = 0x5d4c
	SpindleImpendingFailureGeneralHardDriveFailure          AdditionalSenseCode = 0x5d50
	SpindleImpendingFailureDriveErrorRateTooHigh            AdditionalSenseCode = 0x5d51
	SpindleImpendingFailureDataErrorRateTooHigh             AdditionalSenseCode = 0x5d52
	SpindleImpendingFailureSeekErrorRateTooHigh             AdditionalSenseCode = 0x5d53
	SpindleImpendingFailureTooManyBlockReassigns            AdditionalSenseCode = 0x5d54
	SpindleImpendingFailureAccessTimesTooHigh               AdditionalSenseCode = 0x5d55
	SpindleImpendingFailureStartUnitTimesTooHigh            AdditionalSenseCode = 0x5d56
	SpindleImpendingFailureChannelParametrics               AdditionalSenseCode = 0x5d57
	SpindleImpendingFailureControllerDetected               AdditionalSenseCode = 0x5d58
	SpindleImpendingFailureThroughputPerformance            AdditionalSenseCode = 0x5d59
	SpindleImpendingFailureSeekTimePerformance              AdditionalSenseCode = 0x5d5a
	SpindleImpendingFailureSpinUpRetryCount                 AdditionalSenseCode = 0x5d5b
	SpindleImpendingFailureDriveCalibrationRetryCount       AdditionalSenseCode = 0x5d5c
	FirmwareImpendingFailureGeneralHardDriveFailure         AdditionalSenseCode = 0x5d60
	FirmwareImpendingFailureDriveErrorRateTooHigh           AdditionalSenseCode = 0x5d61
	FirmwareImpendingFailureDataErrorRateTooHigh            AdditionalSenseCode = 0x5d62
	FirmwareImpendingFailureSeekErrorRateTooHigh            AdditionalSenseCode = 0x5d63
	FirmwareImpendingFailureTooManyBlockReassigns           AdditionalSenseCode = 0x5d64
	FirmwareImpendingFailureAccessTimesTooHigh              AdditionalSenseCode = 0x5d65
	FirmwareImpendingFailureStartUnitTimesTooHigh           AdditionalSenseCode = 0x5d66
	FirmwareImpendingFailureChannelParametrics              AdditionalSenseCode = 0x5d67
	FirmwareImpendingFailureControllerDetected              AdditionalSenseCode = 0x5d68
	FirmwareImpendingFailureThroughputPerformance           AdditionalSenseCode = 0x5d69
	FirmwareImpendingFailureSeekTimePerformance             AdditionalSenseCode = 0x5d6a
	FirmwareImpendingFailureSpinUpRetryCount                AdditionalSenseCode = 0x5d6b
	FirmwareImpendingFailureDriveCalibrationRetryCount      AdditionalSenseCode = 0x5d6c
	MediaImpendingFailureEnduranceLimitMet                  AdditionalSenseCode = 0x5d73
	FailurePredictionThresholdExceededfalse                 AdditionalSenseCode = 0x5dff
	LowPowerConditionOn                                     AdditionalSenseCode = 0x5e00
	IdleConditionActivatedByTimer                           AdditionalSenseCode = 0x5e01
	StandbyConditionActivatedByTimer                        AdditionalSenseCode = 0x5e02
	IdleConditionActivatedByCommand                         AdditionalSenseCode = 0x5e03
	StandbyConditionActivatedByCommand                      AdditionalSenseCode = 0x5e04
	IdleBConditionActivatedByTimer                          AdditionalSenseCode = 0x5e05
	IdleBConditionActivatedByCommand                        AdditionalSenseCode = 0x5e06
	IdleCConditionActivatedByTimer                          AdditionalSenseCode = 0x5e07
	IdleCConditionActivatedByCommand                        AdditionalSenseCode = 0x5e08
	StandbyYConditionActivatedByTimer                       AdditionalSenseCode = 0x5e09
	StandbyYConditionActivatedByCommand                     AdditionalSenseCode = 0x5e0a
	PowerStateChangeToActive                                AdditionalSenseCode = 0x5e41
	PowerStateChangeToIdle                                  AdditionalSenseCode = 0x5e42
	PowerStateChangeToStandby                               AdditionalSenseCode = 0x5e43
	PowerStateChangeToSleep                                 AdditionalSenseCode = 0x5e45
	PowerStateChangeToDeviceControl                         AdditionalSenseCode = 0x5e47
	LampFailure                                             AdditionalSenseCode = 0x6000
	VideoAcquisitionError                                   AdditionalSenseCode = 0x6100
	UnableToAcquireVideo                                    AdditionalSenseCode = 0x6101
	OutOfFocus                                              AdditionalSenseCode = 0x6102
	ScanHeadPositioningError                                AdditionalSenseCode = 0x6200
	EndOfUserAreaEncounteredOnThisTrack                     AdditionalSenseCode = 0x6300
	PacketDoesNotFitInAvailableSpace                        AdditionalSenseCode = 0x6301
	IllegalModeForThisTrack                                 AdditionalSenseCode = 0x6400
	InvalidPacketSize                                       AdditionalSenseCode = 0x6401
	VoltageFault                                            AdditionalSenseCode = 0x6500
	AutomaticDocumentFeederCoverUp                          AdditionalSenseCode = 0x6600
	AutomaticDocumentFeederLiftUp                           AdditionalSenseCode = 0x6601
	DocumentJamInAutomaticDocumentFeeder                    AdditionalSenseCode = 0x6602
	DocumentMissFeedAutomaticInDocumentFeeder               AdditionalSenseCode = 0x6603
	ConfigurationFailure                                    AdditionalSenseCode = 0x6700
	ConfigurationOfIncapableLogicalUnitsFailed              AdditionalSenseCode = 0x6701
	AddLogicalUnitFailed                                    AdditionalSenseCode = 0x6702
	ModificationOfLogicalUnitFailed                         AdditionalSenseCode = 0x6703
	ExchangeOfLogicalUnitFailed                             AdditionalSenseCode = 0x6704
	RemoveOfLogicalUnitFailed                               AdditionalSenseCode = 0x6705
	AttachmentOfLogicalUnitFailed                           AdditionalSenseCode = 0x6706
	CreationOfLogicalUnitFailed                             AdditionalSenseCode = 0x6707
	AssignFailureOccurred                                   AdditionalSenseCode = 0x6708
	MultiplyAssignedLogicalUnit                             AdditionalSenseCode = 0x6709
	SetTargetPortGroupsCommandFailed                        AdditionalSenseCode = 0x670a
	AtaDeviceFeatureNotEnabled                              AdditionalSenseCode = 0x670b
	CommandRejected                                         AdditionalSenseCode = 0x670c
	ExplicitBindNotAllowed                                  AdditionalSenseCode = 0x670d
	LogicalUnitNotConfigured                                AdditionalSenseCode = 0x6800
	SubsidiaryLogicalUnitNotConfigured                      AdditionalSenseCode = 0x6801
	DataLossOnLogicalUnit                                   AdditionalSenseCode = 0x6900
	MultipleLogicalUnitFailures                             AdditionalSenseCode = 0x6901
	ParitydataMismatch                                      AdditionalSenseCode = 0x6902
	InformationalReferToLog                                 AdditionalSenseCode = 0x6a00
	StateChangeHasOccurred                                  AdditionalSenseCode = 0x6b00
	RedundancyLevelGotBetter                                AdditionalSenseCode = 0x6b01
	RedundancyLevelGotWorse                                 AdditionalSenseCode = 0x6b02
	RebuildFailureOccurred                                  AdditionalSenseCode = 0x6c00
	RecalculateFailureOccurred                              AdditionalSenseCode = 0x6d00
	CommandToLogicalUnitFailed                              AdditionalSenseCode = 0x6e00
	CopyProtectionKeyExchangeFailureAuthentication          AdditionalSenseCode = 0x6f00
	CopyProtectionKeyExchangeFailureKeyNotPresent           AdditionalSenseCode = 0x6f01
	CopyProtectionKeyExchangeFailureKeyNotEstablished       AdditionalSenseCode = 0x6f02
	ReadOfScrambledSectorWithoutAuthentication              AdditionalSenseCode = 0x6f03
	MediaRegionCodeIsMismatchedToLogicalUnitRegion          AdditionalSenseCode = 0x6f04
	DriveRegionMustBePermanentregionResetCountError         AdditionalSenseCode = 0x6f05
	InsufficientBlockCountForBindingNonceRecording          AdditionalSenseCode = 0x6f06
	ConflictInBindingNonceRecording                         AdditionalSenseCode = 0x6f07
	InsufficientPermission                                  AdditionalSenseCode = 0x6f08
	InvalidDriveHostPairingServer                           AdditionalSenseCode = 0x6f09
	DriveHostPairingSuspended                               AdditionalSenseCode = 0x6f0a
	DecompressionExceptionLongAlgorithmId                   AdditionalSenseCode = 0x7100
	SessionFixationError                                    AdditionalSenseCode = 0x7200
	SessionFixationErrorWritingLeadIn                       AdditionalSenseCode = 0x7201
	SessionFixationErrorWritingLeadOut                      AdditionalSenseCode = 0x7202
	SessionFixationErrorIncompleteTrackInSession            AdditionalSenseCode = 0x7203
	EmptyOrPartiallyWrittenReservedTrack                    AdditionalSenseCode = 0x7204
	NoMoreTrackReservationsAllowed                          AdditionalSenseCode = 0x7205
	RmzExtensionIsNotAllowed                                AdditionalSenseCode = 0x7206
	NoMoreTestZoneExtensionsAreAllowed                      AdditionalSenseCode = 0x7207
	CdControlError                                          AdditionalSenseCode = 0x7300
	PowerCalibrationAreaAlmostFull                          AdditionalSenseCode = 0x7301
	PowerCalibrationAreaIsFull                              AdditionalSenseCode = 0x7302
	PowerCalibrationAreaError                               AdditionalSenseCode = 0x7303
	ProgramMemoryAreaUpdateFailure                          AdditionalSenseCode = 0x7304
	ProgramMemoryAreaIsFull                                 AdditionalSenseCode = 0x7305
	RmapmaIsAlmostFull                                      AdditionalSenseCode = 0x7306
	CurrentPowerCalibrationAreaAlmostFull                   AdditionalSenseCode = 0x7310
	CurrentPowerCalibrationAreaIsFull                       AdditionalSenseCode = 0x7311
	RdzIsFull                                               AdditionalSenseCode = 0x7317
	SecurityError                                           AdditionalSenseCode = 0x7400
	UnableToDecryptData                                     AdditionalSenseCode = 0x7401
	UnencryptedDataEncounteredWhileDecrypting               AdditionalSenseCode = 0x7402
	IncorrectDataEncryptionKey                              AdditionalSenseCode = 0x7403
	CryptographicIntegrityValidationFailed                  AdditionalSenseCode = 0x7404
	ErrorDecryptingData                                     AdditionalSenseCode = 0x7405
	UnknownSignatureVerificationKey                         AdditionalSenseCode = 0x7406
	EncryptionParametersNotUseable                          AdditionalSenseCode = 0x7407
	DigitalSignatureValidationFailure                       AdditionalSenseCode = 0x7408
	EncryptionModeMismatchOnRead                            AdditionalSenseCode = 0x7409
	EncryptedBlockNotRawReadEnabled                         AdditionalSenseCode = 0x740a
	IncorrectEncryptionParameters                           AdditionalSenseCode = 0x740b
	UnableToDecryptParameterList                            AdditionalSenseCode = 0x740c
	EncryptionAlgorithmDisabled                             AdditionalSenseCode = 0x740d
	SaCreationParameterValueInvalid                         AdditionalSenseCode = 0x7410
	SaCreationParameterValueRejected                        AdditionalSenseCode = 0x7411
	InvalidSaUsage                                          AdditionalSenseCode = 0x7412
	DataEncryptionConfigurationPrevented                    AdditionalSenseCode = 0x7421
	SaCreationParameterNotSupported                         AdditionalSenseCode = 0x7430
	AuthenticationFailed                                    AdditionalSenseCode = 0x7440
	ExternalDataEncryptionKeyManagerAccessError             AdditionalSenseCode = 0x7461
	ExternalDataEncryptionKeyManagerError                   AdditionalSenseCode = 0x7462
	ExternalDataEncryptionKeyNotFound                       AdditionalSenseCode = 0x7463
	ExternalDataEncryptionRequestNotAuthorized              AdditionalSenseCode = 0x7464
	ExternalDataEncryptionControlTimeout                    AdditionalSenseCode = 0x746e
	ExternalDataEncryptionControlError                      AdditionalSenseCode = 0x746f
	LogicalUnitAccessNotAuthorized                          AdditionalSenseCode = 0x7471
	SecurityConflictInTranslatedDevice                      AdditionalSenseCode = 0x7479
)

var additionalSenseCodeDesc = map[AdditionalSenseCode]string{
	0x0000: "no additional sense information",
	0x0001: "filemark detected",
	0x0002: "end-of-partition/medium detected",
	0x0003: "setmark detected",
	0x0004: "beginning-of-partition/medium detected",
	0x0005: "end-of-data detected",
	0x0006: "i/o process terminated",
	0x0007: "programmable early warning detected",
	0x0011: "audio play operation in progress",
	0x0012: "audio play operation paused",
	0x0013: "audio play operation successfully completed",
	0x0014: "audio play operation stopped due to error",
	0x0015: "no current audio status to return",
	0x0016: "operation in progress",
	0x0017: "cleaning requested",
	0x0018: "erase operation in progress",
	0x0019: "locate operation in progress",
	0x001a: "rewind operation in progress",
	0x001b: "set capacity operation in progress",
	0x001c: "verify operation in progress",
	0x001d: "ata pass through information available",
	0x001e: "conflicting sa creation request",
	0x001f: "logical unit transitioning to another power condition",
	0x0020: "extended copy information available",
	0x0021: "atomic command aborted due to aca",
	0x0022: "deferred microcode is pending",
	0x0100: "no index/sector signal",
	0x0200: "no seek complete",
	0x0300: "peripheral device write fault",
	0x0301: "no write current",
	0x0302: "excessive write errors",
	0x0400: "logical unit not ready, cause not reportable",
	0x0401: "logical unit is in process of becoming ready",
	0x0402: "logical unit not ready, initializing command required",
	0x0403: "logical unit not ready, manual intervention required",
	0x0404: "logical unit not ready, format in progress",
	0x0405: "logical unit not ready, rebuild in progress",
	0x0406: "logical unit not ready, recalculation in progress",
	0x0407: "logical unit not ready, operation in progress",
	0x0408: "logical unit not ready, long write in progress",
	0x0409: "logical unit not ready, self-test in progress",
	0x040a: "logical unit not accessible, asymmetric access state transition",
	0x040b: "logical unit not accessible, target port in standby state",
	0x040c: "logical unit not accessible, target port in unavailable state",
	0x040d: "logical unit not ready, structure check required",
	0x040e: "logical unit not ready, security session in progress",
	0x0410: "logical unit not ready, auxiliary memory not accessible",
	0x0411: "logical unit not ready, notify (enable spinup) required",
	0x0412: "logical unit not ready, offline",
	0x0413: "logical unit not ready, sa creation in progress",
	0x0414: "logical unit not ready, space allocation in progress",
	0x0415: "logical unit not ready, robotics disabled",
	0x0416: "logical unit not ready, configuration required",
	0x0417: "logical unit not ready, calibration required",
	0x0418: "logical unit not ready, a door is open",
	0x0419: "logical unit not ready, operating in sequential mode",
	0x041a: "logical unit not ready, start stop unit command in progress",
	0x041b: "logical unit not ready, sanitize in progress",
	0x041c: "logical unit not ready, additional power use not yet granted",
	0x041d: "logical unit not ready, configuration in progress",
	0x041e: "logical unit not ready, microcode activation required",
	0x041f: "logical unit not ready, microcode download required",
	0x0420: "logical unit not ready, logical unit reset required",
	0x0421: "logical unit not ready, hard reset required",
	0x0422: "logical unit not ready, power cycle required",
	0x0423: "logical unit not ready, affiliation required",
	0x0424: "depopulation in progress",
	0x0500: "logical unit does not respond to selection",
	0x0600: "no reference position found",
	0x0700: "multiple peripheral devices selected",
	0x0800: "logical unit communication failure",
	0x0801: "logical unit communication time-out",
	0x0802: "logical unit communication parity error",
	0x0803: "logical unit communication crc error (ultra-dma/32)",
	0x0804: "unreachable copy target",
	0x0900: "track following error",
	0x0901: "tracking servo failure",
	0x0902: "focus servo failure",
	0x0903: "spindle servo failure",
	0x0904: "head select fault",
	0x0905: "vibration induced tracking error",
	0x0a00: "error log overflow",
	0x0b00: "warning",
	0x0b01: "warning - specified temperature exceeded",
	0x0b02: "warning - enclosure degraded",
	0x0b03: "warning - background self-test failed",
	0x0b04: "warning - background pre-scan detected medium error",
	0x0b05: "warning - background medium scan detected medium error",
	0x0b06: "warning - non-volatile cache now volatile",
	0x0b07: "warning - degraded power to non-volatile cache",
	0x0b08: "warning - power loss expected",
	0x0b09: "warning - device statistics notification active",
	0x0b0a: "warning - high critical temperature limit exceeded",
	0x0b0b: "warning - low critical temperature limit exceeded",
	0x0b0c: "warning - high operating temperature limit exceeded",
	0x0b0d: "warning - low operating temperature limit exceeded",
	0x0b0e: "warning - high critical humidity limit exceeded",
	0x0b0f: "warning - low critical humidity limit exceeded",
	0x0b10: "warning - high operating humidity limit exceeded",
	0x0b11: "warning - low operating humidity limit exceeded",
	0x0b12: "warning - microcode security at risk",
	0x0b13: "warning - microcode digital signature validation failure",
	0x0b14: "warning - physical element status change",
	0x0c00: "write error",
	0x0c01: "write error - recovered with auto reallocation",
	0x0c02: "write error - auto reallocation failed",
	0x0c03: "write error - recommend reassignment",
	0x0c04: "compression check miscompare error",
	0x0c05: "data expansion occurred during compression",
	0x0c06: "block not compressible",
	0x0c07: "write error - recovery needed",
	0x0c08: "write error - recovery failed",
	0x0c09: "write error - loss of streaming",
	0x0c0a: "write error - padding blocks added",
	0x0c0b: "auxiliary memory write error",
	0x0c0c: "write error - unexpected unsolicited data",
	0x0c0d: "write error - not enough unsolicited data",
	0x0c0e: "multiple write errors",
	0x0c0f: "defects in error window",
	0x0c10: "incomplete multiple atomic write operations",
	0x0c11: "write error - recovery scan needed",
	0x0c12: "write error - insufficient zone resources",
	0x0d00: "error detected by third party temporary initiator",
	0x0d01: "third party device failure",
	0x0d02: "copy target device not reachable",
	0x0d03: "incorrect copy target device type",
	0x0d04: "copy target device data underrun",
	0x0d05: "copy target device data overrun",
	0x0e00: "invalid information unit",
	0x0e01: "information unit too short",
	0x0e02: "information unit too long",
	0x0e03: "invalid field in command information unit",
	0x1000: "id crc or ecc error",
	0x1001: "logical block guard check failed",
	0x1002: "logical block application tag check failed",
	0x1003: "logical block reference tag check failed",
	0x1004: "logical block protection error on recover buffered data",
	0x1005: "logical block protection method error",
	0x1100: "unrecovered read error",
	0x1101: "read retries exhausted",
	0x1102: "error too long to correct",
	0x1103: "multiple read errors",
	0x1104: "unrecovered read error - auto reallocate failed",
	0x1105: "l-ec uncorrectable error",
	0x1106: "circ unrecovered error",
	0x1107: "data re-synchronization error",
	0x1108: "incomplete block read",
	0x1109: "no gap found",
	0x110a: "miscorrected error",
	0x110b: "unrecovered read error - recommend reassignment",
	0x110c: "unrecovered read error - recommend rewrite the data",
	0x110d: "de-compression crc error",
	0x110e: "cannot decompress using declared algorithm",
	0x110f: "error reading upc/ean number",
	0x1110: "error reading isrc number",
	0x1111: "read error - loss of streaming",
	0x1112: "auxiliary memory read error",
	0x1113: "read error - failed retransmission request",
	0x1114: "read error - lba marked bad by application client",
	0x1115: "write after sanitize required",
	0x1200: "address mark not found for id field",
	0x1300: "address mark not found for data field",
	0x1400: "recorded entity not found",
	0x1401: "record not found",
	0x1402: "filemark or setmark not found",
	0x1403: "end-of-data not found",
	0x1404: "block sequence error",
	0x1405: "record not found - recommend reassignment",
	0x1406: "record not found - data auto-reallocated",
	0x1407: "locate operation failure",
	0x1500: "random positioning error",
	0x1501: "mechanical positioning error",
	0x1502: "positioning error detected by read of medium",
	0x1600: "data synchronization mark error",
	0x1601: "data sync error - data rewritten",
	0x1602: "data sync error - recommend rewrite",
	0x1603: "data sync error - data auto-reallocated",
	0x1604: "data sync error - recommend reassignment",
	0x1700: "recovered data with no error correction applied",
	0x1701: "recovered data with retries",
	0x1702: "recovered data with positive head offset",
	0x1703: "recovered data with negative head offset",
	0x1704: "recovered data with retries and/or circ applied",
	0x1705: "recovered data using previous sector id",
	0x1706: "recovered data without ecc - data auto-reallocated",
	0x1707: "recovered data without ecc - recommend reassignment",
	0x1708: "recovered data without ecc - recommend rewrite",
	0x1709: "recovered data without ecc - data rewritten",
	0x1800: "recovered data with error correction applied",
	0x1801: "recovered data with error corr. & retries applied",
	0x1802: "recovered data - data auto-reallocated",
	0x1803: "recovered data with circ",
	0x1804: "recovered data with l-ec",
	0x1805: "recovered data - recommend reassignment",
	0x1806: "recovered data - recommend rewrite",
	0x1807: "recovered data with ecc - data rewritten",
	0x1808: "recovered data with linking",
	0x1900: "defect list error",
	0x1901: "defect list not available",
	0x1902: "defect list error in primary list",
	0x1903: "defect list error in grown list",
	0x1a00: "parameter list length error",
	0x1b00: "synchronous data transfer error",
	0x1c00: "defect list not found",
	0x1c01: "primary defect list not found",
	0x1c02: "grown defect list not found",
	0x1d00: "miscompare during verify operation",
	0x1d01: "miscompare verify of unmapped lba",
	0x1e00: "recovered id with ecc correction",
	0x1f00: "partial defect list transfer",
	0x2000: "invalid command operation code",
	0x2001: "access denied - initiator pending-enrolled",
	0x2002: "access denied - no access rights",
	0x2003: "access denied - invalid mgmt id key",
	0x2004: "illegal command while in write capable state",
	0x2006: "illegal command while in explicit address mode",
	0x2007: "illegal command while in implicit address mode",
	0x2008: "access denied - enrollment conflict",
	0x2009: "access denied - invalid lu identifier",
	0x200a: "access denied - invalid proxy token",
	0x200b: "access denied - acl lun conflict",
	0x200c: "illegal command when not in append-only mode",
	0x200d: "not an administrative logical unit",
	0x200e: "not a subsidiary logical unit",
	0x200f: "not a conglomerate logical unit",
	0x2100: "logical block address out of range",
	0x2101: "invalid element address",
	0x2102: "invalid address for write",
	0x2103: "invalid write crossing layer jump",
	0x2104: "unaligned write command",
	0x2105: "write boundary violation",
	0x2106: "attempt to read invalid data",
	0x2107: "read boundary violation",
	0x2108: "misaligned write command",
	0x2200: "illegal function (use 20 00, 24 00, or 26 00)",
	0x2300: "invalid token operation, cause not reportable",
	0x2301: "invalid token operation, unsupported token type",
	0x2302: "invalid token operation, remote token usage not supported",
	0x2303: "invalid token operation, remote rod token creation not supported",
	0x2304: "invalid token operation, token unknown",
	0x2305: "invalid token operation, token corrupt",
	0x2306: "invalid token operation, token revoked",
	0x2307: "invalid token operation, token expired",
	0x2308: "invalid token operation, token cancelled",
	0x2309: "invalid token operation, token deleted",
	0x230a: "invalid token operation, invalid token length",
	0x2400: "invalid field in cdb",
	0x2401: "cdb decryption error",
	0x2404: "security audit value frozen",
	0x2405: "security working key frozen",
	0x2406: "nonce not unique",
	0x2407: "nonce timestamp out of range",
	0x2408: "invalid xcdb",
	0x2409: "invalid fast format",
	0x2500: "logical unit not supported",
	0x2600: "invalid field in parameter list",
	0x2601: "parameter not supported",
	0x2602: "parameter value invalid",
	0x2603: "threshold parameters not supported",
	0x2604: "invalid release of persistent reservation",
	0x2605: "data decryption error",
	0x2606: "too many target descriptors",
	0x2607: "unsupported target descriptor type code",
	0x2608: "too many segment descriptors",
	0x2609: "unsupported segment descriptor type code",
	0x260a: "unexpected inexact segment",
	0x260b: "inline data length exceeded",
	0x260c: "invalid operation for copy source or destination",
	0x260d: "copy segment granularity violation",
	0x260e: "invalid parameter while port is enabled",
	0x260f: "invalid data-out buffer integrity check value",
	0x2610: "data decryption key fail limit reached",
	0x2611: "incomplete key-associated data set",
	0x2612: "vendor specific key reference not found",
	0x2613: "application tag mode page is invalid",
	0x2614: "tape stream mirroring prevented",
	0x2615: "copy source or copy destination not authorized",
	0x2700: "write protected",
	0x2701: "hardware write protected",
	0x2702: "logical unit software write protected",
	0x2703: "associated write protect",
	0x2704: "persistent write protect",
	0x2705: "permanent write protect",
	0x2706: "conditional write protect",
	0x2707: "space allocation failed write protect",
	0x2708: "zone is read only",
	0x2800: "not ready to ready change, medium may have changed",
	0x2801: "import or export element accessed",
	0x2802: "format-layer may have changed",
	0x2803: "import/export element accessed, medium changed",
	0x2900: "power on, reset, or bus device reset occurred",
	0x2901: "power on occurred",
	0x2902: "scsi bus reset occurred",
	0x2903: "bus device reset function occurred",
	0x2904: "device internal reset",
	0x2905: "transceiver mode changed to single-ended",
	0x2906: "transceiver mode changed to lvd",
	0x2907: "i_t nexus loss occurred",
	0x2a00: "parameters changed",
	0x2a01: "mode parameters changed",
	0x2a02: "log parameters changed",
	0x2a03: "reservations preempted",
	0x2a04: "reservations released",
	0x2a05: "registrations preempted",
	0x2a06: "asymmetric access state changed",
	0x2a07: "implicit asymmetric access state transition failed",
	0x2a08: "priority changed",
	0x2a09: "capacity data has changed",
	0x2a0a: "error history i_t nexus cleared",
	0x2a0b: "error history snapshot released",
	0x2a0c: "error recovery attributes have changed",
	0x2a0d: "data encryption capabilities changed",
	0x2a10: "timestamp changed",
	0x2a11: "data encryption parameters changed by another i_t nexus",
	0x2a12: "data encryption parameters changed by vendor specific event",
	0x2a13: "data encryption key instance counter has changed",
	0x2a14: "sa creation capabilities data has changed",
	0x2a15: "medium removal prevention preempted",
	0x2a16: "zone reset write pointer recommended",
	0x2b00: "copy cannot execute since host cannot disconnect",
	0x2c00: "command sequence error",
	0x2c01: "too many windows specified",
	0x2c02: "invalid combination of windows specified",
	0x2c03: "current program area is not empty",
	0x2c04: "current program area is empty",
	0x2c05: "illegal power condition request",
	0x2c06: "persistent prevent conflict",
	0x2c07: "previous busy status",
	0x2c08: "previous task set full status",
	0x2c09: "previous reservation conflict status",
	0x2c0a: "partition or collection contains user objects",
	0x2c0b: "not reserved",
	0x2c0c: "orwrite generation does not match",
	0x2c0d: "reset write pointer not allowed",
	0x2c0e: "zone is offline",
	0x2c0f: "stream not open",
	0x2c10: "unwritten data in zone",
	0x2c11: "descriptor format sense data required",
	0x2d00: "overwrite error on update in place",
	0x2e00: "insufficient time for operation",
	0x2e01: "command timeout before processing",
	0x2e02: "command timeout during processing",
	0x2e03: "command timeout during processing due to error recovery",
	0x2f00: "commands cleared by another initiator",
	0x2f01: "commands cleared by power loss notification",
	0x2f02: "commands cleared by device server",
	0x2f03: "some commands cleared by queuing layer event",
	0x3000: "incompatible medium installed",
	0x3001: "cannot read medium - unknown format",
	0x3002: "cannot read medium - incompatible format",
	0x3003: "cleaning cartridge installed",
	0x3004: "cannot write medium - unknown format",
	0x3005: "cannot write medium - incompatible format",
	0x3006: "cannot format medium - incompatible medium",
	0x3007: "cleaning failure",
	0x3008: "cannot write - application code mismatch",
	0x3009: "current session not fixated for append",
	0x300a: "cleaning request rejected",
	0x300c: "worm medium - overwrite attempted",
	0x300d: "worm medium - integrity check",
	0x3010: "medium not formatted",
	0x3011: "incompatible volume type",
	0x3012: "incompatible volume qualifier",
	0x3013: "cleaning volume expired",
	0x3100: "medium format corrupted",
	0x3101: "format command failed",
	0x3102: "zoned formatting failed due to spare linking",
	0x3103: "sanitize command failed",
	0x3104: "depopulation failed",
	0x3200: "no defect spare location available",
	0x3201: "defect list update failure",
	0x3300: "tape length error",
	0x3400: "enclosure failure",
	0x3500: "enclosure services failure",
	0x3501: "unsupported enclosure function",
	0x3502: "enclosure services unavailable",
	0x3503: "enclosure services transfer failure",
	0x3504: "enclosure services transfer refused",
	0x3505: "enclosure services checksum error",
	0x3600: "ribbon, ink, or toner failure",
	0x3700: "rounded parameter",
	0x3800: "event status notification",
	0x3802: "esn - power management class event",
	0x3804: "esn - media class event",
	0x3806: "esn - device busy class event",
	0x3807: "thin provisioning soft threshold reached",
	0x3900: "saving parameters not supported",
	0x3a00: "medium not present",
	0x3a01: "medium not present - tray closed",
	0x3a02: "medium not present - tray open",
	0x3a03: "medium not present - loadable",
	0x3a04: "medium not present - medium auxiliary memory accessible",
	0x3b00: "sequential positioning error",
	0x3b01: "tape position error at beginning-of-medium",
	0x3b02: "tape position error at end-of-medium",
	0x3b03: "tape or electronic vertical forms unit not ready",
	0x3b04: "slew failure",
	0x3b05: "paper jam",
	0x3b06: "failed to sense top-of-form",
	0x3b07: "failed to sense bottom-of-form",
	0x3b08: "reposition error",
	0x3b09: "read past end of medium",
	0x3b0a: "read past beginning of medium",
	0x3b0b: "position past end of medium",
	0x3b0c: "position past beginning of medium",
	0x3b0d: "medium destination element full",
	0x3b0e: "medium source element empty",
	0x3b0f: "end of medium reached",
	0x3b11: "medium magazine not accessible",
	0x3b12: "medium magazine removed",
	0x3b13: "medium magazine inserted",
	0x3b14: "medium magazine locked",
	0x3b15: "medium magazine unlocked",
	0x3b16: "mechanical positioning or changer error",
	0x3b17: "read past end of user object",
	0x3b18: "element disabled",
	0x3b19: "element enabled",
	0x3b1a: "data transfer device removed",
	0x3b1b: "data transfer device inserted",
	0x3b1c: "too many logical objects on partition to support operation",
	0x3d00: "invalid bits in identify message",
	0x3e00: "logical unit has not self-configured yet",
	0x3e01: "logical unit failure",
	0x3e02: "timeout on logical unit",
	0x3e03: "logical unit failed self-test",
	0x3e04: "logical unit unable to update self-test log",
	0x3f00: "target operating conditions have changed",
	0x3f01: "microcode has been changed",
	0x3f02: "changed operating definition",
	0x3f03: "inquiry data has changed",
	0x3f04: "component device attached",
	0x3f05: "device identifier changed",
	0x3f06: "redundancy group created or modified",
	0x3f07: "redundancy group deleted",
	0x3f08: "spare created or modified",
	0x3f09: "spare deleted",
	0x3f0a: "volume set created or modified",
	0x3f0b: "volume set deleted",
	0x3f0c: "volume set deassigned",
	0x3f0d: "volume set reassigned",
	0x3f0e: "reported luns data has changed",
	0x3f0f: "echo buffer overwritten",
	0x3f10: "medium loadable",
	0x3f11: "medium auxiliary memory accessible",
	0x3f12: "iscsi ip address added",
	0x3f13: "iscsi ip address removed",
	0x3f14: "iscsi ip address changed",
	0x3f15: "inspect referrals sense descriptors",
	0x3f16: "microcode has been changed without reset",
	0x3f17: "zone transition to full",
	0x3f18: "bind completed",
	0x3f19: "bind redirected",
	0x3f1a: "subsidiary binding changed",
	0x4000: "ram failure (should use 40 nn)",
	0x4100: "data path failure (should use 40 nn)",
	0x4200: "power-on or self-test failure (should use 40 nn)",
	0x4300: "message error",
	0x4400: "internal target failure",
	0x4401: "persistent reservation information lost",
	0x4471: "ata device failed set features",
	0x4500: "select or reselect failure",
	0x4600: "unsuccessful soft reset",
	0x4700: "scsi parity error",
	0x4701: "data phase crc error detected",
	0x4702: "scsi parity error detected during st data phase",
	0x4703: "information unit iucrc error detected",
	0x4704: "asynchronous information protection error detected",
	0x4705: "protocol service crc error",
	0x4706: "phy test function in progress",
	0x477f: "some commands cleared by iscsi protocol event",
	0x4800: "initiator detected error message received",
	0x4900: "invalid message error",
	0x4a00: "command phase error",
	0x4b00: "data phase error",
	0x4b01: "invalid target port transfer tag received",
	0x4b02: "too much write data",
	0x4b03: "ack/nak timeout",
	0x4b04: "nak received",
	0x4b05: "data offset error",
	0x4b06: "initiator response timeout",
	0x4b07: "connection lost",
	0x4b08: "data-in buffer overflow - data buffer size",
	0x4b09: "data-in buffer overflow - data buffer descriptor area",
	0x4b0a: "data-in buffer error",
	0x4b0b: "data-out buffer overflow - data buffer size",
	0x4b0c: "data-out buffer overflow - data buffer descriptor area",
	0x4b0d: "data-out buffer error",
	0x4b0e: "pcie fabric error",
	0x4b0f: "pcie completion timeout",
	0x4b10: "pcie completer abort",
	0x4b11: "pcie poisoned tlp received",
	0x4b12: "pcie ecrc check failed",
	0x4b13: "pcie unsupported request",
	0x4b14: "pcie acs violation",
	0x4b15: "pcie tlp prefix blocked",
	0x4c00: "logical unit failed self-configuration",
	0x4e00: "overlapped commands attempted",
	0x5000: "write append error",
	0x5001: "write append position error",
	0x5002: "position error related to timing",
	0x5100: "erase failure",
	0x5101: "erase failure - incomplete erase operation detected",
	0x5200: "cartridge fault",
	0x5300: "media load or eject failed",
	0x5301: "unload tape failure",
	0x5302: "medium removal prevented",
	0x5303: "medium removal prevented by data transfer element",
	0x5304: "medium thread or unthread failure",
	0x5305: "volume identifier invalid",
	0x5306: "volume identifier missing",
	0x5307: "duplicate volume identifier",
	0x5308: "element status unknown",
	0x5309: "data transfer device error - load failed",
	0x530a: "data transfer device error - unload failed",
	0x530b: "data transfer device error - unload missing",
	0x530c: "data transfer device error - eject failed",
	0x530d: "data transfer device error - library communication failed",
	0x5400: "scsi to host system interface failure",
	0x5500: "system resource failure",
	0x5501: "system buffer full",
	0x5502: "insufficient reservation resources",
	0x5503: "insufficient resources",
	0x5504: "insufficient registration resources",
	0x5505: "insufficient access control resources",
	0x5506: "auxiliary memory out of space",
	0x5507: "quota error",
	0x5508: "maximum number of supplemental decryption keys exceeded",
	0x5509: "medium auxiliary memory not accessible",
	0x550a: "data currently unavailable",
	0x550b: "insufficient power for operation",
	0x550c: "insufficient resources to create rod",
	0x550d: "insufficient resources to create rod token",
	0x550e: "insufficient zone resources",
	0x550f: "insufficient zone resources to complete write",
	0x5510: "maximum number of streams open",
	0x5511: "insufficient resources to bind",
	0x5700: "unable to recover table-of-contents",
	0x5800: "generation does not exist",
	0x5900: "updated block read",
	0x5a00: "operator request or state change input",
	0x5a01: "operator medium removal request",
	0x5a02: "operator selected write protect",
	0x5a03: "operator selected write permit",
	0x5b00: "log exception",
	0x5b01: "threshold condition met",
	0x5b02: "log counter at maximum",
	0x5b03: "log list codes exhausted",
	0x5c00: "rpl status change",
	0x5c01: "spindles synchronized",
	0x5c02: "spindles not synchronized",
	0x5d00: "failure prediction threshold exceeded",
	0x5d01: "media failure prediction threshold exceeded",
	0x5d02: "logical unit failure prediction threshold exceeded",
	0x5d03: "spare area exhaustion prediction threshold exceeded",
	0x5d10: "hardware impending failure general hard drive failure",
	0x5d11: "hardware impending failure drive error rate too high",
	0x5d12: "hardware impending failure data error rate too high",
	0x5d13: "hardware impending failure seek error rate too high",
	0x5d14: "hardware impending failure too many block reassigns",
	0x5d15: "hardware impending failure access times too high",
	0x5d16: "hardware impending failure start unit times too high",
	0x5d17: "hardware impending failure channel parametrics",
	0x5d18: "hardware impending failure controller detected",
	0x5d19: "hardware impending failure throughput performance",
	0x5d1a: "hardware impending failure seek time performance",
	0x5d1b: "hardware impending failure spin-up retry count",
	0x5d1c: "hardware impending failure drive calibration retry count",
	0x5d1d: "hardware impending failure power loss protection circuit",
	0x5d20: "controller impending failure general hard drive failure",
	0x5d21: "controller impending failure drive error rate too high",
	0x5d22: "controller impending failure data error rate too high",
	0x5d23: "controller impending failure seek error rate too high",
	0x5d24: "controller impending failure too many block reassigns",
	0x5d25: "controller impending failure access times too high",
	0x5d26: "controller impending failure start unit times too high",
	0x5d27: "controller impending failure channel parametrics",
	0x5d28: "controller impending failure controller detected",
	0x5d29: "controller impending failure throughput performance",
	0x5d2a: "controller impending failure seek time performance",
	0x5d2b: "controller impending failure spin-up retry count",
	0x5d2c: "controller impending failure drive calibration retry count",
	0x5d30: "data channel impending failure general hard drive failure",
	0x5d31: "data channel impending failure drive error rate too high",
	0x5d32: "data channel impending failure data error rate too high",
	0x5d33: "data channel impending failure seek error rate too high",
	0x5d34: "data channel impending failure too many block reassigns",
	0x5d35: "data channel impending failure access times too high",
	0x5d36: "data channel impending failure start unit times too high",
	0x5d37: "data channel impending failure channel parametrics",
	0x5d38: "data channel impending failure controller detected",
	0x5d39: "data channel impending failure throughput performance",
	0x5d3a: "data channel impending failure seek time performance",
	0x5d3b: "data channel impending failure spin-up retry count",
	0x5d3c: "data channel impending failure drive calibration retry count",
	0x5d40: "servo impending failure general hard drive failure",
	0x5d41: "servo impending failure drive error rate too high",
	0x5d42: "servo impending failure data error rate too high",
	0x5d43: "servo impending failure seek error rate too high",
	0x5d44: "servo impending failure too many block reassigns",
	0x5d45: "servo impending failure access times too high",
	0x5d46: "servo impending failure start unit times too high",
	0x5d47: "servo impending failure channel parametrics",
	0x5d48: "servo impending failure controller detected",
	0x5d49: "servo impending failure throughput performance",
	0x5d4a: "servo impending failure seek time performance",
	0x5d4b: "servo impending failure spin-up retry count",
	0x5d4c: "servo impending failure drive calibration retry count",
	0x5d50: "spindle impending failure general hard drive failure",
	0x5d51: "spindle impending failure drive error rate too high",
	0x5d52: "spindle impending failure data error rate too high",
	0x5d53: "spindle impending failure seek error rate too high",
	0x5d54: "spindle impending failure too many block reassigns",
	0x5d55: "spindle impending failure access times too high",
	0x5d56: "spindle impending failure start unit times too high",
	0x5d57: "spindle impending failure channel parametrics",
	0x5d58: "spindle impending failure controller detected",
	0x5d59: "spindle impending failure throughput performance",
	0x5d5a: "spindle impending failure seek time performance",
	0x5d5b: "spindle impending failure spin-up retry count",
	0x5d5c: "spindle impending failure drive calibration retry count",
	0x5d60: "firmware impending failure general hard drive failure",
	0x5d61: "firmware impending failure drive error rate too high",
	0x5d62: "firmware impending failure data error rate too high",
	0x5d63: "firmware impending failure seek error rate too high",
	0x5d64: "firmware impending failure too many block reassigns",
	0x5d65: "firmware impending failure access times too high",
	0x5d66: "firmware impending failure start unit times too high",
	0x5d67: "firmware impending failure channel parametrics",
	0x5d68: "firmware impending failure controller detected",
	0x5d69: "firmware impending failure throughput performance",
	0x5d6a: "firmware impending failure seek time performance",
	0x5d6b: "firmware impending failure spin-up retry count",
	0x5d6c: "firmware impending failure drive calibration retry count",
	0x5d73: "media impending failure endurance limit met",
	0x5dff: "failure prediction threshold exceeded (false)",
	0x5e00: "low power condition on",
	0x5e01: "idle condition activated by timer",
	0x5e02: "standby condition activated by timer",
	0x5e03: "idle condition activated by command",
	0x5e04: "standby condition activated by command",
	0x5e05: "idle_b condition activated by timer",
	0x5e06: "idle_b condition activated by command",
	0x5e07: "idle_c condition activated by timer",
	0x5e08: "idle_c condition activated by command",
	0x5e09: "standby_y condition activated by timer",
	0x5e0a: "standby_y condition activated by command",
	0x5e41: "power state change to active",
	0x5e42: "power state change to idle",
	0x5e43: "power state change to standby",
	0x5e45: "power state change to sleep",
	0x5e47: "power state change to device control",
	0x6000: "lamp failure",
	0x6100: "video acquisition error",
	0x6101: "unable to acquire video",
	0x6102: "out of focus",
	0x6200: "scan head positioning error",
	0x6300: "end of user area encountered on this track",
	0x6301: "packet does not fit in available space",
	0x6400: "illegal mode for this track",
	0x6401: "invalid packet size",
	0x6500: "voltage fault",
	0x6600: "automatic document feeder cover up",
	0x6601: "automatic document feeder lift up",
	0x6602: "document jam in automatic document feeder",
	0x6603: "document miss feed automatic in document feeder",
	0x6700: "configuration failure",
	0x6701: "configuration of incapable logical units failed",
	0x6702: "add logical unit failed",
	0x6703: "modification of logical unit failed",
	0x6704: "exchange of logical unit failed",
	0x6705: "remove of logical unit failed",
	0x6706: "attachment of logical unit failed",
	0x6707: "creation of logical unit failed",
	0x6708: "assign failure occurred",
	0x6709: "multiply assigned logical unit",
	0x670a: "set target port groups command failed",
	0x670b: "ata device feature not enabled",
	0x670c: "command rejected",
	0x670d: "explicit bind not allowed",
	0x6800: "logical unit not configured",
	0x6801: "subsidiary logical unit not configured",
	0x6900: "data loss on logical unit",
	0x6901: "multiple logical unit failures",
	0x6902: "parity/data mismatch",
	0x6a00: "informational, refer to log",
	0x6b00: "state change has occurred",
	0x6b01: "redundancy level got better",
	0x6b02: "redundancy level got worse",
	0x6c00: "rebuild failure occurred",
	0x6d00: "recalculate failure occurred",
	0x6e00: "command to logical unit failed",
	0x6f00: "copy protection key exchange failure - authentication",
	0x6f01: "copy protection key exchange failure - key not present",
	0x6f02: "copy protection key exchange failure - key not established",
	0x6f03: "read of scrambled sector without authentication",
	0x6f04: "media region code is mismatched to logical unit region",
	0x6f05: "drive region must be permanent/region reset count error",
	0x6f06: "insufficient block count for binding nonce recording",
	0x6f07: "conflict in binding nonce recording",
	0x6f08: "insufficient permission",
	0x6f09: "invalid drive-host pairing server",
	0x6f0a: "drive-host pairing suspended",
	0x7100: "decompression exception long algorithm id",
	0x7200: "session fixation error",
	0x7201: "session fixation error writing lead-in",
	0x7202: "session fixation error writing lead-out",
	0x7203: "session fixation error - incomplete track in session",
	0x7204: "empty or partially written reserved track",
	0x7205: "no more track reservations allowed",
	0x7206: "rmz extension is not allowed",
	0x7207: "no more test zone extensions are allowed",
	0x7300: "cd control error",
	0x7301: "power calibration area almost full",
	0x7302: "power calibration area is full",
	0x7303: "power calibration area error",
	0x7304: "program memory area update failure",
	0x7305: "program memory area is full",
	0x7306: "rma/pma is almost full",
	0x7310: "current power calibration area almost full",
	0x7311: "current power calibration area is full",
	0x7317: "rdz is full",
	0x7400: "security error",
	0x7401: "unable to decrypt data",
	0x7402: "unencrypted data encountered while decrypting",
	0x7403: "incorrect data encryption key",
	0x7404: "cryptographic integrity validation failed",
	0x7405: "error decrypting data",
	0x7406: "unknown signature verification key",
	0x7407: "encryption parameters not useable",
	0x7408: "digital signature validation failure",
	0x7409: "encryption mode mismatch on read",
	0x740a: "encrypted block not raw read enabled",
	0x740b: "incorrect encryption parameters",
	0x740c: "unable to decrypt parameter list",
	0x740d: "encryption algorithm disabled",
	0x7410: "sa creation parameter value invalid",
	0x7411: "sa creation parameter value rejected",
	0x7412: "invalid sa usage",
	0x7421: "data encryption configuration prevented",
	0x7430: "sa creation parameter not supported",
	0x7440: "authentication failed",
	0x7461: "external data encryption key manager access error",
	0x7462: "external data encryption key manager error",
	0x7463: "external data encryption key not found",
	0x7464: "external data encryption request not authorized",
	0x746e: "external data encryption control timeout",
	0x746f: "external data encryption control error",
	0x7471: "logical unit access not authorized",
	0x7479: "security conflict in translated device",
}

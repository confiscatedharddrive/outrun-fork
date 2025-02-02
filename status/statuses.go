package status

const (
	OK                         = int64(0)
	ServerSecurityError        = int64(-19001)
	VersionDifference          = int64(-19002)
	DecryptionFailure          = int64(-19003)
	ParamHashDifference        = int64(-19004)
	ServerNextVersion          = int64(-19990)
	ServerMaintenance          = int64(-19997)
	ServerBusyError            = int64(-19998)
	ServerSystemError          = int64(-19999)
	RequestParamError          = int64(-10100)
	NotAvailablePlayer         = int64(-10101)
	MissingPlayer              = int64(-10102)
	ExpiredSession             = int64(-10103)
	InvalidPassword            = int64(-10104)
	InvalidSerialCode          = int64(-10105)
	UsedSerialCode             = int64(-10106)
	HSPWebAPIError             = int64(-10115)
	ApolloWebAPIError          = int64(-10115)
	DataMismatch               = int64(-30120)
	MasterDataMismatch         = int64(-10121)
	NotEnoughRedRings          = int64(-21030)
	NotEnoughRings             = int64(-20131)
	NotEnoughEnergy            = int64(-20132)
	RouletteUseLimit           = int64(-30401)
	RouletteBoardReset         = int64(-30411)
	CharacterLevelLimit        = int64(-20601)
	AllChaoLevelLimit          = int64(-20602)
	AlreadyInvitedFriend       = int64(-30801)
	AlreadyRequestedEnergy     = int64(-30901)
	AlreadySentEnergy          = int64(-30902)
	ReceiveFailureMessage      = int64(-30910)
	AlreadyExistedPrePurchase  = int64(-11001)
	AlreadyRemovedPrePurchase  = int64(-11002)
	InvalidReceiptData         = int64(-11003)
	AlreadyProcessedReceipt    = int64(-11004)
	EnergyLimitPurchaseTrigger = int64(-21010)
	NotStartEvent              = int64(-10201)
	AlreadyEndEvent            = int64(-10202)
	UsernameUnacceptable       = int64(-40000) // 2.1.0 and later; generic server side rejection of username change
	UsernameTooLong            = int64(-40001) // 2.1.0 and later; username not saved due to length
	UsernameHasNGWord          = int64(-40002) // 2.1.0 and later; username not saved due to ng word
	VersionForApplication      = int64(-999002)
	Timeout                    = int64(-7)
	OtherError                 = int64(-8)
	NotReachable               = int64(-10)
	InvalidResponse            = int64(-20)
	ClientError                = int64(-400)
	InternalServerError        = int64(-500)
	HSPPurchaseError           = int64(-600)
	ServerBusy                 = int64(-700) // why different from other busy?
)

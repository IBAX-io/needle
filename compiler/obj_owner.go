package compiler

// OwnerInfo contains the owner information of the contract.
type OwnerInfo struct {
	StateId  uint32
	Active   bool
	TableId  int64
	WalletId int64
	TokenId  int64
}

func (*OwnerInfo) isObjInfoValue() {}

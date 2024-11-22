// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ReallocStakeList is the `reallocStakeList` instruction.
type ReallocStakeList struct {
	Capacity *uint32

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] adminAuthority
	//
	// [2] = [WRITE] stakeList
	//
	// [3] = [WRITE, SIGNER] rentFunds
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewReallocStakeListInstructionBuilder creates a new `ReallocStakeList` instruction builder.
func NewReallocStakeListInstructionBuilder() *ReallocStakeList {
	nd := &ReallocStakeList{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetCapacity sets the "capacity" parameter.
func (inst *ReallocStakeList) SetCapacity(capacity uint32) *ReallocStakeList {
	inst.Capacity = &capacity
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *ReallocStakeList) SetStateAccount(state ag_solanago.PublicKey) *ReallocStakeList {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ReallocStakeList) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAuthorityAccount sets the "adminAuthority" account.
func (inst *ReallocStakeList) SetAdminAuthorityAccount(adminAuthority ag_solanago.PublicKey) *ReallocStakeList {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(adminAuthority).SIGNER()
	return inst
}

// GetAdminAuthorityAccount gets the "adminAuthority" account.
func (inst *ReallocStakeList) GetAdminAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStakeListAccount sets the "stakeList" account.
func (inst *ReallocStakeList) SetStakeListAccount(stakeList ag_solanago.PublicKey) *ReallocStakeList {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(stakeList).WRITE()
	return inst
}

// GetStakeListAccount gets the "stakeList" account.
func (inst *ReallocStakeList) GetStakeListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRentFundsAccount sets the "rentFunds" account.
func (inst *ReallocStakeList) SetRentFundsAccount(rentFunds ag_solanago.PublicKey) *ReallocStakeList {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rentFunds).WRITE().SIGNER()
	return inst
}

// GetRentFundsAccount gets the "rentFunds" account.
func (inst *ReallocStakeList) GetRentFundsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ReallocStakeList) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ReallocStakeList {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ReallocStakeList) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst ReallocStakeList) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ReallocStakeList,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ReallocStakeList) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ReallocStakeList) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Capacity == nil {
			return errors.New("Capacity parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AdminAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StakeList is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RentFunds is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ReallocStakeList) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ReallocStakeList")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Capacity", *inst.Capacity))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("adminAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     stakeList", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     rentFunds", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj ReallocStakeList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Capacity` param:
	err = encoder.Encode(obj.Capacity)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ReallocStakeList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Capacity`:
	err = decoder.Decode(&obj.Capacity)
	if err != nil {
		return err
	}
	return nil
}

// NewReallocStakeListInstruction declares a new ReallocStakeList instruction with the provided parameters and accounts.
func NewReallocStakeListInstruction(
	// Parameters:
	capacity uint32,
	// Accounts:
	state ag_solanago.PublicKey,
	adminAuthority ag_solanago.PublicKey,
	stakeList ag_solanago.PublicKey,
	rentFunds ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ReallocStakeList {
	return NewReallocStakeListInstructionBuilder().
		SetCapacity(capacity).
		SetStateAccount(state).
		SetAdminAuthorityAccount(adminAuthority).
		SetStakeListAccount(stakeList).
		SetRentFundsAccount(rentFunds).
		SetSystemProgramAccount(systemProgram)
}
// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PartialUnstake is the `partialUnstake` instruction.
type PartialUnstake struct {
	StakeIndex           *uint32
	ValidatorIndex       *uint32
	DesiredUnstakeAmount *uint64

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] validatorManagerAuthority
	//
	// [2] = [WRITE] validatorList
	//
	// [3] = [WRITE] stakeList
	//
	// [4] = [WRITE] stakeAccount
	//
	// [5] = [] stakeDepositAuthority
	//
	// [6] = [] reservePda
	//
	// [7] = [WRITE, SIGNER] splitStakeAccount
	//
	// [8] = [WRITE, SIGNER] splitStakeRentPayer
	//
	// [9] = [] clock
	//
	// [10] = [] rent
	//
	// [11] = [] stakeHistory
	//
	// [12] = [] systemProgram
	//
	// [13] = [] stakeProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPartialUnstakeInstructionBuilder creates a new `PartialUnstake` instruction builder.
func NewPartialUnstakeInstructionBuilder() *PartialUnstake {
	nd := &PartialUnstake{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetStakeIndex sets the "stakeIndex" parameter.
func (inst *PartialUnstake) SetStakeIndex(stakeIndex uint32) *PartialUnstake {
	inst.StakeIndex = &stakeIndex
	return inst
}

// SetValidatorIndex sets the "validatorIndex" parameter.
func (inst *PartialUnstake) SetValidatorIndex(validatorIndex uint32) *PartialUnstake {
	inst.ValidatorIndex = &validatorIndex
	return inst
}

// SetDesiredUnstakeAmount sets the "desiredUnstakeAmount" parameter.
func (inst *PartialUnstake) SetDesiredUnstakeAmount(desiredUnstakeAmount uint64) *PartialUnstake {
	inst.DesiredUnstakeAmount = &desiredUnstakeAmount
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *PartialUnstake) SetStateAccount(state ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *PartialUnstake) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetValidatorManagerAuthorityAccount sets the "validatorManagerAuthority" account.
func (inst *PartialUnstake) SetValidatorManagerAuthorityAccount(validatorManagerAuthority ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(validatorManagerAuthority).SIGNER()
	return inst
}

// GetValidatorManagerAuthorityAccount gets the "validatorManagerAuthority" account.
func (inst *PartialUnstake) GetValidatorManagerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetValidatorListAccount sets the "validatorList" account.
func (inst *PartialUnstake) SetValidatorListAccount(validatorList ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

// GetValidatorListAccount gets the "validatorList" account.
func (inst *PartialUnstake) GetValidatorListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStakeListAccount sets the "stakeList" account.
func (inst *PartialUnstake) SetStakeListAccount(stakeList ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeList).WRITE()
	return inst
}

// GetStakeListAccount gets the "stakeList" account.
func (inst *PartialUnstake) GetStakeListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStakeAccountAccount sets the "stakeAccount" account.
func (inst *PartialUnstake) SetStakeAccountAccount(stakeAccount ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccountAccount gets the "stakeAccount" account.
func (inst *PartialUnstake) GetStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetStakeDepositAuthorityAccount sets the "stakeDepositAuthority" account.
func (inst *PartialUnstake) SetStakeDepositAuthorityAccount(stakeDepositAuthority ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(stakeDepositAuthority)
	return inst
}

// GetStakeDepositAuthorityAccount gets the "stakeDepositAuthority" account.
func (inst *PartialUnstake) GetStakeDepositAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReservePdaAccount sets the "reservePda" account.
func (inst *PartialUnstake) SetReservePdaAccount(reservePda ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(reservePda)
	return inst
}

// GetReservePdaAccount gets the "reservePda" account.
func (inst *PartialUnstake) GetReservePdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSplitStakeAccountAccount sets the "splitStakeAccount" account.
func (inst *PartialUnstake) SetSplitStakeAccountAccount(splitStakeAccount ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(splitStakeAccount).WRITE().SIGNER()
	return inst
}

// GetSplitStakeAccountAccount gets the "splitStakeAccount" account.
func (inst *PartialUnstake) GetSplitStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSplitStakeRentPayerAccount sets the "splitStakeRentPayer" account.
func (inst *PartialUnstake) SetSplitStakeRentPayerAccount(splitStakeRentPayer ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(splitStakeRentPayer).WRITE().SIGNER()
	return inst
}

// GetSplitStakeRentPayerAccount gets the "splitStakeRentPayer" account.
func (inst *PartialUnstake) GetSplitStakeRentPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetClockAccount sets the "clock" account.
func (inst *PartialUnstake) SetClockAccount(clock ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *PartialUnstake) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
func (inst *PartialUnstake) SetRentAccount(rent ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *PartialUnstake) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetStakeHistoryAccount sets the "stakeHistory" account.
func (inst *PartialUnstake) SetStakeHistoryAccount(stakeHistory ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(stakeHistory)
	return inst
}

// GetStakeHistoryAccount gets the "stakeHistory" account.
func (inst *PartialUnstake) GetStakeHistoryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *PartialUnstake) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *PartialUnstake) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetStakeProgramAccount sets the "stakeProgram" account.
func (inst *PartialUnstake) SetStakeProgramAccount(stakeProgram ag_solanago.PublicKey) *PartialUnstake {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(stakeProgram)
	return inst
}

// GetStakeProgramAccount gets the "stakeProgram" account.
func (inst *PartialUnstake) GetStakeProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst PartialUnstake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PartialUnstake,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PartialUnstake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PartialUnstake) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.StakeIndex == nil {
			return errors.New("StakeIndex parameter is not set")
		}
		if inst.ValidatorIndex == nil {
			return errors.New("ValidatorIndex parameter is not set")
		}
		if inst.DesiredUnstakeAmount == nil {
			return errors.New("DesiredUnstakeAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ValidatorManagerAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ValidatorList is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.StakeList is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.StakeAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.StakeDepositAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReservePda is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SplitStakeAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SplitStakeRentPayer is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Clock is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.StakeHistory is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.StakeProgram is not set")
		}
	}
	return nil
}

func (inst *PartialUnstake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PartialUnstake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("          StakeIndex", *inst.StakeIndex))
						paramsBranch.Child(ag_format.Param("      ValidatorIndex", *inst.ValidatorIndex))
						paramsBranch.Child(ag_format.Param("DesiredUnstakeAmount", *inst.DesiredUnstakeAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("validatorManagerAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            validatorList", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                stakeList", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                    stake", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    stakeDepositAuthority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("               reservePda", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("               splitStake", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      splitStakeRentPayer", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                    clock", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                     rent", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("             stakeHistory", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("            systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("             stakeProgram", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj PartialUnstake) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StakeIndex` param:
	err = encoder.Encode(obj.StakeIndex)
	if err != nil {
		return err
	}
	// Serialize `ValidatorIndex` param:
	err = encoder.Encode(obj.ValidatorIndex)
	if err != nil {
		return err
	}
	// Serialize `DesiredUnstakeAmount` param:
	err = encoder.Encode(obj.DesiredUnstakeAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PartialUnstake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StakeIndex`:
	err = decoder.Decode(&obj.StakeIndex)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorIndex`:
	err = decoder.Decode(&obj.ValidatorIndex)
	if err != nil {
		return err
	}
	// Deserialize `DesiredUnstakeAmount`:
	err = decoder.Decode(&obj.DesiredUnstakeAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewPartialUnstakeInstruction declares a new PartialUnstake instruction with the provided parameters and accounts.
func NewPartialUnstakeInstruction(
	// Parameters:
	stakeIndex uint32,
	validatorIndex uint32,
	desiredUnstakeAmount uint64,
	// Accounts:
	state ag_solanago.PublicKey,
	validatorManagerAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	stakeList ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
	stakeDepositAuthority ag_solanago.PublicKey,
	reservePda ag_solanago.PublicKey,
	splitStakeAccount ag_solanago.PublicKey,
	splitStakeRentPayer ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey) *PartialUnstake {
	return NewPartialUnstakeInstructionBuilder().
		SetStakeIndex(stakeIndex).
		SetValidatorIndex(validatorIndex).
		SetDesiredUnstakeAmount(desiredUnstakeAmount).
		SetStateAccount(state).
		SetValidatorManagerAuthorityAccount(validatorManagerAuthority).
		SetValidatorListAccount(validatorList).
		SetStakeListAccount(stakeList).
		SetStakeAccountAccount(stakeAccount).
		SetStakeDepositAuthorityAccount(stakeDepositAuthority).
		SetReservePdaAccount(reservePda).
		SetSplitStakeAccountAccount(splitStakeAccount).
		SetSplitStakeRentPayerAccount(splitStakeRentPayer).
		SetClockAccount(clock).
		SetRentAccount(rent).
		SetStakeHistoryAccount(stakeHistory).
		SetSystemProgramAccount(systemProgram).
		SetStakeProgramAccount(stakeProgram)
}

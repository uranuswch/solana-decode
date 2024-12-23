// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package marinade_finance

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// LiquidUnstake is the `liquidUnstake` instruction.
type LiquidUnstake struct {
	MsolAmount *uint64

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] msolMint
	//
	// [2] = [WRITE] liqPoolSolLegPda
	//
	// [3] = [WRITE] liqPoolMsolLeg
	//
	// [4] = [WRITE] treasuryMsolAccount
	//
	// [5] = [WRITE] getMsolFrom
	//
	// [6] = [SIGNER] getMsolFromAuthority
	//
	// [7] = [WRITE] transferSolTo
	//
	// [8] = [] systemProgram
	//
	// [9] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewLiquidUnstakeInstructionBuilder creates a new `LiquidUnstake` instruction builder.
func NewLiquidUnstakeInstructionBuilder() *LiquidUnstake {
	nd := &LiquidUnstake{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetMsolAmount sets the "msolAmount" parameter.
func (inst *LiquidUnstake) SetMsolAmount(msolAmount uint64) *LiquidUnstake {
	inst.MsolAmount = &msolAmount
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *LiquidUnstake) SetStateAccount(state ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *LiquidUnstake) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMsolMintAccount sets the "msolMint" account.
func (inst *LiquidUnstake) SetMsolMintAccount(msolMint ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(msolMint).WRITE()
	return inst
}

// GetMsolMintAccount gets the "msolMint" account.
func (inst *LiquidUnstake) GetMsolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLiqPoolSolLegPdaAccount sets the "liqPoolSolLegPda" account.
func (inst *LiquidUnstake) SetLiqPoolSolLegPdaAccount(liqPoolSolLegPda ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(liqPoolSolLegPda).WRITE()
	return inst
}

// GetLiqPoolSolLegPdaAccount gets the "liqPoolSolLegPda" account.
func (inst *LiquidUnstake) GetLiqPoolSolLegPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLiqPoolMsolLegAccount sets the "liqPoolMsolLeg" account.
func (inst *LiquidUnstake) SetLiqPoolMsolLegAccount(liqPoolMsolLeg ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(liqPoolMsolLeg).WRITE()
	return inst
}

// GetLiqPoolMsolLegAccount gets the "liqPoolMsolLeg" account.
func (inst *LiquidUnstake) GetLiqPoolMsolLegAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTreasuryMsolAccountAccount sets the "treasuryMsolAccount" account.
func (inst *LiquidUnstake) SetTreasuryMsolAccountAccount(treasuryMsolAccount ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(treasuryMsolAccount).WRITE()
	return inst
}

// GetTreasuryMsolAccountAccount gets the "treasuryMsolAccount" account.
func (inst *LiquidUnstake) GetTreasuryMsolAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetGetMsolFromAccount sets the "getMsolFrom" account.
func (inst *LiquidUnstake) SetGetMsolFromAccount(getMsolFrom ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(getMsolFrom).WRITE()
	return inst
}

// GetGetMsolFromAccount gets the "getMsolFrom" account.
func (inst *LiquidUnstake) GetGetMsolFromAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetGetMsolFromAuthorityAccount sets the "getMsolFromAuthority" account.
func (inst *LiquidUnstake) SetGetMsolFromAuthorityAccount(getMsolFromAuthority ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(getMsolFromAuthority).SIGNER()
	return inst
}

// GetGetMsolFromAuthorityAccount gets the "getMsolFromAuthority" account.
func (inst *LiquidUnstake) GetGetMsolFromAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTransferSolToAccount sets the "transferSolTo" account.
func (inst *LiquidUnstake) SetTransferSolToAccount(transferSolTo ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(transferSolTo).WRITE()
	return inst
}

// GetTransferSolToAccount gets the "transferSolTo" account.
func (inst *LiquidUnstake) GetTransferSolToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *LiquidUnstake) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *LiquidUnstake) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *LiquidUnstake) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *LiquidUnstake {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *LiquidUnstake) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst LiquidUnstake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_LiquidUnstake,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst LiquidUnstake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *LiquidUnstake) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MsolAmount == nil {
			return errors.New("MsolAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MsolMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.LiqPoolSolLegPda is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LiqPoolMsolLeg is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TreasuryMsolAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.GetMsolFrom is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.GetMsolFromAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TransferSolTo is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *LiquidUnstake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("LiquidUnstake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MsolAmount", *inst.MsolAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            msolMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    liqPoolSolLegPda", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      liqPoolMsolLeg", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        treasuryMsol", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         getMsolFrom", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("getMsolFromAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       transferSolTo", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj LiquidUnstake) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MsolAmount` param:
	err = encoder.Encode(obj.MsolAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *LiquidUnstake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MsolAmount`:
	err = decoder.Decode(&obj.MsolAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewLiquidUnstakeInstruction declares a new LiquidUnstake instruction with the provided parameters and accounts.
func NewLiquidUnstakeInstruction(
	// Parameters:
	msolAmount uint64,
	// Accounts:
	state ag_solanago.PublicKey,
	msolMint ag_solanago.PublicKey,
	liqPoolSolLegPda ag_solanago.PublicKey,
	liqPoolMsolLeg ag_solanago.PublicKey,
	treasuryMsolAccount ag_solanago.PublicKey,
	getMsolFrom ag_solanago.PublicKey,
	getMsolFromAuthority ag_solanago.PublicKey,
	transferSolTo ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *LiquidUnstake {
	return NewLiquidUnstakeInstructionBuilder().
		SetMsolAmount(msolAmount).
		SetStateAccount(state).
		SetMsolMintAccount(msolMint).
		SetLiqPoolSolLegPdaAccount(liqPoolSolLegPda).
		SetLiqPoolMsolLegAccount(liqPoolMsolLeg).
		SetTreasuryMsolAccountAccount(treasuryMsolAccount).
		SetGetMsolFromAccount(getMsolFrom).
		SetGetMsolFromAuthorityAccount(getMsolFromAuthority).
		SetTransferSolToAccount(transferSolTo).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}

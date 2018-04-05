package main

import (
	"bytes"
	"encoding/json"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/vm/runtime"
	"github.com/pkg/errors"
	"io"
	"math/big"
	"strings"
	"github.com/kowala-tech/kcoin/core"
	"time"
	"math/rand"
	"github.com/kowala-tech/kcoin/params"
	"github.com/kowala-tech/kcoin/contracts/network/contracts"
	"github.com/kowala-tech/kcoin/accounts/abi"
	"github.com/kowala-tech/kcoin/core/vm"
)

var (
	ErrEmptyMaxNumValidators                        = errors.New("max number of validators is mandatory")
	ErrEmptyUnbondingPeriod                         = errors.New("unbonding period in days is mandatory")
	ErrEmptyWalletAddressValidator                  = errors.New("Wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator                = errors.New("Wallet address of genesis validator is invalid")
	ErrEmptyPrefundedAccounts                       = errors.New("empty prefunded accounts, at least the validator wallet address should be included")
	ErrWalletAddressValidatorNotInPrefundedAccounts = errors.New("prefunded accounts should include genesis validator account")
	ErrInvalidAddressInPrefundedAccounts            = errors.New("address in prefunded accounts is invalid")
	ErrInvalidContractsOwnerAddress                 = errors.New("address used for smart contracts is invalid")
)

type GenerateGenesisCommand struct {
	network                       string
	maxNumValidators              string
	unbondingPeriod               string
	walletAddressGenesisValidator string
	prefundedAccounts             []PrefundedAccount
	consensusEngine               string
	smartContractsOwner           string
	extraData string
}

type PrefundedAccount struct {
	walletAddress string
	balance       int64
}

type validPrefundedAccount struct {
	walletAddress *common.Address
	balance       *big.Int
}

type GenerateGenesisCommandHandler struct {
	w io.Writer
}

func (h *GenerateGenesisCommandHandler) Handle(command GenerateGenesisCommand) error {
	genesis, err := GenerateGenesis(command)
	if err != nil {
		return err
	}

	out, _ := json.MarshalIndent(genesis, "", "  ")

	_, err = h.w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func createMaxNumValidators(s string) (*big.Int, error) {
	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}

	numValidators, ok := new(big.Int).SetString(s, 0)
	if !ok {
		//TODO: Create error
		return nil, errors.New("invalid max num of validators.")
	}

	return numValidators, nil
}

func createUnbondingPeriod(uP string) (*big.Int, error) {
	var text string
	if text = strings.TrimSpace(uP); text == "" {
		return nil, ErrEmptyUnbondingPeriod
	}

	unbondingPeriod, ok := new(big.Int).SetString(text, 0)
	if !ok {
		//TODO: Create error
		return nil, errors.New("invalid max num of validators.")
	}

	return unbondingPeriod, nil
}

func createWalletAddress(wA string) (*common.Address, error) {
	stringAddr := wA

	if text := strings.TrimSpace(wA); text == "" {
		return nil, ErrEmptyWalletAddressValidator
	}

	if strings.HasPrefix(stringAddr, "0x") {
		stringAddr = strings.TrimPrefix(stringAddr, "0x")
	}

	if len(stringAddr) != 40 {
		return nil, ErrInvalidWalletAddressValidator
	}

	bigaddr, _ := new(big.Int).SetString(stringAddr, 16)
	address := common.BigToAddress(bigaddr)

	return &address, nil
}

func createPrefundedAccounts(accounts []PrefundedAccount) ([]*validPrefundedAccount, error) {
	var validAccounts []*validPrefundedAccount

	if len(accounts) == 0 {
		return nil, ErrEmptyPrefundedAccounts
	}

	for _, a := range accounts {
		address, err := createWalletAddress(a.walletAddress)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance := big.NewInt(a.balance)

		validAccount := &validPrefundedAccount{
			walletAddress: address,
			balance:       balance,
		}

		validAccounts = append(validAccounts, validAccount)
	}

	return validAccounts, nil
}

func prefundedIncludesValidatorWallet(
	accounts []*validPrefundedAccount,
	addresses *common.Address,
) bool {
	for _, account := range accounts {
		if bytes.Equal(account.walletAddress.Bytes(), addresses.Bytes()) {
			return true
		}
	}

	return false
}

func createContract(cfg *runtime.Config, code []byte) (*contractData, error) {
	out, addr, _, err := runtime.Create(code, cfg)
	if err != nil {
		return nil, err
	}
	return &contractData{
		addr:    addr,
		code:    out,
		storage: cfg.EVMConfig.Tracer.(*vmTracer).data[addr],
	}, nil
}

type contractData struct {
	addr    common.Address
	code    []byte
	storage map[common.Hash]common.Hash
}

func GenerateGenesis(command GenerateGenesisCommand) (*core.Genesis, error) {
	network, err := NewNetwork(command.network)
	if err != nil {
		return nil, err
	}

	maxNumValidators, err := createMaxNumValidators(command.maxNumValidators)
	if err != nil {
		return nil, err
	}

	unbondingPeriod, err := createUnbondingPeriod(command.unbondingPeriod)
	if err != nil {
		return nil, err
	}

	walletAddressValidator, err := createWalletAddress(command.walletAddressGenesisValidator)
	if err != nil {
		return nil, err
	}

	validPrefundedAccounts, err := createPrefundedAccounts(command.prefundedAccounts)
	if err != nil {
		return nil, err
	}

	if !prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return nil, ErrWalletAddressValidatorNotInPrefundedAccounts
	}

	consensusEngine := TendermintConsensus
	if command.consensusEngine != "" {
		consensusEngine, err = NewConsensusEngine(command.consensusEngine)
		if err != nil {
			return nil, err
		}
	}

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	switch network {
	case MainNetwork:
		genesis.Config.ChainID = params.MainnetChainConfig.ChainID
	case TestNetwork:
		genesis.Config.ChainID = params.TestnetChainConfig.ChainID
	case OtherNetwork:
		genesis.Config.ChainID = new(big.Int).SetUint64(uint64(rand.Intn(65536)))
	}

	switch consensusEngine {
	case TendermintConsensus:
		genesis.Config.Tendermint = &params.TendermintConfig{Rewarded: true}
	}

	extra := ""
	if command.extraData != "" {
		extra = command.extraData
	}

	genesis.ExtraData = make([]byte, 32)
	if len(extra) > 32 {
		extra = extra[:32]
	}

	genesis.ExtraData = append([]byte(extra), genesis.ExtraData[len(extra):]...)

	strAddr := DefaultSmartContractsOwner
	if command.smartContractsOwner != "" {
		strAddr = command.smartContractsOwner
	}

	owner, err := createWalletAddress(strAddr)
	if err != nil {
		return nil, ErrInvalidContractsOwnerAddress
	}

	genesis.Alloc[*owner] = core.GenesisAccount{Balance: new(big.Int).Mul(common.Big1, big.NewInt(params.Ether))}

	//TODO: This maybe will be need to be available to change by the parameters in the command in the future, right now is 0.
	baseDeposit := common.Big0

	electionABI, err := abi.JSON(strings.NewReader(contracts.ElectionContractABI))
	if err != nil {
		return nil, err
	}

	electionParams, err := electionABI.Pack(
		"",
		baseDeposit,
		maxNumValidators,
		unbondingPeriod,
		*walletAddressValidator,
	)
	if err != nil {
		return nil, err
	}

	runtimeCfg := &runtime.Config{
		Origin: *owner,
		EVMConfig: vm.Config{
			Debug:  true,
			Tracer: newVmTracer(),
		},
	}

	contract, err := createContract(runtimeCfg, append(common.FromHex(contracts.ElectionContractBin), electionParams...))
	if err != nil {
		return nil, err
	}

	genesis.Alloc[contract.addr] = core.GenesisAccount{
		Code:    contract.code,
		Storage: contract.storage,
		Balance: new(big.Int).Mul(baseDeposit, new(big.Int).SetUint64(params.Ether)),
	}

	for _, vAccount := range validPrefundedAccounts {
		genesis.Alloc[*vAccount.walletAddress] = core.GenesisAccount{
			Balance: vAccount.balance,
		}
	}

	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}

	return genesis, nil
}

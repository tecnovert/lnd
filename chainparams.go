package lnd

import (
	"github.com/btcsuite/btcd/chaincfg"
	bitcoinCfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	bitcoinWire "github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/keychain"
	litecoinCfg "github.com/ltcsuite/ltcd/chaincfg"
	litecoinWire "github.com/ltcsuite/ltcd/wire"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = bitcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// litecoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type litecoinNetParams struct {
	*litecoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// particlNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type particlNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.TestNet3Params,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinMainNetParams contains parameters specific to the current Bitcoin
// mainnet.
var bitcoinMainNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.MainNetParams,
	rpcPort:  "8334",
	CoinType: keychain.CoinTypeBitcoin,
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinSimNetParams contains parameters specific to the simulation test
// network.
var litecoinSimNetParams = litecoinNetParams{
	Params:   &litecoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var litecoinTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinMainNetParams contains the parameters specific to the current
// Litecoin mainnet.
var litecoinMainNetParams = litecoinNetParams{
	Params:   &litecoinCfg.MainNetParams,
	rpcPort:  "9334",
	CoinType: keychain.CoinTypeLitecoin,
}

// litecoinRegTestNetParams contains parameters specific to a local litecoin
// regtest network.
var litecoinRegTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinRegTestNetParams contains parameters specific to a local bitcoin
// regtest network.
var bitcoinRegTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// particlTestNetParams contains parameters specific to the current
// Particl test network.
var particlTestNetParams = particlNetParams{
	Params:   &bitcoinCfg.ParticlTestNetParams,
	rpcPort:  "51935",
	CoinType: keychain.CoinTypeTestnet,
}

// particlMainNetParams contains the parameters specific to the current
// Particl mainnet.
var particlMainNetParams = particlNetParams{
	Params:   &bitcoinCfg.ParticlMainNetParams,
	rpcPort:  "51735",
	CoinType: keychain.CoinTypeParticl,
}

// particlRegTestNetParams contains parameters specific to a local regtest network.
var particlRegTestNetParams = particlNetParams{
	Params:   &bitcoinCfg.ParticlRegressionNetParams,
	rpcPort:  "51936",
	CoinType: keychain.CoinTypeTestnet,
}

// applyLitecoinParams applies the relevant chain configuration parameters that
// differ for litecoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyLitecoinParams(params *bitcoinNetParams, litecoinParams *litecoinNetParams) {
	params.Name = litecoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(litecoinParams.Net)
	params.DefaultPort = litecoinParams.DefaultPort
	params.CoinbaseMaturity = litecoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], litecoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = litecoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = litecoinParams.ScriptHashAddrID
	params.PrivateKeyID = litecoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = litecoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = litecoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = litecoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], litecoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], litecoinParams.HDPublicKeyID[:])

	params.HDCoinType = litecoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(litecoinParams.Checkpoints))
	for i := 0; i < len(litecoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], litecoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: litecoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = litecoinParams.rpcPort
	params.CoinType = litecoinParams.CoinType
}

// applyParticlParams applies the relevant chain configuration parameters that
// differ for particl to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyParticlParams(params *bitcoinNetParams, particlParams *particlNetParams) {
	params.CoinName = particlParams.CoinName
	params.Name = particlParams.Name
	params.Net = bitcoinWire.BitcoinNet(particlParams.Net)
	params.DefaultPort = particlParams.DefaultPort
	params.CoinbaseMaturity = particlParams.CoinbaseMaturity

	copy(params.GenesisHash[:], particlParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = particlParams.PubKeyHashAddrID
	params.ScriptHashAddrID = particlParams.ScriptHashAddrID
	params.PrivateKeyID = particlParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = particlParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = particlParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = particlParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], particlParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], particlParams.HDPublicKeyID[:])

	params.HDCoinType = particlParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(particlParams.Checkpoints))
	for i := 0; i < len(particlParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], particlParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: particlParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = particlParams.rpcPort
	params.CoinType = particlParams.CoinType
}

// isTestnet tests if the given params correspond to a testnet
// parameter configuration.
func isTestnet(params *bitcoinNetParams) bool {
	switch params.Params.Net {
	case bitcoinWire.TestNet3, bitcoinWire.BitcoinNet(litecoinWire.TestNet4), bitcoinWire.ParticlTestNet:
		return true
	default:
		return false
	}
}

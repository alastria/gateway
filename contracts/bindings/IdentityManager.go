// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindigs

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IdentityManagerABI is the input ABI used to generate the binding from.
const IdentityManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"recoveryKey\",\"type\":\"address\"}],\"name\":\"isRecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOlderOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"recoveryKey\",\"type\":\"address\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createIdentityWithCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newIdManager\",\"type\":\"address\"}],\"name\":\"initiateMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrationNewAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"addOwnerFromRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"recoveryKey\",\"type\":\"address\"}],\"name\":\"changeRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"finalizeMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"cancelMigration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"forwardTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"recoveryKey\",\"type\":\"address\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrationInitiated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"recoveryKey\",\"type\":\"address\"}],\"name\":\"createIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_userTimeLock\",\"type\":\"uint256\"},{\"name\":\"_adminTimeLock\",\"type\":\"uint256\"},{\"name\":\"_adminRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"recoveryKey\",\"type\":\"address\"}],\"name\":\"LogIdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instigator\",\"type\":\"address\"}],\"name\":\"LogOwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instigator\",\"type\":\"address\"}],\"name\":\"LogOwnerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"recoveryKey\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instigator\",\"type\":\"address\"}],\"name\":\"LogRecoveryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newIdManager\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instigator\",\"type\":\"address\"}],\"name\":\"LogMigrationInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newIdManager\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instigator\",\"type\":\"address\"}],\"name\":\"LogMigrationCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newIdManager\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instigator\",\"type\":\"address\"}],\"name\":\"LogMigrationFinalized\",\"type\":\"event\"}]"

// IdentityManagerBin is the compiled bytecode used for deploying new contracts.
const IdentityManagerBin = `608060405234801561001057600080fd5b50604051606080612b5683398101806040528101908080519060200190929190805190602001909291908051906020019092919050505082821015151561005657600080fd5b816000819055508260018190555080600281905550505050612ad98061007d6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806311fe12b3146100eb57806316d390bf1461016657806332967ea0146101c95780633c8ac88e146102445780633dcf59ca1461030d578063422e33f3146103705780635143eea2146103f357806353faa9a914610456578063633b1954146104b95780636f022ac4146104fc57806373b40a5c1461053f578063781f5a83146105f25780637ddc02d414610655578063c778427b146106d0578063d10e73ab14610727578063fbe5ce0a1461078a575b600080fd5b3480156100f757600080fd5b5061014c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107ed565b604051808215151515815260200191505060405180910390f35b34801561017257600080fd5b506101c7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610885565b005b3480156101d557600080fd5b5061022a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610add565b604051808215151515815260200191505060405180910390f35b34801561025057600080fd5b5061030b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610bf2565b005b34801561031957600080fd5b5061036e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f34565b005b34801561037c57600080fd5b506103b1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110e0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103ff57600080fd5b50610454600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611113565b005b34801561046257600080fd5b506104b7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113eb565b005b3480156104c557600080fd5b506104fa600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611664565b005b34801561050857600080fd5b5061053d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a4b565b005b34801561054b57600080fd5b506105f0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611c00565b005b3480156105fe57600080fd5b50610653600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d45565b005b34801561066157600080fd5b506106b6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611fb6565b604051808215151515815260200191505060405180910390f35b3480156106dc57600080fd5b50610711600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506120cb565b6040518082815260200191505060405180910390f35b34801561073357600080fd5b50610788600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506120e3565b005b34801561079657600080fd5b506107eb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506122fa565b005b60008173ffffffffffffffffffffffffffffffffffffffff16600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b816108908133610add565b151561089b57600080fd5b826002544203600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151561092a57600080fd5b42600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506109b58484611fb6565b1515156109c157600080fd5b6001544203600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8672e8532f3edff41d3acf0cd4be6ff900e427461b81d094f0197354471cb3c633604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a350505050565b600080600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054118015610bea575042600054600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540111155b905092915050565b600083600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610c3157600080fd5b610c39612572565b604051809103906000f080158015610c55573d6000803e3d6000fd5b5091506000544203600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555084600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f14e580ab5cd452b772e031536a7c893ec705152c17b3665c6671b382c3ee626689604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a48173ffffffffffffffffffffffffffffffffffffffff1663d7f31eb9856000866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610ec6578082015181840152602081019050610eab565b50505050905090810190601f168015610ef35780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610f1457600080fd5b505af1158015610f28573d6000803e3d6000fd5b50505050505050505050565b81610f3f8133610add565b1515610f4a57600080fd5b81600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610f8757600080fd5b42600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fb7b4557c664a8a4a9a57f5e00f216de044a96f96fc84b70eaa6bd48cb078454133604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a350505050565b60076020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b813373ffffffffffffffffffffffffffffffffffffffff16600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156111ad57600080fd5b826002544203600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151561123c57600080fd5b42600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506112c78484611fb6565b1515156112d357600080fd5b42600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8672e8532f3edff41d3acf0cd4be6ff900e427461b81d094f0197354471cb3c633604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a350505050565b816113f68133610add565b151561140157600080fd5b826002544203600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151561149057600080fd5b42600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561154e57600080fd5b83600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fbd1ad05c16aafa75ac8c1d6b8264f47ad9f3045596f667c9476f3f749c01870933604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a35050505050565b6000816116718133610add565b151561167c57600080fd5b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415801561170e575042600054600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401105b151561171957600080fd5b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690558273ffffffffffffffffffffffffffffffffffffffff16631a695230836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156118bb57600080fd5b505af11580156118cf573d6000803e3d6000fd5b50505050600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090558173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f565ed5a2d5e196b82acb6ff1149b699c2b13169c7d34412d736d3c380de64f3133604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a3505050565b600081611a588133611fb6565b1515611a6357600080fd5b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690558173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f60e805c5650597523aba29fb00a59c856d925c672bce1ea7d579f03aef3a10ce33604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a3505050565b83611c0b8133611fb6565b1515611c1657600080fd5b8473ffffffffffffffffffffffffffffffffffffffff1663d7f31eb98585856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611cd8578082015181840152602081019050611cbd565b50505050905090810190601f168015611d055780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015611d2657600080fd5b505af1158015611d3a573d6000803e3d6000fd5b505050505050505050565b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611d8257600080fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515611e0657600080fd5b6000544203600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f14e580ab5cd452b772e031536a7c893ec705152c17b3665c6671b382c3ee626686604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a4505050565b600080600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541180156120c3575042600154600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540111155b905092915050565b60066020528060005260406000206000915090505481565b600081600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561212257600080fd5b61212a612572565b604051809103906000f080158015612146573d6000803e3d6000fd5b5091506000544203600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f14e580ab5cd452b772e031536a7c893ec705152c17b3665c6671b382c3ee626687604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a450505050565b816123058133610add565b151561231057600080fd5b826002544203600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410151561239f57600080fd5b42600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561245b57600080fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090558273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f5e159cd4447854ae8b4aa048f91c8daad986faab696c3685030e1a5e5a4e8ced33604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a350505050565b60405161052b806125838339019056006080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506104d8806100536000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631a695230146100b25780632f54bf6e146100f55780638da5cb5b14610150578063d7f31eb9146101a7575b3373ffffffffffffffffffffffffffffffffffffffff167f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874346040518082815260200191505060405180910390a2005b3480156100be57600080fd5b506100f3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061023a565b005b34801561010157600080fd5b50610136600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102c7565b604051808215151515815260200191505060405180910390f35b34801561015c57600080fd5b50610165610320565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101b357600080fd5b50610238600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610345565b005b610243336102c7565b151561024e57600080fd5b3073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156102c457806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61034e336102c7565b151561035957600080fd5b8273ffffffffffffffffffffffffffffffffffffffff16828260405180828051906020019080838360005b8381101561039f578082015181840152602081019050610384565b50505050905090810190601f1680156103cc5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af19250505015156103ec57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff167fc1de93dfa06362c6a616cde73ec17d116c0d588dd1df70f27f91b500de207c4183836040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561046c578082015181840152602081019050610451565b50505050905090810190601f1680156104995780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050505600a165627a7a72305820f0043eb0138285c41ec38ac01d038f4983d41dd12198ccad1a5b8be62a2a59700029a165627a7a72305820503e2a0efd33d1753b9a549de205449837415671f2b415ab5a00894c23f586da0029`

// DeployIdentityManager deploys a new Ethereum contract, binding an instance of IdentityManager to it.
func DeployIdentityManager(auth *bind.TransactOpts, backend bind.ContractBackend, _userTimeLock *big.Int, _adminTimeLock *big.Int, _adminRate *big.Int) (common.Address, *types.Transaction, *IdentityManager, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityManagerBin), backend, _userTimeLock, _adminTimeLock, _adminRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}}, nil
}

// IdentityManager is an auto generated Go binding around an Ethereum contract.
type IdentityManager struct {
	IdentityManagerCaller     // Read-only binding to the contract
	IdentityManagerTransactor // Write-only binding to the contract
}

// IdentityManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityManagerSession struct {
	Contract     *IdentityManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityManagerCallerSession struct {
	Contract *IdentityManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityManagerTransactorSession struct {
	Contract     *IdentityManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityManagerRaw struct {
	Contract *IdentityManager // Generic contract binding to access the raw methods on
}

// IdentityManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityManagerCallerRaw struct {
	Contract *IdentityManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityManagerTransactorRaw struct {
	Contract *IdentityManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityManager creates a new instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManager(address common.Address, backend bind.ContractBackend) (*IdentityManager, error) {
	contract, err := bindIdentityManager(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityManager{IdentityManagerCaller: IdentityManagerCaller{contract: contract}, IdentityManagerTransactor: IdentityManagerTransactor{contract: contract}}, nil
}

// NewIdentityManagerCaller creates a new read-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerCaller(address common.Address, caller bind.ContractCaller) (*IdentityManagerCaller, error) {
	contract, err := bindIdentityManager(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerCaller{contract: contract}, nil
}

// NewIdentityManagerTransactor creates a new write-only instance of IdentityManager, bound to a specific deployed contract.
func NewIdentityManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityManagerTransactor, error) {
	contract, err := bindIdentityManager(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &IdentityManagerTransactor{contract: contract}, nil
}

// bindIdentityManager binds a generic wrapper to an already deployed contract.
func bindIdentityManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.IdentityManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.IdentityManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManager *IdentityManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManager *IdentityManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManager *IdentityManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManager.Contract.contract.Transact(opts, method, params...)
}

// IsOlderOwner is a free data retrieval call binding the contract method 0x32967ea0.
//
// Solidity: function isOlderOwner(identity address, owner address) constant returns(bool)
func (_IdentityManager *IdentityManagerCaller) IsOlderOwner(opts *bind.CallOpts, identity common.Address, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityManager.contract.Call(opts, out, "isOlderOwner", identity, owner)
	return *ret0, err
}

// IsOlderOwner is a free data retrieval call binding the contract method 0x32967ea0.
//
// Solidity: function isOlderOwner(identity address, owner address) constant returns(bool)
func (_IdentityManager *IdentityManagerSession) IsOlderOwner(identity common.Address, owner common.Address) (bool, error) {
	return _IdentityManager.Contract.IsOlderOwner(&_IdentityManager.CallOpts, identity, owner)
}

// IsOlderOwner is a free data retrieval call binding the contract method 0x32967ea0.
//
// Solidity: function isOlderOwner(identity address, owner address) constant returns(bool)
func (_IdentityManager *IdentityManagerCallerSession) IsOlderOwner(identity common.Address, owner common.Address) (bool, error) {
	return _IdentityManager.Contract.IsOlderOwner(&_IdentityManager.CallOpts, identity, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x7ddc02d4.
//
// Solidity: function isOwner(identity address, owner address) constant returns(bool)
func (_IdentityManager *IdentityManagerCaller) IsOwner(opts *bind.CallOpts, identity common.Address, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityManager.contract.Call(opts, out, "isOwner", identity, owner)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x7ddc02d4.
//
// Solidity: function isOwner(identity address, owner address) constant returns(bool)
func (_IdentityManager *IdentityManagerSession) IsOwner(identity common.Address, owner common.Address) (bool, error) {
	return _IdentityManager.Contract.IsOwner(&_IdentityManager.CallOpts, identity, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x7ddc02d4.
//
// Solidity: function isOwner(identity address, owner address) constant returns(bool)
func (_IdentityManager *IdentityManagerCallerSession) IsOwner(identity common.Address, owner common.Address) (bool, error) {
	return _IdentityManager.Contract.IsOwner(&_IdentityManager.CallOpts, identity, owner)
}

// IsRecovery is a free data retrieval call binding the contract method 0x11fe12b3.
//
// Solidity: function isRecovery(identity address, recoveryKey address) constant returns(bool)
func (_IdentityManager *IdentityManagerCaller) IsRecovery(opts *bind.CallOpts, identity common.Address, recoveryKey common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityManager.contract.Call(opts, out, "isRecovery", identity, recoveryKey)
	return *ret0, err
}

// IsRecovery is a free data retrieval call binding the contract method 0x11fe12b3.
//
// Solidity: function isRecovery(identity address, recoveryKey address) constant returns(bool)
func (_IdentityManager *IdentityManagerSession) IsRecovery(identity common.Address, recoveryKey common.Address) (bool, error) {
	return _IdentityManager.Contract.IsRecovery(&_IdentityManager.CallOpts, identity, recoveryKey)
}

// IsRecovery is a free data retrieval call binding the contract method 0x11fe12b3.
//
// Solidity: function isRecovery(identity address, recoveryKey address) constant returns(bool)
func (_IdentityManager *IdentityManagerCallerSession) IsRecovery(identity common.Address, recoveryKey common.Address) (bool, error) {
	return _IdentityManager.Contract.IsRecovery(&_IdentityManager.CallOpts, identity, recoveryKey)
}

// MigrationInitiated is a free data retrieval call binding the contract method 0xc778427b.
//
// Solidity: function migrationInitiated( address) constant returns(uint256)
func (_IdentityManager *IdentityManagerCaller) MigrationInitiated(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityManager.contract.Call(opts, out, "migrationInitiated", arg0)
	return *ret0, err
}

// MigrationInitiated is a free data retrieval call binding the contract method 0xc778427b.
//
// Solidity: function migrationInitiated( address) constant returns(uint256)
func (_IdentityManager *IdentityManagerSession) MigrationInitiated(arg0 common.Address) (*big.Int, error) {
	return _IdentityManager.Contract.MigrationInitiated(&_IdentityManager.CallOpts, arg0)
}

// MigrationInitiated is a free data retrieval call binding the contract method 0xc778427b.
//
// Solidity: function migrationInitiated( address) constant returns(uint256)
func (_IdentityManager *IdentityManagerCallerSession) MigrationInitiated(arg0 common.Address) (*big.Int, error) {
	return _IdentityManager.Contract.MigrationInitiated(&_IdentityManager.CallOpts, arg0)
}

// MigrationNewAddress is a free data retrieval call binding the contract method 0x422e33f3.
//
// Solidity: function migrationNewAddress( address) constant returns(address)
func (_IdentityManager *IdentityManagerCaller) MigrationNewAddress(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityManager.contract.Call(opts, out, "migrationNewAddress", arg0)
	return *ret0, err
}

// MigrationNewAddress is a free data retrieval call binding the contract method 0x422e33f3.
//
// Solidity: function migrationNewAddress( address) constant returns(address)
func (_IdentityManager *IdentityManagerSession) MigrationNewAddress(arg0 common.Address) (common.Address, error) {
	return _IdentityManager.Contract.MigrationNewAddress(&_IdentityManager.CallOpts, arg0)
}

// MigrationNewAddress is a free data retrieval call binding the contract method 0x422e33f3.
//
// Solidity: function migrationNewAddress( address) constant returns(address)
func (_IdentityManager *IdentityManagerCallerSession) MigrationNewAddress(arg0 common.Address) (common.Address, error) {
	return _IdentityManager.Contract.MigrationNewAddress(&_IdentityManager.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x16d390bf.
//
// Solidity: function addOwner(identity address, newOwner address) returns()
func (_IdentityManager *IdentityManagerTransactor) AddOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "addOwner", identity, newOwner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x16d390bf.
//
// Solidity: function addOwner(identity address, newOwner address) returns()
func (_IdentityManager *IdentityManagerSession) AddOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.AddOwner(&_IdentityManager.TransactOpts, identity, newOwner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x16d390bf.
//
// Solidity: function addOwner(identity address, newOwner address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) AddOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.AddOwner(&_IdentityManager.TransactOpts, identity, newOwner)
}

// AddOwnerFromRecovery is a paid mutator transaction binding the contract method 0x5143eea2.
//
// Solidity: function addOwnerFromRecovery(identity address, newOwner address) returns()
func (_IdentityManager *IdentityManagerTransactor) AddOwnerFromRecovery(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "addOwnerFromRecovery", identity, newOwner)
}

// AddOwnerFromRecovery is a paid mutator transaction binding the contract method 0x5143eea2.
//
// Solidity: function addOwnerFromRecovery(identity address, newOwner address) returns()
func (_IdentityManager *IdentityManagerSession) AddOwnerFromRecovery(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.AddOwnerFromRecovery(&_IdentityManager.TransactOpts, identity, newOwner)
}

// AddOwnerFromRecovery is a paid mutator transaction binding the contract method 0x5143eea2.
//
// Solidity: function addOwnerFromRecovery(identity address, newOwner address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) AddOwnerFromRecovery(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.AddOwnerFromRecovery(&_IdentityManager.TransactOpts, identity, newOwner)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x6f022ac4.
//
// Solidity: function cancelMigration(identity address) returns()
func (_IdentityManager *IdentityManagerTransactor) CancelMigration(opts *bind.TransactOpts, identity common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "cancelMigration", identity)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x6f022ac4.
//
// Solidity: function cancelMigration(identity address) returns()
func (_IdentityManager *IdentityManagerSession) CancelMigration(identity common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.CancelMigration(&_IdentityManager.TransactOpts, identity)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x6f022ac4.
//
// Solidity: function cancelMigration(identity address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) CancelMigration(identity common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.CancelMigration(&_IdentityManager.TransactOpts, identity)
}

// ChangeRecovery is a paid mutator transaction binding the contract method 0x53faa9a9.
//
// Solidity: function changeRecovery(identity address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerTransactor) ChangeRecovery(opts *bind.TransactOpts, identity common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "changeRecovery", identity, recoveryKey)
}

// ChangeRecovery is a paid mutator transaction binding the contract method 0x53faa9a9.
//
// Solidity: function changeRecovery(identity address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerSession) ChangeRecovery(identity common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.ChangeRecovery(&_IdentityManager.TransactOpts, identity, recoveryKey)
}

// ChangeRecovery is a paid mutator transaction binding the contract method 0x53faa9a9.
//
// Solidity: function changeRecovery(identity address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) ChangeRecovery(identity common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.ChangeRecovery(&_IdentityManager.TransactOpts, identity, recoveryKey)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0xd10e73ab.
//
// Solidity: function createIdentity(owner address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerTransactor) CreateIdentity(opts *bind.TransactOpts, owner common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "createIdentity", owner, recoveryKey)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0xd10e73ab.
//
// Solidity: function createIdentity(owner address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerSession) CreateIdentity(owner common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.CreateIdentity(&_IdentityManager.TransactOpts, owner, recoveryKey)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0xd10e73ab.
//
// Solidity: function createIdentity(owner address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) CreateIdentity(owner common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.CreateIdentity(&_IdentityManager.TransactOpts, owner, recoveryKey)
}

// CreateIdentityWithCall is a paid mutator transaction binding the contract method 0x3c8ac88e.
//
// Solidity: function createIdentityWithCall(owner address, recoveryKey address, destination address, data bytes) returns()
func (_IdentityManager *IdentityManagerTransactor) CreateIdentityWithCall(opts *bind.TransactOpts, owner common.Address, recoveryKey common.Address, destination common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "createIdentityWithCall", owner, recoveryKey, destination, data)
}

// CreateIdentityWithCall is a paid mutator transaction binding the contract method 0x3c8ac88e.
//
// Solidity: function createIdentityWithCall(owner address, recoveryKey address, destination address, data bytes) returns()
func (_IdentityManager *IdentityManagerSession) CreateIdentityWithCall(owner common.Address, recoveryKey common.Address, destination common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.CreateIdentityWithCall(&_IdentityManager.TransactOpts, owner, recoveryKey, destination, data)
}

// CreateIdentityWithCall is a paid mutator transaction binding the contract method 0x3c8ac88e.
//
// Solidity: function createIdentityWithCall(owner address, recoveryKey address, destination address, data bytes) returns()
func (_IdentityManager *IdentityManagerTransactorSession) CreateIdentityWithCall(owner common.Address, recoveryKey common.Address, destination common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.CreateIdentityWithCall(&_IdentityManager.TransactOpts, owner, recoveryKey, destination, data)
}

// FinalizeMigration is a paid mutator transaction binding the contract method 0x633b1954.
//
// Solidity: function finalizeMigration(identity address) returns()
func (_IdentityManager *IdentityManagerTransactor) FinalizeMigration(opts *bind.TransactOpts, identity common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "finalizeMigration", identity)
}

// FinalizeMigration is a paid mutator transaction binding the contract method 0x633b1954.
//
// Solidity: function finalizeMigration(identity address) returns()
func (_IdentityManager *IdentityManagerSession) FinalizeMigration(identity common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.FinalizeMigration(&_IdentityManager.TransactOpts, identity)
}

// FinalizeMigration is a paid mutator transaction binding the contract method 0x633b1954.
//
// Solidity: function finalizeMigration(identity address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) FinalizeMigration(identity common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.FinalizeMigration(&_IdentityManager.TransactOpts, identity)
}

// ForwardTo is a paid mutator transaction binding the contract method 0x73b40a5c.
//
// Solidity: function forwardTo(identity address, destination address, value uint256, data bytes) returns()
func (_IdentityManager *IdentityManagerTransactor) ForwardTo(opts *bind.TransactOpts, identity common.Address, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "forwardTo", identity, destination, value, data)
}

// ForwardTo is a paid mutator transaction binding the contract method 0x73b40a5c.
//
// Solidity: function forwardTo(identity address, destination address, value uint256, data bytes) returns()
func (_IdentityManager *IdentityManagerSession) ForwardTo(identity common.Address, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.ForwardTo(&_IdentityManager.TransactOpts, identity, destination, value, data)
}

// ForwardTo is a paid mutator transaction binding the contract method 0x73b40a5c.
//
// Solidity: function forwardTo(identity address, destination address, value uint256, data bytes) returns()
func (_IdentityManager *IdentityManagerTransactorSession) ForwardTo(identity common.Address, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IdentityManager.Contract.ForwardTo(&_IdentityManager.TransactOpts, identity, destination, value, data)
}

// InitiateMigration is a paid mutator transaction binding the contract method 0x3dcf59ca.
//
// Solidity: function initiateMigration(identity address, newIdManager address) returns()
func (_IdentityManager *IdentityManagerTransactor) InitiateMigration(opts *bind.TransactOpts, identity common.Address, newIdManager common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "initiateMigration", identity, newIdManager)
}

// InitiateMigration is a paid mutator transaction binding the contract method 0x3dcf59ca.
//
// Solidity: function initiateMigration(identity address, newIdManager address) returns()
func (_IdentityManager *IdentityManagerSession) InitiateMigration(identity common.Address, newIdManager common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.InitiateMigration(&_IdentityManager.TransactOpts, identity, newIdManager)
}

// InitiateMigration is a paid mutator transaction binding the contract method 0x3dcf59ca.
//
// Solidity: function initiateMigration(identity address, newIdManager address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) InitiateMigration(identity common.Address, newIdManager common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.InitiateMigration(&_IdentityManager.TransactOpts, identity, newIdManager)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x781f5a83.
//
// Solidity: function registerIdentity(owner address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerTransactor) RegisterIdentity(opts *bind.TransactOpts, owner common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "registerIdentity", owner, recoveryKey)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x781f5a83.
//
// Solidity: function registerIdentity(owner address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerSession) RegisterIdentity(owner common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.RegisterIdentity(&_IdentityManager.TransactOpts, owner, recoveryKey)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x781f5a83.
//
// Solidity: function registerIdentity(owner address, recoveryKey address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) RegisterIdentity(owner common.Address, recoveryKey common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.RegisterIdentity(&_IdentityManager.TransactOpts, owner, recoveryKey)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xfbe5ce0a.
//
// Solidity: function removeOwner(identity address, owner address) returns()
func (_IdentityManager *IdentityManagerTransactor) RemoveOwner(opts *bind.TransactOpts, identity common.Address, owner common.Address) (*types.Transaction, error) {
	return _IdentityManager.contract.Transact(opts, "removeOwner", identity, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xfbe5ce0a.
//
// Solidity: function removeOwner(identity address, owner address) returns()
func (_IdentityManager *IdentityManagerSession) RemoveOwner(identity common.Address, owner common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.RemoveOwner(&_IdentityManager.TransactOpts, identity, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xfbe5ce0a.
//
// Solidity: function removeOwner(identity address, owner address) returns()
func (_IdentityManager *IdentityManagerTransactorSession) RemoveOwner(identity common.Address, owner common.Address) (*types.Transaction, error) {
	return _IdentityManager.Contract.RemoveOwner(&_IdentityManager.TransactOpts, identity, owner)
}

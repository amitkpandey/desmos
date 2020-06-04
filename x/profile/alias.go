package profile

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

const (
	OpWeightMsgSaveProfile   = simulation.OpWeightMsgSaveProfile
	OpWeightMsgDeleteProfile = simulation.OpWeightMsgDeleteProfile
	DefaultGasValue          = simulation.DefaultGasValue
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	MinNameSurnameLength     = types.MinNameSurnameLength
	MaxNameSurnameLength     = types.MaxNameSurnameLength
	MinMonikerLength         = types.MinMonikerLength
	MaxMonikerLength         = types.MaxMonikerLength
	MaxBioLength             = types.MaxBioLength
	ActionSaveProfile        = types.ActionSaveProfile
	ActionDeleteProfile      = types.ActionDeleteProfile
	QuerierRoute             = types.QuerierRoute
	QueryProfile             = types.QueryProfile
	QueryProfiles            = types.QueryProfiles
	EventTypeProfileSaved    = types.EventTypeProfileSaved
	EventTypeProfileDeleted  = types.EventTypeProfileDeleted
	AttributeProfileMoniker  = types.AttributeProfileMoniker
	AttributeProfileCreator  = types.AttributeProfileCreator
	DefaultParamspace        = types.DefaultParamspace
)

var (
	// functions aliases
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	RegisterInvariants          = keeper.RegisterInvariants
	AllInvariants               = keeper.AllInvariants
	ValidProfileInvariant       = keeper.ValidProfileInvariant
	NewHandler                  = keeper.NewHandler
	RandomizedGenState          = simulation.RandomizedGenState
	DecodeStore                 = simulation.DecodeStore
	SimulateMsgSaveProfile      = simulation.SimulateMsgSaveProfile
	SimulateMsgDeleteProfile    = simulation.SimulateMsgDeleteProfile
	RandomProfileData           = simulation.RandomProfileData
	RandomProfile               = simulation.RandomProfile
	RandomMoniker               = simulation.RandomMoniker
	RandomName                  = simulation.RandomName
	RandomSurname               = simulation.RandomSurname
	RandomBio                   = simulation.RandomBio
	RandomProfilePic            = simulation.RandomProfilePic
	RandomProfileCover          = simulation.RandomProfileCover
	GetSimAccount               = simulation.GetSimAccount
	WeightedOperations          = simulation.WeightedOperations
	ProfileStoreKey             = types.ProfileStoreKey
	MonikerStoreKey             = types.MonikerStoreKey
	NewProfile                  = types.NewProfile
	NewPictures                 = types.NewPictures
	NewGenesisState             = types.NewGenesisState
	DefaultGenesisState         = types.DefaultGenesisState
	ValidateGenesis             = types.ValidateGenesis
	RegisterCodec               = types.RegisterCodec
	ParamKeyTable               = types.ParamKeyTable
	NewNameSurnameLenParams     = types.NewNameSurnameLenParams
	DefaultNameSurnameLenParams = types.DefaultNameSurnameLenParams
	NewMonikerLenParams         = types.NewMonikerLenParams
	DefaultMonikerLenParams     = types.DefaultMonikerLenParams
	NewBioLenParams             = types.NewBioLenParams
	DefaultBioLenParams         = types.DefaultBioLenParams
	NewMsgSaveProfile           = types.NewMsgSaveProfile
	NewMsgDeleteProfile         = types.NewMsgDeleteProfile

	// variable aliases
	TxHashRegEx                 = types.TxHashRegEx
	URIRegEx                    = types.URIRegEx
	ProfileStorePrefix          = types.ProfileStorePrefix
	MonikerStorePrefix          = types.MonikerStorePrefix
	ModuleCdc                   = types.ModuleCdc
	DefaultMinNameSurnameLength = types.DefaultMinNameSurnameLength
	DefaultMaxNameSurnameLength = types.DefaultMaxNameSurnameLength
	DefaultMinMonikerLength     = types.DefaultMinMonikerLength
	DefaultMaxMonikerLength     = types.DefaultMaxMonikerLength
	DefaultMaxBioLength         = types.DefaultMaxBioLength
	ParamStoreKeyNameSurnameLen = types.ParamStoreKeyNameSurnameLen
	ParamStoreKeyMonikerLen     = types.ParamStoreKeyMonikerLen
	ParamStoreKeyMaxBioLen      = types.ParamStoreKeyMaxBioLen
)

type (
	Keeper               = keeper.Keeper
	ProfileData          = simulation.ProfileData
	Profile              = types.Profile
	Profiles             = types.Profiles
	Pictures             = types.Pictures
	GenesisState         = types.GenesisState
	NameSurnameLenParams = types.NameSurnameLenParams
	MonikerLenParams     = types.MonikerLenParams
	BioLenParams         = types.BioLenParams
	MsgSaveProfile       = types.MsgSaveProfile
	MsgDeleteProfile     = types.MsgDeleteProfile
)

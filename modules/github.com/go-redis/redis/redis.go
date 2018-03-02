package redis

import (
	original_redis "github.com/go-redis/redis"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("redis")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"NewBoolCmd":               original_redis.NewBoolCmd,
			"NewBoolResult":            original_redis.NewBoolResult,
			"NewBoolSliceCmd":          original_redis.NewBoolSliceCmd,
			"NewBoolSliceResult":       original_redis.NewBoolSliceResult,
			"NewClient":                original_redis.NewClient,
			"NewClusterClient":         original_redis.NewClusterClient,
			"NewClusterSlotsCmd":       original_redis.NewClusterSlotsCmd,
			"NewClusterSlotsCmdResult": original_redis.NewClusterSlotsCmdResult,
			"NewCmd":                   original_redis.NewCmd,
			"NewCmdResult":             original_redis.NewCmdResult,
			"NewCommandsInfoCmd":       original_redis.NewCommandsInfoCmd,
			"NewCommandsInfoCmdResult": original_redis.NewCommandsInfoCmdResult,
			"NewDurationCmd":           original_redis.NewDurationCmd,
			"NewDurationResult":        original_redis.NewDurationResult,
			"NewFailoverClient":        original_redis.NewFailoverClient,
			"NewFloatCmd":              original_redis.NewFloatCmd,
			"NewFloatResult":           original_redis.NewFloatResult,
			"NewGeoLocationCmd":        original_redis.NewGeoLocationCmd,
			"NewGeoLocationCmdResult":  original_redis.NewGeoLocationCmdResult,
			"NewGeoPosCmd":             original_redis.NewGeoPosCmd,
			"NewIntCmd":                original_redis.NewIntCmd,
			"NewIntResult":             original_redis.NewIntResult,
			"NewRing":                  original_redis.NewRing,
			"NewScanCmd":               original_redis.NewScanCmd,
			"NewScanCmdResult":         original_redis.NewScanCmdResult,
			"NewScript":                original_redis.NewScript,
			"NewSliceCmd":              original_redis.NewSliceCmd,
			"NewSliceResult":           original_redis.NewSliceResult,
			"NewStatusCmd":             original_redis.NewStatusCmd,
			"NewStatusResult":          original_redis.NewStatusResult,
			"NewStringCmd":             original_redis.NewStringCmd,
			"NewStringIntMapCmd":       original_redis.NewStringIntMapCmd,
			"NewStringIntMapCmdResult": original_redis.NewStringIntMapCmdResult,
			"NewStringResult":          original_redis.NewStringResult,
			"NewStringSliceCmd":        original_redis.NewStringSliceCmd,
			"NewStringSliceResult":     original_redis.NewStringSliceResult,
			"NewStringStringMapCmd":    original_redis.NewStringStringMapCmd,
			"NewStringStringMapResult": original_redis.NewStringStringMapResult,
			"NewStringStructMapCmd":    original_redis.NewStringStructMapCmd,
			"NewTimeCmd":               original_redis.NewTimeCmd,
			"NewUniversalClient":       original_redis.NewUniversalClient,
			"NewZSliceCmd":             original_redis.NewZSliceCmd,
			"NewZSliceCmdResult":       original_redis.NewZSliceCmdResult,
			"ParseURL":                 original_redis.ParseURL,
			"SetLogger":                original_redis.SetLogger,

			// Var and consts
			"Nil":         original_redis.Nil,
			"TxFailedErr": original_redis.TxFailedErr,

			// Types (value type)
			"BitCount":           func() original_redis.BitCount { return original_redis.BitCount{} },
			"BoolCmd":            func() original_redis.BoolCmd { return original_redis.BoolCmd{} },
			"BoolSliceCmd":       func() original_redis.BoolSliceCmd { return original_redis.BoolSliceCmd{} },
			"Client":             func() original_redis.Client { return original_redis.Client{} },
			"ClusterClient":      func() original_redis.ClusterClient { return original_redis.ClusterClient{} },
			"ClusterNode":        func() original_redis.ClusterNode { return original_redis.ClusterNode{} },
			"ClusterOptions":     func() original_redis.ClusterOptions { return original_redis.ClusterOptions{} },
			"ClusterSlot":        func() original_redis.ClusterSlot { return original_redis.ClusterSlot{} },
			"ClusterSlotsCmd":    func() original_redis.ClusterSlotsCmd { return original_redis.ClusterSlotsCmd{} },
			"Cmd":                func() original_redis.Cmd { return original_redis.Cmd{} },
			"CommandInfo":        func() original_redis.CommandInfo { return original_redis.CommandInfo{} },
			"CommandsInfoCmd":    func() original_redis.CommandsInfoCmd { return original_redis.CommandsInfoCmd{} },
			"Conn":               func() original_redis.Conn { return original_redis.Conn{} },
			"DurationCmd":        func() original_redis.DurationCmd { return original_redis.DurationCmd{} },
			"FailoverOptions":    func() original_redis.FailoverOptions { return original_redis.FailoverOptions{} },
			"FloatCmd":           func() original_redis.FloatCmd { return original_redis.FloatCmd{} },
			"GeoLocation":        func() original_redis.GeoLocation { return original_redis.GeoLocation{} },
			"GeoLocationCmd":     func() original_redis.GeoLocationCmd { return original_redis.GeoLocationCmd{} },
			"GeoPos":             func() original_redis.GeoPos { return original_redis.GeoPos{} },
			"GeoPosCmd":          func() original_redis.GeoPosCmd { return original_redis.GeoPosCmd{} },
			"GeoRadiusQuery":     func() original_redis.GeoRadiusQuery { return original_redis.GeoRadiusQuery{} },
			"IntCmd":             func() original_redis.IntCmd { return original_redis.IntCmd{} },
			"Message":            func() original_redis.Message { return original_redis.Message{} },
			"Options":            func() original_redis.Options { return original_redis.Options{} },
			"Pipeline":           func() original_redis.Pipeline { return original_redis.Pipeline{} },
			"Pong":               func() original_redis.Pong { return original_redis.Pong{} },
			"PubSub":             func() original_redis.PubSub { return original_redis.PubSub{} },
			"Ring":               func() original_redis.Ring { return original_redis.Ring{} },
			"RingOptions":        func() original_redis.RingOptions { return original_redis.RingOptions{} },
			"ScanCmd":            func() original_redis.ScanCmd { return original_redis.ScanCmd{} },
			"ScanIterator":       func() original_redis.ScanIterator { return original_redis.ScanIterator{} },
			"Script":             func() original_redis.Script { return original_redis.Script{} },
			"SliceCmd":           func() original_redis.SliceCmd { return original_redis.SliceCmd{} },
			"Sort":               func() original_redis.Sort { return original_redis.Sort{} },
			"StatusCmd":          func() original_redis.StatusCmd { return original_redis.StatusCmd{} },
			"StringCmd":          func() original_redis.StringCmd { return original_redis.StringCmd{} },
			"StringIntMapCmd":    func() original_redis.StringIntMapCmd { return original_redis.StringIntMapCmd{} },
			"StringSliceCmd":     func() original_redis.StringSliceCmd { return original_redis.StringSliceCmd{} },
			"StringStringMapCmd": func() original_redis.StringStringMapCmd { return original_redis.StringStringMapCmd{} },
			"StringStructMapCmd": func() original_redis.StringStructMapCmd { return original_redis.StringStructMapCmd{} },
			"Subscription":       func() original_redis.Subscription { return original_redis.Subscription{} },
			"TimeCmd":            func() original_redis.TimeCmd { return original_redis.TimeCmd{} },
			"Tx":                 func() original_redis.Tx { return original_redis.Tx{} },
			"UniversalOptions":   func() original_redis.UniversalOptions { return original_redis.UniversalOptions{} },
			"Z":                  func() original_redis.Z { return original_redis.Z{} },
			"ZRangeBy":           func() original_redis.ZRangeBy { return original_redis.ZRangeBy{} },
			"ZSliceCmd":          func() original_redis.ZSliceCmd { return original_redis.ZSliceCmd{} },
			"ZStore":             func() original_redis.ZStore { return original_redis.ZStore{} },

			// Types (pointer type)
			"NewBitCount":         func() *original_redis.BitCount { return &original_redis.BitCount{} },
			"NewClusterNode":      func() *original_redis.ClusterNode { return &original_redis.ClusterNode{} },
			"NewClusterOptions":   func() *original_redis.ClusterOptions { return &original_redis.ClusterOptions{} },
			"NewClusterSlot":      func() *original_redis.ClusterSlot { return &original_redis.ClusterSlot{} },
			"NewCommandInfo":      func() *original_redis.CommandInfo { return &original_redis.CommandInfo{} },
			"NewConn":             func() *original_redis.Conn { return &original_redis.Conn{} },
			"NewFailoverOptions":  func() *original_redis.FailoverOptions { return &original_redis.FailoverOptions{} },
			"NewGeoLocation":      func() *original_redis.GeoLocation { return &original_redis.GeoLocation{} },
			"NewGeoPos":           func() *original_redis.GeoPos { return &original_redis.GeoPos{} },
			"NewGeoRadiusQuery":   func() *original_redis.GeoRadiusQuery { return &original_redis.GeoRadiusQuery{} },
			"NewMessage":          func() *original_redis.Message { return &original_redis.Message{} },
			"NewOptions":          func() *original_redis.Options { return &original_redis.Options{} },
			"NewPipeline":         func() *original_redis.Pipeline { return &original_redis.Pipeline{} },
			"NewPong":             func() *original_redis.Pong { return &original_redis.Pong{} },
			"NewPubSub":           func() *original_redis.PubSub { return &original_redis.PubSub{} },
			"NewRingOptions":      func() *original_redis.RingOptions { return &original_redis.RingOptions{} },
			"NewScanIterator":     func() *original_redis.ScanIterator { return &original_redis.ScanIterator{} },
			"NewSort":             func() *original_redis.Sort { return &original_redis.Sort{} },
			"NewSubscription":     func() *original_redis.Subscription { return &original_redis.Subscription{} },
			"NewTx":               func() *original_redis.Tx { return &original_redis.Tx{} },
			"NewUniversalOptions": func() *original_redis.UniversalOptions { return &original_redis.UniversalOptions{} },
			"NewZ":                func() *original_redis.Z { return &original_redis.Z{} },
			"NewZRangeBy":         func() *original_redis.ZRangeBy { return &original_redis.ZRangeBy{} },
			"NewZStore":           func() *original_redis.ZStore { return &original_redis.ZStore{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}

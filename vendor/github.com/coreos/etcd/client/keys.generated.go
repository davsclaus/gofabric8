// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package client

import (
	"errors"
	"fmt"
	codec1978 "github.com/coreos/etcd/Godeps/_workspace/src/github.com/ugorji/go/codec"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81819 = 1
	codecSelferC_RAW1819  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1819 = 10
	codecSelferValueTypeMap1819   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1819    = 2
	codecSelfer_containerMapValue1819  = 3
	codecSelfer_containerMapEnd1819    = 4
	codecSelfer_containerArrayElem1819 = 6
	codecSelfer_containerArrayEnd1819  = 7
)

var (
	codecSelferBitsize1819                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1819 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1819 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 time.Time
		_ = v0
	}
}

func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(3)
			} else {
				yynn2 = 3
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81819, string(x.Action))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("action"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81819, string(x.Action))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				if x.Node == nil {
					r.EncodeNil()
				} else {
					x.Node.CodecEncodeSelf(e)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("node"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				if x.Node == nil {
					r.EncodeNil()
				} else {
					x.Node.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				if x.PrevNode == nil {
					r.EncodeNil()
				} else {
					x.PrevNode.CodecEncodeSelf(e)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("prevNode"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				if x.PrevNode == nil {
					r.EncodeNil()
				} else {
					x.PrevNode.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1819)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1819)
			}
		}
	}
}

func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym8 := z.DecBinary()
	_ = yym8
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct9 := r.ContainerType()
		if yyct9 == codecSelferValueTypeMap1819 {
			yyl9 := r.ReadMapStart()
			if yyl9 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1819)
			} else {
				x.codecDecodeSelfFromMap(yyl9, d)
			}
		} else if yyct9 == codecSelferValueTypeArray1819 {
			yyl9 := r.ReadArrayStart()
			if yyl9 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
			} else {
				x.codecDecodeSelfFromArray(yyl9, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1819)
		}
	}
}

func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys10Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys10Slc
	var yyhl10 bool = l >= 0
	for yyj10 := 0; ; yyj10++ {
		if yyhl10 {
			if yyj10 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1819)
		yys10Slc = r.DecodeBytes(yys10Slc, true, true)
		yys10 := string(yys10Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1819)
		switch yys10 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				x.Action = string(r.DecodeString())
			}
		case "node":
			if r.TryDecodeAsNil() {
				if x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}
				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if r.TryDecodeAsNil() {
				if x.PrevNode != nil {
					x.PrevNode = nil
				}
			} else {
				if x.PrevNode == nil {
					x.PrevNode = new(Node)
				}
				x.PrevNode.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys10)
		} // end switch yys10
	} // end for yyj10
	z.DecSendContainerState(codecSelfer_containerMapEnd1819)
}

func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj14 int
	var yyb14 bool
	var yyhl14 bool = l >= 0
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		x.Action = string(r.DecodeString())
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		if x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}
		x.Node.CodecDecodeSelf(d)
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		if x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}
		x.PrevNode.CodecDecodeSelf(d)
	}
	for {
		yyj14++
		if yyhl14 {
			yyb14 = yyj14 > l
		} else {
			yyb14 = r.CheckBreak()
		}
		if yyb14 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1819)
		z.DecStructFieldNotFound(yyj14-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
}

func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym18 := z.EncBinary()
		_ = yym18
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep19 := !z.EncBinary()
			yy2arr19 := z.EncBasicHandle().StructToArray
			var yyq19 [8]bool
			_, _, _ = yysep19, yyq19, yy2arr19
			const yyr19 bool = false
			yyq19[1] = x.Dir != false
			yyq19[6] = x.Expiration != nil
			yyq19[7] = x.TTL != 0
			var yynn19 int
			if yyr19 || yy2arr19 {
				r.EncodeArrayStart(8)
			} else {
				yynn19 = 5
				for _, b := range yyq19 {
					if b {
						yynn19++
					}
				}
				r.EncodeMapStart(yynn19)
				yynn19 = 0
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				yym21 := z.EncBinary()
				_ = yym21
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81819, string(x.Key))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("key"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				yym22 := z.EncBinary()
				_ = yym22
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81819, string(x.Key))
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				if yyq19[1] {
					yym24 := z.EncBinary()
					_ = yym24
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq19[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1819)
					r.EncodeString(codecSelferC_UTF81819, string("dir"))
					z.EncSendContainerState(codecSelfer_containerMapValue1819)
					yym25 := z.EncBinary()
					_ = yym25
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				yym27 := z.EncBinary()
				_ = yym27
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81819, string(x.Value))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("value"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				yym28 := z.EncBinary()
				_ = yym28
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81819, string(x.Value))
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("nodes"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				yym31 := z.EncBinary()
				_ = yym31
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("createdIndex"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				yym32 := z.EncBinary()
				_ = yym32
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				yym34 := z.EncBinary()
				_ = yym34
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1819)
				r.EncodeString(codecSelferC_UTF81819, string("modifiedIndex"))
				z.EncSendContainerState(codecSelfer_containerMapValue1819)
				yym35 := z.EncBinary()
				_ = yym35
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				if yyq19[6] {
					if x.Expiration == nil {
						r.EncodeNil()
					} else {
						yym37 := z.EncBinary()
						_ = yym37
						if false {
						} else if yym38 := z.TimeRtidIfBinc(); yym38 != 0 {
							r.EncodeBuiltin(yym38, x.Expiration)
						} else if z.HasExtensions() && z.EncExt(x.Expiration) {
						} else if yym37 {
							z.EncBinaryMarshal(x.Expiration)
						} else if !yym37 && z.IsJSONHandle() {
							z.EncJSONMarshal(x.Expiration)
						} else {
							z.EncFallback(x.Expiration)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq19[6] {
					z.EncSendContainerState(codecSelfer_containerMapKey1819)
					r.EncodeString(codecSelferC_UTF81819, string("expiration"))
					z.EncSendContainerState(codecSelfer_containerMapValue1819)
					if x.Expiration == nil {
						r.EncodeNil()
					} else {
						yym39 := z.EncBinary()
						_ = yym39
						if false {
						} else if yym40 := z.TimeRtidIfBinc(); yym40 != 0 {
							r.EncodeBuiltin(yym40, x.Expiration)
						} else if z.HasExtensions() && z.EncExt(x.Expiration) {
						} else if yym39 {
							z.EncBinaryMarshal(x.Expiration)
						} else if !yym39 && z.IsJSONHandle() {
							z.EncJSONMarshal(x.Expiration)
						} else {
							z.EncFallback(x.Expiration)
						}
					}
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1819)
				if yyq19[7] {
					yym42 := z.EncBinary()
					_ = yym42
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq19[7] {
					z.EncSendContainerState(codecSelfer_containerMapKey1819)
					r.EncodeString(codecSelferC_UTF81819, string("ttl"))
					z.EncSendContainerState(codecSelfer_containerMapValue1819)
					yym43 := z.EncBinary()
					_ = yym43
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				}
			}
			if yyr19 || yy2arr19 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1819)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1819)
			}
		}
	}
}

func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym44 := z.DecBinary()
	_ = yym44
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct45 := r.ContainerType()
		if yyct45 == codecSelferValueTypeMap1819 {
			yyl45 := r.ReadMapStart()
			if yyl45 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1819)
			} else {
				x.codecDecodeSelfFromMap(yyl45, d)
			}
		} else if yyct45 == codecSelferValueTypeArray1819 {
			yyl45 := r.ReadArrayStart()
			if yyl45 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
			} else {
				x.codecDecodeSelfFromArray(yyl45, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1819)
		}
	}
}

func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys46Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys46Slc
	var yyhl46 bool = l >= 0
	for yyj46 := 0; ; yyj46++ {
		if yyhl46 {
			if yyj46 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1819)
		yys46Slc = r.DecodeBytes(yys46Slc, true, true)
		yys46 := string(yys46Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1819)
		switch yys46 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = string(r.DecodeString())
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = bool(r.DecodeBool())
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = string(r.DecodeString())
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				yyv50 := &x.Nodes
				yyv50.CodecDecodeSelf(d)
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				x.CreatedIndex = uint64(r.DecodeUint(64))
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				x.ModifiedIndex = uint64(r.DecodeUint(64))
			}
		case "expiration":
			if r.TryDecodeAsNil() {
				if x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}
				yym54 := z.DecBinary()
				_ = yym54
				if false {
				} else if yym55 := z.TimeRtidIfBinc(); yym55 != 0 {
					r.DecodeBuiltin(yym55, x.Expiration)
				} else if z.HasExtensions() && z.DecExt(x.Expiration) {
				} else if yym54 {
					z.DecBinaryUnmarshal(x.Expiration)
				} else if !yym54 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(x.Expiration)
				} else {
					z.DecFallback(x.Expiration, false)
				}
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				x.TTL = int64(r.DecodeInt(64))
			}
		default:
			z.DecStructFieldNotFound(-1, yys46)
		} // end switch yys46
	} // end for yyj46
	z.DecSendContainerState(codecSelfer_containerMapEnd1819)
}

func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj57 int
	var yyb57 bool
	var yyhl57 bool = l >= 0
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = string(r.DecodeString())
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = bool(r.DecodeBool())
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = string(r.DecodeString())
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		yyv61 := &x.Nodes
		yyv61.CodecDecodeSelf(d)
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		x.CreatedIndex = uint64(r.DecodeUint(64))
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		x.ModifiedIndex = uint64(r.DecodeUint(64))
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		if x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}
		yym65 := z.DecBinary()
		_ = yym65
		if false {
		} else if yym66 := z.TimeRtidIfBinc(); yym66 != 0 {
			r.DecodeBuiltin(yym66, x.Expiration)
		} else if z.HasExtensions() && z.DecExt(x.Expiration) {
		} else if yym65 {
			z.DecBinaryUnmarshal(x.Expiration)
		} else if !yym65 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(x.Expiration)
		} else {
			z.DecFallback(x.Expiration, false)
		}
	}
	yyj57++
	if yyhl57 {
		yyb57 = yyj57 > l
	} else {
		yyb57 = r.CheckBreak()
	}
	if yyb57 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1819)
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		x.TTL = int64(r.DecodeInt(64))
	}
	for {
		yyj57++
		if yyhl57 {
			yyb57 = yyj57 > l
		} else {
			yyb57 = r.CheckBreak()
		}
		if yyb57 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1819)
		z.DecStructFieldNotFound(yyj57-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1819)
}

func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym68 := z.EncBinary()
		_ = yym68
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			h.encNodes((Nodes)(x), e)
		}
	}
}

func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym69 := z.DecBinary()
	_ = yym69
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.decNodes((*Nodes)(x), d)
	}
}

func (x codecSelfer1819) encNodes(v Nodes, e *codec1978.Encoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv70 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1819)
		if yyv70 == nil {
			r.EncodeNil()
		} else {
			yyv70.CodecEncodeSelf(e)
		}
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1819)
}

func (x codecSelfer1819) decNodes(v *Nodes, d *codec1978.Decoder) {
	var h codecSelfer1819
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv71 := *v
	yyh71, yyl71 := z.DecSliceHelperStart()
	var yyc71 bool
	if yyl71 == 0 {
		if yyv71 == nil {
			yyv71 = []*Node{}
			yyc71 = true
		} else if len(yyv71) != 0 {
			yyv71 = yyv71[:0]
			yyc71 = true
		}
	} else if yyl71 > 0 {
		var yyrr71, yyrl71 int
		var yyrt71 bool
		if yyl71 > cap(yyv71) {

			yyrg71 := len(yyv71) > 0
			yyv271 := yyv71
			yyrl71, yyrt71 = z.DecInferLen(yyl71, z.DecBasicHandle().MaxInitLen, 8)
			if yyrt71 {
				if yyrl71 <= cap(yyv71) {
					yyv71 = yyv71[:yyrl71]
				} else {
					yyv71 = make([]*Node, yyrl71)
				}
			} else {
				yyv71 = make([]*Node, yyrl71)
			}
			yyc71 = true
			yyrr71 = len(yyv71)
			if yyrg71 {
				copy(yyv71, yyv271)
			}
		} else if yyl71 != len(yyv71) {
			yyv71 = yyv71[:yyl71]
			yyc71 = true
		}
		yyj71 := 0
		for ; yyj71 < yyrr71; yyj71++ {
			yyh71.ElemContainerState(yyj71)
			if r.TryDecodeAsNil() {
				if yyv71[yyj71] != nil {
					*yyv71[yyj71] = Node{}
				}
			} else {
				if yyv71[yyj71] == nil {
					yyv71[yyj71] = new(Node)
				}
				yyw72 := yyv71[yyj71]
				yyw72.CodecDecodeSelf(d)
			}

		}
		if yyrt71 {
			for ; yyj71 < yyl71; yyj71++ {
				yyv71 = append(yyv71, nil)
				yyh71.ElemContainerState(yyj71)
				if r.TryDecodeAsNil() {
					if yyv71[yyj71] != nil {
						*yyv71[yyj71] = Node{}
					}
				} else {
					if yyv71[yyj71] == nil {
						yyv71[yyj71] = new(Node)
					}
					yyw73 := yyv71[yyj71]
					yyw73.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj71 := 0
		for ; !r.CheckBreak(); yyj71++ {

			if yyj71 >= len(yyv71) {
				yyv71 = append(yyv71, nil) // var yyz71 *Node
				yyc71 = true
			}
			yyh71.ElemContainerState(yyj71)
			if yyj71 < len(yyv71) {
				if r.TryDecodeAsNil() {
					if yyv71[yyj71] != nil {
						*yyv71[yyj71] = Node{}
					}
				} else {
					if yyv71[yyj71] == nil {
						yyv71[yyj71] = new(Node)
					}
					yyw74 := yyv71[yyj71]
					yyw74.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj71 < len(yyv71) {
			yyv71 = yyv71[:yyj71]
			yyc71 = true
		} else if yyj71 == 0 && yyv71 == nil {
			yyv71 = []*Node{}
			yyc71 = true
		}
	}
	yyh71.End()
	if yyc71 {
		*v = yyv71
	}
}

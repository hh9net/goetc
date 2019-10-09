// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Plate.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//状态
type PlateStateReport struct {
	State                string   `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlateStateReport) Reset()         { *m = PlateStateReport{} }
func (m *PlateStateReport) String() string { return proto.CompactTextString(m) }
func (*PlateStateReport) ProtoMessage()    {}
func (*PlateStateReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{0}
}

func (m *PlateStateReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlateStateReport.Unmarshal(m, b)
}
func (m *PlateStateReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlateStateReport.Marshal(b, m, deterministic)
}
func (m *PlateStateReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlateStateReport.Merge(m, src)
}
func (m *PlateStateReport) XXX_Size() int {
	return xxx_messageInfo_PlateStateReport.Size(m)
}
func (m *PlateStateReport) XXX_DiscardUnknown() {
	xxx_messageInfo_PlateStateReport.DiscardUnknown(m)
}

var xxx_messageInfo_PlateStateReport proto.InternalMessageInfo

func (m *PlateStateReport) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type ReqVehRecognize struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVehRecognize) Reset()         { *m = ReqVehRecognize{} }
func (m *ReqVehRecognize) String() string { return proto.CompactTextString(m) }
func (*ReqVehRecognize) ProtoMessage()    {}
func (*ReqVehRecognize) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{1}
}

func (m *ReqVehRecognize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqVehRecognize.Unmarshal(m, b)
}
func (m *ReqVehRecognize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqVehRecognize.Marshal(b, m, deterministic)
}
func (m *ReqVehRecognize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqVehRecognize.Merge(m, src)
}
func (m *ReqVehRecognize) XXX_Size() int {
	return xxx_messageInfo_ReqVehRecognize.Size(m)
}
func (m *ReqVehRecognize) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqVehRecognize.DiscardUnknown(m)
}

var xxx_messageInfo_ReqVehRecognize proto.InternalMessageInfo

type RstVehRecognize struct {
	VehicleInfo          *Car              `protobuf:"bytes,1,opt,name=VehicleInfo,proto3" json:"VehicleInfo,omitempty"`
	VehicleOtherInfo     *OtherVehicleInfo `protobuf:"bytes,2,opt,name=VehicleOtherInfo,proto3" json:"VehicleOtherInfo,omitempty"`
	VehicleFaceInfo      *TotalFaceInfo    `protobuf:"bytes,3,opt,name=VehicleFaceInfo,proto3" json:"VehicleFaceInfo,omitempty"`
	VehicleImg           *Img              `protobuf:"bytes,4,opt,name=VehicleImg,proto3" json:"VehicleImg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RstVehRecognize) Reset()         { *m = RstVehRecognize{} }
func (m *RstVehRecognize) String() string { return proto.CompactTextString(m) }
func (*RstVehRecognize) ProtoMessage()    {}
func (*RstVehRecognize) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{2}
}

func (m *RstVehRecognize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RstVehRecognize.Unmarshal(m, b)
}
func (m *RstVehRecognize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RstVehRecognize.Marshal(b, m, deterministic)
}
func (m *RstVehRecognize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RstVehRecognize.Merge(m, src)
}
func (m *RstVehRecognize) XXX_Size() int {
	return xxx_messageInfo_RstVehRecognize.Size(m)
}
func (m *RstVehRecognize) XXX_DiscardUnknown() {
	xxx_messageInfo_RstVehRecognize.DiscardUnknown(m)
}

var xxx_messageInfo_RstVehRecognize proto.InternalMessageInfo

func (m *RstVehRecognize) GetVehicleInfo() *Car {
	if m != nil {
		return m.VehicleInfo
	}
	return nil
}

func (m *RstVehRecognize) GetVehicleOtherInfo() *OtherVehicleInfo {
	if m != nil {
		return m.VehicleOtherInfo
	}
	return nil
}

func (m *RstVehRecognize) GetVehicleFaceInfo() *TotalFaceInfo {
	if m != nil {
		return m.VehicleFaceInfo
	}
	return nil
}

func (m *RstVehRecognize) GetVehicleImg() *Img {
	if m != nil {
		return m.VehicleImg
	}
	return nil
}

type CarInfo struct {
	Lpn                  string   `protobuf:"bytes,1,opt,name=Lpn,proto3" json:"Lpn,omitempty"`
	LpnScore             string   `protobuf:"bytes,2,opt,name=LpnScore,proto3" json:"LpnScore,omitempty"`
	LpnColor             string   `protobuf:"bytes,3,opt,name=LpnColor,proto3" json:"LpnColor,omitempty"`
	LpnposLeft           string   `protobuf:"bytes,4,opt,name=LpnposLeft,proto3" json:"LpnposLeft,omitempty"`
	LpnposTop            string   `protobuf:"bytes,5,opt,name=LpnposTop,proto3" json:"LpnposTop,omitempty"`
	LpnposRight          string   `protobuf:"bytes,6,opt,name=LpnposRight,proto3" json:"LpnposRight,omitempty"`
	LpnposBottom         string   `protobuf:"bytes,7,opt,name=LpnposBottom,proto3" json:"LpnposBottom,omitempty"`
	Color                string   `protobuf:"bytes,8,opt,name=Color,proto3" json:"Color,omitempty"`
	ColorScore           string   `protobuf:"bytes,9,opt,name=ColorScore,proto3" json:"ColorScore,omitempty"`
	Brand                string   `protobuf:"bytes,10,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Brand0               string   `protobuf:"bytes,11,opt,name=Brand0,proto3" json:"Brand0,omitempty"`
	Brand1               string   `protobuf:"bytes,12,opt,name=Brand1,proto3" json:"Brand1,omitempty"`
	Brand2               string   `protobuf:"bytes,13,opt,name=Brand2,proto3" json:"Brand2,omitempty"`
	Brand3               string   `protobuf:"bytes,14,opt,name=Brand3,proto3" json:"Brand3,omitempty"`
	Vehtype              string   `protobuf:"bytes,15,opt,name=Vehtype,proto3" json:"Vehtype,omitempty"`
	VehClass             string   `protobuf:"bytes,16,opt,name=VehClass,proto3" json:"VehClass,omitempty"`
	Type0                string   `protobuf:"bytes,17,opt,name=Type0,proto3" json:"Type0,omitempty"`
	Type1                string   `protobuf:"bytes,18,opt,name=Type1,proto3" json:"Type1,omitempty"`
	Type2                string   `protobuf:"bytes,19,opt,name=Type2,proto3" json:"Type2,omitempty"`
	Type3                string   `protobuf:"bytes,20,opt,name=Type3,proto3" json:"Type3,omitempty"`
	Subbrand             string   `protobuf:"bytes,21,opt,name=Subbrand,proto3" json:"Subbrand,omitempty"`
	Subbrand0            string   `protobuf:"bytes,22,opt,name=Subbrand0,proto3" json:"Subbrand0,omitempty"`
	Subbrand1            string   `protobuf:"bytes,23,opt,name=Subbrand1,proto3" json:"Subbrand1,omitempty"`
	Subbrand2            string   `protobuf:"bytes,24,opt,name=Subbrand2,proto3" json:"Subbrand2,omitempty"`
	Subbrand3            string   `protobuf:"bytes,25,opt,name=Subbrand3,proto3" json:"Subbrand3,omitempty"`
	Year                 string   `protobuf:"bytes,26,opt,name=Year,proto3" json:"Year,omitempty"`
	Year0                string   `protobuf:"bytes,27,opt,name=Year0,proto3" json:"Year0,omitempty"`
	Year1                string   `protobuf:"bytes,28,opt,name=Year1,proto3" json:"Year1,omitempty"`
	Year2                string   `protobuf:"bytes,29,opt,name=Year2,proto3" json:"Year2,omitempty"`
	Year3                string   `protobuf:"bytes,30,opt,name=Year3,proto3" json:"Year3,omitempty"`
	BrandScore           string   `protobuf:"bytes,31,opt,name=BrandScore,proto3" json:"BrandScore,omitempty"`
	BrandScore0          string   `protobuf:"bytes,32,opt,name=BrandScore0,proto3" json:"BrandScore0,omitempty"`
	BrandScore1          string   `protobuf:"bytes,33,opt,name=BrandScore1,proto3" json:"BrandScore1,omitempty"`
	BrandScore2          string   `protobuf:"bytes,34,opt,name=BrandScore2,proto3" json:"BrandScore2,omitempty"`
	BrandScore3          string   `protobuf:"bytes,35,opt,name=BrandScore3,proto3" json:"BrandScore3,omitempty"`
	Pose                 string   `protobuf:"bytes,36,opt,name=Pose,proto3" json:"Pose,omitempty"`
	CarposLeft           string   `protobuf:"bytes,37,opt,name=CarposLeft,proto3" json:"CarposLeft,omitempty"`
	CarposTop            string   `protobuf:"bytes,38,opt,name=CarposTop,proto3" json:"CarposTop,omitempty"`
	CarposRight          string   `protobuf:"bytes,39,opt,name=CarposRight,proto3" json:"CarposRight,omitempty"`
	CarposBottom         string   `protobuf:"bytes,40,opt,name=CarposBottom,proto3" json:"CarposBottom,omitempty"`
	CarRectScore         string   `protobuf:"bytes,41,opt,name=CarRectScore,proto3" json:"CarRectScore,omitempty"`
	ImgQuality           string   `protobuf:"bytes,42,opt,name=ImgQuality,proto3" json:"ImgQuality,omitempty"`
	IDNumber             string   `protobuf:"bytes,43,opt,name=IDNumber,proto3" json:"IDNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarInfo) Reset()         { *m = CarInfo{} }
func (m *CarInfo) String() string { return proto.CompactTextString(m) }
func (*CarInfo) ProtoMessage()    {}
func (*CarInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{3}
}

func (m *CarInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarInfo.Unmarshal(m, b)
}
func (m *CarInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarInfo.Marshal(b, m, deterministic)
}
func (m *CarInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarInfo.Merge(m, src)
}
func (m *CarInfo) XXX_Size() int {
	return xxx_messageInfo_CarInfo.Size(m)
}
func (m *CarInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CarInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CarInfo proto.InternalMessageInfo

func (m *CarInfo) GetLpn() string {
	if m != nil {
		return m.Lpn
	}
	return ""
}

func (m *CarInfo) GetLpnScore() string {
	if m != nil {
		return m.LpnScore
	}
	return ""
}

func (m *CarInfo) GetLpnColor() string {
	if m != nil {
		return m.LpnColor
	}
	return ""
}

func (m *CarInfo) GetLpnposLeft() string {
	if m != nil {
		return m.LpnposLeft
	}
	return ""
}

func (m *CarInfo) GetLpnposTop() string {
	if m != nil {
		return m.LpnposTop
	}
	return ""
}

func (m *CarInfo) GetLpnposRight() string {
	if m != nil {
		return m.LpnposRight
	}
	return ""
}

func (m *CarInfo) GetLpnposBottom() string {
	if m != nil {
		return m.LpnposBottom
	}
	return ""
}

func (m *CarInfo) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *CarInfo) GetColorScore() string {
	if m != nil {
		return m.ColorScore
	}
	return ""
}

func (m *CarInfo) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CarInfo) GetBrand0() string {
	if m != nil {
		return m.Brand0
	}
	return ""
}

func (m *CarInfo) GetBrand1() string {
	if m != nil {
		return m.Brand1
	}
	return ""
}

func (m *CarInfo) GetBrand2() string {
	if m != nil {
		return m.Brand2
	}
	return ""
}

func (m *CarInfo) GetBrand3() string {
	if m != nil {
		return m.Brand3
	}
	return ""
}

func (m *CarInfo) GetVehtype() string {
	if m != nil {
		return m.Vehtype
	}
	return ""
}

func (m *CarInfo) GetVehClass() string {
	if m != nil {
		return m.VehClass
	}
	return ""
}

func (m *CarInfo) GetType0() string {
	if m != nil {
		return m.Type0
	}
	return ""
}

func (m *CarInfo) GetType1() string {
	if m != nil {
		return m.Type1
	}
	return ""
}

func (m *CarInfo) GetType2() string {
	if m != nil {
		return m.Type2
	}
	return ""
}

func (m *CarInfo) GetType3() string {
	if m != nil {
		return m.Type3
	}
	return ""
}

func (m *CarInfo) GetSubbrand() string {
	if m != nil {
		return m.Subbrand
	}
	return ""
}

func (m *CarInfo) GetSubbrand0() string {
	if m != nil {
		return m.Subbrand0
	}
	return ""
}

func (m *CarInfo) GetSubbrand1() string {
	if m != nil {
		return m.Subbrand1
	}
	return ""
}

func (m *CarInfo) GetSubbrand2() string {
	if m != nil {
		return m.Subbrand2
	}
	return ""
}

func (m *CarInfo) GetSubbrand3() string {
	if m != nil {
		return m.Subbrand3
	}
	return ""
}

func (m *CarInfo) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *CarInfo) GetYear0() string {
	if m != nil {
		return m.Year0
	}
	return ""
}

func (m *CarInfo) GetYear1() string {
	if m != nil {
		return m.Year1
	}
	return ""
}

func (m *CarInfo) GetYear2() string {
	if m != nil {
		return m.Year2
	}
	return ""
}

func (m *CarInfo) GetYear3() string {
	if m != nil {
		return m.Year3
	}
	return ""
}

func (m *CarInfo) GetBrandScore() string {
	if m != nil {
		return m.BrandScore
	}
	return ""
}

func (m *CarInfo) GetBrandScore0() string {
	if m != nil {
		return m.BrandScore0
	}
	return ""
}

func (m *CarInfo) GetBrandScore1() string {
	if m != nil {
		return m.BrandScore1
	}
	return ""
}

func (m *CarInfo) GetBrandScore2() string {
	if m != nil {
		return m.BrandScore2
	}
	return ""
}

func (m *CarInfo) GetBrandScore3() string {
	if m != nil {
		return m.BrandScore3
	}
	return ""
}

func (m *CarInfo) GetPose() string {
	if m != nil {
		return m.Pose
	}
	return ""
}

func (m *CarInfo) GetCarposLeft() string {
	if m != nil {
		return m.CarposLeft
	}
	return ""
}

func (m *CarInfo) GetCarposTop() string {
	if m != nil {
		return m.CarposTop
	}
	return ""
}

func (m *CarInfo) GetCarposRight() string {
	if m != nil {
		return m.CarposRight
	}
	return ""
}

func (m *CarInfo) GetCarposBottom() string {
	if m != nil {
		return m.CarposBottom
	}
	return ""
}

func (m *CarInfo) GetCarRectScore() string {
	if m != nil {
		return m.CarRectScore
	}
	return ""
}

func (m *CarInfo) GetImgQuality() string {
	if m != nil {
		return m.ImgQuality
	}
	return ""
}

func (m *CarInfo) GetIDNumber() string {
	if m != nil {
		return m.IDNumber
	}
	return ""
}

type CarFeature struct {
	GJ                   string   `protobuf:"bytes,1,opt,name=GJ,proto3" json:"GJ,omitempty"`
	NJB                  string   `protobuf:"bytes,2,opt,name=NJB,proto3" json:"NJB,omitempty"`
	TC                   string   `protobuf:"bytes,3,opt,name=TC,proto3" json:"TC,omitempty"`
	AQD                  string   `protobuf:"bytes,4,opt,name=AQD,proto3" json:"AQD,omitempty"`
	DDH                  string   `protobuf:"bytes,5,opt,name=DDH,proto3" json:"DDH,omitempty"`
	ZYB                  string   `protobuf:"bytes,6,opt,name=ZYB,proto3" json:"ZYB,omitempty"`
	CZH                  string   `protobuf:"bytes,7,opt,name=CZH,proto3" json:"CZH,omitempty"`
	CRZ                  string   `protobuf:"bytes,8,opt,name=CRZ,proto3" json:"CRZ,omitempty"`
	XSB                  string   `protobuf:"bytes,9,opt,name=XSB,proto3" json:"XSB,omitempty"`
	JSY                  string   `protobuf:"bytes,10,opt,name=JSY,proto3" json:"JSY,omitempty"`
	LT                   string   `protobuf:"bytes,11,opt,name=LT,proto3" json:"LT,omitempty"`
	CD                   string   `protobuf:"bytes,12,opt,name=CD,proto3" json:"CD,omitempty"`
	JSC                  string   `protobuf:"bytes,13,opt,name=JSC,proto3" json:"JSC,omitempty"`
	HSJ                  string   `protobuf:"bytes,14,opt,name=HSJ,proto3" json:"HSJ,omitempty"`
	XLJ                  string   `protobuf:"bytes,15,opt,name=XLJ,proto3" json:"XLJ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarFeature) Reset()         { *m = CarFeature{} }
func (m *CarFeature) String() string { return proto.CompactTextString(m) }
func (*CarFeature) ProtoMessage()    {}
func (*CarFeature) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{4}
}

func (m *CarFeature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarFeature.Unmarshal(m, b)
}
func (m *CarFeature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarFeature.Marshal(b, m, deterministic)
}
func (m *CarFeature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarFeature.Merge(m, src)
}
func (m *CarFeature) XXX_Size() int {
	return xxx_messageInfo_CarFeature.Size(m)
}
func (m *CarFeature) XXX_DiscardUnknown() {
	xxx_messageInfo_CarFeature.DiscardUnknown(m)
}

var xxx_messageInfo_CarFeature proto.InternalMessageInfo

func (m *CarFeature) GetGJ() string {
	if m != nil {
		return m.GJ
	}
	return ""
}

func (m *CarFeature) GetNJB() string {
	if m != nil {
		return m.NJB
	}
	return ""
}

func (m *CarFeature) GetTC() string {
	if m != nil {
		return m.TC
	}
	return ""
}

func (m *CarFeature) GetAQD() string {
	if m != nil {
		return m.AQD
	}
	return ""
}

func (m *CarFeature) GetDDH() string {
	if m != nil {
		return m.DDH
	}
	return ""
}

func (m *CarFeature) GetZYB() string {
	if m != nil {
		return m.ZYB
	}
	return ""
}

func (m *CarFeature) GetCZH() string {
	if m != nil {
		return m.CZH
	}
	return ""
}

func (m *CarFeature) GetCRZ() string {
	if m != nil {
		return m.CRZ
	}
	return ""
}

func (m *CarFeature) GetXSB() string {
	if m != nil {
		return m.XSB
	}
	return ""
}

func (m *CarFeature) GetJSY() string {
	if m != nil {
		return m.JSY
	}
	return ""
}

func (m *CarFeature) GetLT() string {
	if m != nil {
		return m.LT
	}
	return ""
}

func (m *CarFeature) GetCD() string {
	if m != nil {
		return m.CD
	}
	return ""
}

func (m *CarFeature) GetJSC() string {
	if m != nil {
		return m.JSC
	}
	return ""
}

func (m *CarFeature) GetHSJ() string {
	if m != nil {
		return m.HSJ
	}
	return ""
}

func (m *CarFeature) GetXLJ() string {
	if m != nil {
		return m.XLJ
	}
	return ""
}

type TotalBike struct {
	BikeCount            string      `protobuf:"bytes,1,opt,name=BikeCount,proto3" json:"BikeCount,omitempty"`
	BikeRsds             []*BikeInfo `protobuf:"bytes,2,rep,name=BikeRsds,proto3" json:"BikeRsds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TotalBike) Reset()         { *m = TotalBike{} }
func (m *TotalBike) String() string { return proto.CompactTextString(m) }
func (*TotalBike) ProtoMessage()    {}
func (*TotalBike) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{5}
}

func (m *TotalBike) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalBike.Unmarshal(m, b)
}
func (m *TotalBike) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalBike.Marshal(b, m, deterministic)
}
func (m *TotalBike) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalBike.Merge(m, src)
}
func (m *TotalBike) XXX_Size() int {
	return xxx_messageInfo_TotalBike.Size(m)
}
func (m *TotalBike) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalBike.DiscardUnknown(m)
}

var xxx_messageInfo_TotalBike proto.InternalMessageInfo

func (m *TotalBike) GetBikeCount() string {
	if m != nil {
		return m.BikeCount
	}
	return ""
}

func (m *TotalBike) GetBikeRsds() []*BikeInfo {
	if m != nil {
		return m.BikeRsds
	}
	return nil
}

type BikeInfo struct {
	Posleft              string   `protobuf:"bytes,1,opt,name=Posleft,proto3" json:"Posleft,omitempty"`
	Postop               string   `protobuf:"bytes,2,opt,name=Postop,proto3" json:"Postop,omitempty"`
	Posright             string   `protobuf:"bytes,3,opt,name=Posright,proto3" json:"Posright,omitempty"`
	Posbottom            string   `protobuf:"bytes,4,opt,name=Posbottom,proto3" json:"Posbottom,omitempty"`
	Score                string   `protobuf:"bytes,5,opt,name=Score,proto3" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BikeInfo) Reset()         { *m = BikeInfo{} }
func (m *BikeInfo) String() string { return proto.CompactTextString(m) }
func (*BikeInfo) ProtoMessage()    {}
func (*BikeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{6}
}

func (m *BikeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BikeInfo.Unmarshal(m, b)
}
func (m *BikeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BikeInfo.Marshal(b, m, deterministic)
}
func (m *BikeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BikeInfo.Merge(m, src)
}
func (m *BikeInfo) XXX_Size() int {
	return xxx_messageInfo_BikeInfo.Size(m)
}
func (m *BikeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BikeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BikeInfo proto.InternalMessageInfo

func (m *BikeInfo) GetPosleft() string {
	if m != nil {
		return m.Posleft
	}
	return ""
}

func (m *BikeInfo) GetPostop() string {
	if m != nil {
		return m.Postop
	}
	return ""
}

func (m *BikeInfo) GetPosright() string {
	if m != nil {
		return m.Posright
	}
	return ""
}

func (m *BikeInfo) GetPosbottom() string {
	if m != nil {
		return m.Posbottom
	}
	return ""
}

func (m *BikeInfo) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

type TotalFaceInfo struct {
	FaceCount            string      `protobuf:"bytes,1,opt,name=FaceCount,proto3" json:"FaceCount,omitempty"`
	FaceRsds             []*FaceInfo `protobuf:"bytes,2,rep,name=FaceRsds,proto3" json:"FaceRsds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TotalFaceInfo) Reset()         { *m = TotalFaceInfo{} }
func (m *TotalFaceInfo) String() string { return proto.CompactTextString(m) }
func (*TotalFaceInfo) ProtoMessage()    {}
func (*TotalFaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{7}
}

func (m *TotalFaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalFaceInfo.Unmarshal(m, b)
}
func (m *TotalFaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalFaceInfo.Marshal(b, m, deterministic)
}
func (m *TotalFaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalFaceInfo.Merge(m, src)
}
func (m *TotalFaceInfo) XXX_Size() int {
	return xxx_messageInfo_TotalFaceInfo.Size(m)
}
func (m *TotalFaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalFaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TotalFaceInfo proto.InternalMessageInfo

func (m *TotalFaceInfo) GetFaceCount() string {
	if m != nil {
		return m.FaceCount
	}
	return ""
}

func (m *TotalFaceInfo) GetFaceRsds() []*FaceInfo {
	if m != nil {
		return m.FaceRsds
	}
	return nil
}

type FaceInfo struct {
	FaceRectScore        string   `protobuf:"bytes,1,opt,name=FaceRectScore,proto3" json:"FaceRectScore,omitempty"`
	FaceposLeft          string   `protobuf:"bytes,2,opt,name=FaceposLeft,proto3" json:"FaceposLeft,omitempty"`
	FaceposTop           string   `protobuf:"bytes,3,opt,name=FaceposTop,proto3" json:"FaceposTop,omitempty"`
	FaceposRight         string   `protobuf:"bytes,4,opt,name=FaceposRight,proto3" json:"FaceposRight,omitempty"`
	FaceposBottom        string   `protobuf:"bytes,5,opt,name=FaceposBottom,proto3" json:"FaceposBottom,omitempty"`
	FacekeyPoint1X       string   `protobuf:"bytes,6,opt,name=FacekeyPoint1X,proto3" json:"FacekeyPoint1X,omitempty"`
	FacekeyPoint1Y       string   `protobuf:"bytes,7,opt,name=FacekeyPoint1Y,proto3" json:"FacekeyPoint1Y,omitempty"`
	FacekeyPoint2X       string   `protobuf:"bytes,8,opt,name=FacekeyPoint2X,proto3" json:"FacekeyPoint2X,omitempty"`
	FacekeyPoint2Y       string   `protobuf:"bytes,9,opt,name=FacekeyPoint2Y,proto3" json:"FacekeyPoint2Y,omitempty"`
	FacekeyPoint3X       string   `protobuf:"bytes,10,opt,name=FacekeyPoint3X,proto3" json:"FacekeyPoint3X,omitempty"`
	FacekeyPoint3Y       string   `protobuf:"bytes,11,opt,name=FacekeyPoint3Y,proto3" json:"FacekeyPoint3Y,omitempty"`
	FacekeyPoint4X       string   `protobuf:"bytes,12,opt,name=FacekeyPoint4X,proto3" json:"FacekeyPoint4X,omitempty"`
	FacekeyPoint4Y       string   `protobuf:"bytes,13,opt,name=FacekeyPoint4Y,proto3" json:"FacekeyPoint4Y,omitempty"`
	FacekeyPoint5X       string   `protobuf:"bytes,14,opt,name=FacekeyPoint5X,proto3" json:"FacekeyPoint5X,omitempty"`
	FacekeyPoint5Y       string   `protobuf:"bytes,15,opt,name=FacekeyPoint5Y,proto3" json:"FacekeyPoint5Y,omitempty"`
	FaceFeature          string   `protobuf:"bytes,16,opt,name=FaceFeature,proto3" json:"FaceFeature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceInfo) Reset()         { *m = FaceInfo{} }
func (m *FaceInfo) String() string { return proto.CompactTextString(m) }
func (*FaceInfo) ProtoMessage()    {}
func (*FaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{8}
}

func (m *FaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceInfo.Unmarshal(m, b)
}
func (m *FaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceInfo.Marshal(b, m, deterministic)
}
func (m *FaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceInfo.Merge(m, src)
}
func (m *FaceInfo) XXX_Size() int {
	return xxx_messageInfo_FaceInfo.Size(m)
}
func (m *FaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FaceInfo proto.InternalMessageInfo

func (m *FaceInfo) GetFaceRectScore() string {
	if m != nil {
		return m.FaceRectScore
	}
	return ""
}

func (m *FaceInfo) GetFaceposLeft() string {
	if m != nil {
		return m.FaceposLeft
	}
	return ""
}

func (m *FaceInfo) GetFaceposTop() string {
	if m != nil {
		return m.FaceposTop
	}
	return ""
}

func (m *FaceInfo) GetFaceposRight() string {
	if m != nil {
		return m.FaceposRight
	}
	return ""
}

func (m *FaceInfo) GetFaceposBottom() string {
	if m != nil {
		return m.FaceposBottom
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint1X() string {
	if m != nil {
		return m.FacekeyPoint1X
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint1Y() string {
	if m != nil {
		return m.FacekeyPoint1Y
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint2X() string {
	if m != nil {
		return m.FacekeyPoint2X
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint2Y() string {
	if m != nil {
		return m.FacekeyPoint2Y
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint3X() string {
	if m != nil {
		return m.FacekeyPoint3X
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint3Y() string {
	if m != nil {
		return m.FacekeyPoint3Y
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint4X() string {
	if m != nil {
		return m.FacekeyPoint4X
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint4Y() string {
	if m != nil {
		return m.FacekeyPoint4Y
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint5X() string {
	if m != nil {
		return m.FacekeyPoint5X
	}
	return ""
}

func (m *FaceInfo) GetFacekeyPoint5Y() string {
	if m != nil {
		return m.FacekeyPoint5Y
	}
	return ""
}

func (m *FaceInfo) GetFaceFeature() string {
	if m != nil {
		return m.FaceFeature
	}
	return ""
}

type Img struct {
	PlateImgBuffer       string   `protobuf:"bytes,1,opt,name=PlateImgBuffer,proto3" json:"PlateImgBuffer,omitempty"`
	CarImgBuffer         string   `protobuf:"bytes,2,opt,name=CarImgBuffer,proto3" json:"CarImgBuffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Img) Reset()         { *m = Img{} }
func (m *Img) String() string { return proto.CompactTextString(m) }
func (*Img) ProtoMessage()    {}
func (*Img) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{9}
}

func (m *Img) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Img.Unmarshal(m, b)
}
func (m *Img) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Img.Marshal(b, m, deterministic)
}
func (m *Img) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Img.Merge(m, src)
}
func (m *Img) XXX_Size() int {
	return xxx_messageInfo_Img.Size(m)
}
func (m *Img) XXX_DiscardUnknown() {
	xxx_messageInfo_Img.DiscardUnknown(m)
}

var xxx_messageInfo_Img proto.InternalMessageInfo

func (m *Img) GetPlateImgBuffer() string {
	if m != nil {
		return m.PlateImgBuffer
	}
	return ""
}

func (m *Img) GetCarImgBuffer() string {
	if m != nil {
		return m.CarImgBuffer
	}
	return ""
}

type Car struct {
	CarBaseInfo          *CarInfo    `protobuf:"bytes,1,opt,name=CarBaseInfo,proto3" json:"CarBaseInfo,omitempty"`
	CarFeatureInfo       *CarFeature `protobuf:"bytes,2,opt,name=CarFeatureInfo,proto3" json:"CarFeatureInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Car) Reset()         { *m = Car{} }
func (m *Car) String() string { return proto.CompactTextString(m) }
func (*Car) ProtoMessage()    {}
func (*Car) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{10}
}

func (m *Car) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Car.Unmarshal(m, b)
}
func (m *Car) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Car.Marshal(b, m, deterministic)
}
func (m *Car) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Car.Merge(m, src)
}
func (m *Car) XXX_Size() int {
	return xxx_messageInfo_Car.Size(m)
}
func (m *Car) XXX_DiscardUnknown() {
	xxx_messageInfo_Car.DiscardUnknown(m)
}

var xxx_messageInfo_Car proto.InternalMessageInfo

func (m *Car) GetCarBaseInfo() *CarInfo {
	if m != nil {
		return m.CarBaseInfo
	}
	return nil
}

func (m *Car) GetCarFeatureInfo() *CarFeature {
	if m != nil {
		return m.CarFeatureInfo
	}
	return nil
}

type OtherVehicleInfo struct {
	TotalMotobikeInfo    *TotalBike `protobuf:"bytes,1,opt,name=TotalMotobikeInfo,proto3" json:"TotalMotobikeInfo,omitempty"`
	TotalTribikeInfo     *TotalBike `protobuf:"bytes,2,opt,name=TotalTribikeInfo,proto3" json:"TotalTribikeInfo,omitempty"`
	TotalBikeInfo        *TotalBike `protobuf:"bytes,3,opt,name=TotalBikeInfo,proto3" json:"TotalBikeInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OtherVehicleInfo) Reset()         { *m = OtherVehicleInfo{} }
func (m *OtherVehicleInfo) String() string { return proto.CompactTextString(m) }
func (*OtherVehicleInfo) ProtoMessage()    {}
func (*OtherVehicleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5566925a79f0f4, []int{11}
}

func (m *OtherVehicleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherVehicleInfo.Unmarshal(m, b)
}
func (m *OtherVehicleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherVehicleInfo.Marshal(b, m, deterministic)
}
func (m *OtherVehicleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherVehicleInfo.Merge(m, src)
}
func (m *OtherVehicleInfo) XXX_Size() int {
	return xxx_messageInfo_OtherVehicleInfo.Size(m)
}
func (m *OtherVehicleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherVehicleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OtherVehicleInfo proto.InternalMessageInfo

func (m *OtherVehicleInfo) GetTotalMotobikeInfo() *TotalBike {
	if m != nil {
		return m.TotalMotobikeInfo
	}
	return nil
}

func (m *OtherVehicleInfo) GetTotalTribikeInfo() *TotalBike {
	if m != nil {
		return m.TotalTribikeInfo
	}
	return nil
}

func (m *OtherVehicleInfo) GetTotalBikeInfo() *TotalBike {
	if m != nil {
		return m.TotalBikeInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PlateStateReport)(nil), "pb.PlateStateReport")
	proto.RegisterType((*ReqVehRecognize)(nil), "pb.ReqVehRecognize")
	proto.RegisterType((*RstVehRecognize)(nil), "pb.RstVehRecognize")
	proto.RegisterType((*CarInfo)(nil), "pb.CarInfo")
	proto.RegisterType((*CarFeature)(nil), "pb.CarFeature")
	proto.RegisterType((*TotalBike)(nil), "pb.TotalBike")
	proto.RegisterType((*BikeInfo)(nil), "pb.BikeInfo")
	proto.RegisterType((*TotalFaceInfo)(nil), "pb.TotalFaceInfo")
	proto.RegisterType((*FaceInfo)(nil), "pb.FaceInfo")
	proto.RegisterType((*Img)(nil), "pb.Img")
	proto.RegisterType((*Car)(nil), "pb.Car")
	proto.RegisterType((*OtherVehicleInfo)(nil), "pb.OtherVehicleInfo")
}

func init() { proto.RegisterFile("Plate.proto", fileDescriptor_7d5566925a79f0f4) }

var fileDescriptor_7d5566925a79f0f4 = []byte{
	// 1166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x97, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0xc7, 0x2b, 0xca, 0xa7, 0xc7, 0xf9, 0x70, 0x86, 0xb0, 0x0c, 0xcb, 0xb2, 0x04, 0xb1, 0xec,
	0x66, 0xa1, 0x48, 0x45, 0x12, 0x50, 0x45, 0xed, 0x05, 0x2c, 0xd7, 0x6e, 0xec, 0x32, 0x8b, 0x23,
	0xbb, 0x82, 0x9d, 0x9b, 0xec, 0x9d, 0xd8, 0xae, 0xb5, 0x3d, 0x42, 0x92, 0x0f, 0xe1, 0x19, 0x38,
	0xf1, 0x26, 0x3c, 0x01, 0x2f, 0xc3, 0x99, 0x67, 0xa0, 0xba, 0xa7, 0xf5, 0x35, 0xe6, 0xe4, 0xee,
	0x5f, 0xff, 0x35, 0x6a, 0x8d, 0xba, 0x7b, 0x64, 0x56, 0xef, 0x2d, 0xc2, 0x54, 0x5e, 0x46, 0xb1,
	0x4a, 0x15, 0xb7, 0xa2, 0xb1, 0x7d, 0xc1, 0x1a, 0x88, 0xfa, 0x69, 0x98, 0xca, 0x40, 0x46, 0x2a,
	0x4e, 0xf9, 0x19, 0xdb, 0x45, 0x57, 0x6c, 0x9d, 0x6f, 0x5d, 0xd4, 0x02, 0xed, 0xd8, 0xa7, 0xec,
	0x24, 0x90, 0xbf, 0xdd, 0xca, 0x59, 0x20, 0x27, 0x6a, 0xba, 0x9a, 0xff, 0x2e, 0xed, 0x7f, 0xb6,
	0xd8, 0x49, 0x90, 0xa4, 0x65, 0xc6, 0x5f, 0xb2, 0xfa, 0xad, 0x9c, 0xcd, 0x27, 0x0b, 0xd9, 0x5e,
	0xdd, 0x2b, 0x5c, 0xa2, 0xee, 0xee, 0x5f, 0x46, 0xe3, 0x4b, 0x3f, 0x8c, 0x83, 0x72, 0x8c, 0xff,
	0xc8, 0x1a, 0xe4, 0xfe, 0x92, 0xce, 0x64, 0x8c, 0x7a, 0x0b, 0xf5, 0x67, 0xa0, 0x47, 0x58, 0xd2,
	0x07, 0x1b, 0x6a, 0xfe, 0x8a, 0x9d, 0x10, 0x7b, 0x1d, 0x4e, 0xf4, 0x0d, 0xb7, 0x71, 0x81, 0x53,
	0x58, 0x60, 0xa0, 0xd2, 0x70, 0x91, 0x05, 0x02, 0x53, 0xc9, 0x5f, 0x30, 0x96, 0xad, 0xbe, 0x9c,
	0x8a, 0x9d, 0x22, 0xd1, 0xf6, 0x72, 0x1a, 0x94, 0x42, 0xf6, 0x5f, 0x35, 0xb6, 0xef, 0x87, 0xfa,
	0x8e, 0x0d, 0xb6, 0xdd, 0x8d, 0x56, 0xb4, 0x33, 0x60, 0xf2, 0xc7, 0xec, 0xa0, 0x1b, 0xad, 0xfa,
	0x13, 0x15, 0x4b, 0xcc, 0xbe, 0x16, 0xe4, 0x3e, 0xc5, 0x7c, 0xb5, 0x50, 0x31, 0x26, 0xa6, 0x63,
	0xe8, 0xf3, 0xa7, 0x8c, 0x75, 0xa3, 0x55, 0xa4, 0x92, 0xae, 0xbc, 0x4f, 0xf1, 0xf6, 0xb5, 0xa0,
	0x44, 0xf8, 0x13, 0x56, 0xd3, 0xde, 0x40, 0x45, 0x62, 0x17, 0xc3, 0x05, 0xe0, 0xe7, 0xac, 0xae,
	0x9d, 0x60, 0x3e, 0x9d, 0xa5, 0x62, 0x0f, 0xe3, 0x65, 0xc4, 0x6d, 0x76, 0xa8, 0xdd, 0xa6, 0x4a,
	0x53, 0xb5, 0x14, 0xfb, 0x28, 0xa9, 0x30, 0x78, 0xd3, 0x3a, 0xb9, 0x03, 0xfd, 0xa6, 0xf3, 0xcc,
	0xd0, 0xd0, 0xcf, 0x54, 0xd3, 0x99, 0x15, 0x04, 0xae, 0x6a, 0xc6, 0xe1, 0xea, 0x9d, 0x60, 0xfa,
	0x2a, 0x74, 0xf8, 0x23, 0xb6, 0x87, 0xc6, 0x95, 0xa8, 0x23, 0x26, 0x2f, 0xe7, 0x8e, 0x38, 0x2c,
	0x71, 0x27, 0xe7, 0xae, 0x38, 0x2a, 0x71, 0x37, 0xe7, 0x9e, 0x38, 0x2e, 0x71, 0x8f, 0x0b, 0xb6,
	0x7f, 0x2b, 0x67, 0xe9, 0x43, 0x24, 0xc5, 0x09, 0x06, 0x32, 0x17, 0x76, 0xf9, 0x56, 0xce, 0xfc,
	0x45, 0x98, 0x24, 0xa2, 0xa1, 0x77, 0x39, 0xf3, 0x21, 0xd7, 0xc1, 0x43, 0x24, 0xaf, 0xc4, 0xa9,
	0xce, 0x15, 0x9d, 0x8c, 0x3a, 0x82, 0x17, 0xd4, 0xc9, 0xa8, 0x2b, 0x3e, 0x28, 0xa8, 0x9b, 0x51,
	0x4f, 0x9c, 0x15, 0xd4, 0x83, 0x7b, 0xf6, 0xd7, 0xe3, 0x31, 0x6e, 0xc3, 0x87, 0xfa, 0x9e, 0x99,
	0x0f, 0x6f, 0x2e, 0xb3, 0xaf, 0xc4, 0x23, 0xfd, 0xe6, 0x72, 0x50, 0x8e, 0x3a, 0xe2, 0xa3, 0x6a,
	0xd4, 0x29, 0x47, 0x5d, 0x21, 0xaa, 0x51, 0xb7, 0x1c, 0xf5, 0xc4, 0xc7, 0xd5, 0xa8, 0xc7, 0x39,
	0xdb, 0x19, 0xc9, 0x30, 0x16, 0x8f, 0x31, 0x80, 0x36, 0x64, 0x0f, 0xbf, 0x57, 0xe2, 0x13, 0x9d,
	0x3d, 0x3a, 0x19, 0x75, 0xc4, 0x93, 0x82, 0x3a, 0x19, 0x75, 0xc5, 0xa7, 0x05, 0x75, 0x33, 0xea,
	0x89, 0xa7, 0x05, 0xf5, 0xa0, 0x46, 0xf0, 0xbd, 0xe8, 0x1a, 0xf9, 0x4c, 0xd7, 0x48, 0x41, 0xa0,
	0x3e, 0x0b, 0xef, 0x4a, 0x9c, 0xeb, 0xfa, 0x2c, 0xa1, 0xaa, 0xc2, 0x11, 0x9f, 0x9b, 0x0a, 0xa7,
	0xaa, 0x70, 0x85, 0x6d, 0x2a, 0xdc, 0xaa, 0xc2, 0x13, 0x5f, 0x98, 0x0a, 0xdc, 0x93, 0x9e, 0x4a,
	0xa4, 0x78, 0xa6, 0xf7, 0x04, 0x6c, 0xac, 0xef, 0x30, 0xce, 0x3a, 0xef, 0x4b, 0xaa, 0xef, 0x9c,
	0xc0, 0x2e, 0x6b, 0x0f, 0x3a, 0xef, 0xb9, 0xde, 0xe5, 0x1c, 0xc0, 0x3d, 0xb5, 0xa3, 0x3b, 0xef,
	0x85, 0xbe, 0x67, 0x09, 0x41, 0xe7, 0x69, 0x97, 0x3a, 0xef, 0x42, 0x77, 0x5e, 0x99, 0x91, 0x26,
	0x90, 0x93, 0x54, 0xef, 0xe0, 0xcb, 0x5c, 0x93, 0x33, 0xc8, 0xb3, 0xbd, 0x9c, 0xde, 0xac, 0xc3,
	0xc5, 0x3c, 0x7d, 0x10, 0x5f, 0xe9, 0x3c, 0x0b, 0x02, 0x35, 0xd8, 0x6e, 0xbd, 0x5d, 0x2f, 0xc7,
	0x32, 0x16, 0x5f, 0xeb, 0x1a, 0xcc, 0x7c, 0xfb, 0x4f, 0x0b, 0x1f, 0xf2, 0xb5, 0x0c, 0xd3, 0x75,
	0x2c, 0xf9, 0x31, 0xb3, 0xde, 0x74, 0x68, 0x6a, 0x59, 0x6f, 0x3a, 0x30, 0xc6, 0xde, 0x76, 0x9a,
	0x34, 0xaf, 0xc0, 0x04, 0xc5, 0xc0, 0xa7, 0x21, 0x65, 0x0d, 0x7c, 0x50, 0xfc, 0x74, 0xd3, 0xa2,
	0xb9, 0x04, 0x26, 0x90, 0x56, 0xeb, 0x9a, 0x46, 0x11, 0x98, 0x40, 0xee, 0x46, 0x4d, 0x1a, 0x3e,
	0x60, 0x02, 0xf1, 0xef, 0xae, 0x69, 0xd6, 0x80, 0x89, 0x24, 0xb8, 0xa3, 0x01, 0x03, 0x26, 0x90,
	0x61, 0xbf, 0x49, 0x73, 0x05, 0x4c, 0x20, 0x9d, 0xfe, 0x88, 0xc6, 0x09, 0x98, 0x90, 0x4d, 0x77,
	0x40, 0x83, 0xc4, 0xea, 0x0e, 0xc0, 0xf7, 0x5b, 0x34, 0x40, 0x2c, 0xbf, 0xa5, 0xaf, 0xf0, 0x69,
	0x72, 0x80, 0x09, 0xe4, 0xba, 0xdf, 0xa1, 0x99, 0x01, 0x26, 0xde, 0xa7, 0xdb, 0xa1, 0x61, 0x01,
	0xa6, 0xdd, 0x67, 0x35, 0x3c, 0x13, 0x9a, 0xf3, 0xf7, 0x12, 0xde, 0x32, 0xfc, 0xfa, 0x6a, 0xbd,
	0x4a, 0x69, 0x67, 0x0a, 0xc0, 0x2f, 0xd8, 0x01, 0x38, 0x41, 0xf2, 0x2e, 0x11, 0xd6, 0xf9, 0xf6,
	0x45, 0xdd, 0x3d, 0x84, 0xa3, 0x01, 0x18, 0x9e, 0x26, 0x79, 0xd4, 0xfe, 0x63, 0x4b, 0x4b, 0xf1,
	0x78, 0x10, 0x6c, 0xbf, 0xa7, 0x92, 0x05, 0xd4, 0x95, 0x5e, 0x32, 0x73, 0x61, 0xac, 0xf5, 0x54,
	0x92, 0xaa, 0x88, 0x36, 0x9d, 0x3c, 0x78, 0x89, 0x3d, 0x95, 0xc4, 0x58, 0x4b, 0x74, 0x44, 0x64,
	0x3e, 0xa4, 0xd8, 0x53, 0xc9, 0x58, 0x57, 0x91, 0x7e, 0x13, 0x05, 0xc0, 0x63, 0x1a, 0x6b, 0x67,
	0x97, 0x8e, 0x69, 0x70, 0xec, 0x5f, 0xd9, 0x51, 0xe5, 0xdc, 0x83, 0x45, 0xc0, 0xae, 0x3c, 0x67,
	0x0e, 0xe0, 0x39, 0xc1, 0x31, 0x9f, 0x33, 0x3f, 0x35, 0xf3, 0xa8, 0xfd, 0xef, 0x8e, 0x96, 0xe2,
	0xa2, 0xcf, 0xd8, 0x11, 0x06, 0xf2, 0xfa, 0xd5, 0x0b, 0x57, 0x21, 0xb4, 0x0a, 0x80, 0xac, 0xd3,
	0xf4, 0x83, 0x97, 0x11, 0x94, 0x38, 0xb9, 0xd0, 0x6b, 0xfa, 0xf9, 0x4b, 0x04, 0xda, 0x84, 0x3c,
	0xdd, 0x6d, 0x7a, 0x13, 0x2a, 0x2c, 0xcb, 0xa5, 0xe8, 0xb7, 0xdd, 0x22, 0x97, 0xa2, 0xe1, 0x9e,
	0xb3, 0x63, 0x00, 0xef, 0xe5, 0x43, 0x4f, 0xcd, 0x57, 0xa9, 0x33, 0xa4, 0xb2, 0x35, 0xe8, 0x86,
	0x6e, 0x44, 0xc5, 0x6c, 0x50, 0x53, 0xe7, 0x0e, 0xa9, 0xc4, 0x0d, 0xba, 0xa1, 0x1b, 0x51, 0xe1,
	0x1b, 0xd4, 0xd4, 0x79, 0x43, 0x6a, 0x07, 0x83, 0x6e, 0xe8, 0x46, 0xd4, 0x25, 0x06, 0x35, 0x75,
	0xdf, 0x0e, 0xa9, 0x7b, 0x0c, 0xba, 0xa1, 0x1b, 0x51, 0x53, 0x19, 0xd4, 0xd4, 0x7d, 0x37, 0xa4,
	0x56, 0x33, 0xe8, 0x86, 0x6e, 0x44, 0x0d, 0x68, 0xd0, 0xac, 0x36, 0x68, 0x40, 0xd1, 0xb9, 0x5d,
	0x46, 0xf6, 0x0d, 0xdb, 0x6e, 0x2f, 0xa7, 0xb0, 0x20, 0x7e, 0xa1, 0xb6, 0x97, 0xd3, 0xe6, 0xfa,
	0xfe, 0x5e, 0xc6, 0x54, 0x6b, 0x06, 0xa5, 0x89, 0x5a, 0xa8, 0xac, 0x7c, 0xa2, 0xe6, 0xcc, 0x5e,
	0xb0, 0x6d, 0x3f, 0x8c, 0xf9, 0x37, 0x38, 0xc2, 0x9b, 0x61, 0x52, 0xfe, 0x46, 0xad, 0xd3, 0x37,
	0x2a, 0x96, 0x7d, 0x39, 0xce, 0xbf, 0x67, 0xc7, 0xc5, 0x28, 0x2d, 0x7d, 0xa5, 0x1e, 0xd3, 0x15,
	0x14, 0x09, 0x0c, 0x95, 0xfd, 0xf7, 0x16, 0x6b, 0x98, 0x1f, 0xb1, 0xfc, 0x15, 0x3b, 0xc5, 0xfe,
	0xfc, 0x59, 0xa5, 0x6a, 0x4c, 0x63, 0x83, 0x32, 0x38, 0xca, 0x3f, 0x5a, 0x71, 0xb8, 0x6c, 0xea,
	0xf8, 0x0f, 0xac, 0x81, 0x70, 0x10, 0xcf, 0xf3, 0x6b, 0xad, 0xff, 0xbb, 0x76, 0x43, 0xc6, 0x3d,
	0x9a, 0x0b, 0xd9, 0xa8, 0xa2, 0x0f, 0x65, 0xe3, 0xba, 0xaa, 0x66, 0xbc, 0x87, 0x7f, 0x14, 0xbc,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x60, 0x3c, 0x7f, 0xab, 0x37, 0x0c, 0x00, 0x00,
}

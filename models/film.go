package models

import (
	"fmt"
	"reflect"
	"tableStore/lib"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

// public  final  static  String  OTS_COLUMN_ImageUrl  =  "ImageUrl";
// public  final  static  String  OTS_COLUMN_HospitalID  =  "HospitalID";
// public  final  static  String  OTS_COLUMN_HospitalCode  =  "HospitalCode";
// public  final  static  String  OTS_COLUMN_PatientName  =  "PatientName";
// public  final  static  String  OTS_COLUMN_PatientSex  =  "PatientSex";
// public  final  static  String  OTS_COLUMN_PatientAge  =  "PatientAge";
// public  final  static  String  OTS_COLUMN_PatientBirthDate  =  "PatientBirthDate";
// public  final  static  String  OTS_COLUMN_PatientSize  =  "PatientSize";
// public  final  static  String  OTS_COLUMN_PatientWeight  =  "PatientWeight";
// public  final  static  String  OTS_COLUMN_PatientID  =  "PatientID";
// public  final  static  String  OTS_COLUMN_Modality  =  "Modality";
// public  final  static  String  OTS_COLUMN_InstitutionName  =  "InstitutionName";
// public  final  static  String  OTS_COLUMN_InstitutionAddress  =  "InstitutionAddress";
// public  final  static  String  OTS_COLUMN_ImageType  =  "ImageType";
// public  final  static  String  OTS_COLUMN_InstanceCreationDate  =  "InstanceCreationDate";
// public  final  static  String  OTS_COLUMN_InstanceCreationTime  =  "InstanceCreationTime";
// public  final  static  String  OTS_COLUMN_SOPClassUID  =  "SOPClassUID";
// public  final  static  String  OTS_COLUMN_SOPInstanceUID  =  "SOPInstanceUID";
// public  final  static  String  OTS_COLUMN_StudyDate  =  "StudyDate";
// public  final  static  String  OTS_COLUMN_SeriesDate  =  "SeriesDate";
// public  final  static  String  OTS_COLUMN_AcquisitionDate  =  "AcquisitionDate";
// public  final  static  String  OTS_COLUMN_ContentDate  =  "ContentDate";
// public  final  static  String  OTS_COLUMN_StudyTime  =  "StudyTime";
// public  final  static  String  OTS_COLUMN_SeriesTime  =  "SeriesTime";
// public  final  static  String  OTS_COLUMN_AcquisitionTime  =  "AcquisitionTime";
// public  final  static  String  OTS_COLUMN_ContentTime  =  "ContentTime";
// public  final  static  String  OTS_COLUMN_Manufacturer  =  "Manufacturer";
// public  final  static  String  OTS_COLUMN_StationName  =  "StationName";
// public  final  static  String  OTS_COLUMN_StudyDescription  =  "StudyDescription";
// public  final  static  String  OTS_COLUMN_SeriesDescription  =  "SeriesDescription";
// public  final  static  String  OTS_COLUMN_InstitutionalDepartmentName  =  "InstitutionalDepartmentName";
// public  final  static  String  OTS_COLUMN_PerformingPhysicianName  =  "PerformingPhysicianName";
// public  final  static  String  OTS_COLUMN_ManufacturerModelName  =  "ManufacturerModelName";
// public  final  static  String  OTS_COLUMN_ScanningSequence  =  "ScanningSequence";
// public  final  static  String  OTS_COLUMN_SequenceVariant  =  "SequenceVariant";
// public  final  static  String  OTS_COLUMN_ScanOptions  =  "ScanOptions";
// public  final  static  String  OTS_COLUMN_MRAcquisitionType  =  "MRAcquisitionType";
// public  final  static  String  OTS_COLUMN_SequenceName  =  "SequenceName";
// public  final  static  String  OTS_COLUMN_AngioFlag  =  "AngioFlag";
// public  final  static  String  OTS_COLUMN_SliceTh ickness = "SliceThickness";
// public final static String OTS_COLUMN_RepetitionTime = "RepetitionTime";
// public final static String OTS_COLUMN_EchoTime = "EchoTime";
// public final static String OTS_COLUMN_NumberOfAverages = "NumberOfAverages";
// public final static String OTS_COLUMN_ImagingFrequency = "ImagingFrequency";
// public final static String OTS_COLUMN_ImagedNucleus = "ImagedNucleus";
// public final static String OTS_COLUMN_EchoNumbers = "EchoNumbers";
// public final static String OTS_COLUMN_MagneticFieldStrength = "MagneticFieldStrength";
// public final static String OTS_COLUMN_SpacingBetweenSlices = "SpacingBetweenSlices";
// public final static String OTS_COLUMN_NumberOfPhaseEncodingSteps = "NumberOfPhaseEncodingSteps";
// public final static String OTS_COLUMN_EchoTrainLength = "EchoTrainLength";
// public final static String OTS_COLUMN_PercentSampling = "PercentSampling";
// public final static String OTS_COLUMN_PercentPhaseFieldOfView = "PercentPhaseFieldOfView";
// public final static String OTS_COLUMN_PixelBandwidth = "PixelBandwidth";
// public final static String OTS_COLUMN_DeviceSerialNumber = "DeviceSerialNumber";
// public final static String OTS_COLUMN_SoftwareVersions = "SoftwareVersions";
// public final static String OTS_COLUMN_ProtocolName = "ProtocolName";
// public final static String OTS_COLUMN_TransmitCoilName = "TransmitCoilName";
// public final static String OTS_COLUMN_AcquisitionMatrix = "AcquisitionMatrix";
// public final static String OTS_COLUMN_InPlanePhaseEncodingDirection = "InPlanePhaseEncodingDirection";
// public final static String OTS_COLUMN_FlipAngle = "FlipAngle";
// public final static String OTS_COLUMN_VariableFlipAngleFlag = "VariableFlipAngleFlag";
// public final static String OTS_COLUMN_SAR = "SAR";
// public final static String OTS_COLUMN_PatientPosition = "PatientPosition";
// public final static String OTS_COLUMN_StudyInstanceUID = "StudyInstanceUID";
// public final static String OTS_COLUMN_SeriesInstanceUID = "SeriesInstanceUID";
// public final static String OTS_COLUMN_StudyID = "StudyID";
// public final static String OTS_COLUMN_SeriesNumber = "SeriesNumber";
// public final static String OTS_COLUMN_AcquisitionNumber = "AcquisitionNumber";
// public final static String OTS_COLUMN_InstanceNumber = "InstanceNumber";
// public final static String OTS_COLUMN_ImagePositionPatient = "ImagePositionPatient";
// public final static String OTS_COLUMN_WindowCenter = "WindowCenter";
// public final static String OTS_COLUMN_WindowWidth = "WindowWidth";

// //伊士通  检查部位
// public final static String OTS_COLUMN_BodyPartExamined ="BodyPartExamined";

// public final static String OTHER_PK="otherCode";
// public final static String PK="code";

// Film 图像对象
type Film struct {
	Code                          string
	OtherCode                     string
	ImageURL                      string
	HospitalID                    string
	HospitalCode                  string
	PatientName                   string
	PatientSex                    string
	PatientAge                    string
	PatientBirthDate              string
	PatientSize                   string
	PatientWeight                 string
	PatientID                     string
	Modality                      string
	InstitutionName               string
	InstitutionAddress            string
	ImageType                     string
	InstanceCreationDate          string
	InstanceCreationTime          string
	SOPClassUID                   string
	SOPInstanceUID                string
	StudyDate                     string
	SeriesDate                    string
	AcquisitionDate               string
	ContentDate                   string
	StudyTime                     string
	SeriesTime                    string
	AcquisitionTime               string
	ContentTime                   string
	Manufacturer                  string
	StationName                   string
	StudyDescription              string
	SeriesDescription             string
	InstitutionalDepartmentName   string
	PerformingPhysicianName       string
	ManufacturerModelName         string
	ScanningSequence              string
	SequenceVariant               string
	ScanOptions                   string
	MRAcquisitionType             string
	SequenceName                  string
	AngioFlag                     string
	SliceThickness                string
	RepetitionTime                string
	EchoTime                      string
	NumberOfAverages              string
	ImagingFrequency              string
	ImagedNucleus                 string
	EchoNumbers                   string
	MagneticFieldStrength         string
	SpacingBetweenSlices          string
	NumberOfPhaseEncodingSteps    string
	EchoTrainLength               string
	PercentSampling               string
	PercentPhaseFieldOfView       string
	PixelBandwidth                string
	DeviceSerialNumber            string
	SoftwareVersions              string
	ProtocolName                  string
	TransmitCoilName              string
	AcquisitionMatrix             string
	InPlanePhaseEncodingDirection string
	FlipAngle                     string
	VariableFlipAngleFlag         string
	SAR                           string
	PatientPosition               string
	StudyInstanceUID              string
	SeriesInstanceUID             string
	StudyID                       string
	SeriesNumber                  string
	AcquisitionNumber             string
	InstanceNumber                string
	ImagePositionPatient          string
	WindowCenter                  string
	WindowWidth                   string
}

func init() {
	// InitFilmTable()
}

// InitFilmTable 初始化数据表
func InitFilmTable() {
	tableList, ferr := lib.TsClient.ListTable()
	if ferr != nil {
		fmt.Println("Failed to list table")
		return
	}
	fmt.Println("List table result is")
	for _, table := range tableList.TableNames {
		fmt.Println("TableName: ", table)
		if table == "Film" {
			return
		}
	}
	createtableRequest := new(tablestore.CreateTableRequest)
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = "Film"
	tableMeta.AddPrimaryKeyColumn("Code", tablestore.PrimaryKeyType_STRING)
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createtableRequest.TableMeta = tableMeta
	createtableRequest.TableOption = tableOption
	createtableRequest.ReservedThroughput = reservedThroughput
	_, err := lib.TsClient.CreateTable(createtableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}
}

// AddOneFilm 向数据库中增加Film
func AddOneFilm(film Film) (Code string) {
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "Film"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("Code", film.Code)
	putRowChange.PrimaryKey = putPk
	t := reflect.TypeOf(film)
	v := reflect.ValueOf(film)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			if (t.Field(i).Name != "Film") && (t.Field(i).Name != "Code") {
				putRowChange.AddColumn(t.Field(i).Name, v.Field(i).Interface())
			}
		}
	}
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	_, err := lib.TsClient.PutRow(putRowRequest)
	if err != nil {
		fmt.Println("putrow failed with error:", err)
	} else {
		fmt.Println("putrow finished")
	}
	return film.Code
}

// UpdateOneFilm 向数据库中更新Film
func UpdateOneFilm(film Film) (Code string) {
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = "Film"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("Code", film.Code)
	updateRowChange.PrimaryKey = putPk
	t := reflect.TypeOf(film)
	v := reflect.ValueOf(film)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			if (t.Field(i).Name != "Film") && (t.Field(i).Name != "Code") {
				updateRowChange.PutColumn(t.Field(i).Name, v.Field(i).Interface())
			}
		}
	}
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err := lib.TsClient.UpdateRow(updateRowRequest)
	if err != nil {
		fmt.Println("update failed with error:", err)
	} else {
		fmt.Println("update finished")
	}
	return film.Code
}

// GetOneFilm 获得一个影像的数据
func GetOneFilm(Code string) (film *Film, err error) {
	var result Film
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("Code", Code)
	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = "Film"
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, e := lib.TsClient.GetRow(getRowRequest)
	if len(getResp.Columns) == 0 {
		return nil, nil
	}
	obj := reflect.ValueOf(&result)
	elem := obj.Elem()
	for i := 0; i < len(getResp.Columns); i++ {
		elem.FieldByName(getResp.Columns[i].ColumnName).SetString(getResp.Columns[i].Value.(string)) //反射注入
	}
	return &result, e
}

// DeleteOneFilm 删除一个影像
func DeleteOneFilm(Code string) error {
	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = "Film"
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("Code", Code)
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	_, err := lib.TsClient.DeleteRow(deleteRowReq)
	return err
}

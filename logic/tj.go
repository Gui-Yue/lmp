package logic

import (
	"github.com/influxdata/influxdb1-client/v2"
	"go.uber.org/zap"
	"lmp/dao/influxdb"
)

func DoQueryIRQ() (res []client.Result, err error) {
	// 调用dao层influxdb API
	res, err = influxdb.QueryDB(`select last("duration") from "picknext"`)
	if err != nil {
		zap.L().Error("ERROR in DoQueryIRQ():", zap.Error(err))
		return nil, err
	}
	return
}

func DoQueryCpuUtilize() (res []client.Result, err error) {
	res, err = influxdb.QueryDB(`select last("perce") from "picknext"`)
	if err != nil {
		zap.L().Error("ERROR in DoQueryIRQ():", zap.Error(err))
		return nil, err
	}
	return
}

func DoQueryPickNext() (res []client.Result, err error) {
	res, err = influxdb.QueryDB(`select last("duration") from "picknext"`)
	if err != nil {
		zap.L().Error("ERROR in DoQueryPickNext():", zap.Error(err))
		return nil, err
	}
	return
}

func DoQueryTaskSwitch() (res []client.Result, err error) {
	res, err = influxdb.QueryDB(`select last("duration") from "picknext"`)
	if err != nil {
		zap.L().Error("ERROR in DoQueryTaskSwitch():", zap.Error(err))
		return nil, err
	}
	return
}

func DoQueryHardDiskReadWriteTime() (res []client.Result, err error) {
	res, err = influxdb.QueryDB(`select last("lat") from "picknext"`)
	if err != nil {
		zap.L().Error("ERROR in DoQueryHardDiskReadWriteTime():", zap.Error(err))
		return nil, err
	}
	return
}

func DoQueryWaterMark() (res []client.Result, err error) {
	res, err = influxdb.QueryDB(``)
	if err != nil {
		zap.L().Error("ERROR in DoQueryWaterMark():", zap.Error(err))
		return nil, err
	}
	return
}
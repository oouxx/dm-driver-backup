/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"crypto/tls"
	"dm/security"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	Dm_build_334 = 8192
	Dm_build_335 = 2 * time.Second
)

type dm_build_336 struct {
	dm_build_337 *net.TCPConn
	dm_build_338 *tls.Conn
	dm_build_339 *Dm_build_0
	dm_build_340 *DmConnection
	dm_build_341 security.Cipher
	dm_build_342 bool
	dm_build_343 bool
	dm_build_344 *security.DhKey

	dm_build_345 bool
	dm_build_346 string
	dm_build_347 bool
}

func dm_build_348(dm_build_349 *DmConnection) (*dm_build_336, error) {
	dm_build_350, dm_build_351 := dm_build_353(dm_build_349.dmConnector.host+":"+strconv.Itoa(int(dm_build_349.dmConnector.port)), time.Duration(dm_build_349.dmConnector.socketTimeout)*time.Second)
	if dm_build_351 != nil {
		return nil, dm_build_351
	}

	dm_build_352 := dm_build_336{}
	dm_build_352.dm_build_337 = dm_build_350
	dm_build_352.dm_build_339 = Dm_build_3(Dm_build_615)
	dm_build_352.dm_build_340 = dm_build_349
	dm_build_352.dm_build_342 = false
	dm_build_352.dm_build_343 = false
	dm_build_352.dm_build_345 = false
	dm_build_352.dm_build_346 = ""
	dm_build_352.dm_build_347 = false
	dm_build_349.Access = &dm_build_352

	return &dm_build_352, nil
}

func dm_build_353(dm_build_354 string, dm_build_355 time.Duration) (*net.TCPConn, error) {
	dm_build_356, dm_build_357 := net.DialTimeout("tcp", dm_build_354, dm_build_355)
	if dm_build_357 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_354).throw()
	}

	if tcpConn, ok := dm_build_356.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_335)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_359 *dm_build_336) dm_build_358(dm_build_360 dm_build_736) bool {
	var dm_build_361 = dm_build_359.dm_build_340.dmConnector.compress
	if dm_build_360.dm_build_751() == Dm_build_643 || dm_build_361 == Dm_build_692 {
		return false
	}

	if dm_build_361 == Dm_build_690 {
		return true
	} else if dm_build_361 == Dm_build_691 {
		return !dm_build_359.dm_build_340.Local && dm_build_360.dm_build_749() > Dm_build_689
	}

	return false
}

func (dm_build_363 *dm_build_336) dm_build_362(dm_build_364 dm_build_736) bool {
	var dm_build_365 = dm_build_363.dm_build_340.dmConnector.compress
	if dm_build_364.dm_build_751() == Dm_build_643 || dm_build_365 == Dm_build_692 {
		return false
	}

	if dm_build_365 == Dm_build_690 {
		return true
	} else if dm_build_365 == Dm_build_691 {
		return dm_build_363.dm_build_339.Dm_build_267(Dm_build_651) == 1
	}

	return false
}

func (dm_build_367 *dm_build_336) dm_build_366(dm_build_368 dm_build_736) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_370 := dm_build_368.dm_build_749()

	if dm_build_370 > 0 {

		if dm_build_367.dm_build_358(dm_build_368) {
			var retBytes, err = Compress(dm_build_367.dm_build_339, Dm_build_644, int(dm_build_370), int(dm_build_367.dm_build_340.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_367.dm_build_339.Dm_build_14(Dm_build_644)

			dm_build_367.dm_build_339.Dm_build_55(dm_build_370)

			dm_build_367.dm_build_339.Dm_build_83(retBytes)

			dm_build_368.dm_build_750(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_367.dm_build_339.Dm_build_187(Dm_build_651, 1)
		}

		if dm_build_367.dm_build_343 {
			dm_build_370 = dm_build_368.dm_build_749()
			var retBytes = dm_build_367.dm_build_341.Encrypt(dm_build_367.dm_build_339.Dm_build_294(Dm_build_644, int(dm_build_370)), true)

			dm_build_367.dm_build_339.Dm_build_14(Dm_build_644)

			dm_build_367.dm_build_339.Dm_build_83(retBytes)

			dm_build_368.dm_build_750(int32(len(retBytes)))
		}
	}

	if dm_build_367.dm_build_339.Dm_build_12() > Dm_build_616 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_368.dm_build_745()
	if dm_build_367.dm_build_598(dm_build_368) {
		if dm_build_367.dm_build_338 != nil {
			dm_build_367.dm_build_339.Dm_build_17(0)
			if _, err := dm_build_367.dm_build_339.Dm_build_36(dm_build_367.dm_build_338); err != nil {
				return err
			}
		}
	} else {
		dm_build_367.dm_build_339.Dm_build_17(0)
		if _, err := dm_build_367.dm_build_339.Dm_build_36(dm_build_367.dm_build_337); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_372 *dm_build_336) dm_build_371(dm_build_373 dm_build_736) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_375 := int32(0)
	if dm_build_372.dm_build_598(dm_build_373) {
		if dm_build_372.dm_build_338 != nil {
			dm_build_372.dm_build_339.Dm_build_14(0)
			if _, err := dm_build_372.dm_build_339.Dm_build_30(dm_build_372.dm_build_338, Dm_build_644); err != nil {
				return err
			}

			dm_build_375 = dm_build_373.dm_build_749()
			if dm_build_375 > 0 {
				if _, err := dm_build_372.dm_build_339.Dm_build_30(dm_build_372.dm_build_338, int(dm_build_375)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_372.dm_build_339.Dm_build_14(0)
		if _, err := dm_build_372.dm_build_339.Dm_build_30(dm_build_372.dm_build_337, Dm_build_644); err != nil {
			return err
		}
		dm_build_375 = dm_build_373.dm_build_749()

		if dm_build_375 > 0 {
			if _, err := dm_build_372.dm_build_339.Dm_build_30(dm_build_372.dm_build_337, int(dm_build_375)); err != nil {
				return err
			}
		}
	}

	dm_build_373.dm_build_746()

	dm_build_375 = dm_build_373.dm_build_749()
	if dm_build_375 <= 0 {
		return nil
	}

	if dm_build_372.dm_build_343 {
		ebytes := dm_build_372.dm_build_339.Dm_build_294(Dm_build_644, int(dm_build_375))
		bytes, err := dm_build_372.dm_build_341.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_372.dm_build_339.Dm_build_14(Dm_build_644)
		dm_build_372.dm_build_339.Dm_build_83(bytes)
		dm_build_373.dm_build_750(int32(len(bytes)))
	}

	if dm_build_372.dm_build_362(dm_build_373) {

		dm_build_375 = dm_build_373.dm_build_749()
		cbytes := dm_build_372.dm_build_339.Dm_build_294(Dm_build_644+ULINT_SIZE, int(dm_build_375-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_372.dm_build_340.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_372.dm_build_339.Dm_build_14(Dm_build_644)
		dm_build_372.dm_build_339.Dm_build_83(bytes)
		dm_build_373.dm_build_750(int32(len(bytes)))
	}
	return nil
}

func (dm_build_377 *dm_build_336) dm_build_376(dm_build_378 dm_build_736) (dm_build_379 interface{}, dm_build_380 error) {
	dm_build_380 = dm_build_378.dm_build_740(dm_build_378)
	if dm_build_380 != nil {
		return nil, dm_build_380
	}

	dm_build_380 = dm_build_377.dm_build_366(dm_build_378)
	if dm_build_380 != nil {
		return nil, dm_build_380
	}

	dm_build_380 = dm_build_377.dm_build_371(dm_build_378)
	if dm_build_380 != nil {
		return nil, dm_build_380
	}

	return dm_build_378.dm_build_744(dm_build_378)
}

func (dm_build_382 *dm_build_336) dm_build_381() (*dm_build_1190, error) {

	Dm_build_383 := dm_build_1196(dm_build_382)
	_, dm_build_384 := dm_build_382.dm_build_376(Dm_build_383)
	if dm_build_384 != nil {
		return nil, dm_build_384
	}

	return Dm_build_383, nil
}

func (dm_build_386 *dm_build_336) dm_build_385() error {

	dm_build_387 := dm_build_1057(dm_build_386)
	_, dm_build_388 := dm_build_386.dm_build_376(dm_build_387)
	if dm_build_388 != nil {
		return dm_build_388
	}

	return nil
}

func (dm_build_390 *dm_build_336) dm_build_389() error {

	var dm_build_391 *dm_build_1190
	var err error
	if dm_build_391, err = dm_build_390.dm_build_381(); err != nil {
		return err
	}

	if dm_build_390.dm_build_340.sslEncrypt == 2 {
		if err = dm_build_390.dm_build_594(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_390.dm_build_340.sslEncrypt == 1 {
		if err = dm_build_390.dm_build_594(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_390.dm_build_343 || dm_build_390.dm_build_342 {
		k, err := dm_build_390.dm_build_584()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_391.Dm_build_1194)
		encryptType := dm_build_391.dm_build_1192
		hashType := int(dm_build_391.Dm_build_1193)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_390.dm_build_587(encryptType, sessionKey, dm_build_390.dm_build_340.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_390.dm_build_385(); err != nil {
		return err
	}
	return nil
}

func (dm_build_394 *dm_build_336) Dm_build_393(dm_build_395 *DmStatement) error {
	dm_build_396 := dm_build_1219(dm_build_394, dm_build_395)
	_, dm_build_397 := dm_build_394.dm_build_376(dm_build_396)
	if dm_build_397 != nil {
		return dm_build_397
	}

	return nil
}

func (dm_build_399 *dm_build_336) Dm_build_398(dm_build_400 int32) error {
	dm_build_401 := dm_build_1229(dm_build_399, dm_build_400)
	_, dm_build_402 := dm_build_399.dm_build_376(dm_build_401)
	if dm_build_402 != nil {
		return dm_build_402
	}

	return nil
}

func (dm_build_404 *dm_build_336) Dm_build_403(dm_build_405 *DmStatement, dm_build_406 bool, dm_build_407 int16) (*execRetInfo, error) {
	dm_build_408 := dm_build_1096(dm_build_404, dm_build_405, dm_build_406, dm_build_407)
	dm_build_409, dm_build_410 := dm_build_404.dm_build_376(dm_build_408)
	if dm_build_410 != nil {
		return nil, dm_build_410
	}
	return dm_build_409.(*execRetInfo), nil
}

func (dm_build_412 *dm_build_336) Dm_build_411(dm_build_413 *DmStatement, dm_build_414 int16) (*execRetInfo, error) {
	return dm_build_412.Dm_build_403(dm_build_413, false, Dm_build_696)
}

func (dm_build_416 *dm_build_336) Dm_build_415(dm_build_417 *DmStatement, dm_build_418 []OptParameter) (*execRetInfo, error) {
	dm_build_419, dm_build_420 := dm_build_416.dm_build_376(dm_build_839(dm_build_416, dm_build_417, dm_build_418))
	if dm_build_420 != nil {
		return nil, dm_build_420
	}

	return dm_build_419.(*execRetInfo), nil
}

func (dm_build_422 *dm_build_336) Dm_build_421(dm_build_423 *DmStatement, dm_build_424 int16) (*execRetInfo, error) {
	return dm_build_422.Dm_build_403(dm_build_423, true, dm_build_424)
}

func (dm_build_426 *dm_build_336) Dm_build_425(dm_build_427 *DmStatement, dm_build_428 [][]interface{}) (*execRetInfo, error) {
	dm_build_429 := dm_build_871(dm_build_426, dm_build_427, dm_build_428)
	dm_build_430, dm_build_431 := dm_build_426.dm_build_376(dm_build_429)
	if dm_build_431 != nil {
		return nil, dm_build_431
	}
	return dm_build_430.(*execRetInfo), nil
}

func (dm_build_433 *dm_build_336) Dm_build_432(dm_build_434 *DmStatement, dm_build_435 [][]interface{}, dm_build_436 bool) (*execRetInfo, error) {
	var dm_build_437, dm_build_438 = 0, 0
	var dm_build_439 = len(dm_build_435)
	var dm_build_440 [][]interface{}
	var dm_build_441 = NewExceInfo()
	dm_build_441.updateCounts = make([]int64, dm_build_439)
	var dm_build_442 = false
	for dm_build_437 < dm_build_439 {
		for dm_build_438 = dm_build_437; dm_build_438 < dm_build_439; dm_build_438++ {
			paramData := dm_build_435[dm_build_438]
			bindData := make([]interface{}, dm_build_434.paramCount)
			dm_build_442 = false
			for icol := 0; icol < int(dm_build_434.paramCount); icol++ {
				if dm_build_434.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_433.dm_build_567(bindData, paramData, icol) {
					dm_build_442 = true
					break
				}
			}

			if dm_build_442 {
				break
			}
			dm_build_440 = append(dm_build_440, bindData)
		}

		if dm_build_438 != dm_build_437 {
			tmpExecInfo, err := dm_build_433.Dm_build_425(dm_build_434, dm_build_440)
			if err != nil {
				return nil, err
			}
			dm_build_440 = dm_build_440[0:0]
			dm_build_441.union(tmpExecInfo, dm_build_437, dm_build_438-dm_build_437)
		}

		if dm_build_438 < dm_build_439 {
			tmpExecInfo, err := dm_build_433.Dm_build_443(dm_build_434, dm_build_435[dm_build_438], dm_build_436)
			if err != nil {
				return nil, err
			}

			dm_build_436 = true
			dm_build_441.union(tmpExecInfo, dm_build_438, 1)
		}

		dm_build_437 = dm_build_438 + 1
	}
	for _, i := range dm_build_441.updateCounts {
		if i > 0 {
			dm_build_441.updateCount += i
		}
	}
	return dm_build_441, nil
}

func (dm_build_444 *dm_build_336) Dm_build_443(dm_build_445 *DmStatement, dm_build_446 []interface{}, dm_build_447 bool) (*execRetInfo, error) {

	var dm_build_448 = make([]interface{}, dm_build_445.paramCount)
	for icol := 0; icol < int(dm_build_445.paramCount); icol++ {
		if dm_build_445.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_444.dm_build_567(dm_build_448, dm_build_446, icol) {

			if !dm_build_447 {
				preExecute := dm_build_1085(dm_build_444, dm_build_445, dm_build_445.bindParams)
				dm_build_444.dm_build_376(preExecute)
				dm_build_447 = true
			}

			dm_build_444.dm_build_573(dm_build_445, dm_build_445.bindParams[icol], icol, dm_build_446[icol].(iOffRowBinder))
			dm_build_448[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_449 = make([][]interface{}, 1, 1)
	dm_build_449[0] = dm_build_448

	dm_build_450 := dm_build_871(dm_build_444, dm_build_445, dm_build_449)
	dm_build_451, dm_build_452 := dm_build_444.dm_build_376(dm_build_450)
	if dm_build_452 != nil {
		return nil, dm_build_452
	}
	return dm_build_451.(*execRetInfo), nil
}

func (dm_build_454 *dm_build_336) Dm_build_453(dm_build_455 *DmStatement, dm_build_456 int16) (*execRetInfo, error) {
	dm_build_457 := dm_build_1072(dm_build_454, dm_build_455, dm_build_456)

	dm_build_458, dm_build_459 := dm_build_454.dm_build_376(dm_build_457)
	if dm_build_459 != nil {
		return nil, dm_build_459
	}
	return dm_build_458.(*execRetInfo), nil
}

func (dm_build_461 *dm_build_336) Dm_build_460(dm_build_462 *innerRows, dm_build_463 int64) (*execRetInfo, error) {
	dm_build_464 := dm_build_977(dm_build_461, dm_build_462, dm_build_463, INT64_MAX)
	dm_build_465, dm_build_466 := dm_build_461.dm_build_376(dm_build_464)
	if dm_build_466 != nil {
		return nil, dm_build_466
	}
	return dm_build_465.(*execRetInfo), nil
}

func (dm_build_468 *dm_build_336) Commit() error {
	dm_build_469 := dm_build_824(dm_build_468)
	_, dm_build_470 := dm_build_468.dm_build_376(dm_build_469)
	if dm_build_470 != nil {
		return dm_build_470
	}

	return nil
}

func (dm_build_472 *dm_build_336) Rollback() error {
	dm_build_473 := dm_build_1134(dm_build_472)
	_, dm_build_474 := dm_build_472.dm_build_376(dm_build_473)
	if dm_build_474 != nil {
		return dm_build_474
	}

	return nil
}

func (dm_build_476 *dm_build_336) Dm_build_475(dm_build_477 *DmConnection) error {
	dm_build_478 := dm_build_1139(dm_build_476, dm_build_477.IsoLevel)
	_, dm_build_479 := dm_build_476.dm_build_376(dm_build_478)
	if dm_build_479 != nil {
		return dm_build_479
	}

	return nil
}

func (dm_build_481 *dm_build_336) Dm_build_480(dm_build_482 *DmStatement, dm_build_483 string) error {
	dm_build_484 := dm_build_829(dm_build_481, dm_build_482, dm_build_483)
	_, dm_build_485 := dm_build_481.dm_build_376(dm_build_484)
	if dm_build_485 != nil {
		return dm_build_485
	}

	return nil
}

func (dm_build_487 *dm_build_336) Dm_build_486(dm_build_488 []uint32) ([]int64, error) {
	dm_build_489 := dm_build_1237(dm_build_487, dm_build_488)
	dm_build_490, dm_build_491 := dm_build_487.dm_build_376(dm_build_489)
	if dm_build_491 != nil {
		return nil, dm_build_491
	}
	return dm_build_490.([]int64), nil
}

func (dm_build_493 *dm_build_336) Close() error {
	if dm_build_493.dm_build_347 {
		return nil
	}

	dm_build_494 := dm_build_493.dm_build_337.Close()
	if dm_build_494 != nil {
		return dm_build_494
	}

	dm_build_493.dm_build_340 = nil
	dm_build_493.dm_build_347 = true
	return nil
}

func (dm_build_496 *dm_build_336) dm_build_495(dm_build_497 *lob) (int64, error) {
	dm_build_498 := dm_build_1008(dm_build_496, dm_build_497)
	dm_build_499, dm_build_500 := dm_build_496.dm_build_376(dm_build_498)
	if dm_build_500 != nil {
		return 0, dm_build_500
	}
	return dm_build_499.(int64), nil
}

func (dm_build_502 *dm_build_336) dm_build_501(dm_build_503 *lob, dm_build_504 int32, dm_build_505 int32) ([]byte, error) {
	dm_build_506 := dm_build_995(dm_build_502, dm_build_503, int(dm_build_504), int(dm_build_505))
	dm_build_507, dm_build_508 := dm_build_502.dm_build_376(dm_build_506)
	if dm_build_508 != nil {
		return nil, dm_build_508
	}
	return dm_build_507.([]byte), nil
}

func (dm_build_510 *dm_build_336) dm_build_509(dm_build_511 *DmBlob, dm_build_512 int32, dm_build_513 int32) ([]byte, error) {
	var dm_build_514 = make([]byte, dm_build_513)
	var dm_build_515 int32 = 0
	var dm_build_516 int32 = 0
	var dm_build_517 []byte
	var dm_build_518 error
	for dm_build_515 < dm_build_513 {
		dm_build_516 = dm_build_513 - dm_build_515
		if dm_build_516 > Dm_build_729 {
			dm_build_516 = Dm_build_729
		}
		dm_build_517, dm_build_518 = dm_build_510.dm_build_501(&dm_build_511.lob, dm_build_512, dm_build_516)
		if dm_build_518 != nil {
			return nil, dm_build_518
		}
		if dm_build_517 == nil || len(dm_build_517) == 0 {
			break
		}
		Dm_build_1248.Dm_build_1304(dm_build_514, int(dm_build_515), dm_build_517, 0, len(dm_build_517))
		dm_build_515 += int32(len(dm_build_517))
		dm_build_512 += int32(len(dm_build_517))
		if dm_build_511.readOver {
			break
		}
	}
	return dm_build_514, nil
}

func (dm_build_520 *dm_build_336) dm_build_519(dm_build_521 *DmClob, dm_build_522 int32, dm_build_523 int32) (string, error) {
	var dm_build_524 bytes.Buffer
	var dm_build_525 int32 = 0
	var dm_build_526 int32 = 0
	var dm_build_527 []byte
	var dm_build_528 string
	var dm_build_529 error
	for dm_build_525 < dm_build_523 {
		dm_build_526 = dm_build_523 - dm_build_525
		if dm_build_526 > Dm_build_729/2 {
			dm_build_526 = Dm_build_729 / 2
		}
		dm_build_527, dm_build_529 = dm_build_520.dm_build_501(&dm_build_521.lob, dm_build_522, dm_build_526)
		if dm_build_529 != nil {
			return "", dm_build_529
		}
		if dm_build_527 == nil || len(dm_build_527) == 0 {
			break
		}
		dm_build_528 = Dm_build_1248.Dm_build_1405(dm_build_527, 0, len(dm_build_527), dm_build_521.serverEncoding, dm_build_520.dm_build_340)

		dm_build_524.WriteString(dm_build_528)
		strLen := utf8.RuneCountInString(dm_build_528)
		dm_build_525 += int32(strLen)
		dm_build_522 += int32(strLen)
		if dm_build_521.readOver {
			break
		}
	}
	return dm_build_524.String(), nil
}

func (dm_build_531 *dm_build_336) dm_build_530(dm_build_532 *DmClob, dm_build_533 int, dm_build_534 string, dm_build_535 string) (int, error) {
	var dm_build_536 = Dm_build_1248.Dm_build_1464(dm_build_534, dm_build_535, dm_build_531.dm_build_340)
	var dm_build_537 = 0
	var dm_build_538 = len(dm_build_536)
	var dm_build_539 = 0
	var dm_build_540 = 0
	var dm_build_541 = 0
	var dm_build_542 = dm_build_538/Dm_build_728 + 1
	var dm_build_543 byte = 0
	var dm_build_544 byte = 0x01
	var dm_build_545 byte = 0x02
	for i := 0; i < dm_build_542; i++ {
		dm_build_543 = 0
		if i == 0 {
			dm_build_543 |= dm_build_544
		}
		if i == dm_build_542-1 {
			dm_build_543 |= dm_build_545
		}
		dm_build_541 = dm_build_538 - dm_build_540
		if dm_build_541 > Dm_build_728 {
			dm_build_541 = Dm_build_728
		}

		setLobData := dm_build_1153(dm_build_531, &dm_build_532.lob, dm_build_543, dm_build_533, dm_build_536, dm_build_537, dm_build_541)
		ret, err := dm_build_531.dm_build_376(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_539, nil
		} else {
			dm_build_533 += int(tmp)
			dm_build_539 += int(tmp)
			dm_build_540 += dm_build_541
			dm_build_537 += dm_build_541
		}
	}
	return dm_build_539, nil
}

func (dm_build_547 *dm_build_336) dm_build_546(dm_build_548 *DmBlob, dm_build_549 int, dm_build_550 []byte) (int, error) {
	var dm_build_551 = 0
	var dm_build_552 = len(dm_build_550)
	var dm_build_553 = 0
	var dm_build_554 = 0
	var dm_build_555 = 0
	var dm_build_556 = dm_build_552/Dm_build_728 + 1
	var dm_build_557 byte = 0
	var dm_build_558 byte = 0x01
	var dm_build_559 byte = 0x02
	for i := 0; i < dm_build_556; i++ {
		dm_build_557 = 0
		if i == 0 {
			dm_build_557 |= dm_build_558
		}
		if i == dm_build_556-1 {
			dm_build_557 |= dm_build_559
		}
		dm_build_555 = dm_build_552 - dm_build_554
		if dm_build_555 > Dm_build_728 {
			dm_build_555 = Dm_build_728
		}

		setLobData := dm_build_1153(dm_build_547, &dm_build_548.lob, dm_build_557, dm_build_549, dm_build_550, dm_build_551, dm_build_555)
		ret, err := dm_build_547.dm_build_376(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_553, nil
		} else {
			dm_build_549 += int(tmp)
			dm_build_553 += int(tmp)
			dm_build_554 += dm_build_555
			dm_build_551 += dm_build_555
		}
	}
	return dm_build_553, nil
}

func (dm_build_561 *dm_build_336) dm_build_560(dm_build_562 *lob, dm_build_563 int) (int64, error) {
	dm_build_564 := dm_build_1019(dm_build_561, dm_build_562, dm_build_563)
	dm_build_565, dm_build_566 := dm_build_561.dm_build_376(dm_build_564)
	if dm_build_566 != nil {
		return dm_build_562.length, dm_build_566
	}
	return dm_build_565.(int64), nil
}

func (dm_build_568 *dm_build_336) dm_build_567(dm_build_569 []interface{}, dm_build_570 []interface{}, dm_build_571 int) bool {
	var dm_build_572 = false
	dm_build_569[dm_build_571] = dm_build_570[dm_build_571]

	if binder, ok := dm_build_570[dm_build_571].(iOffRowBinder); ok {
		dm_build_572 = true
		dm_build_569[dm_build_571] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_568.dm_build_340) {
			dm_build_569[dm_build_571] = &lobCtl{lob.buildCtlData()}
			dm_build_572 = false
		}
	} else {
		dm_build_569[dm_build_571] = dm_build_570[dm_build_571]
	}
	return dm_build_572
}

func (dm_build_574 *dm_build_336) dm_build_573(dm_build_575 *DmStatement, dm_build_576 parameter, dm_build_577 int, dm_build_578 iOffRowBinder) error {
	var dm_build_579 = Dm_build_1534()
	dm_build_578.read(dm_build_579)
	var dm_build_580 = 0
	for !dm_build_578.isReadOver() || dm_build_579.Dm_build_1535() > 0 {
		if !dm_build_578.isReadOver() && dm_build_579.Dm_build_1535() < Dm_build_728 {
			dm_build_578.read(dm_build_579)
		}
		if dm_build_579.Dm_build_1535() > Dm_build_728 {
			dm_build_580 = Dm_build_728
		} else {
			dm_build_580 = dm_build_579.Dm_build_1535()
		}

		putData := dm_build_1124(dm_build_574, dm_build_575, int16(dm_build_577), dm_build_579, int32(dm_build_580))
		_, err := dm_build_574.dm_build_376(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_582 *dm_build_336) dm_build_581() ([]byte, error) {
	var dm_build_583 error
	if dm_build_582.dm_build_344 == nil {
		if dm_build_582.dm_build_344, dm_build_583 = security.NewClientKeyPair(); dm_build_583 != nil {
			return nil, dm_build_583
		}
	}
	return security.Bn2Bytes(dm_build_582.dm_build_344.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_585 *dm_build_336) dm_build_584() (*security.DhKey, error) {
	var dm_build_586 error
	if dm_build_585.dm_build_344 == nil {
		if dm_build_585.dm_build_344, dm_build_586 = security.NewClientKeyPair(); dm_build_586 != nil {
			return nil, dm_build_586
		}
	}
	return dm_build_585.dm_build_344, nil
}

func (dm_build_588 *dm_build_336) dm_build_587(dm_build_589 int, dm_build_590 []byte, dm_build_591 string, dm_build_592 int) (dm_build_593 error) {
	if dm_build_589 > 0 && dm_build_589 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_590 != nil {
		dm_build_588.dm_build_341, dm_build_593 = security.NewSymmCipher(dm_build_589, dm_build_590)
	} else if dm_build_589 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_588.dm_build_341, dm_build_593 = security.NewThirdPartCipher(dm_build_589, dm_build_590, dm_build_591, dm_build_592); dm_build_593 != nil {
			dm_build_593 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_593.Error()).throw()
		}
	}
	return
}

func (dm_build_595 *dm_build_336) dm_build_594(dm_build_596 bool) (dm_build_597 error) {
	if dm_build_595.dm_build_338, dm_build_597 = security.NewTLSFromTCP(dm_build_595.dm_build_337, dm_build_595.dm_build_340.dmConnector.sslCertPath, dm_build_595.dm_build_340.dmConnector.sslKeyPath, dm_build_595.dm_build_340.dmConnector.user); dm_build_597 != nil {
		return
	}
	if !dm_build_596 {
		dm_build_595.dm_build_338 = nil
	}
	return
}

func (dm_build_599 *dm_build_336) dm_build_598(dm_build_600 dm_build_736) bool {
	return dm_build_600.dm_build_751() != Dm_build_643 && dm_build_599.dm_build_340.sslEncrypt == 1
}

/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_1247 struct{}

var Dm_build_1248 = &dm_build_1247{}

func (Dm_build_1250 *dm_build_1247) Dm_build_1249(dm_build_1251 []byte, dm_build_1252 int, dm_build_1253 byte) int {
	dm_build_1251[dm_build_1252] = dm_build_1253
	return 1
}

func (Dm_build_1255 *dm_build_1247) Dm_build_1254(dm_build_1256 []byte, dm_build_1257 int, dm_build_1258 int8) int {
	dm_build_1256[dm_build_1257] = byte(dm_build_1258)
	return 1
}

func (Dm_build_1260 *dm_build_1247) Dm_build_1259(dm_build_1261 []byte, dm_build_1262 int, dm_build_1263 int16) int {
	dm_build_1261[dm_build_1262] = byte(dm_build_1263)
	dm_build_1262++
	dm_build_1261[dm_build_1262] = byte(dm_build_1263 >> 8)
	return 2
}

func (Dm_build_1265 *dm_build_1247) Dm_build_1264(dm_build_1266 []byte, dm_build_1267 int, dm_build_1268 int32) int {
	dm_build_1266[dm_build_1267] = byte(dm_build_1268)
	dm_build_1267++
	dm_build_1266[dm_build_1267] = byte(dm_build_1268 >> 8)
	dm_build_1267++
	dm_build_1266[dm_build_1267] = byte(dm_build_1268 >> 16)
	dm_build_1267++
	dm_build_1266[dm_build_1267] = byte(dm_build_1268 >> 24)
	dm_build_1267++
	return 4
}

func (Dm_build_1270 *dm_build_1247) Dm_build_1269(dm_build_1271 []byte, dm_build_1272 int, dm_build_1273 int64) int {
	dm_build_1271[dm_build_1272] = byte(dm_build_1273)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 8)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 16)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 24)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 32)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 40)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 48)
	dm_build_1272++
	dm_build_1271[dm_build_1272] = byte(dm_build_1273 >> 56)
	return 8
}

func (Dm_build_1275 *dm_build_1247) Dm_build_1274(dm_build_1276 []byte, dm_build_1277 int, dm_build_1278 float32) int {
	return Dm_build_1275.Dm_build_1294(dm_build_1276, dm_build_1277, math.Float32bits(dm_build_1278))
}

func (Dm_build_1280 *dm_build_1247) Dm_build_1279(dm_build_1281 []byte, dm_build_1282 int, dm_build_1283 float64) int {
	return Dm_build_1280.Dm_build_1299(dm_build_1281, dm_build_1282, math.Float64bits(dm_build_1283))
}

func (Dm_build_1285 *dm_build_1247) Dm_build_1284(dm_build_1286 []byte, dm_build_1287 int, dm_build_1288 uint8) int {
	dm_build_1286[dm_build_1287] = byte(dm_build_1288)
	return 1
}

func (Dm_build_1290 *dm_build_1247) Dm_build_1289(dm_build_1291 []byte, dm_build_1292 int, dm_build_1293 uint16) int {
	dm_build_1291[dm_build_1292] = byte(dm_build_1293)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 8)
	return 2
}

func (Dm_build_1295 *dm_build_1247) Dm_build_1294(dm_build_1296 []byte, dm_build_1297 int, dm_build_1298 uint32) int {
	dm_build_1296[dm_build_1297] = byte(dm_build_1298)
	dm_build_1297++
	dm_build_1296[dm_build_1297] = byte(dm_build_1298 >> 8)
	dm_build_1297++
	dm_build_1296[dm_build_1297] = byte(dm_build_1298 >> 16)
	dm_build_1297++
	dm_build_1296[dm_build_1297] = byte(dm_build_1298 >> 24)
	return 3
}

func (Dm_build_1300 *dm_build_1247) Dm_build_1299(dm_build_1301 []byte, dm_build_1302 int, dm_build_1303 uint64) int {
	dm_build_1301[dm_build_1302] = byte(dm_build_1303)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 8)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 16)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 24)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 32)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 40)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 48)
	dm_build_1302++
	dm_build_1301[dm_build_1302] = byte(dm_build_1303 >> 56)
	return 3
}

func (Dm_build_1305 *dm_build_1247) Dm_build_1304(dm_build_1306 []byte, dm_build_1307 int, dm_build_1308 []byte, dm_build_1309 int, dm_build_1310 int) int {
	copy(dm_build_1306[dm_build_1307:dm_build_1307+dm_build_1310], dm_build_1308[dm_build_1309:dm_build_1309+dm_build_1310])
	return dm_build_1310
}

func (Dm_build_1312 *dm_build_1247) Dm_build_1311(dm_build_1313 []byte, dm_build_1314 int, dm_build_1315 []byte, dm_build_1316 int, dm_build_1317 int) int {
	dm_build_1314 += Dm_build_1312.Dm_build_1294(dm_build_1313, dm_build_1314, uint32(dm_build_1317))
	return 4 + Dm_build_1312.Dm_build_1304(dm_build_1313, dm_build_1314, dm_build_1315, dm_build_1316, dm_build_1317)
}

func (Dm_build_1319 *dm_build_1247) Dm_build_1318(dm_build_1320 []byte, dm_build_1321 int, dm_build_1322 []byte, dm_build_1323 int, dm_build_1324 int) int {
	dm_build_1321 += Dm_build_1319.Dm_build_1289(dm_build_1320, dm_build_1321, uint16(dm_build_1324))
	return 2 + Dm_build_1319.Dm_build_1304(dm_build_1320, dm_build_1321, dm_build_1322, dm_build_1323, dm_build_1324)
}

func (Dm_build_1326 *dm_build_1247) Dm_build_1325(dm_build_1327 []byte, dm_build_1328 int, dm_build_1329 string, dm_build_1330 string, dm_build_1331 *DmConnection) int {
	dm_build_1332 := Dm_build_1326.Dm_build_1464(dm_build_1329, dm_build_1330, dm_build_1331)
	dm_build_1328 += Dm_build_1326.Dm_build_1294(dm_build_1327, dm_build_1328, uint32(len(dm_build_1332)))
	return 4 + Dm_build_1326.Dm_build_1304(dm_build_1327, dm_build_1328, dm_build_1332, 0, len(dm_build_1332))
}

func (Dm_build_1334 *dm_build_1247) Dm_build_1333(dm_build_1335 []byte, dm_build_1336 int, dm_build_1337 string, dm_build_1338 string, dm_build_1339 *DmConnection) int {
	dm_build_1340 := Dm_build_1334.Dm_build_1464(dm_build_1337, dm_build_1338, dm_build_1339)

	dm_build_1336 += Dm_build_1334.Dm_build_1289(dm_build_1335, dm_build_1336, uint16(len(dm_build_1340)))
	return 2 + Dm_build_1334.Dm_build_1304(dm_build_1335, dm_build_1336, dm_build_1340, 0, len(dm_build_1340))
}

func (Dm_build_1342 *dm_build_1247) Dm_build_1341(dm_build_1343 []byte, dm_build_1344 int) byte {
	return dm_build_1343[dm_build_1344]
}

func (Dm_build_1346 *dm_build_1247) Dm_build_1345(dm_build_1347 []byte, dm_build_1348 int) int16 {
	var dm_build_1349 int16
	dm_build_1349 = int16(dm_build_1347[dm_build_1348] & 0xff)
	dm_build_1348++
	dm_build_1349 |= int16(dm_build_1347[dm_build_1348]&0xff) << 8
	return dm_build_1349
}

func (Dm_build_1351 *dm_build_1247) Dm_build_1350(dm_build_1352 []byte, dm_build_1353 int) int32 {
	var dm_build_1354 int32
	dm_build_1354 = int32(dm_build_1352[dm_build_1353] & 0xff)
	dm_build_1353++
	dm_build_1354 |= int32(dm_build_1352[dm_build_1353]&0xff) << 8
	dm_build_1353++
	dm_build_1354 |= int32(dm_build_1352[dm_build_1353]&0xff) << 16
	dm_build_1353++
	dm_build_1354 |= int32(dm_build_1352[dm_build_1353]&0xff) << 24
	return dm_build_1354
}

func (Dm_build_1356 *dm_build_1247) Dm_build_1355(dm_build_1357 []byte, dm_build_1358 int) int64 {
	var dm_build_1359 int64
	dm_build_1359 = int64(dm_build_1357[dm_build_1358] & 0xff)
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 8
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 16
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 24
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 32
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 40
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 48
	dm_build_1358++
	dm_build_1359 |= int64(dm_build_1357[dm_build_1358]&0xff) << 56
	return dm_build_1359
}

func (Dm_build_1361 *dm_build_1247) Dm_build_1360(dm_build_1362 []byte, dm_build_1363 int) float32 {
	return math.Float32frombits(Dm_build_1361.Dm_build_1377(dm_build_1362, dm_build_1363))
}

func (Dm_build_1365 *dm_build_1247) Dm_build_1364(dm_build_1366 []byte, dm_build_1367 int) float64 {
	return math.Float64frombits(Dm_build_1365.Dm_build_1382(dm_build_1366, dm_build_1367))
}

func (Dm_build_1369 *dm_build_1247) Dm_build_1368(dm_build_1370 []byte, dm_build_1371 int) uint8 {
	return uint8(dm_build_1370[dm_build_1371] & 0xff)
}

func (Dm_build_1373 *dm_build_1247) Dm_build_1372(dm_build_1374 []byte, dm_build_1375 int) uint16 {
	var dm_build_1376 uint16
	dm_build_1376 = uint16(dm_build_1374[dm_build_1375] & 0xff)
	dm_build_1375++
	dm_build_1376 |= uint16(dm_build_1374[dm_build_1375]&0xff) << 8
	return dm_build_1376
}

func (Dm_build_1378 *dm_build_1247) Dm_build_1377(dm_build_1379 []byte, dm_build_1380 int) uint32 {
	var dm_build_1381 uint32
	dm_build_1381 = uint32(dm_build_1379[dm_build_1380] & 0xff)
	dm_build_1380++
	dm_build_1381 |= uint32(dm_build_1379[dm_build_1380]&0xff) << 8
	dm_build_1380++
	dm_build_1381 |= uint32(dm_build_1379[dm_build_1380]&0xff) << 16
	dm_build_1380++
	dm_build_1381 |= uint32(dm_build_1379[dm_build_1380]&0xff) << 24
	return dm_build_1381
}

func (Dm_build_1383 *dm_build_1247) Dm_build_1382(dm_build_1384 []byte, dm_build_1385 int) uint64 {
	var dm_build_1386 uint64
	dm_build_1386 = uint64(dm_build_1384[dm_build_1385] & 0xff)
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 8
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 16
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 24
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 32
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 40
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 48
	dm_build_1385++
	dm_build_1386 |= uint64(dm_build_1384[dm_build_1385]&0xff) << 56
	return dm_build_1386
}

func (Dm_build_1388 *dm_build_1247) Dm_build_1387(dm_build_1389 []byte, dm_build_1390 int) []byte {
	dm_build_1391 := Dm_build_1388.Dm_build_1377(dm_build_1389, dm_build_1390)

	dm_build_1392 := make([]byte, dm_build_1391)
	copy(dm_build_1392[:int(dm_build_1391)], dm_build_1389[dm_build_1390+4:dm_build_1390+4+int(dm_build_1391)])
	return dm_build_1392
}

func (Dm_build_1394 *dm_build_1247) Dm_build_1393(dm_build_1395 []byte, dm_build_1396 int) []byte {
	dm_build_1397 := Dm_build_1394.Dm_build_1372(dm_build_1395, dm_build_1396)

	dm_build_1398 := make([]byte, dm_build_1397)
	copy(dm_build_1398[:int(dm_build_1397)], dm_build_1395[dm_build_1396+2:dm_build_1396+2+int(dm_build_1397)])
	return dm_build_1398
}

func (Dm_build_1400 *dm_build_1247) Dm_build_1399(dm_build_1401 []byte, dm_build_1402 int, dm_build_1403 int) []byte {

	dm_build_1404 := make([]byte, dm_build_1403)
	copy(dm_build_1404[:dm_build_1403], dm_build_1401[dm_build_1402:dm_build_1402+dm_build_1403])
	return dm_build_1404
}

func (Dm_build_1406 *dm_build_1247) Dm_build_1405(dm_build_1407 []byte, dm_build_1408 int, dm_build_1409 int, dm_build_1410 string, dm_build_1411 *DmConnection) string {
	return Dm_build_1406.Dm_build_1501(dm_build_1407[dm_build_1408:dm_build_1408+dm_build_1409], dm_build_1410, dm_build_1411)
}

func (Dm_build_1413 *dm_build_1247) Dm_build_1412(dm_build_1414 []byte, dm_build_1415 int, dm_build_1416 string, dm_build_1417 *DmConnection) string {
	dm_build_1418 := Dm_build_1413.Dm_build_1377(dm_build_1414, dm_build_1415)
	dm_build_1415 += 4
	return Dm_build_1413.Dm_build_1405(dm_build_1414, dm_build_1415, int(dm_build_1418), dm_build_1416, dm_build_1417)
}

func (Dm_build_1420 *dm_build_1247) Dm_build_1419(dm_build_1421 []byte, dm_build_1422 int, dm_build_1423 string, dm_build_1424 *DmConnection) string {
	dm_build_1425 := Dm_build_1420.Dm_build_1372(dm_build_1421, dm_build_1422)
	dm_build_1422 += 2
	return Dm_build_1420.Dm_build_1405(dm_build_1421, dm_build_1422, int(dm_build_1425), dm_build_1423, dm_build_1424)
}

func (Dm_build_1427 *dm_build_1247) Dm_build_1426(dm_build_1428 byte) []byte {
	return []byte{dm_build_1428}
}

func (Dm_build_1430 *dm_build_1247) Dm_build_1429(dm_build_1431 int8) []byte {
	return []byte{byte(dm_build_1431)}
}

func (Dm_build_1433 *dm_build_1247) Dm_build_1432(dm_build_1434 int16) []byte {
	return []byte{byte(dm_build_1434), byte(dm_build_1434 >> 8)}
}

func (Dm_build_1436 *dm_build_1247) Dm_build_1435(dm_build_1437 int32) []byte {
	return []byte{byte(dm_build_1437), byte(dm_build_1437 >> 8), byte(dm_build_1437 >> 16), byte(dm_build_1437 >> 24)}
}

func (Dm_build_1439 *dm_build_1247) Dm_build_1438(dm_build_1440 int64) []byte {
	return []byte{byte(dm_build_1440), byte(dm_build_1440 >> 8), byte(dm_build_1440 >> 16), byte(dm_build_1440 >> 24), byte(dm_build_1440 >> 32),
		byte(dm_build_1440 >> 40), byte(dm_build_1440 >> 48), byte(dm_build_1440 >> 56)}
}

func (Dm_build_1442 *dm_build_1247) Dm_build_1441(dm_build_1443 float32) []byte {
	return Dm_build_1442.Dm_build_1453(math.Float32bits(dm_build_1443))
}

func (Dm_build_1445 *dm_build_1247) Dm_build_1444(dm_build_1446 float64) []byte {
	return Dm_build_1445.Dm_build_1456(math.Float64bits(dm_build_1446))
}

func (Dm_build_1448 *dm_build_1247) Dm_build_1447(dm_build_1449 uint8) []byte {
	return []byte{byte(dm_build_1449)}
}

func (Dm_build_1451 *dm_build_1247) Dm_build_1450(dm_build_1452 uint16) []byte {
	return []byte{byte(dm_build_1452), byte(dm_build_1452 >> 8)}
}

func (Dm_build_1454 *dm_build_1247) Dm_build_1453(dm_build_1455 uint32) []byte {
	return []byte{byte(dm_build_1455), byte(dm_build_1455 >> 8), byte(dm_build_1455 >> 16), byte(dm_build_1455 >> 24)}
}

func (Dm_build_1457 *dm_build_1247) Dm_build_1456(dm_build_1458 uint64) []byte {
	return []byte{byte(dm_build_1458), byte(dm_build_1458 >> 8), byte(dm_build_1458 >> 16), byte(dm_build_1458 >> 24), byte(dm_build_1458 >> 32), byte(dm_build_1458 >> 40), byte(dm_build_1458 >> 48), byte(dm_build_1458 >> 56)}
}

func (Dm_build_1460 *dm_build_1247) Dm_build_1459(dm_build_1461 []byte, dm_build_1462 string, dm_build_1463 *DmConnection) []byte {
	if dm_build_1462 == "UTF-8" {
		return dm_build_1461
	}

	if dm_build_1463 == nil {
		if e := dm_build_1506(dm_build_1462); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1461), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1463.encodeBuffer == nil {
		dm_build_1463.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1463.encode = dm_build_1506(dm_build_1463.getServerEncoding())
		dm_build_1463.transformReaderDst = make([]byte, 4096)
		dm_build_1463.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1463.encode; e != nil {

		dm_build_1463.encodeBuffer.Reset()

		n, err := dm_build_1463.encodeBuffer.ReadFrom(
			Dm_build_1520(bytes.NewReader(dm_build_1461), e.NewEncoder(), dm_build_1463.transformReaderDst, dm_build_1463.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1463.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1465 *dm_build_1247) Dm_build_1464(dm_build_1466 string, dm_build_1467 string, dm_build_1468 *DmConnection) []byte {
	return Dm_build_1465.Dm_build_1459([]byte(dm_build_1466), dm_build_1467, dm_build_1468)
}

func (Dm_build_1470 *dm_build_1247) Dm_build_1469(dm_build_1471 []byte) byte {
	return Dm_build_1470.Dm_build_1341(dm_build_1471, 0)
}

func (Dm_build_1473 *dm_build_1247) Dm_build_1472(dm_build_1474 []byte) int16 {
	return Dm_build_1473.Dm_build_1345(dm_build_1474, 0)
}

func (Dm_build_1476 *dm_build_1247) Dm_build_1475(dm_build_1477 []byte) int32 {
	return Dm_build_1476.Dm_build_1350(dm_build_1477, 0)
}

func (Dm_build_1479 *dm_build_1247) Dm_build_1478(dm_build_1480 []byte) int64 {
	return Dm_build_1479.Dm_build_1355(dm_build_1480, 0)
}

func (Dm_build_1482 *dm_build_1247) Dm_build_1481(dm_build_1483 []byte) float32 {
	return Dm_build_1482.Dm_build_1360(dm_build_1483, 0)
}

func (Dm_build_1485 *dm_build_1247) Dm_build_1484(dm_build_1486 []byte) float64 {
	return Dm_build_1485.Dm_build_1364(dm_build_1486, 0)
}

func (Dm_build_1488 *dm_build_1247) Dm_build_1487(dm_build_1489 []byte) uint8 {
	return Dm_build_1488.Dm_build_1368(dm_build_1489, 0)
}

func (Dm_build_1491 *dm_build_1247) Dm_build_1490(dm_build_1492 []byte) uint16 {
	return Dm_build_1491.Dm_build_1372(dm_build_1492, 0)
}

func (Dm_build_1494 *dm_build_1247) Dm_build_1493(dm_build_1495 []byte) uint32 {
	return Dm_build_1494.Dm_build_1377(dm_build_1495, 0)
}

func (Dm_build_1497 *dm_build_1247) Dm_build_1496(dm_build_1498 []byte, dm_build_1499 string, dm_build_1500 *DmConnection) []byte {
	if dm_build_1499 == "UTF-8" {
		return dm_build_1498
	}

	if dm_build_1500 == nil {
		if e := dm_build_1506(dm_build_1499); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1498), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1500.encodeBuffer == nil {
		dm_build_1500.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1500.encode = dm_build_1506(dm_build_1500.getServerEncoding())
		dm_build_1500.transformReaderDst = make([]byte, 4096)
		dm_build_1500.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1500.encode; e != nil {

		dm_build_1500.encodeBuffer.Reset()

		n, err := dm_build_1500.encodeBuffer.ReadFrom(
			Dm_build_1520(bytes.NewReader(dm_build_1498), e.NewDecoder(), dm_build_1500.transformReaderDst, dm_build_1500.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1500.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1502 *dm_build_1247) Dm_build_1501(dm_build_1503 []byte, dm_build_1504 string, dm_build_1505 *DmConnection) string {
	return string(Dm_build_1502.Dm_build_1496(dm_build_1503, dm_build_1504, dm_build_1505))
}

func dm_build_1506(dm_build_1507 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1507); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1508 struct {
	dm_build_1509 io.Reader
	dm_build_1510 transform.Transformer
	dm_build_1511 error

	dm_build_1512                []byte
	dm_build_1513, dm_build_1514 int

	dm_build_1515                []byte
	dm_build_1516, dm_build_1517 int

	dm_build_1518 bool
}

const dm_build_1519 = 4096

func Dm_build_1520(dm_build_1521 io.Reader, dm_build_1522 transform.Transformer, dm_build_1523 []byte, dm_build_1524 []byte) *Dm_build_1508 {
	dm_build_1522.Reset()
	return &Dm_build_1508{
		dm_build_1509: dm_build_1521,
		dm_build_1510: dm_build_1522,
		dm_build_1512: dm_build_1523,
		dm_build_1515: dm_build_1524,
	}
}

func (dm_build_1526 *Dm_build_1508) Read(dm_build_1527 []byte) (int, error) {
	dm_build_1528, dm_build_1529 := 0, error(nil)
	for {

		if dm_build_1526.dm_build_1513 != dm_build_1526.dm_build_1514 {
			dm_build_1528 = copy(dm_build_1527, dm_build_1526.dm_build_1512[dm_build_1526.dm_build_1513:dm_build_1526.dm_build_1514])
			dm_build_1526.dm_build_1513 += dm_build_1528
			if dm_build_1526.dm_build_1513 == dm_build_1526.dm_build_1514 && dm_build_1526.dm_build_1518 {
				return dm_build_1528, dm_build_1526.dm_build_1511
			}
			return dm_build_1528, nil
		} else if dm_build_1526.dm_build_1518 {
			return 0, dm_build_1526.dm_build_1511
		}

		if dm_build_1526.dm_build_1516 != dm_build_1526.dm_build_1517 || dm_build_1526.dm_build_1511 != nil {
			dm_build_1526.dm_build_1513 = 0
			dm_build_1526.dm_build_1514, dm_build_1528, dm_build_1529 = dm_build_1526.dm_build_1510.Transform(dm_build_1526.dm_build_1512, dm_build_1526.dm_build_1515[dm_build_1526.dm_build_1516:dm_build_1526.dm_build_1517], dm_build_1526.dm_build_1511 == io.EOF)
			dm_build_1526.dm_build_1516 += dm_build_1528

			switch {
			case dm_build_1529 == nil:
				if dm_build_1526.dm_build_1516 != dm_build_1526.dm_build_1517 {
					dm_build_1526.dm_build_1511 = nil
				}

				dm_build_1526.dm_build_1518 = dm_build_1526.dm_build_1511 != nil
				continue
			case dm_build_1529 == transform.ErrShortDst && (dm_build_1526.dm_build_1514 != 0 || dm_build_1528 != 0):

				continue
			case dm_build_1529 == transform.ErrShortSrc && dm_build_1526.dm_build_1517-dm_build_1526.dm_build_1516 != len(dm_build_1526.dm_build_1515) && dm_build_1526.dm_build_1511 == nil:

			default:
				dm_build_1526.dm_build_1518 = true

				if dm_build_1526.dm_build_1511 == nil || dm_build_1526.dm_build_1511 == io.EOF {
					dm_build_1526.dm_build_1511 = dm_build_1529
				}
				continue
			}
		}

		if dm_build_1526.dm_build_1516 != 0 {
			dm_build_1526.dm_build_1516, dm_build_1526.dm_build_1517 = 0, copy(dm_build_1526.dm_build_1515, dm_build_1526.dm_build_1515[dm_build_1526.dm_build_1516:dm_build_1526.dm_build_1517])
		}
		dm_build_1528, dm_build_1526.dm_build_1511 = dm_build_1526.dm_build_1509.Read(dm_build_1526.dm_build_1515[dm_build_1526.dm_build_1517:])
		dm_build_1526.dm_build_1517 += dm_build_1528
	}
}

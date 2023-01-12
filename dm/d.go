/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1530 struct {
	dm_build_1531 *list.List
	dm_build_1532 *dm_build_1584
	dm_build_1533 int
}

func Dm_build_1534() *Dm_build_1530 {
	return &Dm_build_1530{
		dm_build_1531: list.New(),
		dm_build_1533: 0,
	}
}

func (dm_build_1536 *Dm_build_1530) Dm_build_1535() int {
	return dm_build_1536.dm_build_1533
}

func (dm_build_1538 *Dm_build_1530) Dm_build_1537(dm_build_1539 *Dm_build_0, dm_build_1540 int) int {
	var dm_build_1541 = 0
	var dm_build_1542 = 0
	for dm_build_1541 < dm_build_1540 && dm_build_1538.dm_build_1532 != nil {
		dm_build_1542 = dm_build_1538.dm_build_1532.dm_build_1592(dm_build_1539, dm_build_1540-dm_build_1541)
		if dm_build_1538.dm_build_1532.dm_build_1587 == 0 {
			dm_build_1538.dm_build_1574()
		}
		dm_build_1541 += dm_build_1542
		dm_build_1538.dm_build_1533 -= dm_build_1542
	}
	return dm_build_1541
}

func (dm_build_1544 *Dm_build_1530) Dm_build_1543(dm_build_1545 []byte, dm_build_1546 int, dm_build_1547 int) int {
	var dm_build_1548 = 0
	var dm_build_1549 = 0
	for dm_build_1548 < dm_build_1547 && dm_build_1544.dm_build_1532 != nil {
		dm_build_1549 = dm_build_1544.dm_build_1532.dm_build_1596(dm_build_1545, dm_build_1546, dm_build_1547-dm_build_1548)
		if dm_build_1544.dm_build_1532.dm_build_1587 == 0 {
			dm_build_1544.dm_build_1574()
		}
		dm_build_1548 += dm_build_1549
		dm_build_1544.dm_build_1533 -= dm_build_1549
		dm_build_1546 += dm_build_1549
	}
	return dm_build_1548
}

func (dm_build_1551 *Dm_build_1530) Dm_build_1550(dm_build_1552 io.Writer, dm_build_1553 int) int {
	var dm_build_1554 = 0
	var dm_build_1555 = 0
	for dm_build_1554 < dm_build_1553 && dm_build_1551.dm_build_1532 != nil {
		dm_build_1555 = dm_build_1551.dm_build_1532.dm_build_1601(dm_build_1552, dm_build_1553-dm_build_1554)
		if dm_build_1551.dm_build_1532.dm_build_1587 == 0 {
			dm_build_1551.dm_build_1574()
		}
		dm_build_1554 += dm_build_1555
		dm_build_1551.dm_build_1533 -= dm_build_1555
	}
	return dm_build_1554
}

func (dm_build_1557 *Dm_build_1530) Dm_build_1556(dm_build_1558 []byte, dm_build_1559 int, dm_build_1560 int) {
	if dm_build_1560 == 0 {
		return
	}
	var dm_build_1561 = dm_build_1588(dm_build_1558, dm_build_1559, dm_build_1560)
	if dm_build_1557.dm_build_1532 == nil {
		dm_build_1557.dm_build_1532 = dm_build_1561
	} else {
		dm_build_1557.dm_build_1531.PushBack(dm_build_1561)
	}
	dm_build_1557.dm_build_1533 += dm_build_1560
}

func (dm_build_1563 *Dm_build_1530) dm_build_1562(dm_build_1564 int) byte {
	var dm_build_1565 = dm_build_1564
	var dm_build_1566 = dm_build_1563.dm_build_1532
	for dm_build_1565 > 0 && dm_build_1566 != nil {
		if dm_build_1566.dm_build_1587 == 0 {
			continue
		}
		if dm_build_1565 > dm_build_1566.dm_build_1587-1 {
			dm_build_1565 -= dm_build_1566.dm_build_1587
			dm_build_1566 = dm_build_1563.dm_build_1531.Front().Value.(*dm_build_1584)
		} else {
			break
		}
	}
	return dm_build_1566.dm_build_1605(dm_build_1565)
}
func (dm_build_1568 *Dm_build_1530) Dm_build_1567(dm_build_1569 *Dm_build_1530) {
	if dm_build_1569.dm_build_1533 == 0 {
		return
	}
	var dm_build_1570 = dm_build_1569.dm_build_1532
	for dm_build_1570 != nil {
		dm_build_1568.dm_build_1571(dm_build_1570)
		dm_build_1569.dm_build_1574()
		dm_build_1570 = dm_build_1569.dm_build_1532
	}
	dm_build_1569.dm_build_1533 = 0
}
func (dm_build_1572 *Dm_build_1530) dm_build_1571(dm_build_1573 *dm_build_1584) {
	if dm_build_1573.dm_build_1587 == 0 {
		return
	}
	if dm_build_1572.dm_build_1532 == nil {
		dm_build_1572.dm_build_1532 = dm_build_1573
	} else {
		dm_build_1572.dm_build_1531.PushBack(dm_build_1573)
	}
	dm_build_1572.dm_build_1533 += dm_build_1573.dm_build_1587
}

func (dm_build_1575 *Dm_build_1530) dm_build_1574() {
	var dm_build_1576 = dm_build_1575.dm_build_1531.Front()
	if dm_build_1576 == nil {
		dm_build_1575.dm_build_1532 = nil
	} else {
		dm_build_1575.dm_build_1532 = dm_build_1576.Value.(*dm_build_1584)
		dm_build_1575.dm_build_1531.Remove(dm_build_1576)
	}
}

func (dm_build_1578 *Dm_build_1530) Dm_build_1577() []byte {
	var dm_build_1579 = make([]byte, dm_build_1578.dm_build_1533)
	var dm_build_1580 = dm_build_1578.dm_build_1532
	var dm_build_1581 = 0
	var dm_build_1582 = len(dm_build_1579)
	var dm_build_1583 = 0
	for dm_build_1580 != nil {
		if dm_build_1580.dm_build_1587 > 0 {
			if dm_build_1582 > dm_build_1580.dm_build_1587 {
				dm_build_1583 = dm_build_1580.dm_build_1587
			} else {
				dm_build_1583 = dm_build_1582
			}
			copy(dm_build_1579[dm_build_1581:dm_build_1581+dm_build_1583], dm_build_1580.dm_build_1585[dm_build_1580.dm_build_1586:dm_build_1580.dm_build_1586+dm_build_1583])
			dm_build_1581 += dm_build_1583
			dm_build_1582 -= dm_build_1583
		}
		if dm_build_1578.dm_build_1531.Front() == nil {
			dm_build_1580 = nil
		} else {
			dm_build_1580 = dm_build_1578.dm_build_1531.Front().Value.(*dm_build_1584)
		}
	}
	return dm_build_1579
}

type dm_build_1584 struct {
	dm_build_1585 []byte
	dm_build_1586 int
	dm_build_1587 int
}

func dm_build_1588(dm_build_1589 []byte, dm_build_1590 int, dm_build_1591 int) *dm_build_1584 {
	return &dm_build_1584{
		dm_build_1589,
		dm_build_1590,
		dm_build_1591,
	}
}

func (dm_build_1593 *dm_build_1584) dm_build_1592(dm_build_1594 *Dm_build_0, dm_build_1595 int) int {
	if dm_build_1593.dm_build_1587 <= dm_build_1595 {
		dm_build_1595 = dm_build_1593.dm_build_1587
	}
	dm_build_1594.Dm_build_83(dm_build_1593.dm_build_1585[dm_build_1593.dm_build_1586 : dm_build_1593.dm_build_1586+dm_build_1595])
	dm_build_1593.dm_build_1586 += dm_build_1595
	dm_build_1593.dm_build_1587 -= dm_build_1595
	return dm_build_1595
}

func (dm_build_1597 *dm_build_1584) dm_build_1596(dm_build_1598 []byte, dm_build_1599 int, dm_build_1600 int) int {
	if dm_build_1597.dm_build_1587 <= dm_build_1600 {
		dm_build_1600 = dm_build_1597.dm_build_1587
	}
	copy(dm_build_1598[dm_build_1599:dm_build_1599+dm_build_1600], dm_build_1597.dm_build_1585[dm_build_1597.dm_build_1586:dm_build_1597.dm_build_1586+dm_build_1600])
	dm_build_1597.dm_build_1586 += dm_build_1600
	dm_build_1597.dm_build_1587 -= dm_build_1600
	return dm_build_1600
}

func (dm_build_1602 *dm_build_1584) dm_build_1601(dm_build_1603 io.Writer, dm_build_1604 int) int {
	if dm_build_1602.dm_build_1587 <= dm_build_1604 {
		dm_build_1604 = dm_build_1602.dm_build_1587
	}
	dm_build_1603.Write(dm_build_1602.dm_build_1585[dm_build_1602.dm_build_1586 : dm_build_1602.dm_build_1586+dm_build_1604])
	dm_build_1602.dm_build_1586 += dm_build_1604
	dm_build_1602.dm_build_1587 -= dm_build_1604
	return dm_build_1604
}
func (dm_build_1606 *dm_build_1584) dm_build_1605(dm_build_1607 int) byte {
	return dm_build_1606.dm_build_1585[dm_build_1606.dm_build_1586+dm_build_1607]
}

package mysql

// PageHelper
// 目的是做一些通用的限制, 如:
//	1. 参数校验
//	2. limit参数生成
type PageHelper struct {
	pageNo   int
	pageSize int
	offset   int
	total    int
}

const minOffset = 0
const defaultPageNo = 1
const defaultPageSize = 30

func NewPageHelper(po, ps int) *PageHelper {
	if po == 0 {
		po = defaultPageNo
	}
	if ps == 0 {
		ps = defaultPageSize
	}

	return &PageHelper{
		pageNo:   po,
		pageSize: ps,
	}
}

// SetOffset 设置额外起点
// 例如：共81条数据，获取 第3页（每页10条）时
//		offset = 0，则sql的limit为 20, 10
// 		offset = 5，则sql的limit为 25, 10
func (p *PageHelper) SetOffset(off int) {
	if off < minOffset {
		return
	}
	p.offset = off
}

// GetLimitParams 返回sql中limit x, y
func (p *PageHelper) GetLimitParams() (x int, y int) {
	if p.pageSize < 0 {
		return
	}
	x = (p.pageNo-1)*p.pageSize + p.offset
	y = p.pageSize
	return
}

// FromTo 返回数据位置 [x, y] 注意 左右都是闭区间
func (p *PageHelper) FromTo() (x, y int) {
	x = (p.pageNo-1)*p.pageSize + p.offset
	y = x + p.pageSize - 1
	return
}

func (p *PageHelper) HasNext() bool {
	x, _ := p.GetLimitParams()
	if x >= p.total {
		return false
	}

	return true
}

func (p *PageHelper) SetTotal(total int) {
	p.total = total
}

func (p *PageHelper) GetTotal() int {
	total := p.total - p.offset
	if total < 0 {
		total = 0
	}

	return total
}

package main

//边
type side struct {
	t int8
	v int8 //0-平面,1-凸,-1-凹
}

type block struct {
	sides []*side //有四个边
}

func check_end(blocks []*block, w, h int32) bool {
	for i, b := range blocks {
		nbs := get_near_blocks(blocks, int32(i), w, h)
		//左
		if nbs[0] != nil {
			if nbs[0].sides[2].v+b.sides[0].v != 0 {
				return false
			}
		} else {
			if b.sides[0].v != 0 {
				return false
			}
		}
		//上
		if nbs[1] != nil {
			if nbs[0].sides[3].v+b.sides[1].v != 0 {
				return false
			}
		} else {
			if b.sides[1].v != 0 {
				return false
			}
		}
		//右
		if nbs[2] != nil {
			if nbs[2].sides[0].v+b.sides[2].v != 0 {
				return false
			}
		} else {
			if b.sides[2].v != 0 {
				return false
			}
		}
		//下
		if nbs[3] != nil {
			if nbs[3].sides[1].v+b.sides[3].v != 0 {
				return false
			}
		} else {
			if b.sides[3].v != 0 {
				return false
			}
		}
	}
	return true
}
func get_near_blocks(blocks []*block, i, w, h int32) []*block {
	ret := make([]*block, 4, 4)
	x := get_x(i, w, h)
	y := get_x(i, w, h)
	x1 := x - 1
	if x1 >= 0 {
		ret[0] = blocks[get_index(x1, y, w)]
	}
	y1 := y - 1
	if y1 >= 0 {
		ret[1] = blocks[get_index(x, y1, w)]
	}

	x1 = x + 1
	if x1 < w {
		ret[2] = blocks[get_index(x1, y, w)]
	}

	y1 = y + 1
	if y1 < h {
		ret[3] = blocks[get_index(x, y1, w)]
	}
	return ret
}

////4,3 11
////
////

//
func get_x(i, w, h int32) int32 {
	return i % w
}
func get_y(i, w, h int32) int32 {
	return i / w
}
func get_index(x, y, w int32) int32 {
	return y*w + x
}

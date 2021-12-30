package main

const (
	LEFT  int8 = 0 //
	UP    int8 = 1 //
	RIGHT int8 = 2 //
	DOWN  int8 = 3 //
)

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
		if nbs[LEFT] != nil {
			if nbs[LEFT].sides[RIGHT].v+b.sides[LEFT].v != 0 {
				return false
			}
		} else {
			if b.sides[LEFT].v != 0 {
				return false
			}
		}
		//上
		if nbs[UP] != nil {
			if nbs[UP].sides[DOWN].v+b.sides[UP].v != 0 {
				return false
			}
		} else {
			if b.sides[UP].v != 0 {
				return false
			}
		}
		//右
		if nbs[RIGHT] != nil {
			if nbs[RIGHT].sides[LEFT].v+b.sides[RIGHT].v != 0 {
				return false
			}
		} else {
			if b.sides[RIGHT].v != 0 {
				return false
			}
		}
		//下
		if nbs[DOWN] != nil {
			if nbs[DOWN].sides[UP].v+b.sides[DOWN].v != 0 {
				return false
			}
		} else {
			if b.sides[DOWN].v != 0 {
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
		ret[LEFT] = blocks[get_index(x1, y, w)]
	}
	y1 := y - 1
	if y1 >= 0 {
		ret[UP] = blocks[get_index(x, y1, w)]
	}

	x1 = x + 1
	if x1 < w {
		ret[RIGHT] = blocks[get_index(x1, y, w)]
	}

	y1 = y + 1
	if y1 < h {
		ret[DOWN] = blocks[get_index(x, y1, w)]
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

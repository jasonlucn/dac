package dac

type SplayI interface {
	splay(v *BinNode) *BinNode
}

type Splay struct {
	BST
}

func (t *Splay) splay(v *BinNode) *BinNode {
	if v == nil {
		return nil
	}
	var p, g *BinNode
	for {
		p = v.parent
		g = p.parent
		if p == nil || g == nil {
			break
		}
		gg := g.parent
		if v.isLChild() {
			if p.isLChild() { //zig-zig
				g.attachAsLChild(p.rChild)
				
			} else { //zig -zag
			
			}
		} else if v.isRChild() {
			if p.isRChild() { //zag-zag
			
			} else { //zag-zig
					
			}
		}
		if gg == nil {
			v.parent = nil
		} else {
			if g == gg.lChild {
				gg.attachAsLChild(v)
			} else {
				gg.attachAsRChild(v)
			}
		}
		g.updateHeight()
		p.updateHeight()
		v.updateHeight()
	}
	if p == v.parent { //如果p是根，只需要再额外单旋（至多一次）
	}
	v.parent = nil
	return v
}

func (t *Splay) Search(e interface{}) *BinNode {
	panic("implement me")
}

func (t *Splay) Insert(e interface{}) *BinNode {
	panic("implement me")
}

func (t *Splay) Remove(e interface{}) bool {
	panic("implement me")
}

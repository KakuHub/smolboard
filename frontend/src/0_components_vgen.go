package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vjson"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "div", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	vgn.AddAttrList(c.MainAttrs())
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t"}
		vgparent.AppendChild(vgn)
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "main", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t\t"}
			vgparent.AppendChild(vgn)
			{
				var vgcomp vugu.Builder = c.Page
				if vgcomp != nil {
					vgin.BuildEnv.WireComponent(vgcomp)
					vgout.Components = append(vgout.Components, vgcomp)
					vgn = &vugu.VGNode{Component: vgcomp}
					vgparent.AppendChild(vgn)
				}
			}
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t"}
		vgparent.AppendChild(vgn)
		{
			var vgcomp vugu.Builder = c.Footer
			if vgcomp != nil {
				vgin.BuildEnv.WireComponent(vgcomp)
				vgout.Components = append(vgout.Components, vgcomp)
				vgn = &vugu.VGNode{Component: vgcomp}
				vgparent.AppendChild(vgn)
			}
		}
	}
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", Attr: []vugu.VGAttribute(nil)}
	{
		vgn.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\t.main-container {\n\t\tmin-height: 100vh;\n\t\tdisplay: flex;\n\t\tflex-direction: column;\n\t\tjustify-content: space-between;\n\n\t\t/* For loading. */\n\t\ttransition: 100ms linear opacity;\n\t\topacity: 1.0;\n\t}\n\n\t.main-container.loading {\n\t\tpointer-event: none;\n\t\tuser-select: none;\n\t\topacity: 0.4;\n\t}\n\n\tmain {\n\t\tflex: 1;\n\t\tdisplay: flex;\n\t\tflex-direction: column;\n\t\tjustify-content: stretch;\n\t}\n\n\tmain > * {\n\t\tflex: 1;\n\t}\n", Attr: []vugu.VGAttribute(nil)})
	}
	vgout.AppendCSS(vgn)
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ vjson.RawMessage
var _ js.Value

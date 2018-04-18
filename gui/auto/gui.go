package auto

import (
	"korok.io/korok/gfx"
	"korok.io/korok/gui"
)

//	Awesome GUI System
//
type LayoutType int

const (
	Vertical   LayoutType = iota
	Horizontal
	OverLay
)

type ViewType int

const (
	Normal ViewType = iota
	Popup
)

// Options: margin
func Margin(top, left, right, bottom float32) *Options {
	opt := &gLayoutMan.Options
	opt.SetMargin(top, left, right, bottom)
	return opt
}

// Options: gravity
func Gravity(x, y float32) *Options {
	opt := &gLayoutMan.Options
	opt.SetGravity(x, y)
	return opt
}

// Options: Size
func Size(w, h float32) *Options {
	opt := &gLayoutMan.Options
	opt.SetSize(w, h)
	return opt
}

// Widgets: Text
func Text(id gui.ID, text string, style *gui.TextStyle, p *Options) {
	if style == nil {
		style = &gContext.Theme.Text
	}
	gLayoutMan.Text(id, text, style, p)
	return
}

func TextSizeColored(id gui.ID, text string, color uint32, size float32, opt *Options) {
	sty := gContext.Theme.Text
	sty.Color = color
	sty.Size = size
	gLayoutMan.Text(id, text, &sty, opt)
}

// Widgets: InputEditor
func InputText(hint string, style *gui.InputStyle, p *Options) {

}

// Widget: Image
func Image(id gui.ID, tex gfx.Tex2D, style *gui.ImageStyle, p *Options) {
	gLayoutMan.Image(id, tex, style, p)
}

// Widget: Button
func Button(id gui.ID, text string, style *gui.ButtonStyle, p *Options) (event gui.EventType) {
	return gLayoutMan.Button(id, text, style, p)
}

func ImageButton(id gui.ID, normal, pressed gfx.Tex2D, style *gui.ImageButtonStyle, p *Options) gui.EventType{
	return gLayoutMan.ImageButton(id, normal, pressed, style, p)
}

func CheckBox(text string, style *gui.CheckBoxStyle) bool {
	return false
}

// Widget: ProgressBar, Slider
func ProgressBar(fraction float32, style *gui.ProgressBarStyle, p *Options) {

}

func Slider(id gui.ID, value *float32, style *gui.SliderStyle, p *Options) (v gui.EventType){
	return gLayoutMan.Slider(id, value, style, p)
}

// Widget: ListView TODO
func ListView() {

}

// Layout & Group

// Define a view group
func Define(name string, ) {
	gLayoutMan.DefineLayout(name, Normal)
}

func DefineType(name string, xt ViewType) {
	gLayoutMan.DefineLayout(name, xt)
}

func Clear(names ...string) {
	gLayoutMan.Clear(names...)
}

func Layout(id gui.ID, gui func(g *Group), w, h float32, xt LayoutType) {
	gLayoutMan.BeginLayout(id, xt)
	if w != 0 {
		gLayoutMan.current.SetSize(w, h)
	}
	gui(gLayoutMan.current.hGroup)
	gLayoutMan.EndLayout()
}

var gLayoutMan *LayoutMan
var gContext *gui.Context

func init() {
	gContext = gui.DefaultContext()
	gLayoutMan = &LayoutMan{Context: gContext}; gLayoutMan.initialize()
}

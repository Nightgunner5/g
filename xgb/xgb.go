/*
   Copyright 2012 the go.wde authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package xgb

import (
	"fmt"
	"github.com/Nightgunner5/xgb"
	"github.com/Nightgunner5/xgb/xproto"
	"github.com/Nightgunner5/xgbutil"
	"github.com/Nightgunner5/xgbutil/ewmh"
	"github.com/Nightgunner5/xgbutil/icccm"
	"github.com/Nightgunner5/xgbutil/keybind"
	"github.com/Nightgunner5/xgbutil/xgraphics"
	"github.com/Nightgunner5/xgbutil/xwindow"
	"github.com/skelterjohn/go.wde"
	"image"
	"sync"
)

func init() {
	wde.BackendNewWindow = NewWindow
	ch := make(chan struct{})
	wde.BackendRun = func() {
		<-ch
	}
	wde.BackendStop = func() {
		close(ch)
	}
}

const AllEventsMask = xproto.EventMaskKeyPress |
	xproto.EventMaskKeyRelease |
	xproto.EventMaskButtonPress |
	xproto.EventMaskButtonRelease |
	xproto.EventMaskEnterWindow |
	xproto.EventMaskLeaveWindow |
	xproto.EventMaskPointerMotion |
	xproto.EventMaskStructureNotify

type Window struct {
	win           *xwindow.Window
	xu            *xgbutil.XUtil
	conn          *xgb.Conn
	buffer        *xgraphics.Image
	bufferLck     *sync.Mutex
	width, height int
	closed        bool

	events chan interface{}
}

func NewWindow(width, height int) (wde.Window, error) {
	var err error

	w := new(Window)
	w.width, w.height = width, height

	w.xu, err = xgbutil.NewConn()
	if err != nil {
		return nil, err
	}

	w.conn = w.xu.Conn()
	screen := w.xu.Screen()

	w.win, err = xwindow.Generate(w.xu)
	if err != nil {
		return nil, err
	}

	err = w.win.CreateChecked(screen.Root, 600, 500, width, height, 0)
	if err != nil {
		return nil, err
	}

	w.win.Listen(AllEventsMask)

	err = icccm.WmProtocolsSet(w.xu, w.win.Id, []string{"WM_DELETE_WINDOW"})
	if err != nil {
		fmt.Println(err)
		err = nil
	}

	w.bufferLck = &sync.Mutex{}
	w.buffer = xgraphics.New(w.xu, image.Rect(0, 0, width, height))
	w.buffer.XSurfaceSet(w.win.Id)

	keyMap, modMap := keybind.MapsGet(w.xu)
	keybind.KeyMapSet(w.xu, keyMap)
	keybind.ModMapSet(w.xu, modMap)

	w.events = make(chan interface{})

	w.SetIcon(Gordon)
	w.SetIconName("Go")

	go w.handleEvents()

	return w, nil
}

func (w *Window) SetTitle(title string) {
	if w.closed {
		return
	}
	err := ewmh.WmNameSet(w.xu, w.win.Id, title)
	if err != nil {
		// TODO: log
	}
	return
}

func (w *Window) SetSize(width, height int) {
	if w.closed {
		return
	}

	w.win.Resize(width, height)
	w.width, w.height = width, height
	return
}

func (w *Window) Size() (width, height int) {
	if w.closed {
		return
	}
	width, height = w.width, w.height
	return
}

func (w *Window) LockSize(lock bool) {

}

func (w *Window) Show() {
	if w.closed {
		return
	}
	w.win.Map()
}

func (w *Window) Screen() (im wde.Image) {
	if w.closed {
		return
	}
	im = &Image{w.buffer}
	return
}

func (w *Window) FlushImage(bounds ...image.Rectangle) {
	if w.closed {
		return
	}
	if w.buffer.Pixmap == 0 {
		w.bufferLck.Lock()
		if err := w.buffer.XSurfaceSet(w.win.Id); err != nil {
			fmt.Println(err)
		}
		w.bufferLck.Unlock()
	}
	w.buffer.XDraw()
	w.buffer.XPaint(w.win.Id)
}

func (w *Window) Close() (err error) {
	if w.closed {
		return
	}
	w.win.Destroy()
	w.closed = true
	return
}

type Image struct {
	*xgraphics.Image
}

func (buffer Image) CopyRGBA(src *image.RGBA, r image.Rectangle) {
	// clip r against each image's bounds and move sp accordingly (see draw.clip())
	sp := image.ZP
	orig := r.Min
	r = r.Intersect(buffer.Bounds())
	r = r.Intersect(src.Bounds().Add(orig.Sub(sp)))
	dx := r.Min.X - orig.X
	dy := r.Min.Y - orig.Y
	(sp).X += dx
	(sp).Y += dy

	i0 := (r.Min.X - buffer.Rect.Min.X) * 4
	i1 := (r.Max.X - buffer.Rect.Min.X) * 4
	si0 := (sp.X - src.Rect.Min.X) * 4
	yMax := r.Max.Y - buffer.Rect.Min.Y

	y := r.Min.Y - buffer.Rect.Min.Y
	sy := sp.Y - src.Rect.Min.Y
	for ; y != yMax; y, sy = y+1, sy+1 {
		dpix := buffer.Pix[y*buffer.Stride:]
		spix := src.Pix[sy*src.Stride:]

		for i, si := i0, si0; i < i1; i, si = i+4, si+4 {
			dpix[i+0] = spix[si+2]
			dpix[i+1] = spix[si+1]
			dpix[i+2] = spix[si+0]
			dpix[i+3] = spix[si+3]
		}
	}
}

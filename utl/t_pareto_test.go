// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utl

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	"github.com/lei006/gomath/chk"
	"github.com/lei006/gomath/io"
)

func Test_pareto01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("pareto01. compare vectors: Pareto-optimal")

	u := []float64{1, 2, 3, 4, 5, 6}
	v := []float64{1, 2, 3, 4, 5, 6}
	io.Pforan("u = %v\n", u)
	io.Pfblue2("v = %v\n", v)
	uDominates, vDominates := ParetoMin(u, v)
	io.Pfpink("u_dominates = %v\n", uDominates)
	io.Pfpink("v_dominates = %v\n", vDominates)
	if uDominates {
		tst.Errorf("test failed\n")
		return
	}
	if vDominates {
		tst.Errorf("test failed\n")
		return
	}

	v = []float64{1, 1.8, 3, 4, 5, 6}
	io.Pforan("\nu = %v\n", u)
	io.Pfblue2("v = %v\n", v)
	uDominates, vDominates = ParetoMin(u, v)
	io.Pfpink("u_dominates = %v\n", uDominates)
	io.Pfpink("v_dominates = %v\n", vDominates)
	if uDominates {
		tst.Errorf("test failed\n")
		return
	}
	if !vDominates {
		tst.Errorf("test failed\n")
		return
	}

	v = []float64{1, 2.1, 3, 4, 5, 6}
	io.Pforan("\nu = %v\n", u)
	io.Pfblue2("v = %v\n", v)
	uDominates, vDominates = ParetoMin(u, v)
	io.Pfpink("u_dominates = %v\n", uDominates)
	io.Pfpink("v_dominates = %v\n", vDominates)
	if !uDominates {
		tst.Errorf("test failed\n")
		return
	}
	if vDominates {
		tst.Errorf("test failed\n")
		return
	}
}

func Test_pareto02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("pareto02. probabilistic Pareto-optimal")

	rand.Seed(time.Now().UnixNano())

	Φ := []float64{0, 0.25, 0.5, 0.75, 1}
	u := 0.4
	v := 0.6
	for _, φ := range Φ {
		p := ProbContestSmall(u, v, φ)
		io.Pf("u=%v v=%v p(u,v,%5.2f) = %.8f\n", u, v, φ, p)
	}

	ntrials := 1000
	doplot := chk.Verbose
	var buf bytes.Buffer
	var Zu []float64

	U := []float64{1, 2, 3, 4, 5, 6}
	V := []float64{1, 2, 3, 4, 5, 6}
	zu, _ := paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{1, 2, 3, 4, 5, 6}
	V = []float64{2, 2, 3, 4, 5, 6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{1, 2, 3, 4, 5, 6}
	V = []float64{2, 3, 3, 4, 5, 6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{1, 2, 3, 4, 5, 6}
	V = []float64{2, 3, 4, 4, 5, 6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{1, 2, 3, 4, 5, 6}
	V = []float64{2, 3, 4, 5, 5, 6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{1, 2, 3, 4, 5, 6}
	V = []float64{2, 3, 4, 5, 6, 6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{1, 2, 3, 4, 5, 6}
	V = []float64{2, 3, 4, 5, 6, 7}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	if doplot {
		plotParetoTest(&buf, "test_pareto02", Φ, Zu, false, false)
	}
}

func Test_pareto03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("pareto03. probabilistic Pareto-optimal")

	rand.Seed(time.Now().UnixNano())

	Φ := []float64{0, 0.25, 0.5, 0.75, 1}
	u := 0.4
	v := 0.6
	for _, φ := range Φ {
		p := ProbContestSmall(u, v, φ)
		io.Pf("u=%v v=%v p(u,v,%5.2f) = %.8f\n", u, v, φ, p)
	}

	ntrials := 1000
	doplot := chk.Verbose
	var buf bytes.Buffer
	var Zu []float64

	U := []float64{-1, -2, -3, -4, -5, -6}
	V := []float64{-1, -2, -3, -4, -5, -6}
	zu, _ := paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{-2, -2, -3, -4, -5, -6}
	V = []float64{-1, -2, -3, -4, -5, -6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{-2, -3, -3, -4, -5, -6}
	V = []float64{-1, -2, -3, -4, -5, -6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{-2, -3, -4, -4, -5, -6}
	V = []float64{-1, -2, -3, -4, -5, -6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{-2, -3, -4, -5, -5, -6}
	V = []float64{-1, -2, -3, -4, -5, -6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{-2, -3, -4, -5, -6, -6}
	V = []float64{-1, -2, -3, -4, -5, -6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	U = []float64{-2, -3, -4, -5, -6, -7}
	V = []float64{-1, -2, -3, -4, -5, -6}
	zu, _ = paretoTest(U, V, Φ, ntrials)
	Zu = append(Zu, zu...)

	if doplot {
		plotParetoTest(&buf, "test_pareto03", Φ, Zu, false, true)
	}
}

func paretoTest(U, V, Φ []float64, ntrials int) (zu, zv []float64) {
	io.Pforan("\nu = %v\n", U)
	io.Pforan("v = %v\n", V)
	io.PfWhite("%5s%13s%13s\n", "φ", "u wins", "v wins")
	zu = make([]float64, len(Φ))
	zv = make([]float64, len(Φ))
	for i, φ := range Φ {
		uWins := 0
		vWins := 0
		for j := 0; j < ntrials; j++ {
			uDominates := ParetoMinProb(U, V, φ)
			if uDominates {
				uWins++
			} else {
				vWins++
			}
		}
		zu[i] = 100 * float64(uWins) / float64(ntrials)
		zv[i] = 100 * float64(vWins) / float64(ntrials)
		io.Pf("%5.2f%12.2f%%%12.2f%%\n", φ, zu[i], zv[i])
	}
	return
}

func writePythonArray(buf *bytes.Buffer, name string, x []float64) {
	io.Ff(buf, "%s=np.array([", name)
	for i := 0; i < len(x); i++ {
		io.Ff(buf, "%g,", x[i])
	}
	io.Ff(buf, "])\n")
}

func writePythonMatrix(buf *bytes.Buffer, name string, xy [][]float64) {
	io.Ff(buf, "%s=np.array([", name)
	for i := 0; i < len(xy); i++ {
		io.Ff(buf, "[")
		for j := 0; j < len(xy[i]); j++ {
			io.Ff(buf, "%g,", xy[i][j])
		}
		io.Ff(buf, "],\n")
	}
	io.Ff(buf, "])\n")
}

func plotParetoTest(buf *bytes.Buffer, fnkey string, Φ, Zu []float64, show, negative bool) {
	io.Ff(buf, "from gosl import SetForEps, Save\n")
	io.Ff(buf, "from mpl_toolkits.mplot3d import Axes3D\n")
	io.Ff(buf, "import matplotlib.pyplot as plt\n")
	io.Ff(buf, "import numpy as np\n")
	io.Ff(buf, "fig = plt.figure()\n")
	io.Ff(buf, "ax = fig.add_subplot(111, projection='3d')\n")
	writePythonArray(buf, "phi", Φ)
	writePythonArray(buf, "dz", Zu)
	io.Ff(buf, "SetForEps(0.75, 455, mplclose=0, text_usetex=0)\n")
	io.Ff(buf, "n=len(phi)\nx,y=np.meshgrid(phi,np.linspace(0,1,7))\nx=x.flatten()\ny=y.flatten()\nz=np.zeros(n*7)\n")
	io.Ff(buf, "ax.bar3d(x,y,z,dx=0.1*np.ones(n*7),dy=0.1*np.ones(n*7),dz=dz, color='#cee9ff')\n")
	io.Ff(buf, "ax.set_xlabel('$\\phi$')\n")
	io.Ff(buf, "ax.set_zlabel('u-wins [%%]')\n")
	io.Ff(buf, "ax.set_xticks([0,0.25,0.5,0.75,1])\n")
	io.Ff(buf, "ax.set_xticklabels(['0.0','0.25','0.5','0.75','1.0'])\n")
	if negative {
		io.Ff(buf, "ax.set_yticklabels(['u=[-1 -2 -3 -4 -5 -6]', 'u=[-2 -2 -3 -4 -5 -6]', 'u=[-2 -3 -3 -4 -5 -6]', 'u=[-2 -3 -4 -4 -5 -6]', 'u=[-2 -3 -4 -5 -5 -6]', 'u=[-2 -3 -4 -5 -6 -6]', 'u=[-2 -3 -4 -5 -6 -7]'], rotation=-15,verticalalignment='baseline',horizontalalignment='left')\n")
	} else {
		io.Ff(buf, "ax.set_yticklabels(['v=[1 2 3 4 5 6]', 'v=[2 2 3 4 5 6]', 'v=[2 3 3 4 5 6]', 'v=[2 3 4 4 5 6]', 'v=[2 3 4 5 5 6]', 'v=[2 3 4 5 6 6]', 'v=[2 3 4 5 6 7]'], rotation=-15,verticalalignment='baseline',horizontalalignment='left')\n")
	}
	io.Ff(buf, "import matplotlib.patheffects as path_effects\n")
	io.Ff(buf, "for i, xval in enumerate(x): ax.text(xval,y[i],dz[i],'%%.2f'%%dz[i],color='#bf0000',fontsize=10, path_effects=[path_effects.withSimplePatchShadow(offset=(1,-1),shadow_rgbFace='white')])\n")
	if show {
		io.Ff(buf, "plt.show()\n")
	} else {
		io.Ff(buf, "plt.savefig('/tmp/gosl/%s.eps')\n", fnkey)
		io.Ff(buf, "print 'file </tmp/gosl/%s.eps> witten'\n", fnkey)
	}
	io.WriteFileVD("/tmp/gosl", io.Sf("%s.py", fnkey), buf)
}

func Test_pareto04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("pareto04. Pareto front")

	ovs := [][]float64{
		{1.0, 8.0}, {1.8, 7.0}, {2.2, 6.5}, {2.5, 6.0}, {2.8, 5.5}, {3.0, 6.0},
		{3.1, 7.3}, {3.0, 6.0}, {3.5, 4.5}, {3.9, 6.8}, {4.0, 5.5}, {4.1, 3.8},
		{4.3, 6.7}, {4.4, 3.6}, {4.5, 6.0}, {5.0, 3.0}, {5.1, 5.0}, {5.5, 2.8},
		{6.0, 2.5}, {6.5, 7.4}, {6.5, 6.5}, {7.0, 4.0}, {7.0, 2.0}, {8.0, 1.8},
	}

	front := ParetoFront(ovs)
	chk.Ints(tst, "front", front, []int{0, 1, 2, 3, 4, 8, 11, 13, 15, 17, 18, 22, 23})

	if chk.Verbose {

		n := len(front)
		Xp := make([]float64, n)
		Yp := make([]float64, n)
		for k, i := range front {
			Xp[k] = ovs[i][0]
			Yp[k] = ovs[i][1]
		}

		var buf bytes.Buffer
		io.Ff(&buf, "from gosl import SetForEps, Save, Gll\n")
		io.Ff(&buf, "import matplotlib.pyplot as plt\n")
		io.Ff(&buf, "import numpy as np\n")
		io.Ff(&buf, "SetForEps(0.75, 355)\n")
		writePythonMatrix(&buf, "XY", ovs)
		writePythonArray(&buf, "Xp", Xp)
		writePythonArray(&buf, "Yp", Yp)
		io.Ff(&buf, "plt.plot(XY[:,0],XY[:,1],'r.',clip_on=0)\n")
		io.Ff(&buf, "plt.plot(Xp,Yp,'ko',markerfacecolor='none',ms=7,clip_on=0)\n")
		io.Ff(&buf, "Gll('x','y','')\n")
		io.Ff(&buf, "Save('/tmp/gosl/test_pareto04.eps')\n")
		io.Ff(&buf, "print 'file </tmp/gosl/test_pareto04.eps> witten'\n")
		io.WriteFileVD("/tmp/gosl", "test_pareto04.py", &buf)
	}
}

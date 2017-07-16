package golangmobilecrashreproduce

import "github.com/davecgh/go-spew/spew"

func Crasher(ci CrashInterface) {
	spew.Dump(ci)
	ci.FirstFunc()
	spew.Dump("ci.FirstFunc()")
	ci.SecondFunc()
	spew.Dump("ci.SecondFunc()")
	ci.CallMeIfYouCanFunc()
	spew.Dump("ci.CallMeIfYouCanFunc()")
	ci.SeeYouAtAnotherWorld()
	spew.Dump("ci.SeeYouAtAnotherWorld()")

}

type CrashInterface interface {
	FirstFunc()
	SecondFunc()
	CallMeIfYouCanFunc()
	SeeYouAtAnotherWorld()
}

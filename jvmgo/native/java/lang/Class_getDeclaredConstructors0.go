package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
const _constructorConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B)V"

func init() {
	_class(getDeclaredConstructors0, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;")
}

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructors0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jClass := vars.GetRef(0) // this
	publicOnly := vars.GetBoolean(1)

	goClass := jClass.Extra().(*rtc.Class)
	goConstructors := goClass.GetConstructors(publicOnly)
	constructorCount := uint(len(goConstructors))

	constructorClass := rtc.BootLoader().LoadClass("java/lang/reflect/Constructor")
	constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
	constructorArr := constructorClass.NewArray(constructorCount)
	stack := frame.OperandStack()
	stack.PushRef(constructorArr)

	if constructorCount > 0 {
		constructorObjs := constructorArr.Refs()
		thread := frame.Thread()
		for i, goConstructor := range goConstructors {
			constructorObj := constructorClass.NewObjWithExtra(goConstructor)
			constructorObjs[i] = constructorObj

			// call <init>
			newFrame := thread.NewFrame(constructorInitMethod)
			vars := newFrame.LocalVars()
			vars.SetRef(0, constructorObj)                                   // this
			vars.SetRef(1, jClass)                                           // declaringClass
			vars.SetRef(2, getParameterTypeArr(goConstructor))               // parameterTypes
			vars.SetRef(3, getExceptionTypeArr(goConstructor))               // checkedExceptions
			vars.SetInt(4, int32(goConstructor.GetAccessFlags()))            // modifiers
			vars.SetInt(5, int32(0))                                         // todo slot
			vars.SetRef(6, getSignature(&goConstructor.ClassMember))         // signature
			vars.SetRef(7, getAnnotationByteArr(&goConstructor.ClassMember)) // annotations
			vars.SetRef(8, getParameterAnnotationDyteArr(goConstructor))     // parameterAnnotations
			thread.PushFrame(newFrame)
		}
	}
}

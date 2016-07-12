/*
This is comments of main package
*/
package main

import "fmt"

import (
	"funcPackage"
)

//import fmt "fmt" // Package implementing formatted I/O.
//import mat "math" //倒入包，并以'mat'的名称使用这个包

//倒入包方式2
//import math "math"

func main() {

	fmt.Println("The code of Leanr begain.")

	//defineVar.DefineVar()
	//defineVar.DefineVar2()
	//defineVar.DefineVar()

	//baseType.BaseType()

	//typeOperation.TypeConvert()
	//typeOperation.TypeGuess()

	//constPackage.ConstFun()

	//loopPackage.Forloop()
	//loopPackage.Forloop2()

	//condition.IfCondition(1)
	//fmt.Println(condition.IfExpress(2, 4, 10000))
	//condition.SwitchFunc()
	//condition.SwitchFunc2()

	//deferPackage.DeferFun()
	//deferPackage.DeferFunc2()

	//pointPackage.PointFun()

	//structPackage.StructFunc()
	//structPackage.StructPoint()
	//structPackage.StructGrammar()

	//slicePackage.ArrayFunc()
	//slicePackage.SliceFunc()
	//slicePackage.SliceFuncCut()
	//slicePackage.Slicenil()
	//slicePackage.AppendSlice()

	//rangePackage.RangAppendSlice()
	//rangePackage.RangAppendSlice2()

	//mapPackage.MapInit()
	//mapPackage.LoopKeyValue()
	//var value = mapPackage.GetItem("Google")
	//fmt.Println(value)

	//funcPackage.FuncDefine()
	//funcPackage.FuncDefine2()
	//funcPackage.Closure()
	funcPackage.Closure2()

	fmt.Println("The code of Leanr end.")
}

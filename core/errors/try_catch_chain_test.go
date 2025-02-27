package errors

import (
	"errors"
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	// 正常执行
	Try(func() {
		fmt.Println("ok")
	}).Do()

	// try发生异常，走catch
	var errFoo = errors.New("")
	Try(func() {
		panic(errFoo)
	}).Catch(errors.New("bar"), func(err error) {
		fmt.Println("bar")
	}).Catch(errFoo, func(err error) {
		fmt.Println("foo")
	}).Do()

	// try发生异常，走默认catch
	Try(func() {
		panic(errors.New("test"))
	}).Catch(errors.New("bar"), func(err error) {
		fmt.Println("bar")
	}).Catch(errFoo, func(err error) {
		fmt.Println("foo")
	}).DefaultCatch(func(err error) {
		fmt.Println("other")
	}).Do()

	// try未发生异常走else
	Try(func() {
		_ = 100 + 19
	}).DefaultCatch(func(err error) {
		fmt.Println("other")
	}).Else(func() {
		fmt.Println("else")
	}).Do()

	// try发生异常，并且走finally
	Try(func() {
		panic(errors.New("test"))
	}).DefaultCatch(func(err error) {
		fmt.Println("other")
	}).Else(func() {
		fmt.Println("else")
	}).Finally(func() {
		fmt.Println("finally")
	}).Do()

	// try未发生异常，并且走finally
	Try(func() {
		_ = 100 + 19
	}).DefaultCatch(func(err error) {
		fmt.Println("other")
	}).Finally(func() {
		fmt.Println("finally")
	}).Do()

	// 发生panic，尝试捕获错误，但是没有捕获得到，则异常会被向上抛出，即仍然会panic
	Try(func() {
		panic(errors.New("test"))
	}).Catch(errFoo, func(err error) {
		fmt.Println("catch success")
	}).Finally(func() {
		fmt.Println("not catch finally")
	}).Do()

}

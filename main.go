package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func main() {
	e := zero.New()
	e.OnPrefix("ee").Handle(func(ctx *zero.Ctx) {

	})
}

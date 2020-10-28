package server

import (
	"JuneGoBlog/src"
	"testing"
)

func init() {
	src.InitSetting("../../setting.json")
}

func TestGetAbstract(t *testing.T) {
	r := getAbstract(`# 震惊，醋和酱油竟然有如此妙用！

酱油和醋混合是怎么回事呢？酱油和醋相信大家都很熟悉，但是酱油和醋混合是怎么回事呢，下面就让小编带大家一起了解吧。

<!-- more -->

`)
	if r != "酱油和醋混合是怎么回事呢？酱油和醋相信大家都很熟悉，但是酱油和醋混合是怎么回事呢，下面就让小编带大家一起了解吧。" {
		t.Error("Fail")
	}
}

func TestGetAbstract2(t *testing.T) {
	r := getAbstract(`Mutexes do no scale. Atomic loads do.

<!-- more -->

## atomic

atomic 包中提供许多基本数据类型的原子操作，主要可以分为下面几类：
`)
	if r != "Mutexes do no scale. Atomic loads do." {
		t.Error("摘要长度小于设定最大长度时，测试不通过")
	}
}

func TestGetAbstract3(t *testing.T) {
	r := getAbstract(`#test
1234567890123456789012345678901234567890你一定想过。 我们多数情况下都会通过想象来思考问题。如果我问你一个关于 “1 到 100 的数字” 的问题，你脑子里就会下意识的出现一系列画面。例如，我会把它想象成一条从我开始的直线，从数字 1 到 20 然后右转 90 度一直到 1000+。我记得我很小的时候，在我们的幼儿园里，衣帽间里有很多数字，写在墙上，数字 20 恰好在拐角处。你可能有你自己的关于数字的画面。另一个常见的例子是一年四季的视觉展现 —— 有人将之想象成一个盒子，有人将之想象成一个圈。
无论如何，我想用 Go 和 WebGL 把我对于常见的并发模式的具象化尝试展现给大家。这多多少少代表了我对于并发程序的理解。如果能听到我和大家脑海中的画面有什么不同，一定会非常有趣。 我特别想知道 Rob Pike 或者 Sameer Ajmani 脑子里是怎么描绘并发图像的。我打赌我会很感兴趣的。
那么，让我们从一个很基础的 “Hello,Concurrent World” 例子开始我们今天的主题。
Hello, Concurrent world
代码很简单 —— 单个通道，单个 goroutine，单次写入，单次读取。
先进先出顺序很明显了，是吧？我们可以创建一百万个 goroutine，因为它们很轻量，但是对于实现我们的目的来说没有必要。我们来想想其他可以玩的。 例如，常见的消息传递模式。
Fan-In
并发世界中流行的模式之一是所谓的 fan-in 模式。这与 fan-out 模式相反，稍后我们将介绍。简而言之，fan-in 是一项功能，可以从多个输入中读取数据并将其全部多路复用到单个通道中。
`)
	if len(r) != src.Setting.AbstractLen || r[src.Setting.AbstractLen-3:] != "..." {
		t.Error("没有指定摘要且文章长度足够时，测试未通过")
	}
}

func TestGetAbstract4(t *testing.T) {
	r := getAbstract(`#test
1234567890123456789012345678901234567890你一定想过。 我们多数情况下都会通过想象来思考问题。如果我问你一个关于 “1 到 100 的数字” 的问题，你脑子里就会下意识的出现一系列画面。例如，我会把它想象成一条从我开始的直线，从数字 1 到 20 然后右转 90 度一直到 1000+。我记得我很小的时候，在我们的幼儿园里，衣帽间里有很多数字，写在墙上，数字 20 恰好在拐角处。你可能有你自己的关于数字的画面。另一个常见的例子是一年四季的视觉展现 —— 有人将之想象成一个盒子，有人将之想象成一个圈。
无论如何，我想用 Go 和 WebGL 把我对于常见的并发模式的具象化尝试展现给大家。这多多少少代表了我对于并发程序的理解。如果能听到我和大家脑海中的画面有什么不同，一定会非常有趣。 我特别想知道 Rob Pike 或者 Sameer Ajmani 脑子里是怎么描绘并发图像的。我打赌我会很感兴趣的。
那么，让我们从一个很基础的 “Hello,Concurrent World” 例子开始我们今天的主题。
Hello, Concurrent world
代码很简单 —— 单个通道，单个 goroutine，单次写入，单次读取。
先进先出顺序很明显了，是吧？我们可以创建一百万个 goroutine，因为它们很轻量，但是对于实现我们的目的来说没有必要。我们来想想其他可以玩的。 例如，常见的消息传递模式。
Fan-In
并发世界中流行的模式之一是所谓的 fan-in 模式。这与 fan-out 模式相反，稍后我们将介绍。简而言之，fan-in 是一项功能，可以从多个输入中读取数据并将其全部多路复用到单个通道中。
<!-- more -->
vvv
`)
	if len(r) != src.Setting.AbstractLen || r[src.Setting.AbstractLen-3:] != "..." {
		t.Error("指定摘要且摘要长度足够时，测试未通过")
	}
}

func TestGetAbstract5(t *testing.T) {
	r := getAbstract("")
	if len(r) != 0 {
		t.Error("没有指定摘要，文章长度不足，未通过")
	}
}

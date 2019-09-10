package assert

// 忘了从什么地方看来的，不是自己原生实现的
func Assert(condition bool, reason string) {
    if !condition {
        panic(reason)
    }
}

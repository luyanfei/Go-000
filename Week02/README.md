# 作业题目
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
## 作业内容
与数据库打交道的代码是与库的协作代码，按照PPT讲义上的套路，应该使用errors.Wrap来保存堆栈信息，在程序的顶部再根据根错误的类型来做相应处理。
```
func dao() (error) {
    return errors.Wrap(sql.ErrNoRows, "no rows in dao")
}
func service() {
    if errors.Is(err, sql.ErrNoRows) {
        //process no rows error
    }
}
```


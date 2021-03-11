- 边界拷贝slice/maps： 作为函数传入，存储了对他们的引用，不会影响传入的；可对被赋值的进行修改~
- defer释放资源: ns级别
- channel size: 0/1
- 枚举从1开始
- **time.Time**:表示在一个时间点上加上 24 小时并不总是产生一个新的日历日。 因此，在处理时间时始终使用 "time" 包        
  1. Programs using times should typically store and pass them as values,
   not pointers. That is, time variables and struct fields should be of
   type time.Time, not *time.Time
- time.Time：瞬时时间 time.Duration：时间段 、对外部系统使用 time.Time 和 time.Duration 、"time" 包不支持解析闰秒时间戳（8728），也不在计算中考虑闰秒（15190）。
- 错误类型
- 错误包装 (Error Wrapping)

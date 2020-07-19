## java程序启动方式

```bash
java [-options] class [args] 
java [-options] -jar jarfile [args] 
javaw [-options] class [args] 
javaw [-options] -jar jarfile [args]

java 命令行，选项 主类名(jar名) main函数参数

```
javaw 比较适合启动GUI程序，其他的和java是一样的

### options参数

- 标准选项
- 非标准选项

标准选项基本不会轻易变动，非标准选项以-X开头，而非标准选项中的一部分高级选项以-XX开头
```bash
-version  查看版本号
-cp(classpath)  指定classpath
-Xms  指定堆的初始大小
-Xmx  指定最大堆大小
-Xss  指定线程栈大小

```





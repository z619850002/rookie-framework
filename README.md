rookie-framework简介
---------------

rookie-framework 是一个golang开发的websocket框架，当然经过改造也可以作为http服务器使用。具体性能还在测试中。

模块机制
-------------------
* 模块
  模块一般来说每个包内是一个单独的模块，每个模块有一个自己运行的主线程，每次收到其他模块的请求，就会开辟一个携程，之后在携程内处理请求并返回。

* 骨架
  骨架开发运用了装饰器，需要监控的代码段可以由逻辑代码块包裹，之后利用装饰器可以对于代码块的输入输出进行一些想要的操作。
  
* 通讯
  通讯利用channel仿RPC通讯，经过封装已经变成了同步方法，目前还没有提供异步机制。
  
* 响应
  响应时间由模块中的handler决定，这个由用户自定义实现，推荐作为私有内容，这样每个包可以将响应事件封装起来。
  
* 消息
  模块与模块之间传递的消息有专门的buf类，每个buf需要用组合的方式实现获取协议号的方法，这样传输到一个模块之后，由协议号判断具体的类型进行类型转换，再由module执行对应的处理操作即可。

 性能测试
-------------------
* 在gate/gate_test.go中，测试从接收到网络数据开始，到往send channel中送数据为止这个过程中的性能，中间模拟处理过程包括：模块间调用（同步方法），redis存储，Benchmark测试结果大致在(3000-4000)ns/op：
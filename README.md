# ProcessUrl
lab that deal with big data in a small resource machine


`git clone` 后，在src的main下 `go build ./` 即可编译成功

### 目录结构：

```
├── data
│   ├── pygen.py
│   ├── pygen_sorted_result.txt
│   ├── reduce_task_0
│   ├── testurl.txt
│   └── topn.txt
├── readme.md
└── src
    ├── heapsort
    │   ├── heap_sort.go
    │   └── test_test.go
    ├── main
    │   └── main.go
    └── pucommon
        └── common.go
```

`data`目录为数据目录，**将需要处理的url文件放入data下**。pygen.py 是测试数据的生成脚本，`reduce_task_0`为所有url出现的次数结果，`topn.txt`为最后要去topn的结果。

`src` src 为源码文件，`headsort` 为堆排package，提供堆排策略。 `pucommon`主要是common的数据结构。

`main` 为主程序

程序需要三个参数:
    1.url存放文件名

2.需要topN  // (本次需求 top100)

3.map任务的个数 //（本次需求100GB，在1GB内存上执行，最后将map任务设置为100以上）

e.g. `./main testurl.txt 100 110`


### 实现思路：

优于内存远远小于数据，首先将大文件进行hash打散，按照hash后的值分成一个个小文件。再使用reduce进行统计处理。
同时使用小堆排，筛选出top出现次数的url

### 思考:
1. 量化内存的使用

   首先进行读数据，每一个mapF是一个协程，协程初始分配栈8K左右，在使用超过是，调用`morestack()`进行扩充[1]。不过本次实验协程只做消费写入文件的处理。所以应该8K就足够了。调用300个协程应该不到30M的使用。
   主要是reduce的使用的内存会多一些。首先reduceF要维护一个hashmap，go的map使用类似于拉链法，并设置有装载因子阈值。可能会出现旧桶与新桶并存的情况，每次map的扩充是2的指数进行增长[2]。由于只有一个reduceF，那么程序只需要维持一个hashmap就可以。所以reduceF每次处理数据理论上不能高于1GB。由于程序是一次处理一个mapF生成的数据，所以每个mapF的生成数据不能超过1GB。又因为需要维护hashmap这样的数据结构,再考虑到golang的runtime对每次reduceF是否能够及时gc,可能mapF需要300个左右甚至更多，也就是说每一次reduce处理333M数据。同时程序还需要维护一个堆来记录topN，这个小堆是我自己实现的，整个数据结构对内存的开销并不算太大。堆需要维护一个N个数据的slice在内存中，每一份数据为string，int，假设平均大小30Byte。那么维护N个数据的slice大小30N。本次实验中，取top100，即3K左右。

2. 考虑一下有什么场景是当前算法满足不了的，针对这种场景有什么优化方案？

   这里可能会出现mapF结果不均匀的情况，有两种不均匀。一种是因为某个url出现次数本身就很多，不过这个不会影响reduceF对内存的使用，因为出现多少条都只占string+int的大小。还要一种情况就是hash出的不同url在某个号段上比较集中。这种情况我认为是hash算法本身分配不均导致的（即这个hash算法不适合用来做为url的hash）。解决的方法有三种，一个是重新找一个能够更均匀hash的算法。第二就是提升mapF的任务个数。第三就是将比较大的mapF结果每项加盐二次hash，如果还不满足再次加盐（当然，每一次盐都不一样），直到得到reduceF能够处理的大小后结束。
   
[1]https://juejin.im/post/5d9ff459f265da5b8a5160f5
[2]https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/

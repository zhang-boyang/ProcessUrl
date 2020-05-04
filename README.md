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

e.g. `./main testurl.txt testurl.txt 100 110`


### 实现思路：

优于内存远远小于数据，首先将大文件进行打散hash打散，按照hash后的值分成一个个小文件。在使用reduce进行统计处理。
同时使用小堆排，筛选出top出现次数的url
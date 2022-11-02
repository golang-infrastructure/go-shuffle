# 各种洗牌算法的Go语言实现

# 一、支持的洗牌算法

洗牌算法的定义：为有限集合生成随机排序的算法。

目前支持的洗牌算法： 

- Fisher–Yates-Knuth
- Scatology

# 二、Fisher–Yates-Knuth洗牌算法

假设现在有一个数组：

```
[1, 2, 3, 4, 5]
```

从最右边的坐标`len(slice)-1`开始作为`right_index`，每次从`[0, right_index]`随机选择一个下标，将选中的下标的值与`right_index`交换，并将`right_index`减一往左偏移。



# 三、Scatology算法

就是在Fisher–Yates-Knuth的基础上随机选择的时候不再选择最右边的，但是感觉这样子似乎可能会有概率问题？


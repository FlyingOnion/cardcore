## 锄大地
### 牌类型
* ILLEGAL	非法牌型
* SINGLE	单张
* PAIR		对
* TRIPLE	三张
* STRAIGHT	顺
* FLUSH		同花
* SKELETON	骷髅
* KK		金刚
* STRFLUSH	同花顺

<pre>
type Card struct {
	Color int
	Text  string
}

type cddCard struct {
	Card
}
</pre>

`func (c cddCard) LessThan(another cddCard) bool`  
见 `CompareWith`

`func (c cddCard) CompareWith(another cddCard) (result int)`  
比较牌`c`与`another`的大小，`c`和`another`相比大/小/相等时分别返回1, 0, -1

`func (c cddCard) CompareTextWith(another cddCard) int`  
只比较`c`与`another`的牌号大小，不算花色

`func (c cddCard) String() string`  
输出<花色>-<牌号>，如牌红桃3输出`Heart-3`

<pre>
// 牌组类型
type cddCardGroup struct {
	Cards []cddCard
}
</pre>

`func (cg cddCardGroup) Len() int`  
`func (cg cddCardGroup) Less(i, j int) bool`  
`func (cg cddCardGroup) Swap(i, j int)`  
实现`sort.Interface`，将牌组按牌大小排序时用

`func (cg cddCardGroup) validate() (cgType int, err error)`  
判断牌组`cg`是否为一个合法牌组，返回`cgType`牌型及错误信息  
先检查每张牌是否合法，再检查牌数，判断是否为合法单张、对、骷髅及金刚

`func (cg cddCardGroup) LessThan(another cddCardGroup) (bool, error)`  
判断牌组`cg`与`another`的大小情况  
分别调用`validate`检查合法性，然后检查牌数是否一致，一致则可比，否则不可比  
单张、对和三张时直接比较文本返回  
五张牌时比较`cgType`，牌型大的较大，牌型一样分（顺、同花顺）（同花）和（骷髅、金刚）分别比较

`func (cg cddCardGroup) isPair() bool`  
`func (sortedCG cddCardGroup) isTriple() bool`  
`func (sortedCG cddCardGroup) isQuadruple() bool`  
判断牌组（或其中某些牌）是否为2、3、4张牌号一样的，标为`sortedCG`的在调用前需进行排序

`func (cg cddCardGroup) isStraightOrFlush() (bool, bool)`  
`func (sortedCG cddCardGroup) isStraight() bool`  
`func (cg cddCardGroup) isFlush() bool`  
`func (sortedCG cddCardGroup) isSkeleton() bool`  
`func (sortedCG cddCardGroup) isKK() bool`  
判断（5张的）牌组是否为那几个合法牌型

`func (cg cddCardGroup) Text() string`  
返回只包括牌号的`string`

`func (cg cddCardGroup) String() string`  
返回包括花色和牌号的`string`

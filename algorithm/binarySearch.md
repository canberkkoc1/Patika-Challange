# Binary Search

[7, 5, 1, 8, 3, 6, 0, 9, 4, 2] dizisinin Binary-Search-Tree aşamalarını yazınız.

```
1<-5<-6->7->8->9
    
0<-.->3

2<-.->4
```


>6'nın sağındakiler sırasıyla 7,8,9 diye ilerler. Çünkü her bakışımızda seçilen sayı diğerinden büyüktür ve sağ kısma yazarak karşılaştırma yapmaya devam ederiz.

>6'nın solundakiler 5,1 - 1, 0 ve 3 - 3, 2 ve 4 >şeklinde dallanır.
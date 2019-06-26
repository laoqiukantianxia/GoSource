package main
/*
使用WaitGroup值时，模式为：
先统一Add，再并发Done，最后Wait

如果在调用该值的wait方法时，为了增大其计数器的值，
而并发的调用该值的Add方法，很可能会引发panic

*/

func 




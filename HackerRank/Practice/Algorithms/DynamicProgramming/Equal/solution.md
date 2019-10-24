# Equal

First thing to notice is that increasing the chocolate count in all
the registers but a specific one by K is equal to substracting K from this
particular register. Thus, the challenge is reduced to the problem of performing
the least number of _operations_ on the registers so that afterwards
all the registers contain the same value C. An _operation_ consists of either:  
  * subtracting 5 from a chosen register, or  
  * subtracting 2 from a chosen register, or  
  * subtracting 1 from a chosen register.  

The cost of every _operation_ is 1, so the challenge focuses on minimizing
the total count of these operations.

The best way to do this is to find a min & max value contained in the array.

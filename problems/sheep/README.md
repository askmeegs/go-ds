### counting sheep -  analysis


Proving that the only "insomnia" number is zero....
- For any nonzero number...
- You'll cover digit `0` when you get to `10 * N` where `i=10`
- For digits `1-9`... proof by contradiction:
      - Let `P` be the smallest power of 10 greater than N. ie. if `N=3692`, `P=10000` or `10^4`
      - As we multiple N by successive i-values on the way to crossing the `9P` threshold.. or 90000
      - The leftmost digit of N*i *must* take on every `1-9` digit eventually  

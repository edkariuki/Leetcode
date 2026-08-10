func isPalindrome(x int) bool {
    if x<0 || (x%10==0 && x!=0) {
      return false
    }
     
     reversed_half_x:=0

     for x > reversed_half_x {
        reversed_half_x = (reversed_half_x * 10) + (x%10)
        x = x / 10
     }

    return x == reversed_half_x || x == reversed_half_x/10
    
}
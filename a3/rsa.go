package main

import (
   "fmt"
   "big"
   )

var three = big.NewInt(3)
var two = big.NewInt(2)
var zero = big.NewInt(0)

func main() {
   n := new( big.Int )
   n.SetString( "59862884897410580445294975930364810502194059460046784100782961089451413815651385722867179113058394170198337928137605175707021438787385809631292643776317258709958790233150162992821973384914482382094263768626624728641916686410126516762051754901449384230749023005608317630448024255295503829583920557675021174883", 10 )
   fmt.Println( n )

   factorize( n )
}

func factorize( n *big.Int )  {

   factor := new( big.Int )
   j := new( big.Int )
   nTwo := new( big.Int )
   nThree := new( big.Int )

   start := sqrtAprox( n )
   if (j.Rem( start, two )).Cmp( zero ) == 0 {
      start.Add( start, nThree.Neg(three))
   }
   fmt.Println( start )

   i := start

   for {
      if i.Cmp( three ) != 1 {
         break;
      }
      j.Rem( n, i )
      if j.Cmp( zero ) == 0 {
         fmt.Println( "factor" )
         factor = i
         break;
      }
      i.Add( i , nTwo.Neg(two) )
     // fmt.Println( i )
   }

   fmt.Println( factor )
}

func sqrtAprox( n *big.Int ) *big.Int {
   ans := new( big.Int )
   mid := new( big.Int )

   mid.Div( n, two )
   ans.Div( n, mid )


   for mid.Cmp( zero ) != 0 {
      if ans.Cmp( mid ) == -1 {
         mid.Div( mid, two )
         ans.Div( n, mid )
      } else { break }
   }
   return ans
}

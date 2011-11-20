package main

import (
   "fmt"
   "big"
   )

func main() {
   N := new( big.Int )
   N.SetString( "59862884897410580445294975930364810502194059460046784100782961089451413815651385722867179113058394170198337928137605175707021438787385809631292643776317258709958790233150162992821973384914482382094263768626624728641916686410126516762051754901449384230749023005608317630448024255295503829583920557675021174883", 10 )
   //fmt.Println( N )
   factorize( N )
}

func factorize( n *big.Int )  {
   var three = big.NewInt(3)
   var two = big.NewInt(2)
   var zero = big.NewInt(0)

   factor := new( big.Int )
   j := new( big.Int )
   nTwo := new( big.Int )
   nThree := new( big.Int )
   start := new( big.Int )

   start.Div( n, two )
   if (j.Rem( start, two )).Cmp( zero ) == 0 {
      start.Add( start, nThree.Neg(three))
   }
   fmt.Println( start )
   fmt.Println( n )

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
   }

   fmt.Println( factor )
}

func sqrt( n *big.Int ) *big.Int {
   return n 
}

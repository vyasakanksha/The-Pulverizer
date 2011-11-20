package main

import (
   "fmt"
   "big"
   )

func main() {
   N := new( big.Int )
   N.SetString( "59862884897410580445294975930364810502194059460046784100782961089451413815651385722867179113058394170198337928137605175707021438787385809631292643776317258709958790233150162992821973384914482382094263768626624728641916686410126516762051754901449384230749023005608317630448024255295503829583920557675021174883", 10 )
   fmt.Println( N )
   factorize( N )
}

func factorize( n *big.Int )  {
   factor := new( big.Int )
   var end = big.NewInt(3)
   start := sqrt( n )
   fmt.Println( start )

   i := start

   for i.Cmp( end ) == 1 {
      if i.Cmp( n ) == 0 {
         factor = i
         break;
      }
   }

   fmt.Println( "la" )
   fmt.Println( factor )
}

func sqrt( n *big.Int ) *big.Int {
   return n
}

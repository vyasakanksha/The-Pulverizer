package coolFunctions

func ModInverse( bA *big.Int, bN *big.Int ) *big.Int {
   bQ := new( big.Int )
   bR := new( big.Int )

   temp1 := new( big.Int )
   temp2 := new( big.Int )
   temp3 := new( big.Int )
   temp4 := new( big.Int )

   var bX1 = big.NewInt( 1 )
   var bX2 = big.NewInt( 0 )
   var bY1 = big.NewInt( 0 )
   var bY2 = big.NewInt( 1 )

   bQ.Div( bA, bN )
   bR.Rem( bA, bN )

   for ( bR.Cmp( zero ) != 0 ) {
      bQ.Div( bA, bN )
      bR.Rem( bA, bN )
      temp1.Set( bX1 )
      temp2.Set( bX2 )
      bX1.Set( bX2 )
      bY1.Set( bY2 )
      temp3.Mul( bX2, bQ )
      temp4.Mul( bY2, bQ )
      bX2.Sub( temp1, temp3 )
      bY2.Sub( temp2, temp4 )
   }

   for bY2.Cmp(zero) < 0 {
      bY2.Add( bY2, bN )
   }

   return bY2
}


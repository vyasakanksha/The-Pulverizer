package main

import(
   "fmt"
   "os"
   "big"
   "flag"
   "time"
   "rand"
   "strconv"
   )

var e = flag.Bool("e", false, "encrypt text")
var d = flag.Bool("d", false, "decrypt text")

var one = big.NewInt(1)
var two = big.NewInt(2)
var three = big.NewInt(3)

func main( ) {
   var buf []byte
   var ciphertext []big.Int
   //var toWrite string

   flag.Parse()
   if len(os.Args) < 3 {
      fmt.Fprintf( os.Stderr, "RSA: Not enough arguments \n" )
   } else if st, err := os.Stat( os.Args[2] ); err != nil {
      fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
   } else if f, err := os.Open( os.Args[2] ); err != nil {
      fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
   } else {
      buf = make( []byte, st.Size )
      f.Read( buf )
   }

   // toWrite := os.Args[3]

   if( *e ) {
      plaintext := []int(string(buf))
      temp := os.Args[3]
      plength, err := strconv.Atoi(temp); if err != nil {
         fmt.Fprintf( os.Stderr, "%s\n", err )
      }

      fmt.Printf("plength: %d", plength)
      fmt.Println( string( plaintext ))

     // keys
      p       := new( big.Int )
      q       := new( big.Int )
      n       := new( big.Int )
      e       := new( big.Int )
      d       := new( big.Int )
      pMinus1 := new( big.Int )
      qMinus1 := new( big.Int )
      phi     := new( big.Int )

      privateKeyGenerator( p, plength )
      privateKeyGenerator( q, plength )

      if p == q {
         privateKeyGenerator( q, plength )
      }

      fmt.Println( "Private keys" )
      fmt.Println( p )
      fmt.Println( q )

      n.Mul( p, q )

      pMinus1.Sub( p , one )
      qMinus1.Sub( q , one )
      phi.Mul( pMinus1, qMinus1 )

      e = three
      d.ModInverse( e, phi )
      ciphertext = encrypt( plaintext, n, e, len( plaintext ))

   } else if( *d ) {
      ciphertext = make( []big.Int, len( buf ) )
      for i := 0; i < len( buf ); i++ {
         ciphertext[i].SetString( string( buf[i] ), 10)
      }
   } else {
      fmt.Fprintf( os.Stderr, "RSA: No flag set\n" )
      os.Exit(0)
   }

   for i := 0; i < len( ciphertext ); i++ {
      fmt.Println( &ciphertext[i] )
   }
}

func privateKeyGenerator( prime *big.Int, plength int ) {
   var tempB = new( big.Int )
   var mlength = new( big.Int )
   isPrime := false
   plengthB := big.NewInt( int64( plength ) )
   mlength.Exp( two, plengthB, nil )
   source := rand.NewSource( time.Nanoseconds() )
   r := rand.New( source )

   fmt.Println(plengthB)

   for( !isPrime ) {
      tempB.Rand( r, mlength )
      prime.SetBit( tempB, plength, 1)
      isPrime = big.ProbablyPrime( prime, 5 )
   }
}


func encrypt( plaintext []int, n, e *big.Int, textLen int ) []big.Int {
   ciphertext := make( []big.Int, textLen )
   for i := 0; i < textLen; i++ {
      plainB := big.NewInt( int64( plaintext[i] ))
      ciphertext[i].Exp( plainB, e, n )
      fmt.Println( &ciphertext[i] )
   }
   return ciphertext
}

func decrypt( ciphertext []*big.Int, n, d *big.Int, textLen int ) {
   decryptB := make( []big.Int, textLen )
   for i := 0; i < textLen; i++ {
      decryptB[i].Exp( ciphertext[i], d, n )
   }
}


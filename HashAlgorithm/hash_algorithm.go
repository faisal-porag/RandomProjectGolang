package HashAlgorithm

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func GenerateHashAlgorithm() {
	fmt.Println("----------------Small Message----------------")
	message := []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))

	fmt.Println("----------------Large Message----------------")
	message = []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model. This model gives you an experience in which you never navigate to particular pages or even reload a page.  It loads and unloads the various views of our app into the same page itself. If you've ever run popular web apps like Gmail, Facebook, Instagram, or Twitter, you've used a single-page app. In all those apps, the content gets dynamically displayed without requiring you to refresh or navigate to a different page. React gives you a powerful subjective model to work with and supports you to build user interfaces in a declarative and component-driven way.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))
}

/* OUTPUT:
----------------Small Message----------------

Md5: d53c7002ebeeaa872f02efdda82f76f0

Sha1: b56fd6c3eabc1ea6e880ffb7762d31a4b39bbcc9

Sha256: 0e81731d26a2ffd898be00c40bf0ab3c24ec82a0837311701193dd861efdb944

Sha512: 97be88309470c98e8ff840a74567c3948263251f3ac6f232a6e228cebc6daa6d95569a96
6a5f2ba5d694dc644d02c144849e7a181a159b583a2060e2ce3ad375


----------------Large Message----------------

Md5: f52c27fb5fa09afbb717e5ded9ef5707

Sha1: 86b1102f10518ef02089742862695e132ca6b5fe

Sha256: 88c9eca72b7b635458945710d12f6709af44dd02c0a3bb9b16c7e352b6d13f84

Sha512: 5400f78f1642d5b7df7a0c60e22e95ba04a94d16d33d008fc78a5002958f853ecc469bb8
786632c0be47a0c5b44d5f78aa9878abe7f00b688d8712c921aa81b8
*/

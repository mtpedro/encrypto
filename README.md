# encrypto

This is a very simple library for encrypting and decrypting data with a priavte, 32 byte, key.

# example: 
```
package main
import (
	"github.com/mtpedro/encrypto"
	"fmt"
)

func main() {
	data := []byte("this is some data that will be encrypted");
	priv_key := []byte("qwertyuiophsdbuvheurghvnusajdnua");

	encrypted_data, err := encrypto.Encrypt(priv_key, &data); // encrypt
	if err != nil { fmt.Println(err); return };

	fmt.Println(encrypted_data);
	
	decrypted_data, err := encrypto.Decrypt(priv_key, encrypted_data); // decrypt

	if err != nil { fmt.Println(err); return };
	
	fmt.Println(decrypted_data);
}
```

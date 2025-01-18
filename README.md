# Rango: The Smallest Ransomware
It encrypts the txt files that exist in **Documents** folder.\
Don't believe? Download and unzip the moneygo-setup.zip file. Run the install.exe file and see the magic.

---

## Brief about Rango

### Encrption Process
1. [MoneyGo](https://moneygo.netlify.app) Website automatically downloads the rango zip file in victim's device when the victim visits it.
2. When a victim runs the `install.exe` file from zip file, rango starts the file encrption program in victim's device.
3. It is using AES symmetric encryption to encrypt and decrypt the files.
4. Symmetric key for AES is generated randomly using crypto library and only valid for single session.
5. Symmetric key lives only in-memory, and flushed once the encryption completes.
6. Symmetric key is encrypted using RSA asymmetric encryption and shown to the victim once the encryption completes.
7. Program asks the victim to send the email to attacker with the encrypted symmetric key if he/she wants to decrypt the encrypted files.

### Decryption Process
1. Attacker provides the decrypted aes symmetric key to the victim through replying the victim email.
2. Victim needs to run `unistall.exe` from the downloaded zip file.
3. It will start the program on command prompt and ask to provide the aes key.
4. Once the victim enters the aes key, the program will decrypt all the encrypted txt files from Documents folder and close itself.


## Resources
1. https://dev.to/paulwababu/how-to-make-ransomware-with-python-windows-mac-and-linux-142g
2. https://github.com/paulwababu/ransomware-golang
3. https://github.com/crvvdev/cpp-ransomware
4. https://www.tarlogic.com/blog/ransomware-developed-golang/
5. https://medium.com/@andrei_rat/how-to-run-go-on-android-5e8fbd90ddc6
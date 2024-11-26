from enigma import Enigma

rotor1 = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
rotor2 = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
rotor3 = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
reflector = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
plugboard = {"A": "B", "B": "A", "C": "D", "D": "C"}  

enigma = Enigma([rotor1, rotor2, rotor3], reflector, plugboard)
enigma.set_position([0, 0, 0]) 

message = "HELLOENIGMA"
encrypted_message = enigma.encrypt_message(message)
print("Encrypted Message:", encrypted_message)

enigma.set_position([0, 0, 0]) 

decrypted_message = enigma.encrypt_message(encrypted_message)
print("Decrypted Message:", decrypted_message)

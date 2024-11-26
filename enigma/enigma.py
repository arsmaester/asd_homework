class Enigma:
    def __init__(self, rotors, reflector, plugboard):
        self.rotors = rotors
        self.reflector = reflector
        self.plugboard = plugboard
        self.position = [0, 0, 0] 

    def set_position(self, position):
        self.position = position

    def rotate_rotors(self):
        self.position[0] = (self.position[0] + 1) % 26
        if self.position[0] == 0:
            self.position[1] = (self.position[1] + 1) % 26
            if self.position[1] == 0:
                self.position[2] = (self.position[2] + 1) % 26

    def plugboard_swap(self, char):
        return self.plugboard.get(char, char)

    def rotor_forward(self, rotor, pos, char):
        index = (ord(char) - ord('A') + pos) % 26
        return rotor[index]

    def rotor_backward(self, rotor, pos, char):
        index = (rotor.index(char) - pos + 26) % 26
        return chr(index + ord('A'))

    def encrypt_char(self, char):
        if char not in 'ABCDEFGHIJKLMNOPQRSTUVWXYZ':
            return char

        char = self.plugboard_swap(char)

        for i in range(3):
            char = self.rotor_forward(self.rotors[i], self.position[i], char)

        char = self.reflector[ord(char) - ord('A')]

        for i in reversed(range(3)):
            char = self.rotor_backward(self.rotors[i], self.position[i], char)

        char = self.plugboard_swap(char)

        self.rotate_rotors()

        return char

    def encrypt_message(self, message):
        encrypted_message = ""
        for char in message:
            encrypted_message += self.encrypt_char(char)
        return encrypted_message
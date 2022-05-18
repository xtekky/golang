chars = """"""

string = '"'
for char in chars:
    string += str(char.encode())
string += '"'
print(string.replace('b', "").replace("'", ""))

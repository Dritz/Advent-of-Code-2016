import hashlib

code = ""
code2 = ["_"] * 8
validIndexes = ["0", "1", "2", "3", "4", "5", "6", "7"]
doorID = "ojvtpuvg"
#doorID = "wtnhxymk"
i = 0

while len(code) < 8:
  curHash = hashlib.md5((doorID + str(i)).encode('utf-8')).hexdigest()
  if curHash.startswith("00000"):
    code += curHash[5]
    print(code)
  i += 1

i = 0

while "_" in code2:
  curHash = hashlib.md5((doorID + str(i)).encode('utf-8')).hexdigest()
  if curHash.startswith("00000") and curHash[5] in validIndexes:
    if code2[int(curHash[5])] == "_":
      code2[int(curHash[5])] = curHash[6]
      print("".join(code2))
  i += 1

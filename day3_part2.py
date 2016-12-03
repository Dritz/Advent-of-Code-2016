def isTriangle(triList):
  triList = [int(element) for element in triList]
  triList.sort()
  return triList[0] + triList[1] > triList[2]

file = open('day3_input.txt')
count = 0
i = 0
lineBuffer = []
for line in file:
  lineBuffer.append(line.split())
  if i == 2:
    for c in range(0, 3):
      if isTriangle([lineBuffer[0][c], lineBuffer[1][c], lineBuffer[2][c]]):
        count += 1
    i = 0
    lineBuffer = []
  else:
    i += 1

print(count)

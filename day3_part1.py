def isTriangle(triList):
  triList = [int(element) for element in triList]
  triList.sort()
  return triList[0] + triList[1] > triList[2]

file = open('day3_input.txt')
count = 0

for line in file:
  sides = line.split()
  if isTriangle(sides):
    count += 1

print(count)

# Двудольные графы. Раскраски графов
Проверить является ли граф двудольным.

## Формат входных данных:
В первой строке указаны два числа разделенных пробелом: `v` (число вершин) и `e` (число ребер). 
В следующих e строках указаны пары вершин, соединенных ребром. Выполняются ограничения: `10000 ≤ v ≤ 1000, 0 ≤ e ≤ 1000`.

## Формат выходных данных:
Одно слово: `YES`, если граф двудолен, или `NO`, в противном случае.

Sample Input 1:
```
4 2
1 2
3 2
```
Sample Output 1:
```
YES
```
Sample Input 2:
```
3 3
1 2
2 3
3 1
```
Sample Output 2:
```
NO
```
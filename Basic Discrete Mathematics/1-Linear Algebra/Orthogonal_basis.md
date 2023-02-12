# Ортогональный базис

Напишите программу, которая находит наилучшее решение системы линейных алгебраических уравнений методом наименьших квадратов.

## Формат входных данных:

В первой строке задаются два числа: количество уравнений `n` и количество неизвестных `m`. Количество уравнений не меньше количества неизвестных. 
Далее идут `n` строк, каждая из которых содержит `m + 1` число. Первые `m` чисел - это коэффициенты `i-го` уравнения системы, а последнее, `(m + 1) - е` число - коэффициент `b_i`, стоящий в правой части `i-го` уравнения.
## Формат выходных данных:
В качестве результата следует вывести решение системы в виде m чисел, разделенных пробелом.

## Sample Input:
```
4 2
4 2 8
5 2 4
2 6 2
3 0 8
```
## Sample Output:
```
1.6531165311653115 -0.30894308943089427
```

```python
rom math import gcd


def readline():
    return map(int, input().split())


def solve_linear_system(xy):
    # Gaussian elimination
    n = len(xy)
    m = len(xy[0]) - 1
    for (row_idx, row) in enumerate(xy):
        col_idx, a = next((i, v) for (i, v) in enumerate(row) if v)
        for (i, other_row) in enumerate(xy):
            if i == row_idx:
                continue
            b = other_row[col_idx]
            if not b:
                continue
            d = gcd(a, b)
            a1, b1 = a // d, b // d
            for j in range(m + 1):
                other_row[j] = other_row[j] * a1 - row[j] * b1
    solution = [None] * m
    for row in xy:
        # assert there is exactly one non-zero value
        (col_idx, value), = ((i, v) for (i, v) in enumerate(row[:-1]) if v)
        solution[col_idx] = row[-1] / value
    assert all(v is not None for v in solution)
    return solution


def test_solve_linear_system():
    assert solve_linear_system([
        [1, 0, 0, 5],
        [0, 1, 0, 7],
        [0, 0, 1, 9],
    ]) == [5, 7, 9]
    assert solve_linear_system([
        [3, 2, -1, -5],
        [6, -1, 5, 19],
        [5, 3, 0, -4],
    ]) == [1, -3, 2]


def least_squares(xy):
    n = len(xy)
    m = len(xy[0]) - 1
    assert all(len(row) == m + 1 for row in xy), xy
    ab = [
        [sum(xy[k][i] * xy[k][j] for k in range(n)) for j in range(m + 1)]
        for i in range(m)
    ]
    assert len(ab) == m
    assert all(len(row) == m + 1 for row in ab)
    return solve_linear_system(ab)


def main():
    n, m = readline()
    xy = [list(readline()) for __ in range(n)]
    solution = least_squares(xy)
    print(*solution)


if __name__ == '__main__':
    main()
```

```
#include <iostream>
#include <sstream>
#include <iomanip>

int main()
{
	int n, m;
	std::cin >> n >> m;

	long double** a = new long double* [n];
	long double* b = new long double[n];
	for (int i = 0; i < n; i++)
	{
		a[i] = new long double[m];
	}
	//Заполнить матрицу
	for (int i = 0; i < n; i++)
	{
		for (int j = 0; j < m; j++)
		{
			std::cin >> a[i][j];
		}
		std::cin >> b[i];
	}

	long double** a_new = new long double* [n];
	long double* b_new = new long double[n] {0};
	for (int i = 0; i < n; i++)
	{
		a_new[i] = new long double[n] {0};
	}

	for (int j = 0; j < m; j++) {
		for (int k = 0; k < m; k++) {
			for (int i = 0; i < n; i++) {
				a_new[j][k] += a[i][k] * a[i][j];
			}
		}
		for (int i = 0; i < n; i++) {
			b_new[j] += b[i] * a[i][j];
		}
	}

	delete[] b;
	for (int i = 0; i < n; i++) {
		delete[] a[i];
	}

	for (int i = 0; i < n; i++)
	{
		a[i] = new long double[m + 1]{ 0 };
	}

	for (int i = 0; i < m; i++) {
		for (int j = 0; j < m; j++) {
			a[i][j] = a_new[i][j];
		}
		a[i][m] = b_new[i];
	}
	delete[] a_new;
	delete[] b_new;
    
	n = m;

    for (int i = 0; i < n; i++)
    {
        int max = i;
        for (int k = i; k < n; k++) {
            if (abs(a[max][i]) < abs(a[k][i])) {
                max = k;
            }
        }
        long double* tempArr = a[max];
        a[max] = a[i];
        a[i] = tempArr;

        if (abs(a[i][i]) > 0.00000000001) {
            long double temp = a[i][i];
            for (int j = 0; j < m + 1; j++)
            {
                a[i][j] /= temp;
            }
            for (int k = i + 1; k < n; k++)
            {
                long double temp = a[k][i];
                for (int j = i; j < m + 1; j++)
                {
                    a[k][j] -= a[i][j] * temp;;
                }
            }
        }
    }

    bool* correct = new bool[n] { false };
    bool* Dcorrect = new bool[n] { false };
    for (int i = 0; i < n; i++)
    {
        bool detect = true;
        for (int j = 0; j < m; j++)
        {
            if (abs(a[i][j]) > 0.00000000001)
            {
                detect = false;
                break;
            }
        }
        correct[i] = detect;
        detect = true;
        for (int j = 0; j < m + 1; j++)
        {
            if (abs(a[i][j]) > 0.00000000001)
            {
                detect = false;
                break;
            }
        }
        Dcorrect[i] = detect;
    }

    for (int i = 0; i < n; i++) {
        if (correct[i] != Dcorrect[i])
        {
            std::cout << "NO" << std::endl;
            return 0;
        }
    }

    for (int i = 0; i < n; i++)
    {
        int n_new = n;
        int index = 0;
        for (int i = 0; i < n; i++)
        {
            if (correct[i])
            {
                n_new--;
            }
            for (int j = 0; j < m + 1; j++)
            {
                a[index][j] = a[i][j];
            }
            index++;
        }
        n = n_new;
    }

    if (n < m)
    {
        std::cout << "INF" << std::endl;
        return 0;
    }

    long double* Xanswer = new long double[m] {0.0};

    for (int i = n - 1; i >= 0; i--)
    {
        long double temp = a[i][m];
        for (int j = m - 1; j >= i; j--)
        {
            temp -= a[i][j] * Xanswer[j];
        }
        Xanswer[i] = temp;
    }

    for (int i = 0; i < m; i++) {
        std::cout << std::fixed << std::setprecision(16) << Xanswer[i] << " ";
    }
	return 0;
}
```
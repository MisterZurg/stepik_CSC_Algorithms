# Решение систем линейных алгебраических уравнений
Напишите программу, которая решает систему линейных алгебраических уравнений методом Гаусса.

## Формат входных данных:

В первой строке задаются два числа: количество уравнений `n (n ≥ 1)` и количество неизвестных `m (m ≥ 1)`. 
Далее идут `n` строк, каждая из которых содержит `m + 1` число. 
Первые `m` чисел — это коэффициенты `i-го` уравнения системы, а последнее, `(m + 1) - е` число — коэффициент `b_i` i, стоящий в правой части `i-го` уравнения.

## Формат выходных данных:
В первой строке следует вывести слово `YES`, если решение существует и единственно, слово `NO` в случае, если решение не существует, и слово `INF` в случае, когда решений существует бесконечно много. Если решение существует и единственно, то во второй строке следует вывести решение системы в виде m чисел, разделенных пробелом.

## Sample Input 1:
```
3 3
4 2 1 1
7 8 9 1
9 1 3 2
```
## Sample Output 1:
```
YES
0.2608695652173913 0.04347826086956526 -0.1304347826086957
```

## Sample Input 2:
```
2 3
1 3 4 4
2 1 4 5
```
## Sample Output 2:
```
INF
```
## Sample Input 3:
```
3 3
1 3 2 7
2 6 4 8
1 4 3 1
```
## Sample Output 3:
```
NO
```

```
#include <iostream>
#include <math.h>

using namespace std;

int
main()
{
    int n, m;
    cin >> n >> m;

    long double** a = new long double* [n];
    for (int i = 0; i < n; i++)
    {
        a[i] = new long double[m + 1];
    }
    //Заполнить матрицу
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < m + 1; j++)
        {
            cin >> a[i][j];
        }
    }

    //Вычислить верхнюю диагоональную матрицу
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
            cout << "NO" << endl;
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
        cout << "INF" << endl;
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


    cout << "YES" << endl;
    for (int i = 0; i < m; i++) {
        cout << Xanswer[i] << " ";
    }

    return 0;
}
```


```java
 // Gaussian elimination with partial pivoting
    public static double[] lsolve(double[][] A, double[] b) {
        int n = b.length;

        for (int p = 0; p < n; p++) {

            // find pivot row and swap
            int max = p;
            for (int i = p + 1; i < n; i++) {
                if (Math.abs(A[i][p]) > Math.abs(A[max][p])) {
                    max = i;
                }
            }
            double[] temp = A[p]; A[p] = A[max]; A[max] = temp;
            double   t    = b[p]; b[p] = b[max]; b[max] = t;

            // singular or nearly singular
            if (Math.abs(A[p][p]) <= EPSILON) {
                throw new ArithmeticException("Matrix is singular or nearly singular");
            }

            // pivot within A and b
            for (int i = p + 1; i < n; i++) {
                double alpha = A[i][p] / A[p][p];
                b[i] -= alpha * b[p];
                for (int j = p; j < n; j++) {
                    A[i][j] -= alpha * A[p][j];
                }
            }
        }

        // back substitution
        double[] x = new double[n];
        for (int i = n - 1; i >= 0; i--) {
            double sum = 0.0;
            for (int j = i + 1; j < n; j++) {
                sum += A[i][j] * x[j];
            }
            x[i] = (b[i] - sum) / A[i][i];
        }
        return x;
    }
```
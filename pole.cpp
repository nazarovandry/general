#include "stdio.h"

enum
{
    PO = 5,
    NVRJ = PO*PO*PO-PO,
    TYPES = PO*PO*PO+PO*PO+PO+1,
    BLOCKS = (PO*PO+1)*(PO*PO+PO+1)
};

class T
{
    int a;
public:
    T() { a = 0; }
    T(int to) { a = to%PO; }
    T(const T& to) { a = to.a; }
    T operator+(const T& other) const { return T((a + other.a)%PO); }
    T operator-(const T& other) const { return T((a + PO - other.a)%PO); }
    T operator*(const T& other) const { return T((a * other.a)%PO); }
    T op(const T& other, char c) const
    {
        if (c == '-')
            return (a + PO - other.a)%PO;
        else
            return (a + other.a)%PO;
    }
    const T& operator=(const T& other)
    {
        a = other.a;
        return *this;
    }
    const T& operator+=(const T& other)
    {
        a = (a + other.a)%PO;
        return *this;
    }
    const T& operator-=(const T& other)
    {
        a = (a + PO - other.a)%PO;
        return *this;
    }
    const T& operator*=(const T& other)
    {
        a = (a * other.a)%PO;
        return *this;
    }
    const T& neg()
    {
        for (int i=0; i<PO; ++i)
        {
            if (((a*i)%PO) == 1)
            {
                a = i;
                return *this;
            }
        }
        return *this;
    }
    bool operator==(const T& other) const { return a == other.a; }
    bool operator!=(const T& other) const { return a != other.a; }
    bool operator>(const T& other) const { return a > other.a; }
    int operator!() const { return a; }
    ~T() {}
};

struct Nabor
{
    T a11;
    T a12;
    T a21;
    T a22;
};

int take_num(Nabor x)
{
    T a(x.a11);
    T b(x.a12);
    T c(x.a21);
    T d(x.a22);
    T k;
    if (((k=a)>1) || ((a==0) && (((k=b)>1) || ((b==0) && (((k=c)>1)
        || ((c==0) && ((k=d)!=1)))))))
    {
        k.neg();
        a*=k;
        b*=k;
        c*=k;
        d*=k;
    }
    if (a == 0)
    {
        if (b == 0)
        {
            if (c == 0)
                return 0;
            return 1+(!d); 
        }
        return (PO+1)+(!c)*PO+(!d);
    }
    return (PO*PO+PO+1)+(!b)*PO*PO+(!c)*PO+(!d);
}


void print(Nabor n, char k='\n')
{
    printf("[%d%d%d%d]%c", !(n.a11), !(n.a12), !(n.a21), !(n.a22), k);
}

Nabor mul(Nabor x, Nabor y)
{
    Nabor ret;
    ret.a11 = (x.a11)*(y.a11)+(x.a12)*(y.a21);
    ret.a12 = (x.a11)*(y.a12)+(x.a12)*(y.a22);
    ret.a21 = (x.a21)*(y.a11)+(x.a22)*(y.a21);
    ret.a22 = (x.a21)*(y.a12)+(x.a22)*(y.a22);
    return ret;
}

Nabor add(Nabor x, Nabor y, int koef=1)
{
    Nabor ret;
    if (koef==2)
    {
    ret.a11 = (x.a11)+(y.a11)+(y.a11);
    ret.a12 = (x.a12)+(y.a12)+(y.a12);
    ret.a21 = (x.a21)+(y.a21)+(y.a21);
    ret.a22 = (x.a22)+(y.a22)+(y.a22);
    return ret;
    }
    ret.a11 = (x.a11)+(y.a11)*koef;
    ret.a12 = (x.a12)+(y.a12)*koef;
    ret.a21 = (x.a21)+(y.a21)*koef;
    ret.a22 = (x.a22)+(y.a22)*koef;
    return ret;
}

bool rav(Nabor x, Nabor y)
{
    return (x.a11 == y.a11) && (x.a12 == y.a12) && (x.a21 == y.a21) &&
        (x.a22 == y.a22);
}

int main()
{
    Nabor n[TYPES];
	int arr[TYPES][TYPES];
    int ch[BLOCKS+10][2];
    for (int i = 0; i < TYPES; ++i)
	{
	    (n[i]).a11 = ((i < (PO*PO+PO+1)) ? 0 : 1);
        (n[i]).a12 = ((i < (PO*PO+PO+1)) ? ((i+(PO*PO-PO-1))/PO/PO) :
            ((i-PO*PO-PO-1)/PO/PO));
        (n[i]).a21 = ((i == 0) ? 0 : ((i < (PO+1)) ? 1 : (((i-PO-1)/PO)%PO)));
        (n[i]).a22 = ((i == 0) ? 1 : ((i-1)%PO));
	}
    //for (int i=0; i<TYPES; ++i) {printf("%d: ", i); print(n[i], '\n');}
    int k = 0;
    for (int i = 0; i < TYPES; ++i)
    {
        for (int j = 0; j < i; ++j)
            arr[i][j] = arr[j][i];
        arr[i][i] = -1;
        for (int j = (i+1); j < TYPES; ++j)
        {
            int z[PO+1], tmp, first=0, second=0;
            tmp = z[0] = i;
            for (int q = 1; q < PO; ++q)
                z[q] = take_num(add(n[i], n[j], q));
            z[PO] = j;
            //if ((i<5) && (j<5)) for (int q=0;q<=PO;++q) print(n[z[q]], ' ');
            for (int q = 1; q <= PO; ++q)
            {
                if (z[q] < tmp)
                {
                    tmp = z[q];
                    first = q;
                }
            }
            tmp = TYPES;
            for (int q = 0; q <= PO; ++q)
            {
                if ((z[q] < tmp) && (q != first))
                {
                    tmp = z[q];
                    second = q;
                }
            }
            //if((i<5)&&(j<5)) printf("{%d %d %d %d %d %d  %d %d}\n", z[0],
            //    z[1], z[2], z[3], z[4], z[5], first, second);
            if ((first==0) && (second==PO))
            {
                arr[i][j] = k;
                ch[k][0] = z[0];
                ch[k][1] = z[PO];
                k++;
            }
            else
            {
                int l=0;
                while (l < k)
                {
                    if ((ch[l][0] == z[first]) && (ch[l][1] == z[second]))
                    {
                        arr[i][j] = l;
                        l = k;
                    }
                    l++;
                }
            }
        }
    }
    printf("K=%d\n", k);
    
    /*printf("\n");    
    for (int i=0; i<20; ++i)
    {
        for (int j=0; j<20; j++)
            printf("%c ", arr[i][j]+33);
        printf("\n");
    }*/

    int pq[NVRJ];
    int it=0;
    for (int i=0; i<TYPES; ++i)
    {
        if (((n[i]).a11 * (n[i]).a22 - (n[i]).a12 * (n[i]).a21) != 0)
        {
            Nabor obr;
            pq[it] = i;
            it++;
        }
    }

    /*for (int i=0;i<it;i++)
        printf("[%d %d\n %d %d] -- %d\n\n", !(n[pq[i]].a11), !(n[pq[i]].a12),
            !(n[pq[i]].a21), !(n[pq[i]].a22), it);*/

    int h=0;
    int eqv[BLOCKS];
    for (int i = 0; i < BLOCKS; ++i)
        eqv[i] = -1;
    for (int step = 0; step < BLOCKS; ++step)
    {
        if (eqv[step] == -1)
        {
            for (int i = 0; i < NVRJ; ++i)
            {
                for (int j = 0; j < NVRJ; ++j)
                {
                    Nabor tmp0 = mul(mul(n[pq[i]], n[ch[step][0]]), n[pq[j]]);
                    Nabor tmp1 = mul(mul(n[pq[i]], n[ch[step][1]]), n[pq[j]]);
                    eqv[arr[take_num(tmp0)][take_num(tmp1)]] = h;
                }
            }
            eqv[step] = h;
            h++;
        }
    }

    printf("klassov %d\n", h);

    int amount[h];
    int type[h][TYPES];
    for (int i = 0; i < h; ++i)
    {
        amount[i] = 0;
        for (int j = 0; j < TYPES; ++j)
            type[i][j] = 0;
    }
    for (int i = 0; i < BLOCKS; ++i)
    {
        /*if (eqv[i] == 0) {
            print(n[ch[i][0]]);
            print(n[ch[i][1]]);
            printf("NEXT:\n");
        }*/
        amount[eqv[i]]++;
        type[eqv[i]][ch[i][0]]++;
        for (int q = 1; q < PO; ++q)
            type[eqv[i]][take_num(add(n[ch[i][0]], n[ch[i][1]], q))]++;
        type[eqv[i]][ch[i][1]]++;
    }
    for (int i = 0; i < h; ++i)
        printf("[%d] ", amount[i]);
    printf("\n");
    for (int i = 0; i < h; ++i)
    {
        int first = type[i][0], second = type[i][0], kol1=0, kol2=0;
        for (int j = 0; j < TYPES; ++j)
        {
            if (type[i][j] == first)
                kol1++;
            else
            {
                second = type[i][j];
                kol2++;
            }
        }
        printf("'%d': %d    '%d': %d\n", first,  kol1, second, kol2);
    }
    
    printf("L3: %d\n", eqv[arr[PO*PO+PO+1][PO+PO+1]]);
    printf("L4: %d\n", eqv[arr[PO*PO+PO+1][PO+1]]);

    /*for (int i=0;i<TYPES;i++)
        {
        for(int j=0;j<TYPES;j++)
        {
            if (eqv[arr[i][j]] == 3)
            {
                print(n[ch[arr[i][j]][0]], ' ');
                print(n[ch[arr[i][j]][1]]);
            }
        }
    }*/

    //for (int i = 0; i < BLOCKS; ++i)
    //    printf("%d, %d, %d\n", ch[i][0], ch[i][1], eqv[i]);

    /*for (int i=0; i<BLOCKS; ++i)
        printf("%d ", eqv[i]);
    printf("\n");
    printf("\n");*/

    /*for (int i=0; i<TYPES; ++i)
    {
        printf("%d: ", i);
        print(n[i]);
    }
    printf("\n");*/

	return 0;
}
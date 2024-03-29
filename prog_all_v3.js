// language: JavaScript
// Copy and run, for example, here www.programiz.com/javascript/online-compiler/
// to get h_il(x) polynoms (non-zero ones only) change "showWrongTerms" parameter from "false" to "true" in tasks[] definition ("let tasks = ...")

// Язык: JavaScript
// Скопируйте и запустите, например, здесь www.programiz.com/javascript/online-compiler/
// Для получения значений полиномов h_il(x) (только тех, что ненулевые) замените значение параметра "showWrongTerms" с "false" на "true" в tasks[] ("let tasks = ...")

class Task {
    constructor(n, m, p, s, half = 'full', showWrongTerms = false) {
        this.n = n;
        this.m = m;
        this.p = p;
        this.s = s; // description of a task (tasks[])
        this.half = half; // if part of elements of matrix B are zero
        this.showWrongTerms = showWrongTerms;
    }
}

let tasks = [
    new Task(2, 3, 4, `D1 = (-y * A(1, 1) + 4 / 3 * A(1, 3) + 2 / 3 * x * A(2, 2) - 2 / 3 * x * A(2, 3)) * (2 / 3 * B(1, 2) - 7 / 12 * B(1, 3) + B(2, 1) + y * B(2, 3) + y * B(2, 4) - 0.375 * B(3, 1))
D2 = (-y * A(1, 1) - 2 / 3 * A(1, 2) + 1 / 3 * x * A(2, 2)) * (22 / 225 * x * B(1, 1) + 0.375 * B(2, 1) + 0.5 * y * B(2, 2) - 0.5 * y * B(2, 3) + 0.5 * y * B(2, 4) - 0.5 * y * B(3, 2))
D3 = (y * A(1, 1) - 4 / 3 * A(1, 3) + 2 / 3 * x * A(2, 3)) * (2 / 3 * B(1, 2) + 2 / 3 * B(1, 4) + B(2, 1) + y * B(2, 3) + y * B(2, 4) + y * B(3, 2) + y * B(3, 3) - y * B(3, 4))
D4 = (y * A(1, 1) + 4 / 3 * A(1, 2) + 2 / 3 * x * A(2, 3)) * (0.75 * B(1, 3) + 0.25 * B(2, 1) + 0.375 * B(3, 1) - y * B(3, 2))
D5 = (y * A(1, 1) + 4 / 3 * A(1, 2)) * (0.75 * B(1, 3) + 0.25 * B(2, 1) + y * B(2, 2) + 0.375 * B(3, 1) - y * B(3, 2))
D6 = (-y * A(1, 1) + 4 / 3 * A(1, 3) + 0.75 * A(2, 1) - x * A(2, 3)) * (7 / 12 * B(1, 3) + 4 / 9 * B(1, 4) + 0.5 * y * B(3, 2) - y * B(3, 4))
D7 = (-y * A(1, 1) - 2 / 3 * A(1, 2) + 2 / 3 * A(1, 3) + 1 / 3 * x * A(2, 2)) * (0.5 * y * B(2, 3) + 0.5 * y * B(2, 4) + 0.5 * y * B(3, 2))
D8 = (y * A(1, 1) - 2 / 3 * x * A(2, 2)) * (2 / 3 * B(1, 2) + 7 / 12 * B(1, 3) + 0.5 * B(2, 1) - y * B(2, 3) + y * B(2, 4) + 0.375 * B(3, 1))
D9 = (-y * A(1, 1) + 4 / 3 * A(1, 3)) * (2 / 9 * B(1, 4) + 0.375 * B(3, 1) + 0.5 * y * B(3, 2) + y * B(3, 3))
D10 = (-y * A(1, 1) + 0.75 * A(2, 1)) * (2 / 3 * B(1, 2) + 7 / 12 * B(1, 3) - 8 / 9 * B(1, 4) - 0.5 * y * B(2, 2) + y * B(2, 4) + 0.5 * y * B(3, 2) - y * B(3, 4))
D11 = (-y * A(1, 1) - 4 / 3 * A(1, 2) + 2 / 3 * x * A(2, 2)) * (0.5 * y * B(2, 2) + y * B(2, 4))
D12 = (y * A(1, 1)) * (-2 / 9 * B(1, 4) - 0.75 * B(2, 1) - 0.5 * y * B(2, 2) + y * B(2, 3) + 0.375 * B(3, 1) + 0.5 * y * B(3, 2) - y * B(3, 3))
D13 = (-y * A(1, 1) - 4 / 3 * A(1, 2) + 0.75 * A(2, 1) + 1 / 3 * x * A(2, 2)) * (0.5 * y * B(2, 2) - y * B(2, 4))
D14 = (2.25 * A(2, 1) + x * A(2, 3)) * (11 / 36 * B(1, 3) + 4 / 9 * B(1, 4) - 0.5 * y * B(3, 2) + y * B(3, 4))
D15 = (y * A(1, 1) - 2 / 3 * A(1, 3)) * (137 / 900 * x * B(1, 1) + 2 / 9 * B(1, 4) - 0.5 * y * B(2, 3) - 0.5 * y * B(2, 4) + y * B(3, 3))
D16 = (y * A(1, 1) + 2 / 3 * x * A(2, 3)) * (2 / 3 * B(1, 4) - 0.25 * B(2, 1) + 0.75 * B(3, 1) - y * B(3, 3) + y * B(3, 4))
D17 = (2.25 * A(2, 1)) * (4 / 9 * x * B(1, 1) - 2 / 3 * B(1, 2) + 11 / 36 * B(1, 3) + 0.5 * y * B(2, 2) - y * B(2, 4) - 0.5 * y * B(3, 2) + y * B(3, 4))
D18 = (2.25 * A(2, 1) - x * A(2, 2)) * (8 / 9 * B(1, 2) - 0.5 * y * B(2, 2) + y * B(2, 4))

C(1, 1) = -4 * D2 - 4 * D7 + 2 * D9 + 2 * D11 + 2 * D12 + 4 * D15
C(1, 2) = x * (-0.75 * D1 + 32 / 15 * D2 - 0.75 * D4 + 0.75 * D5 + 32 / 15 * D7 + 0.75 * D8 + 13 / 30 * D9 - 16 / 15 * D11 + 13 / 30 * D12 + 26 / 75 * D14 + 13 / 15 * D15 - 26 / 75 * D17 - 26 / 75 * D18)
C(1, 3) = x * (0.375 * D1 + 1.5 * D2 + 0.375 * D4 + 0.375 * D5 - 1.5 * D7 + 0.375 * D8 + 0.75 * D9 - 0.75 * D12)
C(1, 4) = x * (0.75 * D2 + 0.375 * D3 - 0.375 * D6 + 0.75 * D7 + 0.75 * D9 + 0.375 * D10 - 0.75 * D11 + 0.375 * D13 + 0.75 * D15 + 0.375 * D16)
C(2, 1) = y * (D1 + D3 + D4 - D5 + D6 - D8 + D9 - D10 - D11 - D12 - D13 - D14 + D16 + D17 + D18)
C(2, 2) = -1.5 * D4 + 1.5 * D5 + 1.5 * D11 + 1.5 * D13 + 0.5 * D18
C(2, 3) = 0.75 * D1 + 0.75 * D3 + 0.75 * D6 + 0.75 * D8 + 0.75 * D9 + 0.75 * D10 + 0.75 * D12 + 0.25 * D14 - 0.75 * D16 + 0.25 * D17
C(2, 4) = -0.75 * D4 + 0.75 * D5 + 0.75 * D11 + 0.75 * D13 + D14 - D17 - 0.75 * D18`, 'full', false),

    new Task(2, 5, 4, `D1 = (0.25 * y * A(1, 1) + 0.5 * x ^ 1 * A(1, 3) + 2 * x ^ 2 * A(1, 4) - 0.1 * x ^ 2 * A(2, 3) + 2 * x ^ 2 * A(2, 4)) * (4 * x ^ 1 * B(1, 2) + 2 * B(3, 1) + y * B(3, 2))
D2 = (y * A(2, 1) + A(2, 2)) * (2 * B(1, 3) - x ^ 1 * B(2, 1) + y * B(2, 3))
D3 = (-y * A(2, 1)) * (x ^ 2 * B(1, 1) + B(1, 3) + 0.125 * y * B(1, 4) - 0.5 * x ^ 1 * B(2, 1) + y * B(2, 3))
D4 = (y * A(1, 1) + 8 * x ^ 2 * A(1, 4) + 8 * x ^ 2 * A(2, 4)) * (x ^ 2 * B(1, 1) - x ^ 1 * B(1, 2) - 0.5 * B(3, 1) - 0.5 * y * B(3, 2) + 0.125 * y * B(4, 1))
D5 = (-0.25 * y * A(1, 1) + 0.02 * y * A(1, 2) + 0.5 * x ^ 1 * A(1, 3) - 2 * x ^ 2 * A(1, 4) + 0.3 * x ^ 2 * A(2, 3) - 2 * x ^ 2 * A(2, 4)) * (y * B(3, 2))
D6 = (-y * A(2, 1) + A(2, 2)) * (y * B(2, 3))
D7 = (0.05 * y * A(1, 2) - 0.02 * y * A(2, 1)) * (-50 * x ^ 2 * B(1, 2) + y * B(2, 3))
D8 = (-y * A(1, 2) + 0.4 * A(2, 2)) * (-1.25 * x ^ 2 * B(1, 2) + 2.5 * x ^ 2 * B(2, 1) + 2.5 * x ^ 1 * B(2, 2) + 0.05 * y * B(2, 3))
D9 = (y * A(1, 2) + 20 * x ^ 2 * A(2, 3)) * (50 * x ^ 2 * B(2, 1) + y * B(3, 2))
D10 = (y * A(1, 2)) * (1.25 * x ^ 2 * B(1, 2) + 2.5 * x ^ 1 * B(2, 2) - 0.05 * y * B(3, 2))
D11 = (y * A(1, 1) + 0.25 * A(2, 2)) * (4 * B(1, 3) + 2 * x ^ 1 * B(2, 1))
D12 = (-y * A(1, 1)) * (4 * B(1, 3) - 0.25 * y * B(1, 4) + 2 * x ^ 1 * B(2, 1) + 0.25 * y * B(3, 1) + 0.125 * y * B(4, 1))
D13 = (y * A(1, 1) + 4 * x ^ 2 * A(2, 3)) * (0.25 * y * B(3, 1))
D14 = (y * A(1, 1) + 8 * x ^ 2 * A(2, 4)) * (y * B(4, 1))
D15 = (0.5 * y * A(1, 1) + 0.25 * y * A(2, 1)) * (y * B(1, 4))

C(1, 1) = y * D1 + y * D4 - y * D5 + 0.02 * y * D9 - 0.125 * y * D14
C(1, 2) = D1 + D5 + 0.4 * D10 - 2 * x ^ 1 * D13
C(1, 3) = -0.4 * x ^ 2 * D6 + 20 * x ^ 2 * D7 + 0.25 * x ^ 1 * D11
C(1, 4) = 2 * x ^ 2 * D2 + 4 * x ^ 2 * D3 - 2 * x ^ 2 * D6 + 2 * x ^ 2 * D15
C(2, 1) = -0.5 * y * D2 - y * D3 + 0.5 * y * D6 + y * D11 + y * D12 + y * D13 + 0.125 * y * D14 - 0.5 * y * D15
C(2, 2) = -0.02 * y * D6 + y * D7 + y * D8 + 0.05 * y * D9 + y * D10
C(2, 3) = 0.5 * x ^ 1 * D2 + 0.5 * x ^ 1 * D6
C(2, 4) = -4 * x ^ 2 * D2 - 8 * x ^ 2 * D3 + 4 * x ^ 2 * D6`, 'tri', false),

    new Task(2, 4, 4, `D1 = (x ^ 2 * A(1, 2) - y * A(2, 1) - x ^ 2 * A(2, 2) + x ^ 2 * A(2, 3)) * (1.6 * x * B(1, 2) - 1.25 * B(1, 4) + y * B(3, 1))
D2 = (-y * A(1, 1) + 2 * x ^ 2 * A(2, 2)) * (1.5 * x * B(1, 1) - B(1, 3) + y * B(2, 2))
D3 = (-x ^ 2 * A(1, 2) + y * A(2, 1) + x ^ 2 * A(2, 2)) * (x ^ 2 * B(1, 1) + 1.6 * x * B(1, 2) - 1.25 * B(1, 4) + y * B(2, 1) + y * B(3, 1))
D4 = (y * A(1, 1) - 2 * x ^ 2 * A(2, 3)) * (0.75 * x * B(1, 1) - 0.5 * B(1, 3) + 0.05 * y * B(1, 4) - 0.5 * y * B(3, 2))
D5 = (y * A(1, 1) + x ^ 2 * A(1, 2)) * (-y * B(2, 2))
D6 = (-y * A(1, 1) + 15.625 * x * A(2, 1) + 2 * x ^ 2 * A(2, 3)) * (y * B(1, 4))
D7 = (x ^ 2 * A(1, 3) - y * A(2, 1)) * (y * B(3, 1))
D8 = (8 / 3 * x * A(1, 1) + 0.25 * y * A(2, 1)) * (y * B(1, 3))
D9 = (y * A(1, 1)) * (y * B(2, 2) + y * B(3, 2))
D10 = (x ^ 2 * A(1, 2) - y * A(2, 1)) * (y * B(2, 1))
D11 = (y * A(1, 1) + x ^ 2 * A(1, 3)) * (x ^ 2 * B(1, 2) + y * B(3, 2))
D12 = (y * A(2, 1)) * (-0.0625 * y * B(1, 3) + y * B(2, 1) + y * B(3, 1))

C(1, 1) = -2 / 3 * D2 + 2 / 3 * D5 + y * D7 + 0.25 * y * D8 + y * D10 + y * D12
C(1, 2) = -y * D5 - y * D9 + y * D11
C(1, 3) = x * D2 - x * D5
C(1, 4) = -x ^ 2 * D6
C(2, 1) = y * D1 + y * D3 + y * D10
C(2, 2) = -0.625 * D1 + 0.5 * y * D2 + y * D4 + 0.05 * y * D6 + 0.625 * D7 + 0.5 * y * D9
C(2, 3) = 4 * x ^ 2 * D8
C(2, 4) = 0.8 * x * D1 - 0.8 * x * D7`, '2tri', false),
];

// term (c/d)(X^x), for examlpe {c = 2, d = 3, x = 4}: 2/3*(x^4) 
class Term {
    constructor(c = 1, d = 1, x = 0) {
        this.c = c;
        this.d = d;
        this.x = x;
    }
    print() {
        console.log(this.toString());
    }
    toString() {
        let sc = (this.c == 1 && this.d == 1 && this.x != 0 ? '' : this.c);
        let sd = (this.d == 1 ? '' : '/' + this.d);
        let star = ((this.x != 0)
            && (this.c != 1 || this.d != 1) ? '*' : '');
        let sx = (this.x == 0 ? ''
            : (this.x == 1 ? 'x' : '(x^' + this.x + ')'));
        return sc + sd + star + sx;
    }
    mul(c = 1, d = 1, x = 0) {
        this.c *= c;
        this.d *= d;
        this.x += x;
    }
    safeMul(term) {
        return new Term(this.c * term.c, this.d * term.d, this.x + term.x);
    }
    pow(num) {
        //this.c = Math.pow(this.c, num);
        this.x *= num;
    }
    same(term) {
        if (this.c * term.d == this.d * term.c && this.x == term.x)
            return 1
        if (this.c * term.c == this.d * term.d && this.x == -term.x)
            return -1;
        return 0;
    }
}

// Matrix of terms for each a[]b[]
class Mb {
    constructor(m, p) {
        this.B = [-1];
        for (let k = 1; k <= m; k++) {
            this.B.push([-1]);
            for (let l = 1; l <= p; l++) {
                this.B[k].push([]);
            }
        }
    }
    print(i, j, k, l) {
        let str = ('a' + i + j + ' b' + k + l + ' = ');
        for (let ind = 0; ind < this.B[k][l].length; ind++) {
            str += (this.B[k][l][ind].toString() + ' + ');
        }
        if (str.length > 2) {
            str = str.slice(0, -3);
        } else str += '0';
        console.log(str);
    }
}

// Matrix of coeficients a[]b[] for each C(i, j)
class Ma {
    constructor(n, m, p) {
        this.n = n;
        this.m = m;
        this.p = p;
        this.A = [-1];
        for (let i = 1; i <= n; i++) {
            this.A.push([-1]);
            for (let j = 1; j <= m; j++) {
                this.A[i].push(0);
            }
        }
        for (let i = 1; i <= n; i++)
            for (let j = 1; j <= m; j++)
                this.A[i][j] = new Mb(m, p);
    }
    addOnes(half, i, l) {
        for (let j = 1; j <= this.m; j++) {
            if (half == 'tri' && l + j > this.m) continue;
            if (half == '2tri' &&
                (j % 2 == 0 && l + j > this.m
                || j % 2 == 1 && l + j > this.m + 1))
                continue;
            this.A[i][j].B[j][l].push(new Term(-1, 1)); // push -1
        }
    }
    print() {
        for (let i = 1; i <= this.n; i++)
            for (let j = 1; j <= this.m; j++)
                for (let k = 1; k <= this.m; k++)
                    for (let l = 1; l <= this.p; l++)
                        if (this.A[i][j].B[k][l].length > 0)
                            this.A[i][j].print(i, j, k, l);
    }
    reduce(S) {
        for (let i = 1; i <= this.n; i++)
            for (let j = 1; j <= this.m; j++)
                for (let k = 1; k <= this.m; k++)
                    for (let l = 1; l <= this.p; l++) {
                        this.A[i][j].B[k][l] = reducePlus(this.A[i][j].B[k][l]);
                        this.A[i][j].B[k][l] = reduceNod(this.A[i][j].B[k][l]);
                        if (this.A[i][j].B[k][l].length > 0) {
                            let err = this.A[i][j].B[k][l][0];
                            S.terms.push(err.safeMul(err));
                        }
                    }
    }
}

// S(x)
class Sx {
    constructor() {
        this.terms = [];
    }
    print() {
        let str = 'S(x) = ';
        for (let ind = 0; ind < this.terms.length; ind++) {
            str += (this.terms[ind].toString() + ' + ');
        }
        if (str.length > 2) {
            str = str.slice(0, -3);
        } else str += '0';
        console.log(str);
    }
    reduce() {
        this.terms = reducePlus(this.terms);
        this.terms = reduceNod(this.terms);
    }
}

// reduce (ck)/(dk) to c/d
function reduceNod(check) {
    check.forEach(term => {
        let nod = NOD(abs(term.c), term.d);
        term.c /= nod;
        term.d /= nod;
    });
    return check;
}

// reduce a(X^k)+b(X^k) to (a+b)(X^k)
function reducePlus(check) {
    for (let p = 0; p < check.length; p++) {
        for (let q = p + 1; q < check.length; q++) {
            if (check[p].x == check[q].x) {
                let up = check[p].c * check[q].d +
                    check[p].d * check[q].c;
                let down = check[p].d * check[q].d;
                let nod = NOD(abs(up), down)
                check[q].c = up / nod;
                check[q].d = down / nod;
                check[p] = null;
                break;
            }
        }
    }
    let tmp = [];
    for (let p = 0; p < check.length; p++)
        if (check[p] && check[p].c != 0)
            tmp.push(check[p]);
    return tmp;
}

// number "x" from string -> to pair [c, d], where x = c/d
function getNum(num) {
    let i = num.indexOf('.');
    if (i == -1)
        return [Number(num), 1];
    let z = Number(num.slice(0, i));
    let r = num.slice(i+1);
    let down = Math.pow(10, r.length);
    let up = Number(r) + down * z;
    let nod = NOD(up, down);
    return[up / nod, down / nod];
}

function abs(n) {
    return n >= 0 ? n : -n;
}

function NOD(n, m) {
    return m == 0 ? n : NOD(m, n % m);
}

// get terms from "D(par) = ..." string and add to the Matrix
function readD(s, M, par) {
    let Am = [], Bm = []; // coefficients in both brackets for "D(par) = ..."
    let i = s.indexOf('D' + par[0]) + 4;
    let triple = [0, 0, new Term()]; // for converting t*A(i, j) -> [i, j, t]
    let firstMinus = true, num = '', frac = false, pow = false;
    while (i < s.length && s[i] != '\n') {
        i++
        switch (s[i]) {
        case 'x':
            triple[2].mul(1, 1, 1);
            break;
        case 'y':
            triple[2].mul(1, 1, -1);
            break;
        case 'A':
        case 'B':
            firstMinus = false;
            triple[0] = Number(s[i+2]);
            triple[1] = Number(s[i+5]);
            i += 6;
            break;
        case '+':
            Bm.push(triple);
            triple = [0, 0, new Term()];
            break;
        case '-':
            if (firstMinus) {
                triple[2].mul(-1);
            } else {
                Bm.push(triple);
                triple = [0, 0, new Term(-1)];
            }
            break;
        case ')':
            Bm.push(triple);
            if (s[i+1] == '\n')
                break;
            triple = [0, 0, new Term()];
            Am = Bm;
            Bm = [];
            firstMinus = true;
            break;
        case '0': case '1': case '2': case '3': case '4':
        case '5': case '6': case '7': case '8': case '9':
            num = '';
            while (s[i] != ' ') {
                num += s[i];
                i++;
            }
            num = getNum(num);
            if (frac) {
                triple[2].mul(1, num[0]);
                frac = false;
            } else if (pow) {
                triple[2].pow(num[0]);
                pow = false;
            } else {
                triple[2].mul(num[0], num[1]);
            }
            break;
        case '/':
            frac = true;
            break;
        case '^':
            pow = true;
        }
    }
    
    for (a = 0; a < Am.length; a++)
        for (b = 0; b < Bm.length; b++)
            M.A[Am[a][0]][Am[a][1]].B[Bm[b][0]][Bm[b][1]]
                .push(par[1].safeMul(Am[a][2]).safeMul(Bm[b][2]));
}

// for "C(i, j) = ..." string get the list of D(k) used there
function readC(s, x, y) {
    let D = []; // list of D[k]
    let i = s.indexOf('C(' + x + ', ' + y + ')') + 7;
    let pair = [0, new Term()]; // for converting t*D(k) -> [k, t]
    let factor = undefined; // if C(i, j) = factor * (...)
    let firstMinus = true, num = '', frac = false, pow = false;
    while (i < s.length && s[i] != '\n') {
        i++;
        switch (s[i]) {
        case '(':
            factor = pair[1].x;
            pair[1] = new Term();
            break;
        case 'x':
            pair[1].mul(1, 1, 1);
            break;
        case 'y':
            pair[1].mul(1, 1, -1);
            break;
        case 'D':
            firstMinus = false;
            num = '';
            i++;
            while (s[i] >= '0' && s[i] <= '9') {
                num += s[i];
                i++;
            }
            pair[0] = Number(num);
            break;
        case '+':
            D.push(pair);
            pair = [0, new Term()];
            break;
        case '-':
            if (firstMinus) {
                pair[1].mul(-1);
            } else {
                D.push(pair);
                pair = [0, new Term(-1)];
            }
            break;  
        case '0': case '1': case '2': case '3': case '4':
        case '5': case '6': case '7': case '8': case '9':
            num = '';
            while (s[i] != ' ') {
                num += s[i];
                i++;
            }
            num = getNum(num);
            if (frac) {
                pair[1].mul(1, num[0]);
                frac = false;
            } else if (pow) {
                pair[1].pow(num[0]);
                pow = false;
            } else {
                pair[1].mul(num[0], num[1]);
            }
            break;
        case '/':
            frac = true;
            break;
        case '^':
            pow = true;
        }
    }
    D.push(pair);
    if (factor != 0)
        for (j = 0; j < D.length; j++)
            D[j][1].mul(1, 1, factor);
    return D;
}

tasks.forEach(task => {
let S = new Sx(); // S(x)
    console.log('Task <' + task.n + ', ' + task.m + ', ' + task.p + '>');
    for (let i = 1; i <= task.n; i++) {
        for (let j = 1; j <= task.p; j++) {
            D = readC(task.s, i, j);
            
            // create matrix of coeficients a[]b[] for C(i, j)
            let M = new Ma(task.n, task.m, task.p);
            
            for (k = 0; k < D.length; k++) {
                readD(task.s, M, D[k]);
            }
            
            if (task.showWrongTerms)
            console.log('C' + i + j + '  wrong a[]b[] terms:');
            
            // add (-1) terms for a[iw]b[wj] coeficients (w is any)
            M.addOnes(task.half, i, j);
            // reduce terms and get S(x)
            M.reduce(S);
            // print non-zero coeficients
            if (task.showWrongTerms)
                M.print();
        }
    }
    S.reduce();
    S.print();
});
console.log('DONE');

# Join – Union – Agregasi – Subquery – Function (DBMS)

## 1. **Join**
**Join** adalah operasi yang digunakan untuk menggabungkan dua atau lebih tabel berdasarkan kolom terkait antara tabel tersebut.

### Tipe-Tipe Join:
- **Inner Join**: Mengembalikan baris yang memiliki nilai yang cocok di kedua tabel.
  ```sql
  SELECT * 
  FROM table1 
  INNER JOIN table2 ON table1.id = table2.id;
  ```

- **Left (Outer) Join**: Mengembalikan semua baris dari tabel kiri, dan data yang cocok dari tabel kanan. Jika tidak ada kecocokan, nilai dari tabel kanan akan `NULL`.
  ```sql
  SELECT * 
  FROM table1 
  LEFT JOIN table2 ON table1.id = table2.id;
  ```

- **Right (Outer) Join**: Mengembalikan semua baris dari tabel kanan, dan data yang cocok dari tabel kiri. Jika tidak ada kecocokan, nilai dari tabel kiri akan `NULL`.
  ```sql
  SELECT * 
  FROM table1 
  RIGHT JOIN table2 ON table1.id = table2.id;
  ```

- **Full (Outer) Join**: Mengembalikan semua baris dari kedua tabel, dengan nilai `NULL` jika tidak ada kecocokan di salah satu tabel.
  ```sql
  SELECT * 
  FROM table1 
  FULL OUTER JOIN table2 ON table1.id = table2.id;
  ```

## 2. **Union**
**Union** digunakan untuk menggabungkan hasil dari dua atau lebih query SELECT. Hasil penggabungan akan menampilkan data secara vertikal dan menghilangkan duplikasi.

### Contoh penggunaan:
```sql
SELECT column_name FROM table1
UNION
SELECT column_name FROM table2;
```

### Catatan:
- Kedua query harus memiliki jumlah kolom yang sama.
- Tipe data di setiap kolom yang digabungkan harus kompatibel.

## 3. **Agregasi**
Operasi agregasi digunakan untuk melakukan kalkulasi pada sekumpulan nilai dan mengembalikan satu nilai tunggal. Fungsi agregasi yang umum meliputi:

- **COUNT()**: Menghitung jumlah baris.
  ```sql
  SELECT COUNT(*) FROM table_name;
  ```

- **SUM()**: Menjumlahkan nilai dari suatu kolom.
  ```sql
  SELECT SUM(column_name) FROM table_name;
  ```

- **AVG()**: Menghitung rata-rata nilai dari suatu kolom.
  ```sql
  SELECT AVG(column_name) FROM table_name;
  ```

- **MIN()**: Mengembalikan nilai minimum dari suatu kolom.
  ```sql
  SELECT MIN(column_name) FROM table_name;
  ```

- **MAX()**: Mengembalikan nilai maksimum dari suatu kolom.
  ```sql
  SELECT MAX(column_name) FROM table_name;
  ```
## 4. **Subquery**
**Subquery** adalah query di dalam query lain. Subquery dapat berada dalam klausa **SELECT**, **FROM**, atau **WHERE**.

### Contoh penggunaan Subquery:
- **Subquery dalam WHERE**:
  ```sql
  SELECT column_name
  FROM table_name
  WHERE column_name = (SELECT column_name FROM another_table WHERE condition);
  ```

- **Subquery dalam SELECT**:
  ```sql
  SELECT column_name, (SELECT AVG(salary) FROM employees) AS avg_salary
  FROM employees;
  ```

- **Subquery dalam FROM** (dikenal juga sebagai Derived Table):
  ```sql
  SELECT sub.column_name
  FROM (SELECT * FROM table_name WHERE condition) AS sub;
  ```
## **Soal Prioritas 2 (20)**

1. Buatlah sebuah program untuk mencetak sebuah pola berdasarkan faktor bilangan. Jika faktor bilangan merupakan bilangan genap maka gunakan simbol “I” (huruf i besar) dan selain itu menggunakan simbol huruf O.
    
    Input: 26
    
    Output: OIOI
    
    Input: 35
    
    Output: OOOO
    
2. Buatlah sebuah program untuk menentukan prioritas dari sebuah proyek berdasarkan budget, waktu pengerjaan dan tingkat kesulitan. Kriteria dari penentuan proyek adalah menggunakan skor secara keseluruhan. Perhitungan skor dari setiap faktor tersebut dihitung dengan kriteria berikut:
    
    
    | **Budget** | **Skor** |
    | --- | --- |
    | 0 - 50 | 4 |
    | 51 - 80 | 3 |
    | 81 - 100 | 2 |
    |  > 100 | 1 |
    
    | **Waktu Pengerjaan** | **Skor** |
    | --- | --- |
    | 0 - 20 | 10 |
    | 21 - 30 | 5 |
    | 31 - 50 | 2 |
    |  > 50 | 1 |
    
    | **Tingkat Kesulitan** | **Skor** |
    | --- | --- |
    | 0 - 3 | 10 |
    | 4 - 6 | 5 |
    | 8 - 10 | 1 |
    |  > 10 | 0 |
    
    Hasil prioritas proyek diambil dari skor yang sudah dihitung.
    
    | **Skor** | **Hasil** |
    | --- | --- |
    | 24 - 17 | High |
    | 16 - 10 | Medium |
    | 9 - 3 | Low |
    | <= 2 | Impossible |
    
    Test case:
    
    Input:
    
    - Budget: 40
    - Waktu pengerjaan: 25
    - Tingkat kesulitan: 5
    
    Output: Medium
    
    Input:
    
    - Budget: 51
    - Waktu pengerjaan: 10
    - Tingkat kesulitan: 3
    
    Output: High
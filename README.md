# Concurrently

## Concurrent Multiway Merge Sort

- Demonstration clip: <https://youtu.be/uFjDycXaLWE>
- Citation: <https://en.wikipedia.org/wiki/Merge_sort#Parallel_multiway_merge_sort>

```bash
Init records ...
Number of records: 100,000,000; Init time: 1.375897s

Running STD Merge Sort ...
STD Merge Sort: 16.481343s

Running Concurrent Multiway Merge Sort; Threads use: 8/8 ...
Concurrent Mulitway Merge Sort: 3.545462s

Faster: 4.6x times
```

## Test Machine

```bash
            .-/+oossssoo+/-.               lavantien@savaka
        `:+ssssssssssssssssss+:`           ----------------
      -+ssssssssssssssssssyyssss+-         OS: Ubuntu 23.04 x86_64
    .ossssssssssssssssssdMMMNysssso.       Host: HP ProBook 445 G7
   /ssssssssssshdmmNNmmyNMMMMhssssss/      Kernel: 6.2.0-20-generic
  +ssssssssshmydMMMMMMMNddddyssssssss+     Uptime: 20 hours, 44 mins
 /sssssssshNMMMyhhyyyyhmNMMMNhssssssss/    Packages: 2487 (dpkg), 167 (brew), 6 (flatpak), 26 (snap)
.ssssssssdMMMNhsssssssssshNMMMdssssssss.   Shell: zsh 5.9
+sssshhhyNMMNyssssssssssssyNMMMysssssss+   Resolution: 1920x1080
ossyNMMMNyMMhsssssssssssssshmmmhssssssso   DE: GNOME 44.0
ossyNMMMNyMMhsssssssssssssshmmmhssssssso   WM: Mutter
+sssshhhyNMMNyssssssssssssyNMMMysssssss+   WM Theme: Adwaita
.ssssssssdMMMNhsssssssssshNMMMdssssssss.   Theme: Yaru-dark [GTK2/3]
 /sssssssshNMMMyhhyyyyhdNMMMNhssssssss/    Icons: Yaru [GTK2/3]
  +sssssssssdmydMMMMMMMMddddyssssssss+     Terminal: WezTerm
   /ssssssssssshdmNNNNmyNMMMMhssssss/      CPU: AMD Ryzen 7 4700U with Radeon Graphics (8) @ 2.000GHz
    .ossssssssssssssssssdMMMNysssso.       GPU: AMD ATI 05:00.0 Renoir
      -+sssssssssssssssssyyyssss+-         Memory: 2072MiB / 15374MiB
        `:+ssssssssssssssssss+:`
            .-/+oossssoo+/-.

```

# Concurrently

## Concurrent Multiway Merge Sort

```bash
Init records ...
Number of records: 500,000,000; Init time: 8.516280s
Running STD Merge Sort...
STD Merge Sort: 78.424745s
Running Concurrent Multiway Merge Sort; Threads use: 16/20 ...
Concurrent Merge Sort: 17.788759s
Faster: 4.4x times
```

## Test Machine

```bash
            .-/+oossssoo+/-.               lavantien@savaka
        `:+ssssssssssssssssss+:`           ----------------
      -+ssssssssssssssssssyyssss+-         OS: Ubuntu 23.04 x86_64
    .ossssssssssssssssssdMMMNysssso.       Host: MS-7D42 1.0
   /ssssssssssshdmmNNmmyNMMMMhssssss/      Kernel: 6.2.0-20-generic
  +ssssssssshmydMMMMMMMNddddyssssssss+     Uptime: 3 hours, 29 mins
 /sssssssshNMMMyhhyyyyhmNMMMNhssssssss/    Packages: 2530 (dpkg), 162 (brew), 6 (flatpak), 26 (snap)
.ssssssssdMMMNhsssssssssshNMMMdssssssss.   Shell: zsh 5.9
+sssshhhyNMMNyssssssssssssyNMMMysssssss+   Resolution: 3840x2160
ossyNMMMNyMMhsssssssssssssshmmmhssssssso   DE: GNOME 44.0
ossyNMMMNyMMhsssssssssssssshmmmhssssssso   WM: Mutter
+sssshhhyNMMNyssssssssssssyNMMMysssssss+   WM Theme: Adwaita
.ssssssssdMMMNhsssssssssshNMMMdssssssss.   Theme: Yaru-dark [GTK2/3]
 /sssssssshNMMMyhhyyyyhdNMMMNhssssssss/    Icons: Yaru [GTK2/3]
  +sssssssssdmydMMMMMMMMddddyssssssss+     Terminal: WezTerm
   /ssssssssssshdmNNNNmyNMMMMhssssss/      CPU: 12th Gen Intel i7-12700F (20) @ 4.800GHz
    .ossssssssssssssssssdMMMNysssso.       GPU: NVIDIA GeForce RTX 3080 Lite Hash Rate
      -+sssssssssssssssssyyyssss+-         Memory: 4534MiB / 31930MiB
        `:+ssssssssssssssssss+:`
            .-/+oossssoo+/-.

```

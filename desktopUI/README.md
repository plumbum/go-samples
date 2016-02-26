Andalabs UI
===========

[UI](https://github.com/andlabs/ui) минималистичная кросплатформенная обёртка для создания графических интерфейсов.

Для своей работы использует другую библиотеку [LibUI](https://github.com/andlabs/libui) того же автора.

На данный момент библиотека не блещет функционалом, но тем не менее позволяет делать небольшие приложения.

Так же несколько раз у меня вылетала моя программа, и судя по всему это проблема в `libui` а не в обёртке.

Установка LibUI
---------------

Без `libui` работать не будет. Даже собираться не будет.

```
$ cd ~/src/
$ git clone https://github.com/andlabs/libui.git
Cloning into 'libui'...
$ cd libui/
$ make
====== Compiled common/areaevents.c
...
====== Linked out/libui.so.0
$ sudo make install
cp out/libui.so.0 /usr/lib/libui.so.0
ln -fs libui.so.0 /usr/lib/libui.so
cp ui.h ui_unix.h /usr/include/
$ 
```

Не забудьте, что собранная библиотека должна распространяться вместе с приложением

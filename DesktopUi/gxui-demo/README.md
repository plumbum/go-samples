Gxui time example
=================

Небольшой пример использования библиотеки конструирования пользовательских интерфейсов [GXUI](https://github.com/google/gxui).

Следует иметь в виду, что:

  * Библиотека больше не поддерживается разработчиками. Последний актуальный коммит был от 29 сентября 2016 года.
    С другой стороны её [форкнули](https://github.com/nelsam/gxui), так что возможно ещё не всё потеряно ;-)
  * Библиотека использует [OpenGL](https://www.opengl.org/) для рендеринга, так что на устройствах
    без аппаратной поддержки интерфейс будет работать медленно и печально.


Быстрая установка
-----------------

Ставим в систему необходимые библиотеки:

	sudo apt-get install libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev mesa-common-dev libgl1-mesa-dev libxxf86vm-dev

Устанавливаем go библиотеку:

	go get -u github.com/google/gxui/...


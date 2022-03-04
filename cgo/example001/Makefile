main: main.go mylib.a
	go build main.go

mylib.a: mylib.c mylib.h
	gcc -c -o mylib.o mylib.c
	ar rcs mylib.a mylib.o

clean:
	rm *.a *.o main

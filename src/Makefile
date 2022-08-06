redforest: redforest.cpp
	g++ -c interface/MainInterface.cpp 
	g++ -c redforest.cpp
	g++ redforest.o MainInterface.o -L interface -o redforest

clean:
	rm -f redforest *.o 

install:
	cp redforest /usr/local/bin


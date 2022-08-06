redforest: redforest.cpp
	g++ -c interface/MainInterface.cpp 
	g++ -c interface/commands/Command.cpp
	g++ -c interface/commands/Sock.cpp
	g++ -c redforest.cpp
	g++ -o redforest redforest.o MainInterface.o Command.o Sock.o

clean:
	rm -f redforest *.o 

install:
	cp redforest /usr/local/bin


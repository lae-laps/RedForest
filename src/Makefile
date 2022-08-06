redforest: redforest.cpp
	g++ -c interface/MainInterface.cpp 
	g++ -c interface/commands/Command.cpp
	g++ -c interface/Ascii.cpp
	g++ -c lib/Target.cpp
	g++ -c redforest.cpp
	g++ -o redforest redforest.o MainInterface.o Command.o Target.o Ascii.o

release:
	# Copy code from normal make
	make
	strip redforest

clean:
	rm -f redforest *.o 

install:
	cp redforest /usr/local/bin


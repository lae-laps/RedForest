redforest: redforest.cpp
	g++ -c interface/MainInterface.cpp 
	g++ -c interface/commands/Command.cpp
	g++ -c interface/Ascii.cpp
	g++ -c lib/Target.cpp
	g++ -c redforest.cpp
	g++ -o redforest redforest.o MainInterface.o Command.o Target.o Ascii.o

release:
	# run normal make
	make
	# strip debug symbols
	strip redforest

clean:
	rm -f redforest *.o 

install:
	make release
	cp redforest /usr/local/bin


redforest: redforest.cpp
	g++ -c interface/Interface.cpp 
	g++ -c interface/commands/Command.cpp
	g++ -c interface/commands/Sock.cpp
	g++ -c interface/ascii.cpp
	g++ -c lib/Target.cpp
	g++ -c redforest.cpp
	g++ -o redforest redforest.o Interface.o Command.o Target.o ascii.o Sock.o

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

run:
	make
	# Running ./redforest
	./redforest


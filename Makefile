wiki: src/wiki/*.go
	gb build wiki

clean:
	rm bin/wiki
	find pkg -name "*.a" -exec rm -f {} \;

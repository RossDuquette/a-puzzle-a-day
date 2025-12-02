clean:
	@rm -f a-puzzle-a-day

%:
	./docker_run make -f target.mk $@


all:
	@(cd sipgen; $(MAKE))
	@(cd siplib; $(MAKE))

install:
	@(cd sipgen; $(MAKE) install)
	@(cd siplib; $(MAKE) install)
	@test -d $(DESTDIR)/usr/lib/python3/dist-packages || mkdir -p $(DESTDIR)/usr/lib/python3/dist-packages
	cp -f sipconfig.py $(DESTDIR)/usr/lib/python3/dist-packages/sipconfig.py
	cp -f /home/lwh/下载/sip-4.15.5/sipdistutils.py $(DESTDIR)/usr/lib/python3/dist-packages/sipdistutils.py

clean:
	@(cd sipgen; $(MAKE) clean)
	@(cd siplib; $(MAKE) clean)
